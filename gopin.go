package gopin

import (
	"bytes"
	"crypto/tls"
	"errors"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

type PinnedTransport struct {
	http.Transport
	PinnedKey []byte
}

func New(publicKeyInfo []byte, tlsConfig *tls.Config) (*http.Transport, error) {
	dialer := &PinnedTransport{
		PinnedKey: publicKeyInfo,
	}

	if tlsConfig == nil {
		tlsConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
	}

	dialer.Transport = http.Transport{
		Dial:            dialer.Dial,
		TLSClientConfig: tlsConfig,
	}

	return &dialer.Transport, nil
}

func (t *PinnedTransport) Dial(network, address string) (net.Conn, error) {
	conn, err := tls.DialWithDialer(new(net.Dialer), network, address, &tls.Config{InsecureSkipVerify: true})

	if err != nil {
		return nil, err
	}

	defer conn.Close()

	pinnedKeyMatched := false

	// The certificate chain will have the server certificate as the first
	// object and any signing CA's as the additional nodes
	for i, v := range conn.ConnectionState().PeerCertificates {
		if i == 0 {
			if bytes.Equal(t.PinnedKey, v.RawSubjectPublicKeyInfo) {
				pinnedKeyMatched = true
			}
		}
	}

	if pinnedKeyMatched {
		return (&net.Dialer{
			KeepAlive: 30 * time.Second,
			Timeout:   30 * time.Second,
		}).Dial(network, address)
	}

	return nil, errors.New("pin failed")
}

func ReadInTrustedPublicKey(file string) ([]byte, error) {
	fileBytes, err := ioutil.ReadFile(file)

	if err != nil {
		return []byte{}, err
	}

	return fileBytes, nil
}
