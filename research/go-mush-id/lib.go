package main

import (
	"errors"
	"fmt"
	"net"
	"os"
	"time"
)

type MushID struct {
	macAddr string
	pid    int
	seedFn func() string
	hashFn func(string) string
}

func NewMushID(sf func() string, hf func(string) string) MushID {
	mac, err := getMacAddr()
	if err != nil {
		panic(err) // don't do it at home.
	}

	pid := os.Getegid()

	return MushID{
		macAddr: mac,
		pid: pid,
		seedFn: sf,
		hashFn: hf,
	}
}

func (mi *MushID) ID() string {
	rawID := fmt.Sprintf("%s-%d-%v-%s", mi.macAddr, mi.pid, time.Now().Nanosecond(), mi.seedFn())
	return mi.hashFn(rawID)
}

func getMacAddr() (string, error) {
    ifas, err := net.Interfaces()
    if err != nil {
        return "", err
    }

    for _, ifa := range ifas {
        a := ifa.HardwareAddr.String()
        if a != "" {
            return a, nil
        }
    }

	return "", errors.New("can't get mac")
}

func DefaultHashFn(s string) string {
	return s
}

func DefaultSeedFn() string {
	return ""
}
