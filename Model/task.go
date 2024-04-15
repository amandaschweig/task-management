//DigPort Atividade

package model

type Task struct {
	ID          uint
	Title       string
	Description string
	Assignee    User
	Status      string
}
