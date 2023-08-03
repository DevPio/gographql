package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Course struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
	CategoryId  string
}

func NewCourse(db *sql.DB) *Course {
	return &Course{db: db}
}

func (c *Course) CreateCourse(name string, description string, categoryId string) (*Course, error) {
	id := uuid.New().String()

	err := c.db.QueryRow("INSERT INTO courses (id, name, description,categoryId) VALUES ($1,$2,$3,$4) RETURNING id;", id, name, description, categoryId).Scan(&c.ID)

	if err != nil {
		return &Course{}, err
	}

	return c, nil
}

func (c *Course) FindAll() ([]Course, error) {

	result, err := c.db.Query("SELECT * FROM courses;")

	if err != nil {
		return nil, err
	}

	courses := []Course{}
	for result.Next() {
		var id, name, description, category_id string

		result.Scan(&id, &name, &description, &category_id)

		courses = append(courses, Course{ID: id, Name: name, Description: description, CategoryId: category_id})
	}

	return courses, nil
}

func (c *Course) FindByCategoryId(id string) ([]Course, error) {

	result, err := c.db.Query("SELECT * FROM courses WHERE CategoryId = ?;", id)

	if err != nil {
		return nil, err
	}

	defer result.Close()

	courses := []Course{}

	for result.Next() {
		var id, name, description, category_id string
		result.Scan(id, name, description, category_id)
		courses = append(courses, Course{
			ID:          id,
			Name:        name,
			Description: description,
			CategoryId:  category_id,
		})
	}

	return courses, nil
}
