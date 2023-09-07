package main


type User struct {
	ID     uint   `gorm:"primary_key" json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email" gorm:"unique"`
	Passwd string `json:"-"`
	Leads  []Lead `gorm:"foreignkey:CreatedByID"` 
}

type Lead struct {
    ID        uint   `gorm:"primary_key" json:"id"`
    Name      string `json:"name"`
    Email     string `json:"email" gorm:"unique"`
    Phone     string `json:"phone"`
    Notes     string `json:"notes"`
    CreatedByID uint 
    CreatedBy   User `gorm:"foreignkey:CreatedByID"` 
}
