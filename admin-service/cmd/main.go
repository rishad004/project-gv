package main

import (
	"log"

	"github.com/rishad004/project-gv/admin-service/pkg/config"
	"github.com/rishad004/project-gv/admin-service/pkg/di"
)

func main() {

	if err := config.Config(); err != nil {
		log.Fatal(err)
	}

	di.InitGRPC()
}
