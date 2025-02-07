package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/bits-and-blooms/bloom/v3"
	"github.com/nbd-wtf/go-nostr"
)

var (
	blm1 = bloom.NewWithEstimates(10000, 0.9)
	sec1 = "0000000000000000000000000000000000000000000000000000000000000001"
	pub1 = "79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798"
	
	r, _ = nostr.RelayConnect(context.Background(), "wss://jellyfish.land")

	blm2 = bloom.NewWithEstimates(10000, 0.9)
	pub2 = "bd4ae3e67e29964d494172261dc45395c89f6bd2e774642e366127171dfb81f5"
	messages = []string{
		"1dac46d2b92eff5199e97d39fcd8a7541ab04ab1e9c372f2a521156c40cac9ba",
		"264f6959b9030bd0a0532c1d16997dff7d7af74328550af55fd49a11eeaa072b",
		"92a96d54062147c56803b843ebf491ed28b826fd0b512432c7aa2c99f40c9300",
		"c5280452bbc29913bb26d62a5cb9f3e33cf4028d98f558ed9a98835cc9b17f50",
	}
)

func publishSeenList() {
		// adding seen events.
		blm1.Add([]byte(messages[0]))
		blm1.Add([]byte(messages[1]))
		blm1.Add([]byte(messages[2]))

		var buffer bytes.Buffer
		blm1.WriteTo(&buffer)

		evt := nostr.Event{
			CreatedAt: nostr.Timestamp(time.Now().Unix()),
			Kind:      30010,
			Content:   base64.RawStdEncoding.EncodeToString(buffer.Bytes()),
			Tags: nostr.Tags{
				nostr.Tag{"d", pub2},
			},
		}

		evt.Sign(sec1)
		r.Publish(context.Background(), evt)
}

func checkSeenList() {
	// check seen events.
	sub, _ := r.Subscribe(context.Background(), nostr.Filters{
		nostr.Filter{
			Authors: []string{pub1},
			Kinds: []int{30010},
			Tags: nostr.TagMap{
				"d": []string{""},
			},
		},
	})

	var filt string
	for evt := range sub.Events {
		filt = evt.Content
		return
	}

	fmt.Println(filt + "sss")

	b, _ := base64.RawStdEncoding.DecodeString(filt)
	var buffer bytes.Buffer
	buffer.Read(b)

	blm2.ReadFrom(&buffer)

	fmt.Println(blm2.Test([]byte(messages[0]))) // true
	fmt.Println(blm2.Test([]byte(messages[1]))) // true
	fmt.Println(blm2.Test([]byte(messages[2]))) // true
	fmt.Println(blm2.Test([]byte(messages[3]))) // false
}

func main() {
	fmt.Println("start")
	go publishSeenList()
	checkSeenList()
}
