package mocking

import (
	"final-project-olib/model"
	"final-project-olib/model/dto"

	"github.com/stretchr/testify/mock"
)

type TaskRepoMock struct {
	mock.Mock
}

func (a *TaskRepoMock) FindAllBook(page int, size int) ([]model.Book, dto.Paging, error) {
	args := a.Called(page, size)
	return args.Get(0).([]model.Book), args.Get(1).(dto.Paging), args.Error(2)
}

func (a *TaskRepoMock) FindBookById(id string) (model.Book, error) {
	args := a.Called(id)
	return args.Get(0).(model.Book), args.Error(1)
}

func (a *TaskRepoMock) PostNewBook(newTask model.Book) (model.Book, error) {
	args := a.Called(newTask)
	return args.Get(0).(model.Book), args.Error(1)
}

func (a *TaskRepoMock) UpdateBook(id string, book model.Book) (model.Book, error) {
	args := a.Called(id, book)
	return args.Get(0).(model.Book), args.Error(1)
}

func (a *TaskRepoMock) DeleteBook(id string) error {
	args := a.Called(id)
	return args.Error(0)
}

func (a *TaskRepoMock) FindAllBorrower(page int, size int) ([]model.Borrower, dto.Paging, error) {
	args := a.Called(page, size)
	return args.Get(0).([]model.Borrower), args.Get(1).(dto.Paging), args.Error(2)
}

func (a *TaskRepoMock) FindBorrowerById(id string) (model.Borrower, error) {
	args := a.Called(id)
	return args.Get(0).(model.Borrower), args.Error(1)
}
func (a *TaskRepoMock) FindBorrowerByEmail(email string) (model.Borrower, error) {
	args := a.Called(email)
	return args.Get(0).(model.Borrower), args.Error(1)
}
func (a *TaskRepoMock) RegisterBorrower(newTask model.Borrower) (model.Borrower, error) {
	args := a.Called(newTask)
	return args.Get(0).(model.Borrower), args.Error(1)
}

func (a *TaskRepoMock) UpdateBorrower(id string, borrower model.Borrower) (model.Borrower, error) {
	args := a.Called(id, borrower)
	return args.Get(0).(model.Borrower), args.Error(1)
}

func (a *TaskRepoMock) DeleteBorrower(id string) error {
	args := a.Called(id)
	return args.Error(0)
}

func (a *TaskRepoMock) FindAllAdmin(page int, size int) ([]model.Admin, dto.Paging, error) {
	args := a.Called(page, size)
	return args.Get(0).([]model.Admin), args.Get(1).(dto.Paging), args.Error(2)
}

func (a *TaskRepoMock) FindAdminById(id string) (model.Admin, error) {
	args := a.Called(id)
	return args.Get(0).(model.Admin), args.Error(1)
}
func (a *TaskRepoMock) FindAdminByEmail(email string) (model.Admin, error) {
	args := a.Called(email)
	return args.Get(0).(model.Admin), args.Error(1)
}
func (a *TaskRepoMock) RegisterAdmin(newTask model.Admin) (model.Admin, error) {
	args := a.Called(newTask)
	return args.Get(0).(model.Admin), args.Error(1)
}

func (a *TaskRepoMock) UpdateAdmin(id string, admin model.Admin) (model.Admin, error) {
	args := a.Called(id, admin)
	return args.Get(0).(model.Admin), args.Error(1)
}

func (a *TaskRepoMock) DeleteAdmin(id string) error {
	args := a.Called(id)
	return args.Error(0)
}

func (a *TaskRepoMock) FindAllTransaction(page int, size int) ([]model.Transaction, dto.Paging, error) {
	args := a.Called(page, size)
	return args.Get(0).([]model.Transaction), args.Get(1).(dto.Paging), args.Error(2)
}

func (a *TaskRepoMock) FindTransactionById(id string) (model.Transaction, error) {
	args := a.Called(id)
	return args.Get(0).(model.Transaction), args.Error(1)
}

func (a *TaskRepoMock) PostNewTransaction(newTask model.Transaction) (model.Transaction, error) {
	args := a.Called(newTask)
	return args.Get(0).(model.Transaction), args.Error(1)
}

func (a *TaskRepoMock) UpdateTransaction(id string, transaction model.Transaction) (model.Transaction, error) {
	args := a.Called(id, transaction)
	return args.Get(0).(model.Transaction), args.Error(1)
}

func (a *TaskRepoMock) DeleteTransaction(id string) error {
	args := a.Called(id)
	return args.Error(0)
}
