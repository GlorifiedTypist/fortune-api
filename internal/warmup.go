package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

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

func warmCache() error {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	fortunes := getFortunesFromURL()

	for i := 0;  i<=len(fortunes) - 1; i++ {

		key  := fmt.Sprintf("fortune:%d", i)
		_, err = conn.Do("HMSET", key, "fortune", strings.ReplaceAll(fortunes[i], "\n", ""))
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Fortune [%d]: %s", i, strings.ReplaceAll(fortunes[i], "\n", ""))
	}

	return err
}