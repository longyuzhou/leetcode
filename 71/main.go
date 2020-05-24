package main

import (
	"fmt"
	"strings"
)

func simplifyPath(path string) string {
	stack := []string{}

	start := 0
	bytes := []byte(path)
	size := len(bytes)
	next := func() string {
		i := start
		for i < size && bytes[i] == '/' {
			i++
		}
		j := i
		for j < size && bytes[j] != '/' {
			j++
		}
		start = j + 1
		return string(bytes[min(i, size):min(j, size)])
	}

	for {
		name := next()
		if name == "" {
			break
		}
		if name == "." {
			// ignore
		} else if name == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, name)
		}
	}
	return "/" + strings.Join(stack, "/")
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func test(path, expect string) {
	actual := simplifyPath(path)
	if actual != expect {
		fmt.Printf("fail! expect: %s, got %s\n", expect, actual)
	} else {
		fmt.Println("pass")
	}
}

func main() {
	test("/home/", "/home")
	test("/../", "/")
	test("/home//foo/", "/home/foo")
	test("/a/./b/../../c/", "/c")
	test("/a/../../b/../c//.//", "/c")
	test("/./bGWC/tUnSS/../lAqL/QkbTK/./nP", "/bGWC/lAqL/QkbTK/nP")
}
