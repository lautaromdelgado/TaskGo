package models

type Tasks struct {
	Id          uint   `json:"id" gorm:"primary_key"`
	Title       string `json:"titulo"`
	Descripcion string `json:"descripcion",omitempty`
	Nivel       int    `json:"nivel"`
	Idusuario   int    `json:"idusuario"`
}

func (Tasks) TableName() string {
	return "tasks"
}
