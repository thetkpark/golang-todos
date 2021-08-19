package data

import (
	"github.com/hashicorp/go-hclog"
	"github.com/thetkpark/golang-todo/models"
	"gorm.io/gorm"
)

type TodoRepository interface {
	FindAll(userId uint) ([]*models.Todo, error)
	Create(title string, userId uint) (*models.Todo, error)
	FinishTodo(todoId uint, userId uint) (*models.Todo, error)
	FindByID(todoId uint, userId uint) (*models.Todo, error)
	Delete(todoId uint, userId uint) (*models.Todo, error)
}

type GormTodoRepository struct {
	db  *gorm.DB
	log hclog.Logger
}

func NewGormTodoRepository(db *gorm.DB, log hclog.Logger) *GormTodoRepository {
	return &GormTodoRepository{
		db:  db,
		log: log,
	}
}

func (repo *GormTodoRepository) FindAll(userId uint) ([]*models.Todo, error) {
	var todos []*models.Todo
	tx := repo.db.Where(&models.Todo{UserId: userId}).Find(&todos)
	if tx.Error != nil {
		repo.log.Error("unable to fetch all todos", tx.Error.Error())
		return nil, tx.Error
	}
	return todos, nil
}

func (repo *GormTodoRepository) Create(title string, userId uint) (*models.Todo, error) {
	todo := &models.Todo{
		Title:      title,
		IsFinished: false,
		UserId:     userId,
	}

	tx := repo.db.Create(todo)
	if tx.Error != nil {
		repo.log.Error("cannot save todo to db", tx.Error.Error())
		return nil, tx.Error
	}
	return todo, nil
}

func (repo *GormTodoRepository) FinishTodo(todoId uint, userId uint) (*models.Todo, error) {
	todo, err := repo.FindByID(todoId, userId)
	if err != nil {
		return nil, err
	}

	todo.IsFinished = true
	tx := repo.db.Save(todo)
	if tx.Error != nil {
		repo.log.Error("cannot save updated todo", tx.Error.Error())
		return nil, tx.Error
	}

	return todo, nil
}

func (repo *GormTodoRepository) FindByID(todoId uint, userId uint) (*models.Todo, error) {
	var todo models.Todo
	tx := repo.db.Where(&models.Todo{ID: todoId, UserId: userId}).First(&todo)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &todo, nil
}

func (repo *GormTodoRepository) Delete(todoId uint, userId uint) (*models.Todo, error) {
	todo, err := repo.FindByID(todoId, userId)
	if err != nil {
		return nil, err
	}

	tx := repo.db.Delete(todo)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return todo, nil
}
