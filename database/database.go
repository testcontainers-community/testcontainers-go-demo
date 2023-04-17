package database

import (
"database/sql"
"fmt"
_ "github.com/go-sql-driver/mysql"
	"strings"
	"time"
)

type Student struct {
	ID int
	Fname string
	Lname string
	DateOfBirth time.Time
	Email string
	Address string
	Gender string
}

var db *sql.DB

func Connect(user, pass, host, port, database string) (error) {
	// Opening a database connection.
	var err error
	//db, err = sql.Open("mysql", "theuser:thepass@tcp(localhost:3306)/thedb?multiStatements=true&parseTime=true")
	db, err = sql.Open("mysql", strings.Join([]string{ user, ":", pass, "@tcp(", host, ":", port, ")/", database, "?multiStatements=true&parseTime=true"},""))
	if err != nil {
		return err
	}
	fmt.Println("Connected!")
	return nil
}

func Setup()  {
	table := `CREATE TABLE students (
	id bigint NOT NULL AUTO_INCREMENT,
		fname varchar(50) not null,
		lname varchar(50) not null,
		date_of_birth datetime not null,
		email varchar(50) not null,
		address varchar(50) not null,
		gender varchar(50) not null,
		PRIMARY KEY (id)
	);`
	_, err := db.Exec(table)
	if err != nil {
		if strings.Contains(err.Error(),"already exists"){
			return
		}
		panic(err)
	}

	//inserting records
	_, err = db.Exec(records)
	if err != nil {
		panic(err)
	}
}

func Clear()  {
	_, err := db.Exec(`DROP TABLE IF EXISTS students`)
	if err != nil {
		panic(err)
	}
}

func AddStudent(s Student) (int64, error){
	query := "insert into students (fname, lname, date_of_birth, email, gender, address) values (?, ?, ?, ?, ?, ?);"
	result, err := db.Exec(query, s.Fname,s.Lname, s.DateOfBirth, s.Email, s.Gender, s.Address)
	if err != nil {
		return 0, fmt.Errorf("AddStudent Error: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addSudent Error: %v", err)
	}

	return id, nil
}

func FetchStudents() ([]Student, error) {
	// A slice of Students to hold data from returned rows.
	var students []Student

	rows, err := db.Query("SELECT * FROM students")
	if err != nil {
		return nil, fmt.Errorf("FetchStudents %v", err)
	}
	defer rows.Close()

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var s Student
		if err := rows.Scan(&s.ID, &s.Fname, &s.Lname, &s.DateOfBirth, &s.Email, &s.Address, &s.Gender ); err != nil {
			return nil, fmt.Errorf("FetchStudents %v", err)
		}
		students = append(students, s)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("FetchStudents %v", err)
	}
	return students, nil
}

func StudentByID(id int64) (Student, error){
	var st Student

	row := db.QueryRow("SELECT * FROM students WHERE id = ?", id)
	if err := row.Scan(&st.ID, &st.Fname, &st.Lname, &st.DateOfBirth, &st.Email, &st.Address, &st.Gender ); err != nil {
		if err == sql.ErrNoRows {
			return st, fmt.Errorf("studentById %d: no such student", id)
		}
		return st, fmt.Errorf("studentById %d: %v", id, err)
	}
	return st, nil
}
