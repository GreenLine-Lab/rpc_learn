package crypto

import (
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"testing"
)

func TestRSASign(t *testing.T) {

	testPrivateKeyRSAKeyString := "MIICXQIBAAKBgQDVQdH/EJnWDwfOERPknxf/LQbyNJHLmbXNVvJrIkSuKQ4sGC+IUDyiN0ByPHsQnIFX6rQ/5HAreAAWJCKef/WaGTfK5FEvO+NQx8jrdFYQLxkzn1XOhtIm4SCi/nnNUpsr+4ZXSJL/kNQodFlZ6qLLAPupeatHLsj1t/Z3kN+CMwIDAQABAoGBAIj5Q9ZeZVDWSp8S3QIlJZai5yk+lr59NgZz6DTjx9VNTWsJBc66KpfOgThenHTo8TL0711ybpGc/AE8qHfBnI93CXqvDiah1Srvda5bGAdfs15JsEqFSYpTDu3rEV+rqhIXWRGOuxb6Hod8u5CCsOu3VcHw817Eerq0rbklykGRAkEA+SekkyGFXNsR/S7kMvM0Luzey4HukUzyzL1UMh8NXZjCPCsQJTfHlLREvBpUELew+bcUBcGBh6jsjw+TAEVAVQJBANsds4RgUYeUYrfE6sbKV2xAVwjkc7ulroLlR0TrxI+FU+Gpe2rALOvN8rWWlBxxUhl0hgmihX4Z/t+ODbEYIGcCQCEop3x1T5xmA2TgorotJ9q+53/KEQgBZ6bb46KwA0VbmS4MxR9O5x7hRuyJzHpVGBaDyQRjFmwJjUIU3omArekCQH1k3PQXrvC5AVmLjpQe/bLMi0JigrzTTzBHh+awSjecJGnS9PcdbVew74Ht32r7ivFYVtCyulR1Cf/jqoA7f2cCQQCeuepZxLx19Aj+vIuiiTBEccvpPtTrUbC+hYudiiOMKcdQCF/69WB4DXIDQ7RTswyVT057gxTQPkxBkMSiOsH4"
	testPrivateKeyRSAByte, err := base64.StdEncoding.DecodeString(testPrivateKeyRSAKeyString)
	if err != nil {
		t.Fatal(err.Error())
	}

	testPrivateKeyRSA, err := x509.ParsePKCS1PrivateKey(testPrivateKeyRSAByte)
	if err != nil {
		t.Fatal(err.Error())
	}

	priv := Private{
		rsa:   testPrivateKeyRSA,
	}
	pub := priv.Public()

	var TestData = struct {
		Message string
		Flag 	bool
		Number  uint64
	}{
		Message: "Test",
		Flag:    false,
		Number:  15,
	}

	sign := priv.Sign(&TestData)
	fmt.Printf("Sign: %s\n", sign)

	if !pub.VerifySign(&TestData, sign) {
		t.Fatal("Unauthorized")
	}

	fmt.Println("Success")
}