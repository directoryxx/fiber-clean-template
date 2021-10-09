package repository

import (
	"context"

	"github.com/directoryxx/fiber-clean-template/app/infrastructure"
	"github.com/directoryxx/fiber-clean-template/app/rules"
	"github.com/directoryxx/fiber-clean-template/database/gen"
	"github.com/directoryxx/fiber-clean-template/database/gen/Todo"
)

type TodoRepository struct {
	// SQLHandler *gen.Client
	Ctx context.Context
}

func (rr *TodoRepository) Insert(Todo *rules.TodoValidation) (todo *gen.Todo, err error) {
	conn, err := infrastructure.Open()
	if err != nil {
		panic(err)
	}
	create, err := conn.Todo.Create().SetName(Todo.Name).Save(rr.Ctx)
	defer conn.Close()
	return create, err
}

func (rr *TodoRepository) GetAll() (todo []*gen.Todo, err error) {
	conn, err := infrastructure.Open()
	if err != nil {
		panic(err)
	}
	todo, err = conn.Todo.Query().All(rr.Ctx)
	defer conn.Close()
	return todo, err
}

func (rr *TodoRepository) Update(Todo_id int, Todo *rules.TodoValidation) (todo *gen.Todo, err error) {
	conn, err := infrastructure.Open()
	if err != nil {
		panic(err)
	}
	TodoUpdate, errUpdate := conn.Todo.UpdateOneID(Todo_id).SetName(Todo.Name).Save(rr.Ctx)
	defer conn.Close()
	return TodoUpdate, errUpdate
}

func (rr *TodoRepository) Delete(Todo_id int) (err error) {
	conn, err := infrastructure.Open()
	if err != nil {
		panic(err)
	}
	err = conn.Todo.DeleteOneID(Todo_id).Exec(rr.Ctx)
	defer conn.Close()
	return err
}

func (rr *TodoRepository) CountByName(input string) (res int64) {
	conn, err := infrastructure.Open()
	if err != nil {
		panic(err)
	}
	check, _ := conn.Todo.Query().Where(todo.Name(input)).Count(rr.Ctx)
	defer conn.Close()
	return int64(check)
}

func (rr *TodoRepository) FindById(Todo_id int) (TodoData *gen.Todo, err error) {
	conn, err := infrastructure.Open()
	if err != nil {
		panic(err)
	}
	TodoData, errTodoData := conn.Todo.Query().Where(todo.ID(Todo_id)).Only(rr.Ctx)
	defer conn.Close()
	return TodoData, errTodoData
}
