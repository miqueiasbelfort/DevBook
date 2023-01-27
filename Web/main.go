package main

import (
	"fmt"
	"log"
	"net/http"
	"web/src/router"
)

func main() {
	fmt.Println("Rodando WebApp")

	rw := router.GerarRouters()
	log.Fatal(http.ListenAndServe(":3000", rw))
}
