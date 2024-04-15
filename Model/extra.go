// Esse está uma bagunça pois estou apenas estudando e ainda não está finalizado

package model

import "time"

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
		Title:       "Criar tanana tanana",
		Description: "Desenvolver tanana tanana",
		Assignee:    user2,
		DueDate:     time.Now().AddDate(0, 0, 14),
		Status:      "In Progress",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	project.Tasks = append(project.Tasks, task1, task2)

}
