package baseEntity

import "gorm.io/gorm"

type BaseEntity struct {
	gorm.Model
	id int
}
