package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

var (
	DNS403 = []string{"10.202.10.202", "10.202.10.102"}
	SHECAN = []string{"178.22.122.100", "185.51.200.2"}
)

func main() {
	var list []string = []string{}
	var err error

	args := os.Args[1:]

	for _, arg := range args {
		if arg == "1" {
			list = SHECAN
			err = changeDNS(list)
			if err != nil {
				fmt.Printf("error in change dns: %v/n", err)
			} else {
				fmt.Print("dns set successfully:\n", list)
				os.Exit(0)
			}
		} else if arg == "2" {
			list = DNS403
			err = changeDNS(list)
			if err != nil {
				fmt.Printf("error in change dns: %v/n", err)
			} else {
				fmt.Print("dns set successfully:\n", list)
				os.Exit(0)
			}
		} else {
			break
		}
	}

	list = args
	err = changeDNS(list)
	if err != nil {
		fmt.Printf("error in change dns: %v/n", err)
	}
	fmt.Print("dns set successfully:\n", list)
	os.Exit(0)
}

func changeDNS(DNSs []string) error {
	path := "/etc/resolv.conf"

	newContent := "managed by kehiy dns manager CLI\n"
	for _, dns := range DNSs {
		newContent += fmt.Sprintf("nameserver %s\n", dns)
	}

	err := ioutil.WriteFile(path, []byte(newContent), 0644)
	if err != nil {
		return err
	}

	return nil
}
