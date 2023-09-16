package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"os"
)

func main() {
	pemPath := "D:\\workspace\\go_project\\study\\daydayup\\base\\cryptox\\certfile\\privkey.pem"
	rootPEM, err := os.ReadFile(pemPath)
	if err != nil {
		return
	}
	block, rest := pem.Decode(rootPEM)
	if block == nil || block.Type != "PUBLIC KEY" {
		fmt.Println(block.Type) // RSA PRIVATE KEY
		log.Fatal("failed to decode PEM block containing public key")
	}
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Got a %T, with remaining data: %q", pub, rest)
}
