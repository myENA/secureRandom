package secureRandom

import (
	crand "crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"math/big"
	mrand "math/rand"
	"strings"
)

// sanitized alphabet
const saniBet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

// package random number generator
var prng *mrand.Rand

// init the package random number generator
func init() {
	var prime *big.Int // random secure prime
	var err error      // error holder

	// generate a 63 bit secure prime
	if prime, err = crand.Prime(crand.Reader, 63); err != nil {
		panic(fmt.Sprintf("secureRandom - failed to generate 63 bit prime seed: %s", err.Error()))
	}

	// initialize package random number generator
	prng = mrand.New(mrand.NewSource(prime.Int64()))
}

// sanizize returns the passed string with all '-' and '_' characters replaced
func sanitize(s string) string {
	var alen = len(saniBet) // length of the sanitized alphabet
	var split []string      // split input string

	// split input string for looping
	split = strings.Split(s, "")

	// loop over split string and replace as needed
	for idx, char := range split {
		ridx := prng.Intn(alen)
		if char == "-" {
			split[idx] = saniBet[ridx : ridx+1]
		}
		if char == "_" {
			split[idx] = saniBet[ridx : ridx+1]
		}
	}

	// return the string - no error
	return strings.Join(split, "")
}

// New builds a new secure salt of the given length
func New(n int) (string, error) {
	var randomBytes []byte // random bytes
	var randomSize int     // random byte length
	var err error          // error holder

	// why ... just why
	if n < 2 {
		return "", errors.New("Insufficient length requested: 'n' must be >= 2")
	}

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
	return sanitize(base64.RawURLEncoding.EncodeToString(randomBytes)), nil
}
