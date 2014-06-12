package gopin

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"io/ioutil"
	"net/http"
)

type PinnedHttpTransport struct {
	Transport *http.Transport
	State     bool
}

type PinnedConnectionState struct {
	PublicKeyInfo   []byte
	ConnectionState tls.ConnectionState
}

func New(address string, pinnedPublicKey []byte) (*PinnedHttpTransport, error) {
	pinnedTransport := &PinnedHttpTransport{
		State:     false,
		Transport: &http.Transport{},
	}

	if len(pinnedPublicKey) == 0 {
		return pinnedTransport, errors.New("Empty public key")
	}

	remoteTlsState, err := FetchRemotePublicKey(address)

	if err != nil {
		return pinnedTransport, err
	}

	if bytes.Equal(remoteTlsState.PublicKeyInfo, pinnedPublicKey) {
		// The keys matched so set the state as true
		pinnedTransport.State = true

		tlsConfig := &tls.Config{
			Certificates: make([]tls.Certificate, 1),
			RootCAs:      x509.NewCertPool(),
		}

		rawCertificates := make([][]byte, len(remoteTlsState.ConnectionState.PeerCertificates))

		for i, v := range remoteTlsState.ConnectionState.PeerCertificates {
			rawCertificates[i] = v.Raw

			if i > 0 {
				tlsConfig.RootCAs.AddCert(v)
			}
		}

		tlsConfig.Certificates[0] = tls.Certificate{
			Certificate: rawCertificates,
		}

		tlsConfig.ServerName = remoteTlsState.ConnectionState.ServerName
		tlsConfig.CipherSuites = []uint16{remoteTlsState.ConnectionState.CipherSuite}

		pinnedTransport.Transport.TLSClientConfig = tlsConfig
	}

	return pinnedTransport, nil
}

// expects input to be IP||HOSTNAME:PORT
func FetchRemotePublicKey(host string) (PinnedConnectionState, error) {
	pinnedConnectionState := PinnedConnectionState{
		PublicKeyInfo: make([]byte, 0),
	}

	// using a config that does no verification to avoid
	// failing on self-signed certs
	unsafeTlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}

	conn, err := tls.Dial("tcp", host, unsafeTlsConfig)

	if err != nil {
		return pinnedConnectionState, err
	}

	// since we aren't writing or reading directly from this session
	// we need to call Handshake() manually
	err = conn.Handshake()

	if err != nil {
		return pinnedConnectionState, err
	}

	connState := conn.ConnectionState()

	// The certificate chain will have the server certificate as the first
	// object and any signing CA's as the additional nodes
	for i, v := range connState.PeerCertificates {
		if i == 0 {
			pinnedConnectionState.PublicKeyInfo = v.RawSubjectPublicKeyInfo
		}
	}

	if err = conn.Close(); err != nil {
		return pinnedConnectionState, err
	}

	pinnedConnectionState.ConnectionState = connState

	return pinnedConnectionState, nil
}

func ReadInTrustedPublicKey(file string) (PinnedConnectionState, error) {
	fileBytes, err := ioutil.ReadFile(file)

	if err != nil {
		return PinnedConnectionState{}, err
	}

	state := PinnedConnectionState{
		PublicKeyInfo: fileBytes,
	}

	return state, nil
}
