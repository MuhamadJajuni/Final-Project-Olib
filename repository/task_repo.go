package repository

import (
	"database/sql"
	"final-project-olib/model"
	"final-project-olib/model/dto"
	"log"
	"math"
	"time"
)

type taskRepo struct {
	db *sql.DB
}

// DeleteAdmin implements TaskRepo.
func (a *taskRepo) DeleteAdmin(id string) error {
	query := "DELETE FROM admin WHERE id = $1;"
	_, err := a.db.Exec(query, id)
	return err
}

// DeleteBorrower implements TaskRepo.
func (a *taskRepo) DeleteBorrower(id string) error {
	query := "DELETE FROM borrower WHERE id = $1;"
	_, err := a.db.Exec(query, id)
	return err
}

// DeleteTransaction implements TaskRepo.
func (a *taskRepo) DeleteTransaction(id string) error {
	query := "DELETE FROM transaction WHERE id = $1;"
	_, err := a.db.Exec(query, id)
	return err
}

// FindAdminById implements TaskRepo.
func (a *taskRepo) FindAdminById(id string) (model.Admin, error) {
	var task model.Admin
	err := a.db.QueryRow("SELECT id, name, email, created_at, updated_at FROM admin WHERE id=$1", id).Scan(&task.Id, &task.Name, &task.Email, &task.CreatedAt, &task.UpdatedAt)

	if err != nil {
		return model.Admin{}, err
	}
	return task, nil
}
func (a *taskRepo) FindAdminByEmail(email string) (model.Admin, error) {
	var task model.Admin
	err := a.db.QueryRow("SELECT id, name, email, created_at, updated_at FROM admin WHERE email=$1", email).Scan(&task.Id, &task.Name, &task.Email, &task.CreatedAt, &task.UpdatedAt)

	if err != nil {
		return model.Admin{}, err
	}
	return task, nil
}

// FindAllAdmin implements TaskRepo.
func (a *taskRepo) FindAllAdmin(page int, size int) ([]model.Admin, dto.Paging, error) {
	var listData []model.Admin
	var rows *sql.Rows
	var err error

	offset := (page - 1) * size
	rows, err = a.db.Query("SELECT id, name, email, created_at, updated_at FROM admin limit $1 offset $2;", size, offset)
	if err != nil {
		return nil, dto.Paging{}, err
	}
	totalRows := 0
	err = a.db.QueryRow("SELECT COUNT(*) FROM admin").Scan(&totalRows)
	if err != nil {
		return nil, dto.Paging{}, err
	}

	for rows.Next() {
		var task model.Admin
		err := rows.Scan(&task.Id, &task.Name, &task.Email, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			log.Println(err.Error())
		}
		listData = append(listData, task)
	}

	paging := dto.Paging{
		Page:       page,
		Size:       size,
		TotalRows:  totalRows,
		TotalPages: int(math.Ceil(float64(totalRows) / float64(size))),
	}
	return listData, paging, nil
}

// FindAllBorrower implements TaskRepo.
func (a *taskRepo) FindAllBorrower(page int, size int) ([]model.Borrower, dto.Paging, error) {
	var listData []model.Borrower
	var rows *sql.Rows
	var err error

	offset := (page - 1) * size
	rows, err = a.db.Query("SELECT id, name, email, created_at, updated_at, address FROM borrower limit $1 offset $2;", size, offset)
	if err != nil {
		return nil, dto.Paging{}, err
	}
	totalRows := 0
	err = a.db.QueryRow("SELECT COUNT(*) FROM borrower").Scan(&totalRows)
	if err != nil {
		return nil, dto.Paging{}, err
	}

	for rows.Next() {
		var task model.Borrower
		err := rows.Scan(&task.Id, &task.Name, &task.Email, &task.CreatedAt, &task.UpdatedAt, &task.Address)
		if err != nil {
			log.Println(err.Error())
		}
		listData = append(listData, task)
	}

	paging := dto.Paging{
		Page:       page,
		Size:       size,
		TotalRows:  totalRows,
		TotalPages: int(math.Ceil(float64(totalRows) / float64(size))),
	}
	return listData, paging, nil
}

// FindAllTransaction implements TaskRepo.
func (a *taskRepo) FindAllTransaction(page int, size int) ([]model.Transaction, dto.Paging, error) {
	var listData []model.Transaction
	var rows *sql.Rows
	var err error

	offset := (page - 1) * size
	rows, err = a.db.Query("SELECT * FROM transaction limit $1 offset $2;", size, offset)
	if err != nil {
		return nil, dto.Paging{}, err
	}
	totalRows := 0
	err = a.db.QueryRow("SELECT COUNT(*) FROM transaction;").Scan(&totalRows)
	if err != nil {
		return nil, dto.Paging{}, err
	}

	for rows.Next() {
		var task model.Transaction
		err := rows.Scan(&task.Id, &task.Book_id, &task.Borrower_id, &task.Status, &task.CreatedAt, &task.UpdatedAt, &task.Admin_id)
		if err != nil {
			log.Println(err.Error())
		}
		listData = append(listData, task)
	}

	paging := dto.Paging{
		Page:       page,
		Size:       size,
		TotalRows:  totalRows,
		TotalPages: int(math.Ceil(float64(totalRows) / float64(size))),
	}
	return listData, paging, nil
}

// FindBorrowerById implements TaskRepo.
func (a *taskRepo) FindBorrowerById(id string) (model.Borrower, error) {
	var task model.Borrower
	err := a.db.QueryRow("SELECT id, name, email, created_at, updated_at, address FROM borrower WHERE id=$1", id).Scan(&task.Id, &task.Name, &task.Email, &task.CreatedAt, &task.UpdatedAt, &task.Address)

	if err != nil {
		return model.Borrower{}, err
	}
	return task, nil
}
func (a *taskRepo) FindBorrowerByEmail(email string) (model.Borrower, error) {
	var task model.Borrower
	err := a.db.QueryRow("SELECT id, name, email, created_at, updated_at, address FROM borrower WHERE email=$1", email).Scan(&task.Id, &task.Name, &task.Email, &task.CreatedAt, &task.UpdatedAt, &task.Address)

	if err != nil {
		return model.Borrower{}, err
	}
	return task, nil
}

// FindTransactionById implements TaskRepo.
func (a *taskRepo) FindTransactionById(id string) (model.Transaction, error) {
	var task model.Transaction
	err := a.db.QueryRow("SELECT * FROM transaction WHERE id=$1", id).Scan(&task.Id, &task.Book_id, &task.Borrower_id, &task.Status, &task.CreatedAt, &task.UpdatedAt, &task.Admin_id)

	if err != nil {
		return model.Transaction{}, err
	}
	return task, nil
}

// PostNewAdmin implements TaskRepo.
func (a *taskRepo) RegisterAdmin(newTask model.Admin) (model.Admin, error) {

	query := "INSERT INTO admin (name, email, password) VALUES ($1, $2, $3) RETURNING id;"

	var taskId string
	err := a.db.QueryRow(query, newTask.Name, newTask.Email, newTask.Password).Scan(&taskId)

	if err != nil {
		return newTask, err
	}

	newTask.Id = taskId
	return newTask, nil
}

// PostNewBorrower implements TaskRepo.
func (a *taskRepo) RegisterBorrower(newTask model.Borrower) (model.Borrower, error) {

	query := "INSERT INTO borrower (name, email, password, address) VALUES ($1, $2, $3, $4) RETURNING id;"

	var taskId string
	err := a.db.QueryRow(query, newTask.Name, newTask.Email, newTask.Password, newTask.Address).Scan(&taskId)

	if err != nil {
		return newTask, err
	}

	newTask.Id = taskId
	return newTask, nil
}

// PostNewTransaction implements TaskRepo.
func (a *taskRepo) PostNewTransaction(newTask model.Transaction) (model.Transaction, error) {

	query := "INSERT INTO book (book_id, borrower_id, status, admin_id) VALUES ($1, $2, $3) RETURNING id;"

	var taskId string
	err := a.db.QueryRow(query, newTask.Book_id, newTask.Borrower_id, newTask.Status, newTask.Admin_id).Scan(&taskId)

	if err != nil {
		return newTask, err
	}

	newTask.Id = taskId
	return newTask, nil
}

// UpdateAdmin implements TaskRepo.
func (a *taskRepo) UpdateAdmin(id string, admin model.Admin) (model.Admin, error) {
	query := "UPDATE admin SET name = $1, email = $2, password = $3, updated_at = $4 WHERE id = $5 RETURNING id, name, email, created_at, updated_at;"
	err := a.db.QueryRow(query, admin.Name, admin.Email, admin.Password, time.Now(), id).Scan(&admin.Id, &admin.Name, &admin.Email, &admin.CreatedAt, &admin.UpdatedAt)
	if err != nil {
		return admin, err
	}
	return admin, nil
}

// UpdateBorrower implements TaskRepo.
func (a *taskRepo) UpdateBorrower(id string, borrower model.Borrower) (model.Borrower, error) {
	query := "UPDATE borrower SET name = $1, email = $2, password = $3, address = $4, updated_at = $5 WHERE id = $6 RETURNING id, name, email,created_at,updated_at,address;"
	err := a.db.QueryRow(query, borrower.Name, borrower.Email, borrower.Password, borrower.Address, time.Now(), id).Scan(&borrower.Id, &borrower.Name, &borrower.Email, &borrower.CreatedAt, &borrower.UpdatedAt, &borrower.Address)
	if err != nil {
		return borrower, err
	}
	return borrower, nil
}

// UpdateTransaction implements TaskRepo.
func (a *taskRepo) UpdateTransaction(id string, transaction model.Transaction) (model.Transaction, error) {
	query := "UPDATE transaction SET status = $1, updated_at = $2 WHERE id = $3 RETURNING id, book_id, borrower_id,status, created_at, updated_at, admin_id;"
	err := a.db.QueryRow(query, transaction.Status, time.Now(), id).Scan(&transaction.Id, &transaction.Book_id, &transaction.Borrower_id, transaction.CreatedAt, transaction.UpdatedAt, &transaction.Admin_id)
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (a *taskRepo) FindAllBook(page int, size int) ([]model.Book, dto.Paging, error) {
	var listData []model.Book
	var rows *sql.Rows
	var err error

	offset := (page - 1) * size
	rows, err = a.db.Query("SELECT * FROM book limit $1 offset $2;", size, offset)
	if err != nil {
		return nil, dto.Paging{}, err
	}
	totalRows := 0
	err = a.db.QueryRow("SELECT COUNT(*) FROM book").Scan(&totalRows)
	if err != nil {
		return nil, dto.Paging{}, err
	}

	for rows.Next() {
		var task model.Book
		err := rows.Scan(&task.Id, &task.Title, &task.Author, &task.Release_Year, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			log.Println(err.Error())
		}
		listData = append(listData, task)
	}

	paging := dto.Paging{
		Page:       page,
		Size:       size,
		TotalRows:  totalRows,
		TotalPages: int(math.Ceil(float64(totalRows) / float64(size))),
	}
	return listData, paging, nil
}

func (a *taskRepo) FindBookById(id string) (model.Book, error) {
	var task model.Book
	err := a.db.QueryRow("SELECT * FROM book WHERE id=$1", id).Scan(&task.Id, &task.Title, &task.Author, &task.Release_Year, &task.CreatedAt, &task.UpdatedAt)

	if err != nil {
		return model.Book{}, err
	}
	return task, nil
}

func (a *taskRepo) PostNewBook(newTask model.Book) (model.Book, error) {

	query := "INSERT INTO book (title, author, release_year) VALUES ($1, $2, $3) RETURNING id;"

	var taskId string
	err := a.db.QueryRow(query, newTask.Title, newTask.Author, newTask.Release_Year).Scan(&taskId)

	if err != nil {
		return newTask, err
	}

	newTask.Id = taskId
	return newTask, nil
}

func (a *taskRepo) UpdateBook(id string, updatedBook model.Book) (model.Book, error) {
	query := "UPDATE book SET title = $1, author = $2, release_year = $3, updated_at = $4 WHERE id = $5 RETURNING id, title, author, release_year, created_at, updated_at;"
	err := a.db.QueryRow(query, updatedBook.Title, updatedBook.Author, updatedBook.Release_Year, time.Now(), id).Scan(&updatedBook.Id, &updatedBook.Title, &updatedBook.Author, &updatedBook.Release_Year, &updatedBook.CreatedAt, &updatedBook.UpdatedAt)
	if err != nil {
		return updatedBook, err
	}
	return updatedBook, nil
}

func (a *taskRepo) DeleteBook(id string) error {
	query := "DELETE FROM book WHERE id = $1;"
	_, err := a.db.Exec(query, id)
	return err
}

type TaskRepo interface {
	FindAllBook(page int, size int) ([]model.Book, dto.Paging, error)
	FindBookById(id string) (model.Book, error)
	PostNewBook(newTask model.Book) (model.Book, error)
	UpdateBook(id string, book model.Book) (model.Book, error)
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

func NewtaskRepo(database *sql.DB) TaskRepo {
	return &taskRepo{db: database}
}
