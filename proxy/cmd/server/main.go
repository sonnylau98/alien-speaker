package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	"../.."
	"../cmd"
	"../../server"
)
   
var version = "master"

func main() {
	log.SetFlags(log.Lshortfile)

	port, err := strconv.Atoi(os.Getenv("LIGHTSOCKS_SERVER_PORT"))
	if err != nil {
		port = 7448
	}
		
	config := &cmd.Config{
		ListenAddr: fmt.Sprint(":%d", port),
	}
	config.ReadConfig()
	config.SaveConfig()
	
	lsServer, err := local.NewLsServer(config.ListenAddr)
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
