package core

import (
	"strconv"
	"strings"
	"time"
)

type Command string

const (
	DeleteCMD Command = "DEL"
	SetCMD    Command = "SET"
	GetCMD    Command = "GET"
	HaveCMD   Command = "HAV"
)

type Message struct {
	Cmd   Command
	Key   []byte
	Value []byte
	TTL   time.Duration
}

// TODO: refactor this hell!
func Execute(rawCmd []byte, s *Server) []byte {
	var response []byte
	var err error
	var success bool
	rawStr := string(rawCmd)
	parts := strings.Split(rawStr, " ")
	if len(parts) == 0 {
		return []byte("invalid command")
	}

	cmd := parts[0]
	switch cmd {
	case string(GetCMD):
		response, err = s.cache.Get([]byte(parts[1]))
	case string(SetCMD):
		ttl, errconv := strconv.Atoi(parts[3])
		if errconv != nil {
			break
		}
		err = s.cache.Set([]byte(parts[1]), []byte(parts[2]), time.Duration(ttl))
	case string(DeleteCMD):
		success = s.cache.Delete([]byte(parts[1]))
	case string(HaveCMD):
		success = s.cache.Have([]byte(parts[1]))
	default:
		response, success, err = []byte("invalid command"), false, ErrorExecute
	}
	if err != nil {
		response = []byte("invalid command")
	}
	if success {
		response = []byte("success")
	}

	return response
}
