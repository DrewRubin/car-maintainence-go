package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	//argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	//Ask users for
	//fmt.Println(argsWithProg)
	numArguments := len(argsWithoutProg)

	if numArguments == 0 {
		fmt.Println("This is a tool to keep records of your vehicle maintenance")
		fmt.Println("Useful commands: ")
		fmt.Println("\t--add-vehicle")
		fmt.Println("\t--remove-vehicle")
		fmt.Println("\t--add-record")
		fmt.Println("\t--remove-record")
		fmt.Println("\t--view-vehicles")
	}
	fmt.Println(argsWithoutProg[0])

	if numArguments > 2 {
		fmt.Println("Too many arguments.")
	}
	if argsWithoutProg[0] == "--add-vehicle" {
		addVehicle()
	} else if argsWithoutProg[0] == "--remove-vehicle" {

	} else if argsWithoutProg[0] == "--add-record" {

	} else if argsWithoutProg[0] == "--remove-record" {

	} else if argsWithoutProg[0] == "--view-vehicles" {

	} else {

	}
}

func addVehicle() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Make?")
	vehicleMake, _ := reader.ReadString('\n')
	fmt.Println("Model?")
	vehicleModel, _ := reader.ReadString('\n')
	fmt.Println("Year?")
	vehicleYear, _ := reader.ReadString('\n')
	fmt.Println("Mileage?")
	vehicleMileage, _ := reader.ReadString('\n')
	fmt.Println("Registration number (tag)")
	vehicleTag, _ := reader.ReadString('\n')

	os.Create("./vehicles.db")
	database, _ := sql.Open("sqlite3", "./vehicles.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS vehicles (make TEXT, model TEXT, year TEXT, mileage TEXT, tag TEXT)")
	statement.Exec()
	statement, _ = database.Prepare("INSERT INTO vehicles (make, model, year, mileage, tag) VALUES (?,?,?,?,?)")
	statement.Exec(vehicleMake, vehicleModel, vehicleYear, vehicleMileage, vehicleTag)
}
