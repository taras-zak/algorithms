package main

import (
	"fmt"
	"strings"
)

func alienOrder(words []string) string {
	adj := make(map[byte]map[byte]struct{})
	indegree := make(map[byte]int)
	for _, word := range words {
		for i := range word {
			adj[word[i]] = make(map[byte]struct{})
			indegree[word[i]] = 0
		}
	}
	for i := 0; i < len(words)-1; i++ {
		curr := words[i]
		next := words[i+1]
		for j := 0; j < len(curr) && j < len(next); j++ {
			if curr[j] != next[j] {
				if _, ok := adj[curr[j]][next[j]]; !ok {
					adj[curr[j]][next[j]] = struct{}{}
					indegree[next[j]]++
				}
				break
			}
			if len(curr) > len(next) && j == len(next)-1 {
				return ""
			}
		}
	}

	res := &strings.Builder{}
	queue := make([]byte, 0)
	for k, v := range indegree {
		if v == 0 {
			queue = append(queue, k)
		}
	}
	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]
		res.WriteByte(cur)
		for k := range adj[cur] {
			indegree[k]--
			if indegree[k] == 0 {
				queue = append(queue, k)
			}
		}
	}
	if res.Len() != len(indegree) {
		return ""
	}
	return res.String()
}

func main() {
	fmt.Println(alienOrder([]string{"wrt", "wrf", "er", "ett", "rftt"}))
	fmt.Println(alienOrder([]string{"abc", "ab"}))
}
