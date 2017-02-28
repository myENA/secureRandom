package secureRandom

import (
	"testing"
)

func TestNew(t *testing.T) {
	var lengths = []int{2, 3, 6, 7, 8, 15, 16, 23, 24, 31, 32, 63, 64, 127, 128} // test lengths
	var salt string                                                              // salt string
	var err error                                                                // error holder

	for _, l := range lengths {
		if salt, err = New(l); err != nil {
			t.Errorf("New(%d) failed: %s", l, err.Error())
		}
		if len(salt) != l {
			t.Errorf("New(%d) failed: len(salt) != %d", l, l)
		}
		t.Logf("New(%d) = %s (%d)", l, salt, len(salt))
	}
}
