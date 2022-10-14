package main

import (
	"RPCServer/procedures"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

func RPCStart() {
	job := new(procedures.Jobs)
	err := rpc.Register(job)
	if err != nil {
		log.Fatal(err)
	}
	rpc.HandleHTTP()

	l, e := net.Listen("tcp", ":4466")
	if e != nil {
		log.Fatalf("Couldn't start listening on port 4466. Error %s", e)
	}
	log.Println("Serving RPC handler")
	err = http.Serve(l, nil)
	if err != nil {
		log.Fatalf("Error serving: %s", err)
	}
}
