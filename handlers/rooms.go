package handlers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/somuthink/museid_admin/db"
)

func GtRoomsHX(c *fiber.Ctx) error {
	return c.Render("rooms", fiber.Map{"Rooms": db.GtRooms()})
}

func DlRoomHX(c *fiber.Ctx) error {
	roomId, err := strconv.Atoi(c.FormValue("id"))
	if err != nil {
		return err
	}
	db.DlRoom(roomId)

	return nil
}

func UpRoomHX(c *fiber.Ctx) error {
	roomId, err := strconv.Atoi(c.FormValue("id"))
	fmt.Println(roomId)
	if err != nil {
		return err
	}

	db.UpRooms(c.FormValue("roomName"), c.FormValue("roomLink"), roomId)

	return nil
}
