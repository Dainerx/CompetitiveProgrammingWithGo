package main

import (
	"fmt"
)

type contest struct {
	before int
	after  int
}

func main() {
	var n int
	fmt.Scanf("%d\n", &n)
	interested := []contest{}
	for index := 0; index < n; index++ {
		var name string
		var before, after int
		fmt.Scanf("%s %d %d\n", &name, &before, &after)
		if before >= 2400 {
			add := contest{before, after}
			interested = append(interested, add)
		}
	}

	found := false
	for _, contest := range interested {
		if contest.after > contest.before {
			found = true
			break
		}
	}

	if found == false {
		fmt.Print("NO")
	} else {
		fmt.Print("YES")
	}

}
