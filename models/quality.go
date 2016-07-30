package models

import (
	"github.com/jinzhu/gorm"
)

type Quality struct {
	Name string

	gorm.Model
}
