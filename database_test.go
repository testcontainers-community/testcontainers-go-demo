package main

import (
	"fmt"
	"studentAPI/database"
	"testing"
	"time"
)

func TestA(t *testing.T) {

	//variable to hold student ID returned after the creation of the record
	var sID int64

	//test case for AddStudent database function
	t.Run(fmt.Sprintf("AddStudent Test"), func(t *testing.T) {
		s := database.Student{
			Fname:       "Leon",
			Lname:       "Ashling",
			DateOfBirth: time.Date(1994, time.August, 14, 23, 51, 42, 0, time.UTC),
			Email:       "lashling5@senate.gov",
			Address:     "39 Kipling Pass",
			Gender:      "Male",
		}

		//adding student record to table
		var err error
		sID, err = database.AddStudent(s)
		if err != nil {
			fmt.Println(err)
		}
		t.Logf("Created student successfully with ID: %v \n", sID)

	})

	//test case for StudentByID database function
	t.Run(fmt.Sprintf("StudentByID Test"), func(t *testing.T) {
		//selecting student by ID
		st, err := database.StudentByID(sID)
		if err != nil {
			fmt.Println(err)
		}
		t.Logf("Retreived student by ID successfully with ID: %v \n", sID)
		t.Logf("Student Details: %v \n", st)

	})

	// test case for FetchStudents database function
	t.Run(fmt.Sprintf("FetchStudents Test"), func(t *testing.T) {
		//retrieving all records
		students, err := database.FetchStudents()
		// check if errors were returned
		if err != nil {
			t.Error(err)
		}
		t.Logf("Fetched all students data successfully with total of %v records \n", len(students))
	})

}
