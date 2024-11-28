package main

import (
	"log"

	"github.com/rishad004/project-gv/apiGateway/pkg/config"
	"github.com/rishad004/project-gv/apiGateway/pkg/di"
)

func main() {

	if err := config.Config(); err != nil {
		log.Fatal(err)
	}

	di.ClientConnect()
}
