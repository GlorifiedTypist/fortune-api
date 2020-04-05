package command

import (
	"encoding/json"
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gomodule/redigo/redis"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/go-openapi/loads"
	"github.com/spf13/cobra"
	"github.com/heptiolabs/healthcheck"

	"fortune-api/pkg/swagger/server/restapi"
	"fortune-api/pkg/swagger/server/restapi/operations"
)

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.AddCommand(serverRunCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "API server commands.",
}

var serverRunCmd = &cobra.Command{
	Use:   "run",
	Short: "Starts the API server.",
	RunE:  runServer,
}
type FortuneJSON struct {
	fortune string `json:"fortune"`
}

func runServer(cmd *cobra.Command, args []string) error {

	// Init healthchecks
	health := healthcheck.NewHandler()
	health.AddLivenessCheck("goroutine-threshold", healthcheck.GoroutineCountCheck(0))

	adminMux := http.NewServeMux()

	go http.ListenAndServe("0.0.0.0:9402", adminMux)

	adminMux.HandleFunc("/live", health.LiveEndpoint)
	adminMux.HandleFunc("/ready", health.ReadyEndpoint)

	// Init application
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewFortuneAPI(swaggerSpec)
	server := restapi.NewServer(api)

	defer func() {
		if err := server.Shutdown(); err != nil {
			log.Fatalln(err)
		}
	}()

	server.Port = 8080

	api.GetFortuneHandler = operations.GetFortuneHandlerFunc(Fortune)

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

	return err

}

//Fortune route returns fortune
func Fortune(params operations.GetFortuneParams) middleware.Responder {

	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	dbsize, err := redis.Int(conn.Do("DBSIZE"))
	log.Printf("DBSIZE: %d", dbsize)

	rand.Seed(time.Now().UnixNano())
	key := fmt.Sprintf("fortune:%d", rand.Intn(dbsize))

	fortune, err := redis.String(conn.Do("HGET", key, "fortune"))
	if err != nil {
		log.Fatal(err)
	}

	f := operations.GetFortuneOKBody{fortune}
	message, err := json.Marshal(f)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(message))

	return operations.NewGetFortuneOK().WithPayload(&f)
}