package main

import (
	"github.com/prayogatriady/sawer-grpc/config"
	"github.com/prayogatriady/sawer-grpc/core"
)

func main() {
	db, _ := config.InitMySQL()
	core.InitializeGRPCServer(db)
}
