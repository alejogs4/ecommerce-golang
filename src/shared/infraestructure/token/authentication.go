package token

import (
	"crypto/rsa"
	"io/ioutil"
	"sync"

	"github.com/dgrijalva/jwt-go"
)

var (
	signingKey        *rsa.PrivateKey
	verifyKey         *rsa.PublicKey
	certificateLoader sync.Once
)

// LoadCertificates both public and private certicates to works as helpers for jwt generation
func LoadCertificates(publicCertificate, privateCertificate string) error {
	var err error
	certificateLoader.Do(func() {
		err = loadCertificates(publicCertificate, privateCertificate)
	})

	return err
}

func loadCertificates(publicCertificate, privateCertificate string) error {
	signingKeyFile, err := ioutil.ReadFile(privateCertificate)
	if err != nil {
		return err
	}

	verifyKey, err := ioutil.ReadFile(publicCertificate)
	if err != nil {
		return err
	}

	return parseRSAKeys(verifyKey, signingKeyFile)
}

func parseRSAKeys(verifyKeyBytes, signingKeyBytes []byte) error {
	var err error

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyKeyBytes)
	if err != nil {
		return err
	}

	signingKey, err = jwt.ParseRSAPrivateKeyFromPEM(signingKeyBytes)
	return err
}
