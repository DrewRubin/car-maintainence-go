package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"

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
		addRecord()
	} else if argsWithoutProg[0] == "--remove-record" {

	} else if argsWithoutProg[0] == "--view-vehicles" {
		viewVehicles()
	} else if argsWithoutProg[0] == "--view-record" {
		viewRecord()
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
	viewVehicles()
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Which vehicle number do you want to remove?")
	vehicleNumber, _ := reader.ReadString('\n')

	if _, err := os.Stat("./vehicles.db"); os.IsNotExist(err) {
		os.Create("./vehicles.db")
	}
	database, _ := sql.Open("sqlite3", "./vehicles.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS vehicles (id INTEGER PRIMARY KEY, make TEXT, model TEXT, year INT, mileage INT, tag TEXT)")
	statement.Exec()
	// statement, _ = database.Prepare("INSERT INTO vehicles (make, model, year, mileage, tag) VALUES (?,?,?,?,?)")
	statement, _ = database.Prepare("DELETE FROM vehicles WHERE id=?")
	statement.Exec(vehicleNumber)
	database.Close()

}
func viewVehicles() {
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
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 2, '\t', tabwriter.Debug|tabwriter.AlignRight)
	fmt.Fprintln(w, "Number:\tMake:\tModel:\tYear:\tMileage:\tTag:")

	for rows.Next() {
		rows.Scan(&id, &make, &model, &year, &mileage, &tag)
		fmt.Fprintln(w, strconv.Itoa(id)+") "+"\t"+make+"\t"+model+"\t"+strconv.Itoa(year)+"\t"+strconv.Itoa(mileage)+"\t"+tag)

	}
	fmt.Fprintln(w)
	w.Flush()
	database.Close()

}
func viewRecord() {
	viewVehicles()
	fmt.Println("Which vehicle number do you want to view records for?")
	reader := bufio.NewReader(os.Stdin)
	vehicleNumber, _ := reader.ReadString('\n')
	if _, err := os.Stat("./vehicles.db"); os.IsNotExist(err) {
		os.Create("./vehicles.db")
	}

	database, _ := sql.Open("sqlite3", "./vehicles.db")

	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS vehicles (id INTEGER PRIMARY KEY, make TEXT, model TEXT, year INT, mileage INT, tag TEXT)")
	statement.Exec()
	statement, _ = database.Prepare("CREATE TABLE IF NOT EXISTS records (id INTEGER PRIMARY KEY, vehicleid INTEGER, date TEXT, mileage INT, cost INT, description TEXT)")
	statement.Exec()
	matchingRecords, _ := database.Query("SELECT id, date, mileage, cost, description FROM records WHERE vehicleid=?", vehicleNumber)

	var id int
	var date string
	var mileage int
	var cost int
	var description string
	if matchingRecords == nil {
		return
	}
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 2, '\t', tabwriter.Debug|tabwriter.AlignRight)
	fmt.Fprintln(w, "id:\tDate:\tMileage:\tCost:\tDescription:")
	for matchingRecords.Next() {
		matchingRecords.Scan(&id, &date, &mileage, &cost, &description)
		fmt.Fprintln(w, strconv.Itoa(id)+") "+"\t"+date+"\t"+strconv.Itoa(mileage)+"\t$"+strconv.Itoa(cost)+"\t"+description)
	}
	fmt.Fprintln(w)
	w.Flush()
	database.Close()

}
func addRecord() {
	viewVehicles()
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Which vehicle number do you want to add a record for?")
	vehicleNumber, _ := reader.ReadString('\n')
	fmt.Println("Record date (YYYY-MM-DD):")
	recordDate, _ := reader.ReadString('\n')
	recordDate = recordDate[:len(recordDate)-1]
	fmt.Println("Mileage:")
	mileageString, _ := reader.ReadString('\n')
	mileageString = mileageString[:len(mileageString)-1]
	mileage, _ := strconv.Atoi(mileageString)

	fmt.Println("Cost of repair (integer):")
	costString, _ := reader.ReadString('\n')
	costString = costString[:len(costString)-1]
	cost, _ := strconv.Atoi(costString)

	fmt.Println("Description of service/repair:")
	description, _ := reader.ReadString('\n')

	if _, err := os.Stat("./vehicles.db"); os.IsNotExist(err) {
		os.Create("./vehicles.db")
	}
	database, _ := sql.Open("sqlite3", "./vehicles.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS records (id INTEGER PRIMARY KEY, vehicleid INTEGER, date TEXT, mileage INT, cost INT, description TEXT)")
	statement.Exec()
	statement, _ = database.Prepare("INSERT INTO records (vehicleid, date, mileage, cost, description) VALUES (?,?,?,?,?)")
	statement.Exec(vehicleNumber, recordDate, mileage, cost, description)
	database.Close()
}
