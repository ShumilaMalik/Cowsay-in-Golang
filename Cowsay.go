package main

import (
	"fmt"
	"bufio"
	"os"
	//"strings"
	//"regexp"
	//"strconv"
)
const inputSize int = 60

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text()) // Println will add back the final '\n'
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}