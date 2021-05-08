package sandsdk

import (
	"crypto"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
)

func (s *Service) PayNotifyVerifySign(notify *PayNotify) error {
	if notify == nil {
		return nil
	}

	sign, err := base64.StdEncoding.DecodeString(notify.Sign)
	if err != nil {
		return err
	}

	publicKeyData, err := os.ReadFile(s.SandPublicKeyPath)
	if err != nil {
		return err
	}

	block, _ := pem.Decode([]byte(GeneratePem(publicKeyData)))
	if block == nil {
		return errors.New("bad key data: pem decode")
	}

	if got, want := block.Type, "CERTIFICATE"; got != want {
		return fmt.Errorf("unknown key type %q, want %q", got, want)
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return err
	}

	hash := crypto.Hash.New(crypto.SHA1)
	hash.Write([]byte(notify.dataStr))
	hashed := hash.Sum(nil)

	return rsa.VerifyPKCS1v15(cert.PublicKey.(*rsa.PublicKey), crypto.SHA1, hashed, sign)
}
