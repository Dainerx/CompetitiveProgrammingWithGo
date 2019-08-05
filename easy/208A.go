package main
 
import (
	"fmt"
	"strings"
)
 
func main() {
	const WUB = "WUB"
	var song string
	fmt.Scanf("%s", &song)
	words := strings.Split(song, WUB)
	var result string
	for _, word := range words {
		if len(word) == 0 {
			continue
		} else {
			result += word + " "
		}
	}
	fmt.Println(strings.TrimSpace(result))
}
