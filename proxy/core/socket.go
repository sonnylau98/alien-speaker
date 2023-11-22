package core

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

//read source to buffer, write buffer to destination
func (socket *Socket) Copy(dst *net.TCPConn, src *net.TCPConn) error {
	buf := make([]byte, BufSize)
	for {
		readCount, errRead := src.Read(buf)
		if errRead != nil {
			if errRead != io.EOF {
				return errRead
			}
				return nil
		}
		if readCount > 0 {
			writeCount, errWrite := dst.Write(buf[0:readCount])
			if errWrite != nil {
				return errWrite
			}
			if readCount != writeCount {
				return io.ErrShortWrite
			}
		}
	}
}

func (socket *Socket) DialRemote() (*net.TCPConn, error) {
	remoteConn, err := net.DialTCP("tcp", nil, socket.RemoteAddr)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to connect remote server %s: %s",socket.RemoteAddr, err))
	}
	return remoteConn, nil
}
