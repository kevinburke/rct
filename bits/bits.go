// Functions for manipulating bits
package bits

func On(n int, pos uint) bool {
	val := n & (1 << pos)
	return (val > 0)
}

// Sets the bit at pos in the integer n.
func Set(n int, pos uint) int {
	n |= (1 << pos)
	return n
}

// Only set the bit if the conditional is true.
func SetCond(n int, pos uint, question bool) int {
	if question {
		n |= (1 << pos)
	}
	return n
}

// Clears the bit at pos in n.
func Clear(n int, pos uint) int {
	mask := ^(1 << pos)
	n &= mask
	return n
}
