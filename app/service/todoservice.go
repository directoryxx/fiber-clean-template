package service

import (
	"github.com/directoryxx/fiber-clean-template/app/repository"
	"github.com/directoryxx/fiber-clean-template/app/rules"
	"github.com/directoryxx/fiber-clean-template/database/gen"
)

type TodoService struct {
	TodoRepository repository.TodoRepository
}

func (rs TodoService) GetAll() (Todos []*gen.Todo) {
	TodoData, _ := rs.TodoRepository.GetAll()

	return TodoData
}

func (rs TodoService) CreateTodo(Todo *rules.TodoValidation) (user *gen.Todo, err error) {
	data, err := rs.TodoRepository.Insert(Todo)

	return data, err
}

func (rs TodoService) UpdateTodo(Todo_id int, Todo *rules.TodoValidation) (user *gen.Todo, err error) {
	data, err := rs.TodoRepository.Update(Todo_id, Todo)

	return data, err
}

func (rs TodoService) CheckDuplicateNameTodo(name string) int64 {
	data := rs.TodoRepository.CountByName(name)

	return data
}

func (rs TodoService) GetById(Todo_id int) (user *gen.Todo) {
	TodoData, _ := rs.TodoRepository.FindById(Todo_id)

	return TodoData
}

func (rs TodoService) DeleteTodo(Todo_id int) error {
	deleteTodo := rs.TodoRepository.Delete(Todo_id)

	return deleteTodo
}
