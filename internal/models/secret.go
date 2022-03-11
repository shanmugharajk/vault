package models

type Secret struct {
	Key   string `gorm:"primarykey"`
	Value string
}
