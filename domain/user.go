package domain

type User struct {
	ID   uint   `gorm:"column:id"`
	Name string `gorm:"column:name"`
	// Name string `gorm:"column:name;length:80"`
}
