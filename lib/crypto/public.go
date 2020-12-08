package crypto

import (
	"crypto/ecdsa"
	"crypto/rsa"
)

type Public struct {
	rsa   *rsa.PublicKey
	ecdsa *ecdsa.PublicKey
}

func (p *Public) VerifySign(data interface{}, signature string) bool {
	return false
}
