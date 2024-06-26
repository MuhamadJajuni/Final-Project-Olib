package repository

import (
	"database/sql"
	"final-project-olib/model"
	"final-project-olib/model/dto"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func setupTestDB(t *testing.T) (*sql.DB, sqlmock.Sqlmock, func()) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	return db, mock, func() {
		db.Close()
	}
}

func TestFindAdminById(t *testing.T) {
	db, mock, teardown := setupTestDB(t)
	defer teardown()

	repo := NewtaskRepo(db)

	expectedAdmin := model.Admin{
		Id:        "1",
		Name:      "John Doe",
		Email:     "johndoe@example.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	rows := sqlmock.NewRows([]string{"id", "name", "email", "created_at", "updated_at"}).
		AddRow(expectedAdmin.Id, expectedAdmin.Name, expectedAdmin.Email, expectedAdmin.CreatedAt, expectedAdmin.UpdatedAt)

	mock.ExpectQuery("SELECT (.+) FROM admin WHERE id=\\$1").
		WithArgs(expectedAdmin.Id).
		WillReturnRows(rows)

	admin, err := repo.FindAdminById(expectedAdmin.Id)

	assert.NoError(t, err)
	assert.Equal(t, expectedAdmin, admin)
}

func TestFindAdminByEmail(t *testing.T) {
	db, mock, teardown := setupTestDB(t)
	defer teardown()

	repo := NewtaskRepo(db)

	expectedAdmin := model.Admin{
		Id:        "1",
		Name:      "John Doe",
		Email:     "johndoe@example.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	rows := sqlmock.NewRows([]string{"id", "name", "email", "created_at", "updated_at"}).
		AddRow(expectedAdmin.Id, expectedAdmin.Name, expectedAdmin.Email, expectedAdmin.CreatedAt, expectedAdmin.UpdatedAt)

	mock.ExpectQuery("SELECT (.+) FROM admin WHERE email=\\$1").
		WithArgs(expectedAdmin.Email).
		WillReturnRows(rows)

	admin, err := repo.FindAdminByEmail(expectedAdmin.Email)

	assert.NoError(t, err)
	assert.Equal(t, expectedAdmin, admin)
}

func TestFindBorrowerById(t *testing.T) {
	db, mock, teardown := setupTestDB(t)
	defer teardown()

	repo := NewtaskRepo(db)

	expectedBorrower := model.Borrower{
		Id:        "1",
		Name:      "John Doe",
		Email:     "johndoe@example.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Address:   "LA",
	}

	rows := sqlmock.NewRows([]string{"id", "name", "email", "created_at", "updated_at", "address"}).
		AddRow(expectedBorrower.Id, expectedBorrower.Name, expectedBorrower.Email, expectedBorrower.CreatedAt, expectedBorrower.UpdatedAt, expectedBorrower.Address)

	mock.ExpectQuery("SELECT (.+) FROM borrower WHERE id=\\$1").
		WithArgs(expectedBorrower.Id).
		WillReturnRows(rows)

	admin, err := repo.FindBorrowerById(expectedBorrower.Id)

	assert.NoError(t, err)
	assert.Equal(t, expectedBorrower, admin)
}

func TestFindBorrowerByEmail(t *testing.T) {
	db, mock, teardown := setupTestDB(t)
	defer teardown()

	repo := NewtaskRepo(db)

	expectedBorrower := model.Borrower{
		Id:        "1",
		Name:      "John Doe",
		Email:     "johndoe@example.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Address:   "LA",
	}

	rows := sqlmock.NewRows([]string{"id", "name", "email", "created_at", "updated_at", "address"}).
		AddRow(expectedBorrower.Id, expectedBorrower.Name, expectedBorrower.Email, expectedBorrower.CreatedAt, expectedBorrower.UpdatedAt, expectedBorrower.Address)

	mock.ExpectQuery("SELECT (.+) FROM borrower WHERE email=\\$1").
		WithArgs(expectedBorrower.Email).
		WillReturnRows(rows)

	admin, err := repo.FindBorrowerByEmail(expectedBorrower.Email)

	assert.NoError(t, err)
	assert.Equal(t, expectedBorrower, admin)
}

func TestFindTransactionById(t *testing.T) {
	db, mock, teardown := setupTestDB(t)
	defer teardown()

	repo := NewtaskRepo(db)

	expectedTransaction := model.Transaction{
		Id:          "1",
		Book_id:     "1",
		Borrower_id: "1",
		Status:      "Delivered",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Admin_id:    "1",
	}

	rows := sqlmock.NewRows([]string{"id", "book_id", "borrower_id", "status", "created_at", "updated_at", "admin_id"}).
		AddRow(expectedTransaction.Id, expectedTransaction.Book_id, expectedTransaction.Borrower_id, expectedTransaction.Status, expectedTransaction.CreatedAt, expectedTransaction.UpdatedAt, expectedTransaction.Admin_id)

	mock.ExpectQuery("SELECT (.+) FROM transaction WHERE id=\\$1").
		WithArgs(expectedTransaction.Id).
		WillReturnRows(rows)

	admin, err := repo.FindTransactionById(expectedTransaction.Id)

	assert.NoError(t, err)
	assert.Equal(t, expectedTransaction, admin)
}

func TestFindBookById(t *testing.T) {
	db, mock, teardown := setupTestDB(t)
	defer teardown()

	repo := NewtaskRepo(db)

	expectedBook := model.Book{
		Id:           "1",
		Title:        "John Doe",
		Author:       "johndoe",
		Release_Year: 2014,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	rows := sqlmock.NewRows([]string{"id", "title", "author", "release_year", "created_at", "updated_at"}).
		AddRow(expectedBook.Id, expectedBook.Title, expectedBook.Author, expectedBook.Release_Year, expectedBook.CreatedAt, expectedBook.UpdatedAt)

	mock.ExpectQuery("SELECT (.+) FROM book WHERE id=\\$1").
		WithArgs(expectedBook.Id).
		WillReturnRows(rows)

	admin, err := repo.FindBookById(expectedBook.Id)

	assert.NoError(t, err)
	assert.Equal(t, expectedBook, admin)
}

func TestDeleteAdmin(t *testing.T) {
	db, mock, teardown := setupTestDB(t)
	defer teardown()

	repo := NewtaskRepo(db)

	adminId := "1"
	mock.ExpectExec("DELETE FROM admin WHERE id=\\$1").
		WithArgs(adminId).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.DeleteAdmin(adminId)

	assert.NoError(t, err)
}

func TestDeleteBorrower(t *testing.T) {
	db, mock, teardown := setupTestDB(t)
	defer teardown()

	repo := NewtaskRepo(db)

	adminId := "1"
	mock.ExpectExec("DELETE FROM borrower WHERE id=\\$1").
		WithArgs(adminId).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.DeleteBorrower(adminId)

	assert.NoError(t, err)
}

func TestDeleteTransaction(t *testing.T) {
	db, mock, teardown := setupTestDB(t)
	defer teardown()

	repo := NewtaskRepo(db)

	adminId := "1"
	mock.ExpectExec("DELETE FROM transaction WHERE id=\\$1").
		WithArgs(adminId).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.DeleteTransaction(adminId)

	assert.NoError(t, err)
}

func TestDeleteBook(t *testing.T) {
	db, mock, teardown := setupTestDB(t)
	defer teardown()

	repo := NewtaskRepo(db)

	adminId := "1"
	mock.ExpectExec("DELETE FROM Book WHERE id=\\$1").
		WithArgs(adminId).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.DeleteBook(adminId)

	assert.NoError(t, err)
}
func TestFindAllAdmin(t *testing.T) {
	db, mock, teardown := setupTestDB(t)
	defer teardown()

	repo := NewtaskRepo(db)

	expectedAdmins := []model.Admin{
		{
			Id:        "1",
			Name:      "John Doe",
			Email:     "johndoe@example.com",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Id:        "2",
			Name:      "Jane Doe",
			Email:     "janedoe@example.com",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	rows := sqlmock.NewRows([]string{"id", "name", "email", "created_at", "updated_at"})
	for _, admin := range expectedAdmins {
		rows.AddRow(admin.Id, admin.Name, admin.Email, admin.CreatedAt, admin.UpdatedAt)
	}

	mock.ExpectQuery("SELECT (.+) FROM admin limit \\$1 offset \\$2").
		WithArgs(10, 0).
		WillReturnRows(rows)

	mock.ExpectQuery("SELECT COUNT\\(\\*\\) FROM admin").
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(len(expectedAdmins)))

	admins, paging, err := repo.FindAllAdmin(1, 10)

	assert.NoError(t, err)
	assert.Equal(t, expectedAdmins, admins)
	assert.Equal(t, dto.Paging{Page: 1, Size: 10, TotalRows: len(expectedAdmins), TotalPages: 1}, paging)
}

func TestFindAllBorrower(t *testing.T) {
	db, mock, teardown := setupTestDB(t)
	defer teardown()

	repo := NewtaskRepo(db)

	expectedAdmins := []model.Borrower{
		{
			Id:        "1",
			Name:      "John Doe",
			Email:     "johndoe@example.com",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Address:   "LA",
		},
		{
			Id:        "2",
			Name:      "Jane Doe",
			Email:     "janedoe@example.com",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Address:   "LA",
		},
	}

	rows := sqlmock.NewRows([]string{"id", "name", "email", "created_at", "updated_at", "address"})
	for _, admin := range expectedAdmins {
		rows.AddRow(admin.Id, admin.Name, admin.Email, admin.CreatedAt, admin.UpdatedAt, admin.Address)
	}

	mock.ExpectQuery("SELECT (.+) FROM borrower limit \\$1 offset \\$2").
		WithArgs(10, 0).
		WillReturnRows(rows)

	mock.ExpectQuery("SELECT COUNT\\(\\*\\) FROM borrower").
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(len(expectedAdmins)))

	admins, paging, err := repo.FindAllBorrower(1, 10)

	assert.NoError(t, err)
	assert.Equal(t, expectedAdmins, admins)
	assert.Equal(t, dto.Paging{Page: 1, Size: 10, TotalRows: len(expectedAdmins), TotalPages: 1}, paging)
}

func TestFindAllTransaction(t *testing.T) {
	db, mock, teardown := setupTestDB(t)
	defer teardown()

	repo := NewtaskRepo(db)

	expectedAdmins := []model.Transaction{
		{
			Id:          "1",
			Book_id:     "1",
			Borrower_id: "1",
			Status:      "Delivered",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Admin_id:    "1",
		},
		{
			Id:          "2",
			Book_id:     "2",
			Borrower_id: "2",
			Status:      "Delivered",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Admin_id:    "2",
		},
	}

	rows := sqlmock.NewRows([]string{"id", "book_id", "borrower_id", "status", "created_at", "updated_at", "admin_id"})
	for _, admin := range expectedAdmins {
		rows.AddRow(admin.Id, admin.Book_id, admin.Borrower_id, admin.Status, admin.CreatedAt, admin.UpdatedAt, admin.Admin_id)
	}

	mock.ExpectQuery("SELECT (.+) FROM transaction limit \\$1 offset \\$2").
		WithArgs(10, 0).
		WillReturnRows(rows)

	mock.ExpectQuery("SELECT COUNT\\(\\*\\) FROM transaction").
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(len(expectedAdmins)))

	admins, paging, err := repo.FindAllTransaction(1, 10)

	assert.NoError(t, err)
	assert.Equal(t, expectedAdmins, admins)
	assert.Equal(t, dto.Paging{Page: 1, Size: 10, TotalRows: len(expectedAdmins), TotalPages: 1}, paging)
}

func TestFindAllBook(t *testing.T) {
	db, mock, teardown := setupTestDB(t)
	defer teardown()

	repo := NewtaskRepo(db)

	expectedAdmins := []model.Book{
		{
			Id:           "1",
			Title:        "John Doe",
			Author:       "johndoe",
			Release_Year: 2014,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
		{
			Id:           "2",
			Title:        "Jane Doe",
			Author:       "janedoe",
			Release_Year: 2015,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
	}

	rows := sqlmock.NewRows([]string{"id", "title", "author", "release_year", "created_at", "updated_at"})
	for _, admin := range expectedAdmins {
		rows.AddRow(admin.Id, admin.Title, admin.Author, admin.Release_Year, admin.CreatedAt, admin.UpdatedAt)
	}

	mock.ExpectQuery("SELECT (.+) FROM book limit \\$1 offset \\$2").
		WithArgs(10, 0).
		WillReturnRows(rows)

	mock.ExpectQuery("SELECT COUNT\\(\\*\\) FROM book").
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(len(expectedAdmins)))

	admins, paging, err := repo.FindAllBook(1, 10)

	assert.NoError(t, err)
	assert.Equal(t, expectedAdmins, admins)
	assert.Equal(t, dto.Paging{Page: 1, Size: 10, TotalRows: len(expectedAdmins), TotalPages: 1}, paging)
}

// Add similar tests for Borrowers, Books, and Transactions...
