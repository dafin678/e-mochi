package models

//unused
type Mochi struct {
	Id           int    `json:"id" gorm:"primaryKey"`
	Flavour      string `json:"flavour"`
	Price        string `json:"price"`
	IceCreamType string `json:"icecreamtype"`
}
