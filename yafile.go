package main

import (
	"time"
)

const LOAD_NOT_PARSED int = 1
const LOAD_PARSED int = 2
const LOAD_PARSE_IN_PROGRESS int = 3
const FILE_OPEN_ERROR int = 6

const YafileInsertSql string = "INSERT INTO yafile (name, added, status) VALUES (:name, :added, :status)"
const YafileUpdateStatusSql string = "UPDATE yafile SET status = :status WHERE name = :name"

type Yafile struct {
	//Id          int    `db:"id"`
	Name        string `db:"name"`
	Added       string `db:"added"`
	Description string `db:"description"`
	Status      int    `db:"status"`
}

func YafileMake(name string, status int) Yafile {
	yafile := Yafile{}

	yafile.Name = name
	yafile.Status = status

	yafile.Description = ""
	yafile.Added = time.Now().Format("2006-01-02 15:04:05")

	return yafile
}

func (yafile *Yafile) SetLoad() {
	yafile.Status = LOAD_NOT_PARSED
}

func (yafile *Yafile) SetParsed() {
	yafile.Status = LOAD_PARSED
}

func (yafile *Yafile) SetInProgress() {
	yafile.Status = LOAD_PARSE_IN_PROGRESS
}

func (yafile *Yafile) SetFileError() {
	yafile.Status = FILE_OPEN_ERROR
}
