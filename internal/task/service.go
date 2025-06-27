package task

import "errors"

// Lista todas as tarefas
func ListTasks() []Task {
    return GetAll()
}

// Cria uma nova tarefa, validando entrada
func AddTask(t Task) (Task, error) {
    if t.Text == "" {
        return Task{}, errors.New("texto da tarefa é obrigatório")
    }
    return Create(t), nil
}

// Remove uma tarefa por ID (se houver)
func RemoveTask(id int) error {
	ok := Delete(id)
	if !ok {
		return errors.New("tarefa não encontrada")
	}
	return nil
}

