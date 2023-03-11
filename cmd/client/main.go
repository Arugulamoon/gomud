package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"

	"github.com/reiver/go-telnet"
)

func main() {
	caCert, err := ioutil.ReadFile("server.crt")
	if err != nil {
		log.Fatal(err)
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	tlsConfig := &tls.Config{
		RootCAs: caCertPool,
	}

	var caller telnet.Caller = telnet.StandardCaller
	telnet.DialToAndCallTLS("localhost:7324", caller, tlsConfig)
}
