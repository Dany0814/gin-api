package main

import (
	"log"

	"github.com/Dany0814/gin-api/cmd/api/domain"
)

func main() {
	if err := domain.Run(); err != nil {
		log.Fatal(err)
	}
}
