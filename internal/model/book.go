package model

type Book struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Author      string `json:"author"`
}

type BookInput struct {
	Name        string `json:"name" validate:"required,gte=2"`
	Description string `json:"description" validate:"required,gte=6"`
	Author      string `json:"author" validate:"required,gte=2"`
}

func (b BookInput) Validate() error {
	return validate.Struct(b)
}
