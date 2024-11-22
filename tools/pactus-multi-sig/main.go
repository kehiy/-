package main

import (
	"fmt"

	"github.com/pactus-project/pactus/crypto/bls"
)

func main() {	
	pk1, err := bls.PublicKeyFromString("public1pjcsydma9gj299mvkp27an6lman2p8pgkjcnz3t8xh05ndj7zyp45u7z09mv0s2d3p63uk9jzkqrr2rs38ahfs9vh4fmuxzq3k28x38tledxq7xa664j6av6wlg8nz2tgc0vlne5k5mkah7mj6ptv54wvaskw3gjs")
	if err != nil {
		panic(err)
	}

	pk2, err := bls.PublicKeyFromString("public1p4vwfujre7nncezc96j96ru7s3p6qxljaqps093zj2vd4j7az83a730km2rdr3cyqxjtu0q774jjg69enpxml36aaetkmp9zvfcflvuga6uxqv9z0a6qc2z0r2e3n49sm8qtq2l8qcd2pw4upl0306fl64vknsqnm")
	if err != nil {
		panic(err)
	}

	pk3 := bls.PublicKeyAggregate(pk1, pk2)

	fmt.Println(pk3.String())
	fmt.Println(pk3.AccountAddress().String())
}
