package flipit

import "errors"

//
// There are two points to note about this function:
//
//  1. We iterate in reverse, because the right most char in the string
//     is at the bottom of the stack, and things should be flipped
//     "above" the bottom of the stack.
//
//  2. We add one to the number of "flips" each time the current char
//     does not equal the last, because (1) if the bottom of the stack
//     is already a +, it can be ignored because changes will only
//     need to happen "above" it in the stack, and (2) if the chars
//     are contiguous, no further "flips" need to take place at that
//     layer.
//
//     iteration | stack state | flips
//     ----------+-------------+------
//               | -+          | 0
//     1         | ++          | 1
//     2         | ++          | 1
//
//     iteration | stack state | flips
//     ----------+-------------+------
//               | +-          | 0
//     1         | -+          | 1
//     2         | ++          | 2
//
//     iteration | stack state | flips
//     ----------+-------------+------
//               | --+-        | 0
//     1         | +-++        | 1
//     2         | -+++        | 2
//     3         | ++++        | 3
//     4         | ++++        | 3
//
func FlipCount(stack string) (int, error) {
	flips := 0
	lastCake := "+"
	stackLength := len(stack)

	if stackLength < 1 {
		return -1, errors.New("stack cannot be an empty string")
	}

	for i := stackLength - 1; i > -1; i-- {
		cake := string(stack[i])

		if cake != lastCake {
			flips += 1
		}

		lastCake = cake
	}

	return flips, nil
}
