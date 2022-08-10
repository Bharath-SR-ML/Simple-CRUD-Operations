// This program deals with student database.
// Student Table consists of name and rollno
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Student struct {
	rollno string `json:"rollno"`
	name   string `json:"name"`
}

func main() {
	var ch int
	fmt.Println("Student Database")
	db, err := sql.Open("mysql", "root:root1234@tcp(127.0.0.1:3306)/sample")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Connected to Student Table.")

	fmt.Println("Press\n1.INSERT\n2.DISPLAY\n3.UPDATE\n4.DELETE")
	fmt.Print("Enter your choice : ")
	fmt.Scan(&ch)

	if ch == 1 {
		var name, rollno string
		fmt.Println("Enter the name and rollno : ")
		fmt.Scan(&name)
		fmt.Scan(&rollno)
		stmt := fmt.Sprintf("INSERT INTO student values ('%v','%v')", rollno, name)
		insert, err1 := db.Query(stmt)
		if err1 != nil {
			panic(err1.Error())
		}
		defer insert.Close()

		fmt.Println("Inserted Successfully.")
	} else if ch == 2 {
		list, err := db.Query("select * from student")

		if err != nil {
			panic(err.Error())
		}
		for list.Next() {
			var s1 Student
			err = list.Scan(&s1.rollno, &s1.name)
			if err != nil {
				panic(err.Error())
			}
			fmt.Println(s1.name, "-", s1.rollno)
		}
	} else if ch == 3 {
		var name, rollno string
		fmt.Println("Enter the RollNo : ")
		fmt.Scan(&rollno)
		fmt.Println("Enter the new name : ")
		fmt.Scan(&name)
		stmt := fmt.Sprintf("update student set name='%v' where rollno='%v'", name, rollno)
		up, err := db.Prepare(stmt)
		if err != nil {
			panic(err.Error())
		}
		a, e := up.Exec()
		if e != nil {
			panic(e.Error())
		}
		rows, e1 := a.RowsAffected()
		if e1 != nil {
			panic(e1.Error())
		}
		fmt.Print("Rows Affected : ", rows)
	} else if ch == 4 {
		var rollno string
		fmt.Println("Enter the RollNo : ")
		fmt.Scan(&rollno)
		stmt := fmt.Sprintf("delete from student where rollno='%v'", rollno)
		list, err1 := db.Prepare(stmt)
		if err1 != nil {
			panic(err1.Error())
		}
		res, e := list.Exec()
		if e != nil {
			panic(e.Error())
		}
		rows, e1 := res.RowsAffected()
		if e1 != nil {
			panic(e1.Error())
		}
		fmt.Println("Rows Affected : ", rows)

	} else {
		fmt.Println("Please select the correct option.")
	}

}
