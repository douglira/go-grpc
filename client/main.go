package main

import (
	"github.com/douglira/go-grpc/client/routes"
)

func main() {
	r := routes.GetRouter()
	r.Run(":5000")
}
