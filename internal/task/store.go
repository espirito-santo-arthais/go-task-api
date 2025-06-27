package task

import "sync"

// Armazenamento em memória
var (
	tasks  []Task
	nextID = 1
	mu     sync.Mutex
)

// GetAll retorna todas as tarefas armazenadas
func GetAll() []Task {
	mu.Lock()
	defer mu.Unlock()

	// Retorna uma cópia para evitar manipulação externa do slice original
	tasksCopy := make([]Task, len(tasks))
	copy(tasksCopy, tasks)
	return tasksCopy
}

// Create armazena uma nova tarefa com ID gerado automaticamente
func Create(t Task) Task {
	mu.Lock()
	defer mu.Unlock()

	t.ID = nextID
	nextID++
	tasks = append(tasks, t)
	return t
}

// Delete remove uma tarefa pelo ID (se existir)
func Delete(id int) bool {
	mu.Lock()
	defer mu.Unlock()

	for i, t := range tasks {
		if t.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return true
		}
	}
	return false
}
