package main

import (
	"TCPChat/server"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Println("[USAGE]: ./TCPChat $port")
		return
	}
	if len(os.Args) == 1 {
		server := server.NewServer(":8989")
		fmt.Println("Listening on the port :8989")
		log.Fatal(server.Start())
	} else {
		for _, ch := range os.Args[1] {
			if ch < '0' || ch > '9' {
				fmt.Println("wrong port")
				os.Exit(0)
			}
		}
		port := ":" + os.Args[1]
		server := server.NewServer(port)
		fmt.Println("Listening on the port", port)
		log.Fatal(server.Start())
	}

}
