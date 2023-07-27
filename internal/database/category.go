package database

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type Category struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
}

func (category Category) Create(name string, description string) (Category, error) {
	id := uuid.New().String()

	fmt.Println(name, description)
	_, erro := category.db.Exec("INSERT INTO categories (id,name,description) VALUES ($1, $2, $3)", id, name, description)

	if erro != nil {
		return Category{}, erro
	}

	return Category{ID: id, Name: name, Description: description}, nil

}

func NewCategory(db *sql.DB) *Category {

	return &Category{db: db}
}

func (c *Category) FindAll() ([]Category, error) {

	result, err := c.db.Query("SELECT * FROM categories;")

	if err != nil {
		return nil, err
	}

	defer result.Close()
	categories := []Category{}

	for result.Next() {
		var id, name, description string

		result.Scan(&id, &name, &description)

		categories = append(categories, Category{ID: id, Name: name, Description: description})
	}

	return categories, nil
}
