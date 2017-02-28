package secureRandom

import (
	crand "crypto/rand"
	"encoding/base64"
	"math/big"
	mrand "math/rand"
	"strings"
)

// sanizize returns the passed string with all '-' and '_' characters replaced
func sanitize(s string) (string, error) {
	var alpha = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789" // our sanitized alphabeb
	var alen = 62                                                                // length of the above alphabet
	var prime *big.Int                                                           // secure prime
	var split []string                                                           // split input string
	var char string                                                              // individual character
	var idx, aidx int                                                            // string index
	var err error                                                                // error holder

	// generate a 128 bit secure prime
	if prime, err = crand.Prime(crand.Reader, 128); err != nil {
		return s, err
	}

	// seed random number generator
	mrand.Seed(prime.Int64())

	// split input string for looping
	split = strings.Split(s, "")

	// loop over split string and replace as needed
	for idx, char = range split {
		aidx = mrand.Intn(alen)
		if char == "-" {
			split[idx] = alpha[aidx : aidx+1]
		}
		if char == "_" {
			split[idx] = alpha[aidx : aidx+1]
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
