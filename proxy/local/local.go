package local

import (
	"github.com/sonnylau98/alien-speaker/proxy/core"
	"log"
	"net"
)

type LsLocal struct {
	*core.Socket
}

/*
func New(listenAddr *net.TCPAddr, remoteAddr *net.TCPAddr) *LsLocal {
	return &LsLocal{
		Socket: &core.Socket{
			ListenAddr:listenAddr,
			RemoteAddr: remoteAddr,
		},
	}
}
*/

func NewLsLocal(listenAddr, remoteAddr string) (*LsLocal, error) {
	
	structListenAddr, err := net.ResolveTCPAddr("tcp", listenAddr)
	if err != nil {
		return nil, err
	}
	
	structRemoteAddr, err := net.ResolveTCPAddr("tcp", remoteAddr)
	if err != nil {
		return nil, err
	}

	return &LsLocal{
		Socket: &core.Socket{ // Set the embedded Socket directly
			ListenAddr: structListenAddr,
			RemoteAddr: structRemoteAddr,
		},
	}, nil	
}

//obtain a listener from address, use the listener to accept TCP; So get user connection, then handle it.
func (local *LsLocal) Listen(didListen func(listenAddr *net.TCPAddr)) error {
	listener, err := net.ListenTCP("tcp", local.ListenAddr)
	if err != nil {
		return err
	} 

	defer listener.Close()

	if didListen != nil {
		didListen(listener.Addr())
	}

	for {
		userConn, err := listener.AcceptTCP()
		if err != nil {
			log.Println(err)
			continue
		}
		userConn.SetLinger(0)
		go local.handleConn(userConn)
	}
	return nil
}

func (local *LsLocal) handleConn(userConn *net.TCPConn) {
	defer userConn.Close()

	proxyServer, err := local.DialRemote()
	if err != nil {
		log.Println(err)
		return
	}
	defer proxyServer.Close()
	proxyServer.SetLinger(0)

	go func() {
		err := local.Copy(userConn, proxyServer)
		if err != nil {
			userConn.Close()
			proxyServer.Close()
		}
	}()
	
	local.Copy(proxyServer, userConn)
}
