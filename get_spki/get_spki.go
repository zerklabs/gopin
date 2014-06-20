package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	var (
		certFile = flag.String("file", "", "PEM encoded certificate to parse")
		rawCert  []byte
		err      error
	)

	flag.Parse()

	if len(*certFile) == 0 {
		rawCert, err = ioutil.ReadAll(os.Stdin)

		if err != nil {
			log.Fatal(err)
		}
	}

	if len(rawCert) == 0 {

		if len(*certFile) == 0 {
			println("Either pipe in the certificate file or use --file to specify one")
			return
		}

		rawCert, err = ioutil.ReadFile(*certFile)

		if err != nil {
			log.Fatal(err)
		}
	}

	blk, _ := pem.Decode(rawCert)
	cert, err := x509.ParseCertificate(blk.Bytes)

	sigAlg := signatureAlgorithm(cert.SignatureAlgorithm)
	sigAlgHash := signatureAlgorithmHashOnly(cert.SignatureAlgorithm)
	pkAlg := publicKeyAlgorithm(cert.PublicKeyAlgorithm)

	fmt.Printf("Signature Algorithm: %s\n", sigAlg)
	fmt.Printf("Public Key Algorithm: %s\n", pkAlg)

	if err != nil {
		log.Fatal(err)
	}

	hashBytes := make([]byte, 0)

	switch sigAlgHash {
	case "md5":
		sum := md5.Sum(cert.RawSubjectPublicKeyInfo)
		for _, v := range sum {
			hashBytes = append(hashBytes, v)
		}
	case "sha1":
		sum := sha1.Sum(cert.RawSubjectPublicKeyInfo)
		for _, v := range sum {
			hashBytes = append(hashBytes, v)
		}
	case "sha256":
		sum := sha256.Sum256(cert.RawSubjectPublicKeyInfo)
		for _, v := range sum {
			hashBytes = append(hashBytes, v)
		}
	case "sha384":
		sum := sha512.Sum384(cert.RawSubjectPublicKeyInfo)
		for _, v := range sum {
			hashBytes = append(hashBytes, v)
		}
	case "sha512":
		sum := sha512.Sum512(cert.RawSubjectPublicKeyInfo)
		for _, v := range sum {
			hashBytes = append(hashBytes, v)
		}
	default:
		fmt.Println("Unsupported signature algorithm")
		return
	}

	encodedLen := base64.StdEncoding.EncodedLen(len(hashBytes))
	encodedSPKI := make([]byte, encodedLen)
	base64.StdEncoding.Encode(encodedSPKI, hashBytes)

	fmt.Printf("HSTS: %s/%s\n", sigAlgHash, string(encodedSPKI))
}

func publicKeyAlgorithm(algo x509.PublicKeyAlgorithm) string {
	switch algo {
	case x509.DSA:
		return "dsa"
	case x509.RSA:
		return "rsa"
	case x509.ECDSA:
		return "ecdsa"
	}

	return ""
}

func signatureAlgorithm(algo x509.SignatureAlgorithm) string {
	switch algo {
	case x509.MD5WithRSA:
		return "md5RSA"
	case x509.DSAWithSHA1:
		return "sha1DSA"
	case x509.DSAWithSHA256:
		return "sha256DSA"
	case x509.SHA1WithRSA:
		return "sha1RSA"
	case x509.SHA256WithRSA:
		return "sha256RSA"
	case x509.SHA384WithRSA:
		return "sha384RSA"
	case x509.SHA512WithRSA:
		return "sha512RSA"
	}

	return ""
}

func signatureAlgorithmHashOnly(algo x509.SignatureAlgorithm) string {
	switch algo {
	case x509.DSAWithSHA1:
		return "sha1"
	case x509.DSAWithSHA256:
		return "sha256"
	case x509.MD5WithRSA:
		return "md5"
	case x509.SHA1WithRSA:
		return "sha1"
	case x509.SHA256WithRSA:
		return "sha256"
	case x509.SHA384WithRSA:
		return "sha384"
	case x509.SHA512WithRSA:
		return "sha512"
	}

	return ""
}
