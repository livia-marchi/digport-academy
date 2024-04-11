package gestaodetarefas

type Project struct {
	ID         string
	Name       string  //nome do projeto
	Owner      string  //responsável pelo projeto
	Completion float64 //porcentagem de conclusão do projeto
	Task       []Task
}
