package payload

import uuid "github.com/satori/go.uuid"

type BookRequest struct {
	Title   string  `json:"title" form:"title"`
	Author  string  `json:"author" form:"author"`
	Image   string  `json:"image" form:"image"`
	Price   float32 `json:"price" form:"title"`
	Summary string  `json:"summary" form:"summary" gorm:"type:text"`
	Stock   uint    `json:"stock" form:"stock"`
}
type BookResponse struct {
	ID      uuid.UUID `json:"id" form:"id"`
	Title   string    `json:"title" form:"title"`
	Author  string    `json:"author" form:"author"`
	Image   string    `json:"image" form:"image"`
	Price   float32   `json:"price" form:"title"`
	Summary string    `json:"summary" form:"summary" gorm:"type:text"`
	Stock   uint      `json:"stock" form:"stock"`
	Sold    uint      `json:"sold" form:"sold"`
}
