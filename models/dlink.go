package models

import (
	"github.com/jinzhu/gorm"
)

type Dlink struct {
	Url  string
	Type string

	gorm.Model
}
