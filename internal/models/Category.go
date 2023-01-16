package internal

import (
	"database/sql"
	"time"
)

type Category struct {
	ID          uint
	Name        string
	ActivatedAt sql.NullTime
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
