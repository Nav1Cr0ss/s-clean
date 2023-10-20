package encrypt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
)

var pubKey []byte

func SetPublicKey(key []byte) {
	pubKey = key
}

func RsaEncrypt(origData []byte, key []byte) ([]byte, error) {
	block, _ := pem.Decode(key)
	if block == nil {
		return nil, errors.New("public key error")
	}

	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	pub := pubInterface.(*rsa.PublicKey)

	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

func String(input string) (string, error) {
	var output string

	if len(input) == 4 || len(input) == 0 {
		return output, nil
	}

	if len(pubKey) == 0 {
		return output, nil
	}

	encryptedSSN, err := RsaEncrypt([]byte(input), pubKey)
	if err != nil {
		return output, fmt.Errorf("failed to encrypt string: %w", err)
	}

	output = base64.StdEncoding.EncodeToString(encryptedSSN)

	return output, nil
}
