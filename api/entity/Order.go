package entity

import "github.com/golang/protobuf/ptypes/timestamp"

type Order struct {
	Id          uint32 `gorm:"primary_key;auto_increment" json:"id"`
	ArrivalDate timestamp.Timestamp `json:"arrival_date"`
	DepartureDate timestamp.Timestamp `json:"departure_date"`
	UserId     uint32 `gorm:"not null;" json:"user_id"`
	User User `json:"user"`
	RoomId      uint32 `gorm:"not null;" json:"room_id"`
	Room  Room `json:"room;auto_preload"`
	Adults      uint `json:"adults"`
	Child       uint `json:"child"`
}
