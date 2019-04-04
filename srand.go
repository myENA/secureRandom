package secureRandom

import (
	"crypto/rand"
	"encoding/base32"
	"errors"
)

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
	// a base32 string of the requested length
	randomSize = (n * 5) / 8 + 1

	// init random bytes
	randomBytes = make([]byte, randomSize)

	// read random bytes
	if _, err = rand.Read(randomBytes); err != nil {
		return "", err
	}

	// return encoded salt
	return base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(randomBytes)[0:n], nil
}
