package crypto

import (
	"bytes"
	"crypto"
	"crypto/ecdsa"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/gob"
)

type Public struct {
	rsa   *rsa.PublicKey
	ecdsa *ecdsa.PublicKey
}

func (p *Public) VerifySign(data interface{}, signature string) bool {

	if p.rsa != nil {
		var buf bytes.Buffer
		enc := gob.NewEncoder(&buf)
		if err := enc.Encode(data); err != nil {
			return false
		}

		signatureByte, err := base64.StdEncoding.DecodeString(signature)
		if err != nil {
			return false
		}

		hash := sha256.Sum256(buf.Bytes())
		err = rsa.VerifyPKCS1v15(p.rsa, crypto.SHA256, hash[:], signatureByte)
		if err != nil {
			return false
		}

		return true
	}
	return false
}
