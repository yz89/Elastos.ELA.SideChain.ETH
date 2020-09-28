package main

import (
	"fmt"
	"github.com/elastos/Elastos.ELA.SideChain.ETH/common"
	"testing"
)

//var sig = secp256k1.sign(msgHash, privateKey)
//var ret = {}
//ret.r = sig.signature.slice(0, 32)
//ret.s = sig.signature.slice(32, 64)
//ret.v = sig.recovery + 27

func Test(t *testing.T) {
	sig := common.Hex2Bytes("adbcf5d5fa424b96634325dc5e5187a42dab6fe7a142cb3f2d050a7ec675903b6dd13b0488bd6446bbefe37f484c5fec5cf348ab3bca076cbb292b8efd0040d22a")

	hash := common.Hex2Bytes("d90e1d8d976aaf11abe35554ebc37f020fe5c2e2124ff3ca451704013725661f")
	signer := ecrecover(hash, sig)

	fmt.Println(signer)
}
