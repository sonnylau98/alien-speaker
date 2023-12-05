package server

import (
	
	"github.com/sonnylau98/alien-speaker/proxy/core"
	"log"

	"encoding/binary"
	"net"
)

type LsServer struct {
	*core.Socket
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

// listen local
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

	buf := make([]byte, 256)

	_, err := localConn.Read(buf)

	if err != nil || buf[0] != 0x05 {
		return
	}

	localConn.Write([]byte{0x05, 0x00})
	
	n, err := localConn.Read(buf)

	if err != nil || n < 7 {
		return
	}

	// Only support: CONNECT X'01'
	if buf[1] != 0x01 {
		return
	}

	var dIP []byte

	switch buf[3] {
	case 0x01:
		// IP V4 address: X'01'
		dIP = buf[4 : 4+net.IPv4len]
	case 0x03:
		// DOMAINNAME: X'03'
		ipAddr, err := net.ResolveIPAddr("ip", string(buf[5:n-2]))
		if err != nil {
			return
		}
		dIP = ipAddr.IP
	case 0x04:
		// IP V6  address: X'04'
		dIP = buf[4 : 4+net.IPv6len]
	default:
		return
	}

	dPort := buf[n-2:]
	dstAddr := &net.TCPAddr{
		IP: dIP,
		Port: int(binary.BigEndian.Uint16(dPort)),
	}
	
	dstServer, err := net.DialTCP("tcp", nil, dstAddr)
	if err != nil {
		log.Println(err)
		return
	} else {
		defer dstServer.Close()
		dstServer.SetLinger(0)

		localConn.Write([]byte{0x05, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})
	}

	// local to destination
	go func() {
		err := lsServer.Copy(dstServer, localConn)
		if err != nil {
			localConn.Close()
			dstServer.Close()
		}
	}()
	
	lsServer.Copy(localConn, dstServer)
}
