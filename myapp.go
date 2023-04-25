package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Print("Enter a number: ")
    scanner.Scan()
    input := scanner.Text()
    number, err := strconv.Atoi(input)
    if err != nil {
        fmt.Println("Invalid input.")
        os.Exit(1)
    }
    fmt.Printf("The square of %d is %d\n", number, square(number))
}

func square(n int) int {
    return n * n
}
