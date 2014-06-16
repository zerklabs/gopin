package gopin

import (
	"bytes"
	"crypto/tls"
	"errors"
	"net"
)

type PinnedDialer struct {
	TLSConfig *tls.Config
	PinnedKey []byte
}

func New(publicKeyInfo []byte, config *tls.Config) *PinnedDialer {
	return &PinnedDialer{PinnedKey: publicKeyInfo, TLSConfig: config}
}

func (t *PinnedDialer) Dial(network, address string) (net.Conn, error) {
	c, err := tls.Dial(network, address, t.TLSConfig)

	if err != nil {
		println(err)
		return nil, err
	}

	// The certificate chain will have the server certificate as the first
	// object and any signing CA's as the additional nodes
	for i, v := range c.ConnectionState().PeerCertificates {
		if i == 0 {
			if bytes.Equal(t.PinnedKey, v.RawSubjectPublicKeyInfo) {
				return tls.Dial(network, address, t.TLSConfig)
			}
		}
	}

	return nil, errors.New("Pinned key did not match the destination")
}
