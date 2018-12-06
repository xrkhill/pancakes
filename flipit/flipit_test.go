package flipit

import "testing"

func TestFlipCount(t *testing.T) {
	tables := []struct {
		stack         string
		expectedFlips int
	}{
		{"", -1},
		{"-", 1},
		{"-+", 1},
		{"+-", 2},
		{"+++", 0},
		{"--+-", 3},
	}

	for _, table := range tables {
		flips, err := FlipCount(table.stack)

		if flips != table.expectedFlips {
			t.Errorf("Expected flips to be %d, got %d (err: %s)", table.expectedFlips, flips, err)
		}
	}
}
