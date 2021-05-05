package kyu7

// Find the total sum of internal angles (in degrees) in an n-sided simple polygon. N will be greater than 2.
func Angle(n int) int {
	return 180 * (n - 2)
}
