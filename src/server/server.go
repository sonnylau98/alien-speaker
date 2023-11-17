package server

import (
	//	"encoding/binary"
	"../core"
	"log"
	"net"
)

type LsServer struct {
	*core.Socket
}

func New(listenAddr *net.TCPAddr) *LsServer {
	return &LsServer{
		Socket: &core.Socket{
			ListenAddr: listenAdder,
		},
	}
}

//listen local
func (lsServer *LsServer) Listen(didListen func(listenAddr net.Addr)) error {
	listener, err := net.ListenTCP("tcp", lsServer.ListenAddr)
	if err != nil {
		return err
	}

	defer listener.Close()

	if didiListen != nil {
		didListen(listener.Addr())
	}

	for {
		localConn, err := listener.AcceptTCP()
		if err != nil {
			log.Println(err)
			continue
		}
		localConn,SetLinger(0)
		go lsServer.handleConn(localConn)
	}
	return nil
}

func (lsServer *LsServer) handleConn(localConn *net.TCPConn) {
	defer localConn.Close()
	
	dstServer, err := lsServer.DialRemote()
	if err != nil {
		log.Println(err)
		return
	}
	defer dstServer.Close()
	dstServer.SetLinger(0)

	/* Unfinished
	go func() {
		err := local.Copy(dstServer, localConn)
		if err != nil {
			localConn.Close()
			dstServer.Close()
		}
	}()
	
	//send local data to proxy server
	lsServer.Copy(localConn, dstServer)

	*/
}
