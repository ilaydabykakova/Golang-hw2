package main

import "fmt"

//This function divides a by greatest divisible  power of b
func DivisibleMax(x, y int) int {
	for x%y == 0 {
		x = x / y
	}
	return x
}

//Function to check if a number is ugly or not
func Ugly(number int) int {

	number = DivisibleMax(number, 2)
	number = DivisibleMax(number, 3)
	number = DivisibleMax(number, 5)

	if number == 1 {
		return 1 //if number is 1 return again
	} else {
		return 0 //if number is not 1 return main
	}
}

//Function to get the nth ugly number
func getUgly(n int) int {

	var i int = 1     //Counting "1" number too.
	var count int = 1 // ugly number count

	// Check for all integers untill ugly count becomes n
	for n > count {
		i++
		if Ugly(i) == 1 {
			count++
		}
	}
	return i
}

func main() {

	//Test above all functions
	fmt.Println("\nFind n’th Ugly Numbers\n")

	fmt.Println("7th Ugly Number :  ", getUgly(7))

	fmt.Println("10th Ugly Number : ", getUgly(10))

	fmt.Println("15th Ugly Number : ", getUgly(15))

	fmt.Println("150th Ugly Number :", getUgly(150))

}
