package server

import (
	"github.com/sonnylau98/alien-speaker/proxy/core"
	"log"
	"net"
)

type LsServer struct {
	*core.Socket
}

func New(listenAddr *net.TCPAddr) *LsServer {
	return &LsServer{
		Socket: &core.Socket{
			ListenAddr: listenAddr,
		},
	}
}

func NewLsServer(listenAddr string) (*LsServer, error) {
	
	structListenAddr, err := net.ResolveTCPAddr("tcp", listenAddr)
	if err != nil {
		return nil, err
	}
	
	return &LsServer{
		Socket: &core.Socket{
			ListenAddr: structListenAddr,
		},
	}, nil	
}

//listen local
func (lsServer *LsServer) Listen(didListen func(listenAddr *net.TCPAddr)) error {
	listener, err := net.ListenTCP("tcp", lsServer.ListenAddr)
	if err != nil {
		return err
	}

	defer listener.Close()

	if didListen != nil {
		go didListen(listener.Addr().(*net.TCPAddr))
	}

	for {
		localConn, err := listener.AcceptTCP()
		if err != nil {
			log.Println(err)
			continue
		}
		localConn.SetLinger(0)
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

	go func() {
		err := lsServer.Copy(localConn, dstServer)
		if err != nil {
			localConn.Close()
			dstServer.Close()
		}
	}()
	
	lsServer.Copy(dstServer, localConn)
}
