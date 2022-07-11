package domain

import (
	"github.com/Dany0814/gin-api/internal/platform/server"
)

const (
	host = "localhost"
	port = 8080
)

func Run() error {
	srv := server.New(host, port)
	return srv.Run()
}
