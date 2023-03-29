package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin123"
	dbname   = "db-go-sql"
)

var (
	db  *sql.DB
	err error
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// memeriksa/verifikasi info psqlInfo benar
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	// connecting to database
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database")

	// CreateEmployee()
	// GetEmployees()
	// UpdateEmployee()
	DeleteEmployee()
}

type Employee struct {
	Id        int
	Full_name string
	Email     string
	Age       int
	Division  string
}

func CreateEmployee() {
	var employee = Employee{}

	sqlStatement := `
    INSERT INTO employees (full_name, email, age, division) VALUES ($1, $2, $3, $4)
    returning *
  `

	err = db.QueryRow(sqlStatement, "febrianto", "febri5@gmail.com", 23, "Developer").Scan(&employee.Id, &employee.Full_name, &employee.Email, &employee.Age, &employee.Division)

	if err != nil {
		panic(err)
	}

	fmt.Printf("New Employee Data : %+v \n", employee) // output : New Employee Data : {Id:3 Full_name:febrianto Email:febri2@gmail.com Age:23 Division:Developer}
}

func GetEmployees() {
	var results = []Employee{}

	sqlStatement := `SELECT * FROM employees`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var employee = Employee{}

		err = rows.Scan(&employee.Id, &employee.Full_name, &employee.Email, &employee.Age, &employee.Division)

		if err != nil {
			panic(err)
		}

		results = append(results, employee)
	}

	fmt.Println("Employees :", results)
	// contoh output : [{1 febrianto febri@gmail.com 23 Developer} {2 febrianto febri2@gmail.com 23 Developer} ]
}

func UpdateEmployee() {
	sqlStatement := `
		UPDATE employees 
		SET full_name = $2, email = $3, division = $4, age = $5 
		WHERE id = $1;
	`

	res, err := db.Exec(sqlStatement, 1, "febri update", "febriupdated@gmail.com", "desain", 23)

	if err != nil {
		panic(err)
	}

	count, err := res.RowsAffected()

	if err != nil {
		panic(err)
	}

	fmt.Println("Update data amount :", count) // output : update data amount : 1
}

func DeleteEmployee() {
	sqlStatement := `
		DELETE FROM employees
		WHERE id = $1;
	`
	// jika id tidak ada , maka tidak akan trigger error, knp ya?
	res, err := db.Exec(sqlStatement, 8)

	if err != nil {
		panic(err)
	}

	count, err := res.RowsAffected()

	if err != nil {
		panic(err)
	}

	fmt.Println("Deleted data amount :", count)
}
