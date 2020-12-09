package crypto

import (
	"bytes"
	"crypto"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/gob"
)

type Private struct {
	rsa   *rsa.PrivateKey
	ecdsa *ecdsa.PrivateKey
}

func (p *Private) Public() *Public {
	pub := Public{}

	if p.rsa != nil {
		pub.rsa = p.rsa.Public().(*rsa.PublicKey)
	}

	return &pub
}

func (p *Private) Sign(data interface{}) string {

	if p.rsa != nil {
		var buf bytes.Buffer
		enc := gob.NewEncoder(&buf)
		if err := enc.Encode(data); err != nil {
			return ""
		}

		hash := sha256.Sum256(buf.Bytes())
		sign, err := rsa.SignPKCS1v15(rand.Reader, p.rsa, crypto.SHA256, hash[:])
		if err != nil {
			return ""
		}

		return base64.StdEncoding.EncodeToString(sign)
	}

	return ""
}
