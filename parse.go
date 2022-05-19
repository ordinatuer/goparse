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
	//insertQuery := "INSERT INTO corruption VALUES((SELECT MAX(id)+1 FROM corruption)"
	insertQuery := "INSERT INTO corruption VALUES"
	valuesQuery := "(%d"
	i := 2
	id := 1
	bulkSize := 100

	for i < 22 {
		//valuesQuery += fmt.Sprintf(", $%d", i)
		valuesQuery += ", '%s'"
		i++
	}
	valuesQuery += ")"

	values := ""
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

			//fmt.Println(line[1])

			// stmt, err := db.Prepare(insertQuery)
			// if err != nil {
			// 	fmt.Println("Prepare query error", err)
			// 	break
			// }
			id++
			//_, err = stmt.Exec(
			values += fmt.Sprintf(valuesQuery,
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

			if id % bulkSize == 0 {
				_, err = db.Exec(insertQuery + values)
				if err != nil {
					fmt.Println("Insert query error", err)
				}

				fmt.Println("Ins 100 \n")

				//fmt.Println(insertQuery + values)
				values = ""
			} else {
				values += ","
			}

			// if err != nil {
			// 	fmt.Println("Execute query error", err)
			// }
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