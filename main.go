package main

import (
	"fmt"
	"gosql/db"
	"gosql/models"
	"gosql/repositories"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

func printAllHelp() {
	fmt.Println("The commands are:")
	fmt.Println()
	fmt.Printf("\tinsert:\tinsert employee data\n")
	fmt.Printf("\tread:\tread employee data\n")
	fmt.Printf("\tupdate:\tupdate employee data\n")
	fmt.Printf("\tdelete:\tdelete employee data\n")
	fmt.Println()
	fmt.Println("use help <topic> for more information about that topic")
	os.Exit(0)
}

func printInsertHelp() {
	fmt.Printf("usage:\tgo run main.go insert <\"Full Name\"> <Email> <Age> <Division> \n")
	os.Exit(0)
}

func printReadHelp() {
	fmt.Println("usage:")
	fmt.Printf("\tread all \t:\tgo run main.go read all\n")
	fmt.Printf("\tread data by ID :\tgo run main.go read <ID>\n")
	os.Exit(0)
}

func main() {
	database := db.ConnectDB()
	employeeRepo := repositories.NewEmployeeRepo(database)

	arguments := os.Args[1:]
	if len(arguments) < 1 {
		printAllHelp()
	}

	switch arguments[0] {
	case "read":
		if arguments[1] == "all" {
			employees, err := employeeRepo.GetEmployees()
			if err != nil {
				msg := fmt.Sprintf("Failed to fetch data: %s", err.Error())
				panic(msg)
			}

			for _, employee := range employees {
				employee.Print()
			}
		} else {
			ID, err := strconv.Atoi(arguments[1])
			if err != nil {
				panic("Please insert integer to fetch employee data By ID")
			}

			employee, err := employeeRepo.GetEmployeeByID(ID)
			if err != nil {
				msg := fmt.Sprintf("Failed to fetch data: %s", err.Error())
				panic(msg)
			}

			employee.Print()
		}

	case "insert":
		age, err := strconv.Atoi(arguments[3])
		if err != nil {
			panic("Please insert integer for paramater age")
		}

		employee := models.Employee{
			FullName: arguments[1],
			Email:    arguments[2],
			Age:      age,
			Division: arguments[4],
		}

		err = employeeRepo.CreateEmployee(&employee)
		if err != nil {
			msg := fmt.Sprintf("Insert failed: %s", err.Error())
			panic(msg)
		}
		fmt.Println("Insert success")

	case "update":
		ID, err := strconv.Atoi(arguments[1])
		if err != nil {
			panic("Please insert integer for paramater ID")
		}

		age, err := strconv.Atoi(arguments[4])
		if err != nil {
			panic("Please insert integer for paramater age")
		}

		employee := models.Employee{
			FullName: arguments[2],
			Email:    arguments[3],
			Age:      age,
			Division: arguments[5],
		}

		err = employeeRepo.UpdateEmployeeByID(ID, employee)
		if err != nil {
			msg := fmt.Sprintf("Failed to update data: %s", err.Error())
			panic(msg)
		}
		fmt.Println("Data updated succesfully")
	case "delete":
		ID, err := strconv.Atoi(arguments[1])
		if err != nil {
			panic("Please insert integer for paramater ID")
		}

		err = employeeRepo.DeleteEmployeeByID(ID)
		if err != nil {
			msg := fmt.Sprintf("Failed to delete data: %s", err.Error())
			panic(msg)
		}

	case "help":
		if len(arguments[1:]) < 1 {
			printAllHelp()
		}

		if arguments[1] == "insert" {
			printInsertHelp()
		} else if arguments[1] == "read" {
			printReadHelp()
		} else if arguments[1] == "update" {

		} else if arguments[1] == "delete" {

		}
	default:
		printAllHelp()
	}
}
