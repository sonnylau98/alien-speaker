package main

import (
	"fmt"
	"log"
	"net"

	"github.com/sonnylau98/alien-speaker/proxy/cmd"
	"github.com/sonnylau98/alien-speaker/proxy/local"
)

const (
	DefaultListenAddr = ":7448"
)

var version = "master"

func main() {
	log.SetFlags(log.Lshortfile)

	config := &cmd.Config{
		ListenAddr: DefaultListenAddr,
	}
	config.ReadConfig()
	config.SaveConfig()
	
	lsLocal, err := local.NewLsLocal(config.ListenAddr, config.RemoteAddr)
	if err != nil {
		log.Fatalln(err)
	}
	log.Fatalln(lsLocal.Listen(func(listenAddr *net.TCPAddr) {
		log.Println(fmt.Sprintf(
			`alien-speaker:%s successfully started, configs as below:
                        local listen address: %s
                        remote server address: %s
                        `,version, listenAddr, config.RemoteAddr))
	}))
}
