package main

import (
	"fmt"
	"strconv"
)

const InsertSql string = `INSERT INTO corruption
(ya_id,
first_name,
full_name,
email,
phone_number,
address_city,
address_street,
address_house,
address_entrance,
address_floor,
address_office,
address_comment,
address_doorcode,
location,
location_latitude,
location_longitude,
amount_charged,
user_id,
user_agent,
created_at) VALUES(
:ya_id,
:first_name,
:full_name,
:email,
:phone_number,
:address_city,
:address_street,
:address_house,
:address_entrance,
:address_floor,
:address_office,
:address_comment,
:address_doorcode,
:location,
:location_latitude,
:location_longitude,
:amount_charged,
:user_id,
:user_agent,
:created_at
)`

type Corruption struct {
	YaId int64 `db:"ya_id"`
	FirstName string `db:"first_name"`
	FullName string `db:"full_name"`
	Email string `db:"email"`
	PhoneNumber int64 `db:"phone_number"`
	City string `db:"address_city"`
	Street string `db:"address_street"`
	House string `db:"address_house"`
	Entrance string `db:"address_entrance"`
	Floor string `db:"address_floor"`
	Office string `db:"address_office"`
	AddressComment string `db:"address_comment"`
	Doorcode string `db:"address_doorcode"`
	Location string `db:"location"`
	Latitude string `db:"location_latitude"`
	Longitude string `db:"location_longitude"`
	AmountCharged string `db:"amount_charged"`
	UserId int64 `db:"user_id"`
	UserAgent string `db:"user_agent"`
	CreatedAt string `db:"created_at"`
}

func MakeCorruption(line []string) Corruption {
	return Corruption{
		strToInt(line[0]),
		line[1],
		line[2],
		line[3],
		strToInt(line[4]),
		line[5],
		line[6],
		line[7],
		line[8],
		line[9],
		line[10],
		line[11],
		line[18],
		("(" + line[12] + "," + line[13] + ")"),
		line[12],
		line[13],
		line[14],
		strToInt(line[15]),
		line[16],
		line[17]}
}

func strToInt(str string) int64 {
	intFromStr, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Convert values to int error |", err)
		return 0
	}

	return int64(intFromStr)
}
