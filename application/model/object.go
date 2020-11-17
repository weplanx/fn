package model

type Object struct {
	ID    uint64
	Key   string `gorm:"size:200;unique;not null"`
	Value []byte `gorm:"not null"`
}
