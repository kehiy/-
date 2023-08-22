package core

import (
	"log"
	"net"
)

type ServerOpts struct {
	ListenAddr string
}

type Server struct {
	ServerOpts

	cache ICache
}

func NewServer(opts ServerOpts, c ICache) *Server {
	return &Server{
		ServerOpts: opts,
		cache:      c,
	}
}

func (s *Server) Start() {
	ln, err := net.Listen("tcp", s.ListenAddr)
	if err != nil {
		log.Panicf("error starting the server: %s\n", err)
	}

	log.Printf("server starting on port: %s\n", s.ListenAddr)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("can not accept connection from: %s\n error: %s\n",
				conn.RemoteAddr(), err)
			continue
		}
		go s.handleConn(conn)
	}
}

func (s *Server) handleConn(conn net.Conn) {
	defer func() {
		conn.Close()
	}()

	buf := make([]byte, 2048)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Printf("can not read from: %s\n error: %s\n",
				conn.RemoteAddr(), err)
			break
		}

		response := Execute(buf[:n], s)
		_, err = conn.Write(response)
		if err != nil {
			log.Printf("can not write to: %s\n error: %s\n",
				conn.RemoteAddr(), err)
			break
		}
	}
}
