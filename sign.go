package sandsdk

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
)

func (s *Service) Sign(data []byte) (string, error) {
	bytes, err := os.ReadFile(s.AppPrivateKeyPath)
	if err != nil {
		return "", err
	}

	block, _ := pem.Decode(bytes)
	if block == nil {
		return "", errors.New("bad key data: pem decode")
	}

	if got, want := block.Type, "RSA PRIVATE KEY"; got != want {
		return "", fmt.Errorf("unknown key type %q, want %q", got, want)
	}

	private, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}

	h := crypto.Hash.New(crypto.SHA1)
	h.Write(data)
	hashed := h.Sum(nil)

	sign, err := rsa.SignPKCS1v15(rand.Reader, private, crypto.SHA1, hashed)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(sign), nil
}
