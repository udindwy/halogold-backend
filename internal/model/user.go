package model

// User merepresentasikan tabel users di database.
type User struct {
	ID           uint          `gorm:"primaryKey;autoIncrement" json:"id"`
	Nama         string        `gorm:"type:varchar(100);not null" json:"nama"`
	Email        string        `gorm:"type:varchar(100);unique;not null" json:"email"`
	Transactions []Transaction `gorm:"foreignKey:UserID" json:"transactions,omitempty"`
}
