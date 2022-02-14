package table

//-------------------------------------//
//SQL struct
//-------------------------------------//

type Users struct {
	ID       string `gorm:"primaryKey"`
	Email    string
	Name     string
	Password string
}
