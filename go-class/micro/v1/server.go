package v1

import (
	"fmt"
	"net"
	"reflect"
)

type Server struct {
	services map[string]*reflectionStub
}

func (s *Server) Start(address string) error {
	ln, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Printf("accept connetction got error: %v", err)
		}
		go s.handleConnection(conn)
	}
}

func (s *Server) handleConnection(conn net.Conn) {
	for {
		ReadMsg(conn)
	}
}

type reflectionStub struct {
	s     Service
	value reflect.Value
}
