package local

import (
	"../core"
	"log"
	"net"
)

type LsLocal struct {
	*core.Socket
}

func New(listenAddr *net.TCPAddr, remoteAddr *net.TCPAddr) *LsLocal {
	return &LsLocal{
		Socket: &core.Socket{
			ListenAddr:listenAddr,
			RemoteAddr: remoteAddr,
		},
	}
}

//obtain a listener from address, use the listener to accept TCP; So get user connection, then handle it.
func (local *LsLocal) Listen(didListen func(listenAddr net.Addr)) error {
	listener, err := net.ListenTCP("tcp", local.ListenAddr)
	if err != nil {
		return err
	} 

	deter listener.Close()

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

//dial the proxy server, read proxy's data to local data;
func (local *LsLocal) handleConn(userConn *net.TCPConn) {
	defer userConn.Close()

	proxyServer, err := local.DialRemote()
	if err:= nil {
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
	
	//send local data to proxy server
	local.Copy(proxyServer, userConn)
}
