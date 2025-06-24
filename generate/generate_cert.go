package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"log"
	"math/big"
	"net"
	"os"
	"strings"
	"time"
)

func write_to_file(bytes []byte, pemType string, filename string) {
	pemBlock := &pem.Block{
		Type:  pemType, // Or "RSA PUBLIC KEY" for RSA specific
		Bytes: bytes,
	}
	pemData := pem.EncodeToMemory(pemBlock)
	err := os.WriteFile(filename, pemData, 0644)
	if err != nil {
		log.Fatalf("Failed to write public key to file: %v", err)
	}
	fmt.Printf("The %s string is:\n%s\n", pemType, string(pemData))
}

const ServerCommonName = "centos7gpdb7"
const ClientCommonName = "testssl"

var serverHosts = "centos7gpdb7"
var serverIPs = "127.0.0.1"
var clientHosts = "r7centosgpdb56"
var clientIPs = "10.121.171.28"

// ... (functions for generating keys, certificates, and writing to file)

func main() {
	// 1. Generate Root CA
	caPrivKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalf("failed to generate CA private key: %v", err)
	}

	caTemplate := &x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			Organization: []string{"GPDB CA"},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(365 * 24 * time.Hour), // 1 year validity
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageCRLSign,
		BasicConstraintsValid: true,
		IsCA:                  true,
	}

	caBytes, err := x509.CreateCertificate(rand.Reader, caTemplate, caTemplate, &caPrivKey.PublicKey, caPrivKey)
	if err != nil {
		log.Fatalf("failed to create CA certificate: %v", err)
	}
	caprivBytes, err := x509.MarshalPKCS8PrivateKey(caPrivKey)
	if err != nil {
		log.Fatalf("Unable to marshal ca private key: %v", err)
	}
	write_to_file(caBytes, "CERTIFICATE", "root.ca")
	write_to_file(caprivBytes, "PRIVATE KEY", "root.key")
	// Write caBytes and caPrivKey to "root.ca" and "root.key" files

	// 2. Generate Server Certificate
	serverPrivKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalf("failed to generate server private key: %v", err)
	}

	serverTemplate := &x509.Certificate{
		SerialNumber: big.NewInt(2),
		Subject: pkix.Name{
			Organization: []string{"GPDB CA"},
			CommonName:   ServerCommonName,
		},
		NotBefore:   time.Now(),
		NotAfter:    time.Now().Add(365 * 24 * time.Hour),
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		KeyUsage:    x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment,
		DNSNames:    []string{"localhost"}, // Adjust for your server's domain/IP
	}
	if serverHosts != "" {
		hosts := strings.Split(serverHosts, ",")
		for _, h := range hosts {
			serverTemplate.DNSNames = append(serverTemplate.DNSNames, h)
		}
	}
	if serverIPs != "" {
		ips := strings.Split(serverIPs, ",")
		for _, ip := range ips {
			serverTemplate.IPAddresses = append(serverTemplate.IPAddresses, net.IP(ip))
		}
	}

	serverBytes, err := x509.CreateCertificate(rand.Reader, serverTemplate, caTemplate, &serverPrivKey.PublicKey, caPrivKey)
	if err != nil {
		log.Fatalf("failed to create server certificate: %v", err)
	}
	serverprivBytes, err := x509.MarshalPKCS8PrivateKey(serverPrivKey)
	if err != nil {
		log.Fatalf("Unable to marshal ca private key: %v", err)
	}
	write_to_file(serverBytes, "CERTIFICATE", "server.crt")
	write_to_file(serverprivBytes, "PRIVATE KEY", "server.key")
	// Write serverBytes and serverPrivKey to "server.crt" and "server.key" files

	// 3. Generate Client Certificate
	clientPrivKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalf("failed to generate client private key: %v", err)
	}

	clientTemplate := &x509.Certificate{
		SerialNumber: big.NewInt(3),
		Subject: pkix.Name{
			Organization: []string{"GPDB CA"},
			CommonName:   ClientCommonName,
		},
		NotBefore:   time.Now(),
		NotAfter:    time.Now().Add(365 * 24 * time.Hour),
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth},
		KeyUsage:    x509.KeyUsageDigitalSignature,
	}
	if clientHosts != "" {
		hosts := strings.Split(clientHosts, ",")
		for _, h := range hosts {
			clientTemplate.DNSNames = append(clientTemplate.DNSNames, h)
		}
	}
	if clientIPs != "" {
		ips := strings.Split(clientIPs, ",")
		for _, ip := range ips {
			clientTemplate.IPAddresses = append(clientTemplate.IPAddresses, net.IP(ip))
		}
	}

	clientBytes, err := x509.CreateCertificate(rand.Reader, clientTemplate, caTemplate, &clientPrivKey.PublicKey, caPrivKey)
	if err != nil {
		log.Fatalf("failed to create client certificate: %v", err)
	}
	clientprivBytes, err := x509.MarshalPKCS8PrivateKey(clientPrivKey)
	if err != nil {
		log.Fatalf("Unable to marshal ca private key: %v", err)
	}
	write_to_file(clientBytes, "CERTIFICATE", "client.crt")
	write_to_file(clientprivBytes, "PRIVATE KEY", "client.key")
	// Write clientBytes and clientPrivKey to "client.crt" and "client.key" files
}
