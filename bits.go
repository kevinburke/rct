// Functions for manipulating bits
package main

func hasBit(n int, pos uint) bool {
	val := n & (1 << pos)
	return (val > 0)
}

// Sets the bit at pos in the integer n.
func setBit(n int, pos uint) int {
	n |= (1 << pos)
	return n
}

// Clears the bit at pos in n.
func clearBit(n int, pos uint) int {
	mask := ^(1 << pos)
	n &= mask
	return n
}
