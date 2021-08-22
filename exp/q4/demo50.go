package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Enter function main.")

	//panic1()

	//panic2()

	panic3()

	fmt.Println("Exit function main.")
}

// recover 函数错处用法
func panic1() {
	fmt.Println("Enter function panic1.")
	// 引发panic。
	panic(errors.New("something wrong"))
	// recover函数的错误用法。
	p := recover()
	fmt.Printf("panic: %s\n", p)
	fmt.Println("Exit function panic1.")
}

// recover 函数错处用法
func panic2() {
	fmt.Println("Enter function panic2.")
	// recover函数的错误用法。
	fmt.Printf("no panic: %v\n", recover())
	// 引发panic。
	panic(errors.New("something wrong"))
	fmt.Println("Exit function panic2.")
}

// recover 函数错处用法
func panic3() {
	fmt.Println("Enter function panic3.")
	defer func() {
		fmt.Println("Enter defer function.")

		// recover函数的正确用法。
		if p := recover(); p != nil {
			fmt.Printf("panic: %s\n", p)
		}

		fmt.Println("Exit defer function.")
	}()
	// 引发panic。
	panic(errors.New("something wrong"))
	fmt.Println("Exit function panic3.")
}
