package main

import (
	"crypto/ed25519"
	"encoding/base64"
	"fmt"
	"os"
)

func main() {
	public, private, err := ed25519.GenerateKey(nil)
	if err != nil {
		fmt.Println("error occured", err)
		os.Exit(1)
	}
	private_base64 := base64.StdEncoding.EncodeToString([]byte(private))
	public_base64 := base64.StdEncoding.EncodeToString([]byte(public))
	username := "@" + public_base64 + ".ed25519"
	fmt.Printf("Generated account keypair (base64 encoded):\nPublic key:\t%v\nPrivate key:\t%v\nSSB username:\t%v\n", public_base64, private_base64, username)
}
