package gestaodetarefas

type Task struct {
	ID        string
	Name      string   //nome da tarefa
	Status    string   //se a tarefa foi concluida, está em andamento ou “a fazer”
	Priority  string   //alta, média ou baixa
	DueDate   string   //data de entrega da tarefa
	Summary   string   //descrição da tarefa
	BlockedBy string   //tarefa que bloqueia a execução de uma outra tarefa
	Assignee  Assignee //usuário responsável pela execução da tarefa
	Project   Project  //projeto ao qual essa tarefa pertence
}

type Status string

const (
	High   Status = "High"
	Medium Status = "Medium"
)
