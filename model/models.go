package model

type FinancialQuotes struct {
	ID    int    `gorm:"primary_key;auto_increment" json:"id"`
	Quote string `gorm:"not null; unique" json:"quote"`
}

type LifeQuotes struct {
	ID    int    `gorm:"primary_key;auto_increment" json:"id"`
	Quote string `gorm:"not null; unique" json:"quote"`
}
type LoveQuotes struct {
	ID    int    `gorm:"primary_key;auto_increment" json:"id"`
	Quote string `gorm:"not null; unique" json:"quote"`
}
