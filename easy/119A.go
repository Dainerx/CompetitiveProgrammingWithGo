package main
 
import "fmt"
 
func gcd(a, b int) int {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	for a != b {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}
	return a
}
func main() {
	var a, b, n int
	fmt.Scanf("%d %d %d", &a, &b, &n)
	for {
		stonesTaken := gcd(a, n)
		if stonesTaken > n {
			fmt.Println("1")
			return
		}
		n -= stonesTaken
		stonesTaken = gcd(b, n)
		if stonesTaken > n {
			fmt.Println("0")
			return
		}
		n -= stonesTaken
	}
}
