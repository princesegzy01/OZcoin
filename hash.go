package zebracoin

import (
	"crypto/sha256"
	"math/big"
)

type SHA256Sum [32]byte

func Hash(data []byte) SHA256Sum {
	return sha256.Sum256(data)
}

func HashToPt(data []byte) ECCPoint {
	h := Hash(data)
	x, y := CURVE.Params().ScalarBaseMult(h[:])

	return ECCPoint{x, y}
}

func HashPt(m []byte, p ECCPoint) SHA256Sum {
	data := []byte{}
	data = append(data, m...)
	data = append(data, p.X.Bytes()...)
	data = append(data, p.Y.Bytes()...)

	return sha256.Sum256(data)
}

func HashPtIdx(m []byte, p ECCPoint, i, j uint64) SHA256Sum {
	data := []byte{}
	data = append(data, m...)
	data = append(data, p.X.Bytes()...)
	data = append(data, p.Y.Bytes()...)
	iInt, jInt := &big.Int{}, &big.Int{}
	iInt.SetUint64(i)
	jInt.SetUint64(j)
	data = append(data, iInt.Bytes()...)
	data = append(data, jInt.Bytes()...)

	return sha256.Sum256(data)
}