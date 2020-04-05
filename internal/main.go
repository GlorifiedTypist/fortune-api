package main

import (
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gomodule/redigo/redis"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/go-openapi/loads"

	"pkg/swagger/pkg/swagger/server/restapi"
	"pkg/swagger/pkg/swagger/server/restapi/operations"
)

func main() {

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

	log.Println("Warming cache")
	warmCache()
	log.Println("Warming cache, done")

	server.Port = 8080

	api.GetHealthzHandler = operations.GetHealthzHandlerFunc(Health)
	api.GetFortuneHandler = operations.GetFortuneHandlerFunc(Fortune)

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

}

//Health route returns OK
func Health(params operations.GetHealthzParams) middleware.Responder {
	return operations.NewGetHealthzOK().WithPayload("OK")
}

//Fortune route returns fortune
func Fortune(params operations.GetFortuneParams) middleware.Responder {

	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	dbsize, err := redis.String(conn.Do("DBSIZE"))
	log.Println("DBSIZE: %d", dbsize)

	rand.Seed(time.Now().UnixNano())

	i, err := strconv.Atoi(dbsize)
	if err != nil {
		log.Fatalln(err)
	}

	key := fmt.Sprintf("fortune:%d", rand.Intn(i))
	fortune, err := redis.String(conn.Do("HGET", key, "fortune"))
	if err != nil {
		log.Fatal(err)
	}

	return operations.NewGetFortuneOK().WithPayload(fortune)
}