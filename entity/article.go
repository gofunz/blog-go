package entity

type User struct {
	FirstName string `json:"firstname" binding:"required"`
	LastName  string `json:"lastname" binding:"required"`
	Age       uint8  `json:"age" binding:"gte=1,lte=130"`
	Email     string `json:"email" binding:"required,email"`
}

type Article struct {
	Title       string `json:"title" binding:"min=2,max=20" validate:"is-cool"`
	Description string `json:"description" binding:"max=50"`
	Content     string `json:"content" binding:"required,max=1000"`
	Author      User   `json:"author" binding:"required"`
}
