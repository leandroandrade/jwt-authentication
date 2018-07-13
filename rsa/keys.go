package rsa

import (
	"crypto/rsa"
	"os"
	"bufio"
	"encoding/pem"
	"crypto/x509"
	"log"
)

var PrivateKey *rsa.PrivateKey
var PublicKey *rsa.PublicKey

func init() {
	log.Println("loading RSA keys...")

	loadPrivateKey()
	loadPublicKey()

	log.Println("keys loaded sucessfull!")
}

func loadPrivateKey() {
	privateKeyFile, err := os.Open(`c:\temp\private_key`)
	if privateKeyFile != nil {
		defer privateKeyFile.Close()
	}

	if err != nil {
		log.Fatal("ERR: cannot load private key file: ", err)
	}

	pemfileinfo, err := privateKeyFile.Stat()
	if err != nil {
		log.Fatal("ERR: cannot load private key fileinfo: ", err)
	}

	pembytes := make([]byte, pemfileinfo.Size())

	buffer := bufio.NewReader(privateKeyFile)
	if _, err = buffer.Read(pembytes); err != nil {
		log.Fatal("ERR: cannot write private key buffer: ", err)
	}

	data, _ := pem.Decode([]byte(pembytes))

	PrivateKey, err = x509.ParsePKCS1PrivateKey(data.Bytes)
	if err != nil {
		log.Fatal("ERR: cannot ParsePKCS1PrivateKey private key bytes: ", err)
	}
}

func loadPublicKey() {
	publicKeyFile, err := os.Open(`c:\temp\public_key.pub`)
	if publicKeyFile != nil {
		defer publicKeyFile.Close()
	}

	if err != nil {
		log.Fatal("ERR: cannot load public key file: ", err)
	}

	pemfileinfo, err := publicKeyFile.Stat()
	if err != nil {
		log.Fatal("ERR: cannot load public key fileinfo: ", err)
	}

	pembytes := make([]byte, pemfileinfo.Size())

	buffer := bufio.NewReader(publicKeyFile)
	if _, err = buffer.Read(pembytes); err != nil {
		log.Fatal("ERR: cannot write public key buffer: ", err)
	}

	data, _ := pem.Decode([]byte(pembytes))

	publicKeyImported, err := x509.ParsePKIXPublicKey(data.Bytes)
	if err != nil {
		log.Fatal("ERR: cannot ParsePKIXPublicKey public key bytes: ", err)
	}

	var ok bool
	PublicKey, ok = publicKeyImported.(*rsa.PublicKey)
	if !ok {
		log.Fatal("ERR: cannot publicKeyImported public key: ", err)
	}
}
