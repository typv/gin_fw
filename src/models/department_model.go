package models

type Department struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"column:name"`
	Users []User `gorm:"foreignKey:DepartmentID"`
}

func (Department) TableName() string {
	return "departments"
}
