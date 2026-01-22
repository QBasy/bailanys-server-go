package domain

import "time"

type Channel struct {
	ID        string
	Name      string
	CreatedBy string
	CreatedAt time.Time
}
