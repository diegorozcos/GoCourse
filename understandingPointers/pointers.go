package main

import "fmt"

func main() {
	age := 32 // Value 32 stored somewhere in memory

	agePointer := &age

	fmt.Println("Age:", *agePointer)
	editAgeToAdultYears(agePointer)
	fmt.Println(age)
}

func editAgeToAdultYears(age *int) {
	*age = *age - 18
}
