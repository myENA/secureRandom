package secureRandom

import (
	crand "crypto/rand"
	"encoding/base64"
	"math/big"
	mrand "math/rand"
	"strings"
)

// sanitized alphabet
const saniBet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

// sanizize returns the passed string with all '-' and '_' characters replaced
func sanitize(s string) (string, error) {
	var alen = len(saniBet) // length of the sanitized alphabet
	var prime *big.Int      // random secure prime
	var split []string      // split input string
	var err error           // error holder

	// generate a 128 bit secure prime
	if prime, err = crand.Prime(crand.Reader, 128); err != nil {
		return s, err
	}

	// seed random number generator
	mrand.Seed(prime.Int64())

	// split input string for looping
	split = strings.Split(s, "")

	// loop over split string and replace as needed
	for idx, char := range split {
		ridx := mrand.Intn(alen)
		if char == "-" {
			split[idx] = saniBet[ridx : ridx+1]
		}
		if char == "_" {
			split[idx] = saniBet[ridx : ridx+1]
		}
	}

	// return the string - no error
	return strings.Join(split, ""), nil
}

// New builds a new secure salt of the given length
func New(n int) (string, error) {
	var randomBytes []byte // random bytes
	var randomSize int     // random byte length
	var err error          // error holder

	// calculate random byte length needed to generate
	// a base64 string of the requested length
	randomSize = (6*n-5)/8 + 1

	// init random bytes
	randomBytes = make([]byte, randomSize)

	// seed random bytes
	if _, err = crand.Read(randomBytes); err != nil {
		return "", err
	}

	// return sanitized salt
	return sanitize(base64.RawURLEncoding.EncodeToString(randomBytes))
}
