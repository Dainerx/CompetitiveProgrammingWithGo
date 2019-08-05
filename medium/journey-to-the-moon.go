package main

import "fmt"

var sizeCC int64

func dfs(g [][]int, visited []bool, src int) {
	visited[src] = true
	adj := g[src]
	for i := 0; i < len(adj); i++ {
		v := adj[i]
		if visited[v] == false {
			sizeCC++
			dfs(g, visited, v)
		}
	}
}

func main() {
	var n, p int
	fmt.Scanf("%d %d", &n, &p)
	g := make([][]int, n)
	visited := make([]bool, n)
	for i := range visited {
		visited[i] = false
	}
	for i := 0; i < p; i++ {
		var u, v int
		fmt.Scanf("%d %d", &u, &v)
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}
	cc := make([]int64, 0)
	for i := 0; i < n; i++ {
		if visited[i] == false {
			sizeCC = 1
			dfs(g, visited, i)
			cc = append(cc, sizeCC)
		}
	}
	var result int64 = 0
	for i := 0; i < len(cc); i++ {
		r := cc[i]
		for j := i + 1; j < len(cc); j++ {
			result += r * cc[j]
		}
	}
	fmt.Println(result)
}
