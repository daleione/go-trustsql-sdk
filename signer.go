package trustsql

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/sha256"
	"encoding/base64"
	"math/big"

	secp256k1 "github.com/toxeus/go-secp256k1"
)

// Sign 签名
func Sign(privateKey []byte, data []byte) string {
	secp256k1.Start()
	// privKey := PrivateKeyFromBytes(privateKey)
	dataSHA256 := sha256.Sum256(data)

	/*
		r, s, err := ecdsa.Sign(rand.Reader, privKey, dataSHA256[:])
		if err != nil {
			log.Panic(err)
		}
		signature := append(r.Bytes(), s.Bytes()...)
	*/

	// val, success := secp256k1.Sign(dataSHA256, privateKey, nil)
	var dupPrivKey [32]byte
	copy(dupPrivKey[:32], privateKey[:])
	val, _ := secp256k1.Sign(dataSHA256, dupPrivKey, nil)
	secp256k1.Stop()
	return string(base64.StdEncoding.EncodeToString(val))
}

// Verify 验证签名
func Verify(pubkey, sig, data []byte) bool {
	secp256k1.Start()
	dataSHA256 := sha256.Sum256(data)
	return secp256k1.Verify(dataSHA256, sig, pubkey)
}

func PrivateKeyFromBytes(privateKey []byte) *ecdsa.PrivateKey {
	curve := elliptic.P256()
	x, y := curve.ScalarBaseMult(privateKey)

	privKey := &ecdsa.PrivateKey{
		PublicKey: ecdsa.PublicKey{
			Curve: curve,
			X:     x,
			Y:     y,
		},
		D: new(big.Int).SetBytes(privateKey),
	}

	return privKey
}
