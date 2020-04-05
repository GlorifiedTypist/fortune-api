package command

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/heptiolabs/healthcheck"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)


func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.AddCommand(initWarmupCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialises the local cache.",
}

var initWarmupCmd = &cobra.Command{
	Use:   "cache",
	Short: "Initialises the local cache.",
	RunE:  runInit,
}


func runInit(cmd *cobra.Command, args []string) error {

	var redis redis.Conn
	redis = connectToDatabase()

	health := healthcheck.NewHandler()
	health.AddReadinessCheck("redis", healthcheck.TCPDialCheck("localhost:6379", 1*time.Second))

	adminMux := http.NewServeMux()

	go http.ListenAndServe("0.0.0.0:9402", adminMux)

	adminMux.HandleFunc("/live", health.LiveEndpoint)
	adminMux.HandleFunc("/ready", health.ReadyEndpoint)

	log.Println("Begin warming cache.")
	fortunes := getFortunesFromURL()

	for i := 0;  i<=len(fortunes) - 1; i++ {

		time.Sleep(100 * time.Millisecond)

		key  := fmt.Sprintf("fortune:%d", i)
		redis.Do("HMSET", key, "fortune", strings.ReplaceAll(fortunes[i], "\n", ""))

		log.Printf("Fortune [%d]: %s", i, strings.ReplaceAll(fortunes[i], "\n", ""))

		health.AddReadinessCheck("cache-position", healthcheck.Async(func() error {
			if len(fortunes) -1 != i {
				err := fmt.Errorf("Cache not ready, size: %d, position: %d", len(fortunes), i)

				return err
			}
			return nil
		}, 50*time.Millisecond))
	}

	log.Println("Warming cache, completed.")

	return nil
}

func getFortunesFromURL() []string {
	res, err := http.Get("https://raw.githubusercontent.com/ruanyf/fortunes/master/data/fortunes")
	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return strings.Split(string(body), "%")
}

func connectToDatabase() redis.Conn {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	return conn
}
