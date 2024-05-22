package rsa

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"math/big"
)

var ErrWrongType = errors.New("error wrong type: excpected int64 or bigint")

type Int interface {
	big.Int | int
}

func Big(x int64) *big.Int {
	return big.NewInt(x)
}

func toBigInt(n interface{}) *big.Int {
	switch v := n.(type) {
	case int:
		return Big(int64(v))
	case *big.Int:
		return v
	default:
		panic(ErrWrongType)
	}
}

func Cmp(x *big.Int, y interface{}) int {
	return x.Cmp(toBigInt(y))
}

func GE(x *big.Int, y interface{}) bool {
	return Cmp(x, y) == 1
}

func LE(x *big.Int, y interface{}) bool {
	return Cmp(x, y) == -1
}

func LEQ(x *big.Int, y interface{}) bool {
	return Cmp(x, y) <= 0
}

func EQ(x *big.Int, y interface{}) bool {
	return Cmp(x, y) == 0
}

func IsOdd(x *big.Int) bool {
	temp := new(big.Int)
	temp.And(x, Big(1))
	return EQ(temp, 1)
}

func IsEven(x *big.Int) bool {
	return !IsOdd(x)
}

func RandomBigint(a, b interface{}) *big.Int {
	x := toBigInt(a)
	y := toBigInt(b)

	if EQ(x, y) {
		return Big(0).Set(x)
	}
	max := Sub(y, x)
	r, err := rand.Int(rand.Reader, max)
	if err != nil {
		panic(err)
	}

	return r.Add(r, x)
}

func Add(x *big.Int, y interface{}) *big.Int {
	return Big(0).Add(x, toBigInt(y))
}

func Sub(x *big.Int, y interface{}) *big.Int {
	return Big(0).Sub(x, toBigInt(y))
}

func Mul(x *big.Int, y interface{}) *big.Int {
	return Big(0).Mul(x, toBigInt(y))
}

func Div(x *big.Int, y interface{}) *big.Int {
	return Big(0).Div(x, toBigInt(y))
}

func Mod(x *big.Int, y interface{}) *big.Int {
	return Big(0).Mod(x, toBigInt(y))
}

func Sqrt(x *big.Int) *big.Int {
	return Big(0).Sqrt(x)
}

func Inc(x *big.Int) {
	x.Add(x, Big(1))
}

func GCD(x, y *big.Int) *big.Int {
	return Big(0).GCD(nil, nil, x, y)
}

func FromString(s string) *big.Int {
	r, _ := Big(0).SetString(s, 10)
	return r
}

func ToBase64(n *big.Int) string {
	return base64.RawStdEncoding.EncodeToString([]byte(n.String()))
}

func GenerateRandomPrime(a, b *big.Int) *big.Int {
	start := Big(1)
	end := Big(1)

	if a != nil {
		start.Set(a)
	} else {
		start.Lsh(start, 300)
	}

	if b != nil {
		end.Set(b)
	} else {
		end.Lsh(end, 333)
	}

	r := Big(10)

	for ok := true; ok; ok = !r.ProbablyPrime(10) {
		r = RandomBigint(start, end)
	}

	return r
}

func LCM(x, y *big.Int) *big.Int {
	product := Mul(x, y)
	product.Abs(product)

	gcd := GCD(x, y)
	return product.Div(product, gcd)
}

func Pow(x, y, m *big.Int) *big.Int {
	return Big(0).Exp(x, y, m)
}

func PrimeWithLenBits(n, len int) []*big.Int {
	result := make([]*big.Int, 0)
	counter := 0
	start := Big(1)
	start.Lsh(start, uint(n-1))
	end := Big(1)
	end.Lsh(start, uint(n))

	for ; LE(start, end); Inc(start) {
		if start.ProbablyPrime(10) {
			result = append(result, Big(0).Set(start))
			counter++
			if len != 0 && counter == len {
				break
			}
		}
	}

	return result
}
