package models

import "gorm.io/datatypes"

type Task struct {
	ID     uint           `json:"id"`
	UserID int            `json:"userid"`
	Name   string         `json:"name"`
	Desc   string         `json:"desc"`
	Status bool           `json:"status"`
	Tag    datatypes.JSON `json:"tag"`
}
