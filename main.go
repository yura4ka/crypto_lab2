package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/yura4ka/crypto_lab2/rsa"
)

var isBenchmark = flag.Bool("b", false, "")

var aliceRSA *rsa.RSA
var bobRSA *rsa.RSA

func genereteBobMessage(message string) (string, string) {
	bobMessageEnc := rsa.Encrypt(message, aliceRSA.GetPublicKey())
	bobSignature := bobRSA.Sign(message)
	return bobMessageEnc, bobSignature
}

func genereteAliceMessage(message string) (string, string) {
	aliceMessageEnc := rsa.Encrypt(message, bobRSA.GetPublicKey())
	aliceSignature := aliceRSA.Sign(message)
	return aliceMessageEnc, aliceSignature
}

func receiveBobMessage(message, signature string) (string, bool) {
	dec := aliceRSA.Decrypt(message)
	isValidSign := rsa.Verify(dec, signature, bobRSA.GetPublicKey())
	return dec, isValidSign
}

func receiveAliceMessage(message, signature string) (string, bool) {
	dec := bobRSA.Decrypt(message)
	isValidSign := rsa.Verify(dec, signature, aliceRSA.GetPublicKey())
	return dec, isValidSign
}

func bencmarkLen() {
	fmt.Println("Starting bencmark...")
	file, _ := os.OpenFile("results.csv", os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()

	message := "Test message"
	for i := 100; i <= 1000; i++ {
		if i%100 == 0 {
			fmt.Println("Bencmarking for i =", i)
		}
		primes := rsa.PrimeWithLenBits(i, 3)
		start := time.Now()
		r := rsa.NewRSA(primes[0], primes[2])
		enc := rsa.Encrypt(message, r.GetPublicKey())
		r.Decrypt(enc)
		content := fmt.Sprintf("%d,%d\n", i, time.Since(start).Milliseconds())
		file.Write([]byte(content))
	}
	fmt.Println("Done!")
}

func main() {
	flag.Parse()

	if isBenchmark != nil && *isBenchmark {
		bencmarkLen()
		return
	}

	aliceRSA = rsa.NewRSA(nil, nil)
	bobRSA = rsa.NewRSA(nil, nil)

	messages := []string{
		"Hi Bob!",
		"Hi Allice!",
		"How are you doing?",
		"I'm good, thanks",
		"You're welcome",
		"What?",
	}

	for i, v := range messages {
		if i%2 == 0 {
			fmt.Printf("Alice: Send message: %s\n\n", v)
			m, s := genereteAliceMessage(v)
			// Alice send m, s

			// Bob receive m, s
			rm, isValid := receiveAliceMessage(m, s)
			fmt.Printf("Bob: Received message: %s\nValid signature: %v\n\n", rm, isValid)
		} else {
			fmt.Printf("Bob: Send message: %s\n\n", v)
			m, s := genereteBobMessage(v)
			// Bob send m, s

			// Alice receive m, s
			rm, isValid := receiveBobMessage(m, s)
			fmt.Printf("Alice: Received message: %s\nValid signature: %v\n\n", rm, isValid)
		}
	}
}
