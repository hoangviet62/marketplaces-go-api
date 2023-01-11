type User struct {
	ID          uint
	Name        string
	Email       *string
	Age         uint8
	Birthday    *time.Time
	ActivatedAt sql.NullTime
	CreatedAt   time.Time
	UpdatedAt   time.Time
}