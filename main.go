package main

import (
	"fmt"
	"log"
	"net"

	protos "github.com/shivam/translate2/protos/translation2"
	"github.com/shivam/translate2/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// create new gRPC server
	s := grpc.NewServer()

	// create new instance of Translation server
	trans := server.NewTranslation()

	// register reflection API https://github.com/grpc/grpc/blob/master/doc/server-reflection.md
	reflection.Register(s)

	// register it to the grpc server
	protos.RegisterTranslation2Server(s, trans)

	// create socket to listen to requests
	port := "3001"
	tl, err := net.Listen("tcp", "localhost:" + port)

	/*
		db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/testdb")
		if err != nil {
			panic(err.Error())
		}else{
			defer db.Close()
			fmt.Println("DB Connected Success!")
		}
	*/
	
	
	if err != nil {
		log.Fatal(fmt.Println("Error starting tcp listener on port --" + port, err))
	} else {
		fmt.Println("Server running on port --" + port)
	}

	// start listening
	s.Serve(tl)
}
