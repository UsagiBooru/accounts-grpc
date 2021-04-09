package main

import (
	"log"
	"net"

	"github.com/UsagiBooru/accounts-grpc/gen"
	"github.com/UsagiBooru/accounts-grpc/impl"
	"google.golang.org/grpc"
)

func main() {
	listenPort, err := net.Listen("tcp", ":19003")
	if err != nil {
		log.Fatalln(err)
	}
	server := grpc.NewServer()
	accountsService := &impl.AccountsService{}
	gen.RegisterAccountsServer(server, accountsService)
	server.Serve(listenPort)
}
