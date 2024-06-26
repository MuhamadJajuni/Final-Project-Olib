package usecase

import (
	"final-project-olib/model"
	"final-project-olib/model/dto"
	"final-project-olib/repository"
)

type taskUseCase struct {
	repo repository.TaskRepo
}

// DeleteAdmin implements TaskUseCase.
func (a *taskUseCase) DeleteAdmin(id string) error {
	return a.repo.DeleteAdmin(id)
}

// DeleteBorrower implements TaskUseCase.
func (a *taskUseCase) DeleteBorrower(id string) error {
	return a.repo.DeleteBorrower(id)
}

// DeleteTransaction implements TaskUseCase.
func (a *taskUseCase) DeleteTransaction(id string) error {
	return a.repo.DeleteTransaction(id)
}

// FindAdminById implements TaskUseCase.
func (a *taskUseCase) FindAdminById(id string) (model.Admin, error) {
	return a.repo.FindAdminById(id)
}
func (a *taskUseCase) FindAdminByEmail(email string) (model.Admin, error) {
	return a.repo.FindAdminById(email)
}

// FindAllAdmin implements TaskUseCase.
func (a *taskUseCase) FindAllAdmin(page int, size int) ([]model.Admin, dto.Paging, error) {
	return a.repo.FindAllAdmin(page, size)
}

// FindAllBorrower implements TaskUseCase.
func (a *taskUseCase) FindAllBorrower(page int, size int) ([]model.Borrower, dto.Paging, error) {
	return a.repo.FindAllBorrower(page, size)
}

// FindAllTransaction implements TaskUseCase.
func (a *taskUseCase) FindAllTransaction(page int, size int) ([]model.Transaction, dto.Paging, error) {
	return a.repo.FindAllTransaction(page, size)
}

// FindBorrowerById implements TaskUseCase.
func (a *taskUseCase) FindBorrowerById(id string) (model.Borrower, error) {
	return a.repo.FindBorrowerById(id)
}
func (a *taskUseCase) FindBorrowerByEmail(email string) (model.Borrower, error) {
	return a.repo.FindBorrowerById(email)
}

// FindTransactionById implements TaskUseCase.
func (a *taskUseCase) FindTransactionById(id string) (model.Transaction, error) {
	return a.repo.FindTransactionById(id)
}

// PostNewAdmin implements TaskUseCase.
func (a *taskUseCase) RegisterAdmin(newTask model.Admin) (model.Admin, error) {
	return a.repo.RegisterAdmin(newTask)
}

// PostNewBorrower implements TaskUseCase.
func (a *taskUseCase) RegisterBorrower(newTask model.Borrower) (model.Borrower, error) {
	return a.repo.RegisterBorrower(newTask)
}

// PostNewTransaction implements TaskUseCase.
func (a *taskUseCase) PostNewTransaction(newTask model.Transaction) (model.Transaction, error) {
	return a.repo.PostNewTransaction(newTask)
}

// UpdateAdmin implements TaskUseCase.
func (a *taskUseCase) UpdateAdmin(id string, admin model.Admin) (model.Admin, error) {
	return a.repo.UpdateAdmin(id, admin)
}

// UpdateBorrower implements TaskUseCase.
func (a *taskUseCase) UpdateBorrower(id string, borrower model.Borrower) (model.Borrower, error) {
	return a.repo.UpdateBorrower(id, borrower)
}

// UpdateTransaction implements TaskUseCase.
func (a *taskUseCase) UpdateTransaction(id string, transaction model.Transaction) (model.Transaction, error) {
	return a.repo.UpdateTransaction(id, transaction)
}

// FindAll implements taskUseCase.
func (a *taskUseCase) FindAllBook(page int, size int) ([]model.Book, dto.Paging, error) {
	return a.repo.FindAllBook(page, size)
}

// FindById implements taskUseCase.
func (a *taskUseCase) FindBookById(id string) (model.Book, error) {
	return a.repo.FindBookById(id)
}

func (a *taskUseCase) PostNewBook(newTask model.Book) (model.Book, error) {
	return a.repo.PostNewBook(newTask)
}

// UpdateBook implements taskUseCase.
func (a *taskUseCase) UpdateBook(id string, updatedBook model.Book) (model.Book, error) {
	return a.repo.UpdateBook(id, updatedBook)
}

// DeleteBook implements taskUseCase.
func (a *taskUseCase) DeleteBook(id string) error {
	return a.repo.DeleteBook(id)
}

type TaskUseCase interface {
	FindAllBook(page int, size int) ([]model.Book, dto.Paging, error)
	FindBookById(id string) (model.Book, error)
	PostNewBook(newTask model.Book) (model.Book, error)
	UpdateBook(id string, updatedBook model.Book) (model.Book, error)
	DeleteBook(id string) error

	FindAllBorrower(page int, size int) ([]model.Borrower, dto.Paging, error)
	FindBorrowerById(id string) (model.Borrower, error)
	FindBorrowerByEmail(email string) (model.Borrower, error)
	RegisterBorrower(newTask model.Borrower) (model.Borrower, error)
	UpdateBorrower(id string, borrower model.Borrower) (model.Borrower, error)
	DeleteBorrower(id string) error

	FindAllAdmin(page int, size int) ([]model.Admin, dto.Paging, error)
	FindAdminById(id string) (model.Admin, error)
	FindAdminByEmail(email string) (model.Admin, error)
	RegisterAdmin(newTask model.Admin) (model.Admin, error)
	UpdateAdmin(id string, admin model.Admin) (model.Admin, error)
	DeleteAdmin(id string) error

	FindAllTransaction(page int, size int) ([]model.Transaction, dto.Paging, error)
	FindTransactionById(id string) (model.Transaction, error)
	PostNewTransaction(newTask model.Transaction) (model.Transaction, error)
	UpdateTransaction(id string, transaction model.Transaction) (model.Transaction, error)
	DeleteTransaction(id string) error
}

func NewTaskUseCase(repo repository.TaskRepo) TaskUseCase {
	return &taskUseCase{repo: repo}
}
