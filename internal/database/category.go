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

func (c *Category) FindById(id string) Category {
	result := c.db.QueryRow("SELECT * FROM categories WHERE id = ?;", id)

	var currentId, name, description string
	result.Scan(&currentId, &name, &description)

	var category = Category{
		ID:          id,
		Name:        name,
		Description: description,
	}

	return category
}

func (c *Category) FindByCourseId(course Course) ([]Category, error) {
	result, err := c.db.Query("SELECT * FROM courses WHERE categoryId = ?;", c.ID)

	if err != nil {

		return nil, err
	}

	categories := []Category{}

	for result.Next() {
		var id, name, description string

		result.Scan(&id, &name, &description)

		categories = append(categories, Category{
			ID:          id,
			Name:        name,
			Description: description,
		})
	}

	return categories, nil
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
