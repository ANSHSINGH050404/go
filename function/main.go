package main

import "fmt"

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}

	return a/b , nil
}

func main() {
	result, err := divide(33, 11)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Result:", result)

	var a string = "hello"
b := "world"

	fmt.Println(a, b)

}