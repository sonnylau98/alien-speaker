package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	"github.com/sonnylau98/alien-speaker/proxy/cmd"
	"github.com/sonnylau98/alien-speaker/proxy/server"
)
   
var version = "master"

func main() {
	log.SetFlags(log.Lshortfile)

	port, err := strconv.Atoi(os.Getenv("LIGHTSOCKS_SERVER_PORT"))
	if err != nil {
		port = 7448
	}
		
	config := &cmd.Config{
		ListenAddr: fmt.Sprintf(":%d", port),
	}
	config.ReadConfig()
	config.SaveConfig()
	
	lsServer, err := server.NewLsServer(config.ListenAddr)
	if err != nil {
		log.Fatalln(err)
	}
	log.Fatalln(lsServer.Listen(func(listenAddr *net.TCPAddr) {
		log.Println(fmt.Sprintf(
			`alien-speaker:%s successfully started, configs as below:
                        server listen address: %s
                        `,version, listenAddr))
	}))
}
