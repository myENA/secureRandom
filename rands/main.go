package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	srand "github.com/myENA/secureRandom"
)

func main() {
	var err error    // error holder
	var length = 16  // length with default
	var rands string // random string

	// check arguments
	if len(os.Args) > 1 {
		if length, err = strconv.Atoi(os.Args[1]); err != nil {
			fmt.Printf("[ERROR] Invalid length: %s", err.Error())
			os.Exit(1)
		}
	}

	// generate string and check error
	if rands, err = srand.New(length); err != nil {
		fmt.Printf("[ERROR] Failed to generate random string: %s", err.Error())
		os.Exit(1)
	}

	// seed pseudo random generator
	rand.Seed(time.Now().UnixNano())

	// semi-random mixed case
	for idx, r := range rands {
		if rand.Intn(2) == 1 {
			// replace in string
			rands = rands[:idx] + strings.ToLower(string(r)) + rands[idx+1:]
		}
	}

	fmt.Printf("%s\n", rands)
}
