package db

import (
	_ "github.com/go-sql-driver/mysql"
)

type Room struct {
	Id                               int
	Url                              string
	Name                             string
	Center_x, Center_y, End_x, End_y int
}

func GtRooms() []Room {
	rooms_res := []Room{}
	res, err := DB.Query("SELECT `id`,  `url`, `name`, ST_X(center), ST_Y(center) ,  ST_X(end) , ST_Y(end)  FROM `rooms` ")
	if err != nil {
		panic(err)
	}
	for res.Next() {

		var rm Room

		res.Scan(&rm.Id, &rm.Url, &rm.Name, &rm.Center_x, &rm.Center_y, &rm.End_x, &rm.End_y)

		rooms_res = append(rooms_res, rm)
	}

	return rooms_res
}

func UpRooms(roomName, roomLink string, roomId int) {
	DB.Exec("UPDATE `rooms` SET `url` = ?, `name` = ? WHERE `id` = ? ", roomLink, roomName, roomId)
}

func DlRoom(roomId int) {
	DB.Exec("DELETE FROM rooms WHERE `rooms`.`id` = ?", roomId)
}
