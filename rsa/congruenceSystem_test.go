package rsa

import "testing"

func TestSolveCongruenceSystem(t *testing.T) {
	r := SolveCongruenceSystem([]CongruenceParams{
		{Big(2), Big(5)},
		{Big(3), Big(7)},
	})

	if !EQ(r, Big(17)) {
		t.Errorf("got %v", r.String())
	}
}
