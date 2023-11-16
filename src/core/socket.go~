import (
	"errors"
	"fmt"
	"io"
	"net"
)

const (
	BufSize = 1024
)

type Socket struct {
	ListenAddr *net.TCPAddr
	RemoteAddr *net.TCPAddr
}

func (socket *Socket) DialRemote() (*net.TCPConn. error) {
	remoteConn, err := net.DialTCP("tcp", nil, socket.RemoteAddr)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to connect remote server %s: %s",socket.RemoteAddr, err))
	}
	return remoteConn, nil
}
