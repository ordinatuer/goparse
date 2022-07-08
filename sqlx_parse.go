package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const batchSize int64 = 100

// file headers
// id,first_name,full_name,email,phone_number,address_city,address_street,address_house,address_entrance,address_floor,address_office,address_comment,location_latitude,location_longitude,amount_charged,user_id,user_agent,created_at,address_doorcode

const csvRes string = ".csv"
const csvDir string = "./csv/"

func main() {
	_t := time.Now() ////

	db := dbConnect()
	flist := getCsvList(csvDir)
	yafilesList := []Yafile{}

	for _, file := range flist {
		y := YafileMake(file, LOAD_NOT_PARSED)
		yafilesList = append(yafilesList, y)

		go dataInsert(y, db, _t)
	}

	// all files add with status LOAD_NOT_PARSED
	_, err := db.NamedExec(YafileInsertSql, yafilesList)
	if err != nil {
		fmt.Println("FIles log error", err)
	}

	var input string
	fmt.Scanln(&input)
	fmt.Println(time.Since(_t)) ////
}

func dataInsert(file Yafile, db *sqlx.DB, t time.Time) {
	fileOpen, err := os.Open(csvDir + file.Name)
	if err != nil {
		file.SetFileError()
		_, err := db.NamedExec(YafileUpdateStatusSql, file)
		if err != nil {
			fmt.Println("Files (error) log update error", err)
		}

		fmt.Println("File open error |", file.Name, err)
		return
	}
	defer fileOpen.Close()

	reader := csv.NewReader(fileOpen)

	// first line - headers
	_, err = reader.Read()
	if err != nil {
		file.SetFileError()
		_, err := db.NamedExec(YafileUpdateStatusSql, file)
		if err != nil {
			fmt.Println("Files (CSV reading error) log update error", err)
		}

		fmt.Println("CSV reading error (header) |", file.Name, err)
		return
	}

	corr := []Corruption{}
	var cid int64 = 1

	file.SetInProgress()
	_, err = db.NamedExec(YafileUpdateStatusSql, file)
	if err != nil {
		fmt.Println("Files (parsing in progress) log update error", err)
	}

	for {
		line, err := reader.Read()
		if err != nil {
			l := len(line)
			if l == 0 {
				fmt.Printf("File %s parsed and load\n", file.Name)
				break
			}

			fmt.Println("CSV reading error (data string) |", file.Name, l)
			break
		}

		corr = append(corr, MakeCorruption(line))

		cid++

		if cid%batchSize == 0 {
			_, err := db.NamedExec(InsertSql, corr)
			if err != nil {
				fmt.Println("NamedExec error |", err, cid)
				break
			}

			corr = []Corruption{}
		}

	}

	if 0 < len(corr) {
		_, err := db.NamedExec(InsertSql, corr)
		if err != nil {
			fmt.Println("NamedExec error |", err, cid)
		}
	}

	file.SetParsed()
	_, err = db.NamedExec(YafileUpdateStatusSql, file)
	if err != nil {
		fmt.Println("Files (mark file as parsed) log update error", err)
	}

	fmt.Println(time.Since(t))
}

func dbConnect() *sqlx.DB {
	db, err := sqlx.Connect("postgres", "postgresql://smf_gps_user:smfgpspassword@10.20.0.4:5432/smf_gps_db?sslmode=disable")
	if err != nil {
		fmt.Println("Connect db error ", err)
	}

	return db
}

// список .CSV фалов в директории
func getCsvList(directory string) map[int]string {
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		panic(err)
	}

	filesList := make(map[int]string)
	var i int = 0

	for _, file := range files {
		if strings.Contains(file.Name(), csvRes) {
			filesList[i] = file.Name()
			i++
		}
	}

	return filesList
}
