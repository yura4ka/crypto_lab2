package rsa

import "math/big"

// x = c mod m

type CongruenceParams struct {
	C, M *big.Int
}

func SolveCongruenceSystem(data []CongruenceParams) *big.Int {
	M := Big(1)
	for _, v := range data {
		M.Mul(M, v.M)
	}

	result := Big(0)
	for _, v := range data {
		mi := Div(M, v.M)
		inv := Big(0).ModInverse(mi, v.M)
		mi.Mul(mi, inv).Mul(mi, v.C)
		result.Add(result, mi).Mod(result, M)
	}

	return result
}
