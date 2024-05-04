package main

import (
	"fmt"
	"hotels/tasks"
)

func main() {
	fmt.Println("Task for programming number 1. Correct form")
	fmt.Println(tasks.GetStringInCorrectForm(1))
	fmt.Println(tasks.GetStringInCorrectForm(2))
	fmt.Println(tasks.GetStringInCorrectForm(5))
	fmt.Println(tasks.GetStringInCorrectForm(10))
	fmt.Println(tasks.GetStringInCorrectForm(11))
	fmt.Println(tasks.GetStringInCorrectForm(25))
	fmt.Println(tasks.GetStringInCorrectForm(41))
	fmt.Println(tasks.GetStringInCorrectForm(1048))
	fmt.Println("Task for programming number 2. Common divisors.")
	fmt.Println(tasks.CalculaeCommonDivisors([]int{42,12,18}))
	fmt.Println(tasks.CalculaeCommonDivisors([]int{42,12,18,15}))
	fmt.Println(tasks.CalculaeCommonDivisors([]int{42,12,18,24}))
	fmt.Println("Task for programming number 3. Prime numbers.")
	fmt.Println(tasks.FindPrimes(11, 20))
	fmt.Println(tasks.FindPrimes(2, 97))
	fmt.Println("Task for programming number 4. Table of multiplications.")
	t := tasks.New(5)
	t.Fill()
	t.Print()
	fmt.Println("10x10")
	t = tasks.New(10)
	t.Fill()
	t.Print()
}