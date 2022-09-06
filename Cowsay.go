package main

import (
	"fmt"
	"bufio"
	"os"
	//"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	fmt.Printf("Hallo Techstarter: %q", input)


}