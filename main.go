package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strconv"

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
		removeVehicle()
	} else if argsWithoutProg[0] == "--add-record" {

	} else if argsWithoutProg[0] == "--remove-record" {

	} else if argsWithoutProg[0] == "--view-vehicles" {
		viewVehicles()
	} else {

	}
}

func addVehicle() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Make?")
	vehicleMake, _ := reader.ReadString('\n')
	vehicleMake = vehicleMake[:len(vehicleMake)-1]
	fmt.Println("Model?")
	vehicleModel, _ := reader.ReadString('\n')
	vehicleModel = vehicleModel[:len(vehicleModel)-1]
	fmt.Println("Year?")
	vehicleYear, _ := reader.ReadString('\n')
	vehicleYear = vehicleYear[:len(vehicleYear)-1]
	fmt.Println("Mileage?")
	vehicleMileage, _ := reader.ReadString('\n')
	vehicleMileage = vehicleMileage[:len(vehicleMileage)-1]
	fmt.Println("Registration number (tag)")
	vehicleTag, _ := reader.ReadString('\n')
	vehicleTag = vehicleTag[:len(vehicleTag)-1]

	if _, err := os.Stat("./vehicles.db"); os.IsNotExist(err) {
		os.Create("./vehicles.db")
	}

	database, _ := sql.Open("sqlite3", "./vehicles.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS vehicles (id INTEGER PRIMARY KEY, make TEXT, model TEXT, year INT, mileage INT, tag TEXT)")
	statement.Exec()
	statement, _ = database.Prepare("INSERT INTO vehicles (make, model, year, mileage, tag) VALUES (?,?,?,?,?)")
	statement.Exec(vehicleMake, vehicleModel, vehicleYear, vehicleMileage, vehicleTag)
	database.Close()
}

func removeVehicle() {
	// viewVehicles()
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Println("Which vehicle number do you want to remove?")
	// vehicleNumber, _ := reader.ReadString('\n')

}
func viewVehicles() {
	fmt.Println("Number:\tMake:\tModel:\tYear:\tMileage:\tTag:")
	if _, err := os.Stat("./vehicles.db"); os.IsNotExist(err) {
		os.Create("./vehicles.db")
	}

	database, _ := sql.Open("sqlite3", "./vehicles.db")

	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS vehicles (id INTEGER PRIMARY KEY, make TEXT, model TEXT, year INT, mileage INT, tag TEXT)")
	statement.Exec()

	rows, _ := database.Query("SELECT * FROM vehicles")
	var id int
	var make string
	var model string
	var year int
	var mileage int
	var tag string
	for rows.Next() {
		rows.Scan(&id, &make, &model, &year, &mileage, &tag)
		fmt.Println(strconv.Itoa(id) + ") " + "\t" + make + "\t" + model + "\t" + strconv.Itoa(year) + "\t" + strconv.Itoa(mileage) + "\t\t" + tag)

	}
	database.Close()

}
