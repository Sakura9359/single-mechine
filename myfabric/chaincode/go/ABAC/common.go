package main

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/asn1"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"math/big"
)

func GenerateHash(message string) string {
	h := sha256.New()
	h.Write([]byte(message))
	key := hex.EncodeToString(h.Sum(nil))
	return key
}

func ActionParse(level int) (int, int, int, int, error) {
	if level > 15 {
		return 0, 0, 0, 0, errors.New("action level is not legal")
	}
	a := level / 8
	level %= 8
	b := level / 4
	level %= 4
	c := level / 2
	d := level % 2
	return a, b, c, d, nil
}

func Verify(ap *AP) bool {
	policy := ap.Policy
	policyAsBytes, _ := json.Marshal(policy)
	signature, _ := hex.DecodeString(ap.Sign)
	R, S, _ := UnmarshalECDSASignature(signature)
	flag := ecdsa.Verify(decodePub(ap.Policy.SubjectA.PK), policyAsBytes, R, S)
	return flag
}

func encode(privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey) (string, string) {
	x509Encoded, _ := x509.MarshalECPrivateKey(privateKey)
	pemEncoded := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: x509Encoded})
	x509EncodedPub, _ := x509.MarshalPKIXPublicKey(publicKey)
	pemEncodedPub := pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: x509EncodedPub})
	return string(pemEncoded), string(pemEncodedPub)
}

func decode(pemEncoded string, pemEncodedPub string) (*ecdsa.PrivateKey, *ecdsa.PublicKey) {
	block, _ := pem.Decode([]byte(pemEncoded))
	x509Encoded := block.Bytes
	privateKey, _ := x509.ParseECPrivateKey(x509Encoded)
	blockPub, _ := pem.Decode([]byte(pemEncodedPub))
	x509EncodedPub := blockPub.Bytes
	genericpublicKey, _ := x509.ParsePKIXPublicKey(x509EncodedPub)
	publicKey := genericpublicKey.(*ecdsa.PublicKey)
	return privateKey, publicKey
}

func decodePub(pemEncodedPub string) *ecdsa.PublicKey {
	blockPub, _ := pem.Decode([]byte(pemEncodedPub))
	x509EncodedPub := blockPub.Bytes
	genericpublicKey, _ := x509.ParsePKIXPublicKey(x509EncodedPub)
	publicKey := genericpublicKey.(*ecdsa.PublicKey)
	return publicKey
}

func MarshalECDSASignature(r, s *big.Int) ([]byte, error) {
	return asn1.Marshal(Sign{r, s})
}

func UnmarshalECDSASignature(raw []byte) (*big.Int, *big.Int, error) {
	// Unmarshal
	sig := new(Sign)
	_, err := asn1.Unmarshal(raw, sig)
	if err != nil {
		return nil, nil, fmt.Errorf("failed unmashalling signature [%s]", err)
	}

	// Validate sig
	if sig.R == nil {
		return nil, nil, errors.New("invalid signature, R must be different from nil")
	}
	if sig.S == nil {
		return nil, nil, errors.New("invalid signature, S must be different from nil")
	}

	if sig.R.Sign() != 1 {
		return nil, nil, errors.New("invalid signature, R must be larger than zero")
	}
	if sig.S.Sign() != 1 {
		return nil, nil, errors.New("invalid signature, S must be larger than zero")
	}

	return sig.R, sig.S, nil
}
