package main

import (
	"github.com/kehiy/akam/core"
)

func main() {
	opts := core.ServerOpts{
		ListenAddr: ":3000",
	}

	s := core.NewServer(opts, core.New())
	s.Start()
}
