package main

import "fmt"

func getHint(secret string, guess string) string {

	return ""
}

func test(secret, guess, expect string) {
	actual := getHint(secret, guess)
	if actual != expect {
		fmt.Printf("fail! secret: %s, guess: %s, expect: %s\n", secret, guess, expect)
	}
}

func main() {
	test("1807", "7810", "1A3B")
}
