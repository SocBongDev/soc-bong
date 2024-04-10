package common

import "time"

type BaseEntity struct {
	Id        int
	CreatedAt time.Time
	UpdatedAt time.Time
}
