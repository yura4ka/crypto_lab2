package rsa

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"math/big"
)

type PublicKey struct {
	N, E, P, Q *big.Int
}

type RSA struct {
	_publicKey  PublicKey
	_privateKey *big.Int
}

func NewRSA(p, q *big.Int) *RSA {
	if p == nil {
		p = GenerateRandomPrime(nil, nil)
	}
	if q == nil {
		q = GenerateRandomPrime(nil, nil)
	}

	n := Mul(p, q)
	lambda := LCM(Sub(p, 1), Sub(q, 1))
	e := Big(65537)
	d := Big(0).ModInverse(e, lambda)

	return &RSA{_publicKey: PublicKey{n, e, Big(0).Set(p), Big(0).Set(q)}, _privateKey: d}
}

func Encrypt(message string, key PublicKey) string {
	m, _ := Big(0).SetString(hex.EncodeToString([]byte(message)), 16)
	return Big(0).Exp(m, key.E, key.N).String()
}

func (r RSA) Decrypt(message string) string {
	m := FromString(message)
	p, q, d := r._publicKey.P, r._publicKey.Q, r._privateKey
	// dec := Big(0).Exp(m, d, n)

	congrunces := []CongruenceParams{
		{Pow(Mod(m, p), Mod(d, Sub(p, 1)), p), p},
		{Pow(Mod(m, q), Mod(d, Sub(q, 1)), q), q},
	}
	dec := SolveCongruenceSystem(congrunces)

	result, _ := hex.DecodeString(dec.Text(16))
	return string(result)
}

func (r RSA) Sign(message string) string {
	hash := sha256.New()
	hash.Write([]byte(message))
	h := base64.URLEncoding.EncodeToString(hash.Sum(nil))
	m := Big(0).SetBytes([]byte(h))
	return Pow(m, r._privateKey, r._publicKey.N).String()
}

func Verify(message, signature string, key PublicKey) bool {
	hash := sha256.New()
	hash.Write([]byte(message))
	h := base64.URLEncoding.EncodeToString(hash.Sum(nil))
	m := Big(0).SetBytes([]byte(h))
	s := FromString(signature)
	return EQ(Pow(s, key.E, key.N), m)
}

func (r RSA) GetPublicKey() PublicKey {
	return PublicKey{r._publicKey.N, r._publicKey.E, r._publicKey.P, r._publicKey.Q}
}
