//DigPort Atividade

package main

import (
	"time"
)

type User struct {
	ID       uint
	Username string
	Email    string
}

type Project struct {
	ID          uint
	Name        string
	Description string
	Users       []User
	Tasks       []Task
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Task struct {
	ID          uint
	Title       string
	Description string
	Assignee    User
	DueDate     time.Time
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func main() {

	user1 := User{ID: 1, Username: "user1", Email: "user1@example.com"}
	user2 := User{ID: 2, Username: "user2", Email: "user2@example.com"}

	project := Project{
		ID:          1,
		Name:        "Projeto Exemplo",
		Description: "Um projeto de exemplo para demonstração",
		Users:       []User{user1, user2},
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	task1 := Task{
		ID:          1,
		Title:       "Implementar autenticação",
		Description: "Implementar sistema de autenticação utilizando JWT",
		Assignee:    user1,
		DueDate:     time.Now().AddDate(0, 0, 7),
		Status:      "To Do",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	task2 := Task{
		ID:          2,
		Title:       "Criar API REST",
		Description: "Desenvolver API RESTful para interação com o frontend",
		Assignee:    user2,
		DueDate:     time.Now().AddDate(0, 0, 14),
		Status:      "In Progress",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	project.Tasks = append(project.Tasks, task1, task2)

}
