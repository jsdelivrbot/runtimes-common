package types

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/GoogleCloudPlatform/runtimes-common/tuf/scheme"
)

var VALID_CRYPTO_SCHEMES = map[string]bool{
	scheme.ECDSA256: true,
	scheme.RSA256:   false, // Not implemented.
}

var CryptoSchemes []string
var ImplementedSchemes []string

type CryptoScheme struct {
	Scheme string
}

func (scheme *CryptoScheme) String() string {
	return scheme.Scheme
}

func (scheme *CryptoScheme) Set(s string) error {
	value, ok := VALID_CRYPTO_SCHEMES[s]
	if ok && value {
		scheme.Scheme = s
		return nil
	}
	if !ok {
		return fmt.Errorf(`%s is not a valid CryptoScheme.
		Please Provide one of %s`, s, strings.Join(CryptoSchemes, ", "))
	}
	return fmt.Errorf(`%s is not a Not Implemented Yet!
		Please Provide one of %s`, s, strings.Join(ImplementedSchemes, ", "))
}

func (scheme *CryptoScheme) Type() string {
	return "types.CryptoScheme"
}

func NewCryptoScheme(val string) *CryptoScheme {
	value, ok := VALID_CRYPTO_SCHEMES[val]
	if ok && value {
		return &CryptoScheme{Scheme: val}
	}
	return nil
}

func (scheme *CryptoScheme) Store(filename string) error {
	schemeJson, err := json.Marshal(scheme)
	if err != nil {
		return fmt.Errorf("Error while marshalling json %s", err.Error())
	}
	return ioutil.WriteFile(filename, schemeJson, 0644)
}

func init() {
	for k, v := range VALID_CRYPTO_SCHEMES {
		CryptoSchemes = append(CryptoSchemes, k)
		if v {
			ImplementedSchemes = append(ImplementedSchemes, k)
		}
	}
}
