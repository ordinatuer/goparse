package main

import (
	"flist"

	"fmt"
	"os"
	"time"
	"encoding/csv"

	"database/sql"
	_ "github.com/lib/pq"
)

func main() {
	_t := time.Now() ////

	db, err := dbConnect()
	defer db.Close()
	if err != nil {
		fmt.Println("DB connection error")
		return
	}

	flist := flist.GetList(".")

	id := 1
	row, err := db.Query("SELECT MAX(id)+1 FROM corruption")
	defer row.Close()
	if err != nil {
		fmt.Println("Query max(id) error", err)
	} else {
		row.Next()
		err = row.Scan(&id)
		if err != nil {
			fmt.Println("Scan max(id) error", err) // id = 1
		}
	}

	insertQuery := "INSERT INTO corruption VALUES($1"
	i := 2
	for i < 22 {
		insertQuery += fmt.Sprintf(", $%d", i)
		i++
	}
	insertQuery += ")"

	stmt, err := db.Prepare(insertQuery)
	if err != nil {
		fmt.Println("Prepare query error", err, "\n", insertQuery)
		return
	}
	defer stmt.Close()

	for _, file := range flist {
		fileOpen, err := os.Open(file)
		if err != nil {
			panic(err)
		}
		defer fileOpen.Close()

		reader := csv.NewReader(fileOpen)
		reader.Comma = ','

		reader.Read()

		for {
			line, e := reader.Read()
			if e != nil {
				fmt.Println(file, e)
				break
			}

			id++
			_, err = stmt.Exec(
				id,
				line[0],
				line[1],
				line[2],
				line[3],
				line[4],
				line[5],
				line[6],
				line[7],
				line[8],
				line[9],
				line[10],
				line[11],
				line[18],
				fmt.Sprintf("(%s, %s)", line[12], line[13]),
				line[12],
				line[13],
				line[14],
				line[15],
				line[16],
				line[17])

			if err != nil {
				fmt.Println("Execute query error", err)
			}
		}
	}

	fmt.Println(time.Now().Sub(_t)) ////
}

func dbConnect() (*sql.DB, error) {
	const pgConnectUrl = "postgresql://smf_gps_user:smfgpspassword@10.20.0.4:5432/smf_gps_db?sslmode=disable"

	db, err := sql.Open("postgres", pgConnectUrl)
	if err != nil {
		fmt.Println("Open db error ", err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Connection db error ", err)
	}

	return db, err
}