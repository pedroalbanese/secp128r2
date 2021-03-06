package secp128r2

import "crypto/elliptic"
import "math/big"

func strbig(s string) (i *big.Int) {
	i = new(big.Int)
	i.SetString(s,0)
	return
}

var secp128r2 = &elliptic.CurveParams{
	P: strbig("0x00fffffffdffffffffffffffffffffffff"), // Prime
	N: strbig("0x003fffffff7fffffffbe0024720613b5a3"), // Order
	B: strbig("0x005eeefca380d02919dc2c6558bb6d8a5d"), // B
	Gx: strbig("0x7b6aa5d85e572983e6fb32a7cdebc140"),  // Generator X
	Gy: strbig("0x27b6916a894d3aee7106fe805fc34b44"),  // Generator Y
	BitSize: 128,
	Name: "secp128r2",
}

// Secp128r2() returns a Curve which implements Secp128r2
func Secp128r2() elliptic.Curve { return secp128r2 }
