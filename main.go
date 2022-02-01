package secp128r2

import "crypto/elliptic"
import "math/big"

func strbig(s string) (i *big.Int) {
	i = new(big.Int)
	i.SetString(s,0)
	return
}

var secp128r2 = &elliptic.CurveParams{
	P: strbig("0xfffffffdffffffffffffffffffffffff"), // Prime
	N: strbig("0xd6031998d1b3bbfebf59cc9bbff9aee1"), // Order
	B: strbig("0x5eeefca380d02919dc2c6558bb6d8a5d"), // B
	Gx: strbig("0x7b6aa5d85e572983e6fb32a7cdebc140"),  // Generator X
	Gy: strbig("0x27b6916a894d3aee7106fe805fc34b44"),  // Generator Y
	BitSize: 128,
	Name: "secp128r2",
}

// Secp128r2() returns a Curve which implements Secp128r2
func Secp128r2() elliptic.Curve { return secp128r2 }
