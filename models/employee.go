package models

import "fmt"

type Employee struct {
	ID       int
	FullName string
	Email    string
	Age      int
	Division string
}

func (e *Employee) Print() {
	fmt.Println("ID:", e.ID)
	fmt.Println("Full Name:", e.FullName)
	fmt.Println("Email:", e.Email)
	fmt.Println("Age:", e.Age)
	fmt.Println("Division:", e.Division)
	fmt.Println()
}
