//DigPort Atividade

package model

type Project struct {
	ID          uint
	Name        string
	Description string
	status      string
	Duedate     string
	Category    string
}

type Task struct {
	ID          uint
	Title       string
	Description string
	Assignee    User
	Status      string
}
