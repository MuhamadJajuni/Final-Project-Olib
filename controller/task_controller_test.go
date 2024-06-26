package controller

import (
	"bytes"
	"encoding/json"
	"final-project-olib/model"
	"final-project-olib/model/dto"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockTaskUseCase struct {
	mock.Mock
}

func (m *MockTaskUseCase) FindAllBook(page, size int) ([]model.Book, dto.Paging, error) {
	args := m.Called(page, size)
	return args.Get(0).([]model.Book), args.Get(1).(dto.Paging), args.Error(2)
}

func (m *MockTaskUseCase) FindBookById(id string) (model.Book, error) {
	args := m.Called(id)
	return args.Get(0).(model.Book), args.Error(1)
}

func (m *MockTaskUseCase) PostNewBook(book model.Book) (model.Book, error) {
	args := m.Called(book)
	return args.Get(0).(model.Book), args.Error(1)
}

func (m *MockTaskUseCase) UpdateBook(id string, book model.Book) (model.Book, error) {
	args := m.Called(id, book)
	return args.Get(0).(model.Book), args.Error(1)
}

func (m *MockTaskUseCase) DeleteBook(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockTaskUseCase) FindAllBorrower(page, size int) ([]model.Borrower, dto.Paging, error) {
	args := m.Called(page, size)
	return args.Get(0).([]model.Borrower), args.Get(1).(dto.Paging), args.Error(2)
}

func (m *MockTaskUseCase) FindBorrowerById(id string) (model.Borrower, error) {
	args := m.Called(id)
	return args.Get(0).(model.Borrower), args.Error(1)
}

func (m *MockTaskUseCase) FindBorrowerByEmail(email string) (model.Borrower, error) {
	args := m.Called(email)
	return args.Get(0).(model.Borrower), args.Error(1)
}

func (m *MockTaskUseCase) RegisterBorrower(book model.Borrower) (model.Borrower, error) {
	args := m.Called(book)
	return args.Get(0).(model.Borrower), args.Error(1)
}

func (m *MockTaskUseCase) UpdateBorrower(id string, book model.Borrower) (model.Borrower, error) {
	args := m.Called(id, book)
	return args.Get(0).(model.Borrower), args.Error(1)
}

func (m *MockTaskUseCase) DeleteBorrower(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockTaskUseCase) FindAllAdmin(page, size int) ([]model.Admin, dto.Paging, error) {
	args := m.Called(page, size)
	return args.Get(0).([]model.Admin), args.Get(1).(dto.Paging), args.Error(2)
}

func (m *MockTaskUseCase) FindAdminById(id string) (model.Admin, error) {
	args := m.Called(id)
	return args.Get(0).(model.Admin), args.Error(1)
}

func (m *MockTaskUseCase) FindAdminByEmail(email string) (model.Admin, error) {
	args := m.Called(email)
	return args.Get(0).(model.Admin), args.Error(1)
}
func (m *MockTaskUseCase) RegisterAdmin(admin model.Admin) (model.Admin, error) {
	args := m.Called(admin)
	return args.Get(0).(model.Admin), args.Error(1)
}

func (m *MockTaskUseCase) UpdateAdmin(id string, admin model.Admin) (model.Admin, error) {
	args := m.Called(id, admin)
	return args.Get(0).(model.Admin), args.Error(1)
}

func (m *MockTaskUseCase) DeleteAdmin(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockTaskUseCase) FindAllTransaction(page, size int) ([]model.Transaction, dto.Paging, error) {
	args := m.Called(page, size)
	return args.Get(0).([]model.Transaction), args.Get(1).(dto.Paging), args.Error(2)
}

func (m *MockTaskUseCase) FindTransactionById(id string) (model.Transaction, error) {
	args := m.Called(id)
	return args.Get(0).(model.Transaction), args.Error(1)
}

func (m *MockTaskUseCase) PostNewTransaction(newTask model.Transaction) (model.Transaction, error) {
	args := m.Called(newTask)
	return args.Get(0).(model.Transaction), args.Error(1)
}

func (m *MockTaskUseCase) UpdateTransaction(id string, transaction model.Transaction) (model.Transaction, error) {
	args := m.Called(id, transaction)
	return args.Get(0).(model.Transaction), args.Error(1)
}

func (m *MockTaskUseCase) DeleteTransaction(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

// Add similar mock methods for Borrowers, Admins, and Transactions...

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	return router
}

func TestListHandler(t *testing.T) {
	mockUseCase := new(MockTaskUseCase)
	mockUseCase.On("FindAllBook", 1, 10).Return([]model.Book{
		{Id: "1", Title: "Book 1", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}, dto.Paging{Page: 1, Size: 10, TotalRows: 1, TotalPages: 1}, nil)

	router := setupRouter()
	taskController := NewTaskController(mockUseCase, router.Group("/"))
	taskController.Routing()

	req, _ := http.NewRequest("GET", "/books?page=1&size=10", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	mockUseCase.AssertExpectations(t)
}

func TestGetByIdHandler(t *testing.T) {
	mockUseCase := new(MockTaskUseCase)
	mockBook := model.Book{Id: "1", Title: "Book 1", CreatedAt: time.Now(), UpdatedAt: time.Now()}
	mockUseCase.On("FindBookById", "1").Return(mockBook, nil)

	router := setupRouter()
	taskController := NewTaskController(mockUseCase, router.Group("/"))
	taskController.Routing()

	req, _ := http.NewRequest("GET", "/books/1", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	mockUseCase.AssertExpectations(t)
}

func TestPostHandler(t *testing.T) {
	mockUseCase := new(MockTaskUseCase)
	mockBook := model.Book{Title: "New Book"}
	mockUseCase.On("PostNewBook", mockBook).Return(mockBook, nil)

	router := setupRouter()
	taskController := NewTaskController(mockUseCase, router.Group("/"))
	taskController.Routing()

	body, _ := json.Marshal(mockBook)
	req, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)
	mockUseCase.AssertExpectations(t)
}

func TestUpdateHandler(t *testing.T) {
	mockUseCase := new(MockTaskUseCase)
	mockBook := model.Book{Title: "Updated Book"}
	mockUseCase.On("UpdateBook", "1", mockBook).Return(mockBook, nil)

	router := setupRouter()
	taskController := NewTaskController(mockUseCase, router.Group("/"))
	taskController.Routing()

	body, _ := json.Marshal(mockBook)
	req, _ := http.NewRequest("PUT", "/books/1", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	mockUseCase.AssertExpectations(t)
}

func TestDeleteHandler(t *testing.T) {
	mockUseCase := new(MockTaskUseCase)
	mockUseCase.On("DeleteBook", "1").Return(nil)

	router := setupRouter()
	taskController := NewTaskController(mockUseCase, router.Group("/"))
	taskController.Routing()

	req, _ := http.NewRequest("DELETE", "/books/1", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	mockUseCase.AssertExpectations(t)
}

func TestListBorrowerHandler(t *testing.T) {
	mockUseCase := new(MockTaskUseCase)
	mockUseCase.On("FindAllBorrower", 1, 10).Return([]model.Borrower{
		{Id: "1", Name: "Book 1", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}, dto.Paging{Page: 1, Size: 10, TotalRows: 1, TotalPages: 1}, nil)

	router := setupRouter()
	taskController := NewTaskController(mockUseCase, router.Group("/"))
	taskController.Routing()

	req, _ := http.NewRequest("GET", "/borrower?page=1&size=10", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	mockUseCase.AssertExpectations(t)
}

func TestGetBorrowerByIdHandler(t *testing.T) {
	mockUseCase := new(MockTaskUseCase)
	mockBook := model.Borrower{Id: "1", Name: "Book 1", CreatedAt: time.Now(), UpdatedAt: time.Now()}
	mockUseCase.On("FindBorrowerById", "1").Return(mockBook, nil)

	router := setupRouter()
	taskController := NewTaskController(mockUseCase, router.Group("/"))
	taskController.Routing()

	req, _ := http.NewRequest("GET", "/borrower/1", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	mockUseCase.AssertExpectations(t)
}

func TestPostBorrowerHandler(t *testing.T) {
	mockUseCase := new(MockTaskUseCase)
	mockBook := model.Borrower{Name: "New Book"}
	mockUseCase.On("PostNewBorrower", mockBook).Return(mockBook, nil)

	router := setupRouter()
	taskController := NewTaskController(mockUseCase, router.Group("/"))
	taskController.Routing()

	body, _ := json.Marshal(mockBook)
	req, _ := http.NewRequest("POST", "/borrower", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)
	mockUseCase.AssertExpectations(t)
}

func TestUpdateBorrowerHandler(t *testing.T) {
	mockUseCase := new(MockTaskUseCase)
	mockBook := model.Borrower{Name: "Updated Book"}
	mockUseCase.On("UpdateBook", "1", mockBook).Return(mockBook, nil)

	router := setupRouter()
	taskController := NewTaskController(mockUseCase, router.Group("/"))
	taskController.Routing()

	body, _ := json.Marshal(mockBook)
	req, _ := http.NewRequest("PUT", "/borrower/1", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	mockUseCase.AssertExpectations(t)
}

func TestDeleteBorrowerHandler(t *testing.T) {
	mockUseCase := new(MockTaskUseCase)
	mockUseCase.On("DeleteBorrower", "1").Return(nil)

	router := setupRouter()
	taskController := NewTaskController(mockUseCase, router.Group("/"))
	taskController.Routing()

	req, _ := http.NewRequest("DELETE", "/borrower/1", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	mockUseCase.AssertExpectations(t)
}

func TestListAdminHandler(t *testing.T) {
	mockUseCase := new(MockTaskUseCase)
	mockUseCase.On("FindAllAdmin", 1, 10).Return([]model.Admin{
		{Id: "1", Name: "Book 1", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}, dto.Paging{Page: 1, Size: 10, TotalRows: 1, TotalPages: 1}, nil)

	router := setupRouter()
	taskController := NewTaskController(mockUseCase, router.Group("/"))
	taskController.Routing()

	req, _ := http.NewRequest("GET", "/admin?page=1&size=10", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	mockUseCase.AssertExpectations(t)
}

func TestGetAdminByIdHandler(t *testing.T) {
	mockUseCase := new(MockTaskUseCase)
	mockBook := model.Admin{Id: "1", Name: "Book 1", CreatedAt: time.Now(), UpdatedAt: time.Now()}
	mockUseCase.On("FindAdminById", "1").Return(mockBook, nil)

	router := setupRouter()
	taskController := NewTaskController(mockUseCase, router.Group("/"))
	taskController.Routing()

	req, _ := http.NewRequest("GET", "/admin/1", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	mockUseCase.AssertExpectations(t)
}

func TestPostAdminHandler(t *testing.T) {
	mockUseCase := new(MockTaskUseCase)
	mockBook := model.Admin{Name: "New Book"}
	mockUseCase.On("PostNewAdmin", mockBook).Return(mockBook, nil)

	router := setupRouter()
	taskController := NewTaskController(mockUseCase, router.Group("/"))
	taskController.Routing()

	body, _ := json.Marshal(mockBook)
	req, _ := http.NewRequest("POST", "/admin", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)
	mockUseCase.AssertExpectations(t)
}

func TestUpdateAdminHandler(t *testing.T) {
	mockUseCase := new(MockTaskUseCase)
	mockBook := model.Admin{Name: "Updated Book"}
	mockUseCase.On("UpdateAdmin", "1", mockBook).Return(mockBook, nil)

	router := setupRouter()
	taskController := NewTaskController(mockUseCase, router.Group("/"))
	taskController.Routing()

	body, _ := json.Marshal(mockBook)
	req, _ := http.NewRequest("PUT", "/admin/1", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	mockUseCase.AssertExpectations(t)
}

func TestDeleteAdminHandler(t *testing.T) {
	mockUseCase := new(MockTaskUseCase)
	mockUseCase.On("DeleteAdmin", "1").Return(nil)

	router := setupRouter()
	taskController := NewTaskController(mockUseCase, router.Group("/"))
	taskController.Routing()

	req, _ := http.NewRequest("DELETE", "/admin/1", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	mockUseCase.AssertExpectations(t)
}

func TestListTransactionHandler(t *testing.T) {
	mockUseCase := new(MockTaskUseCase)
	mockUseCase.On("FindAllTransaction", 1, 10).Return([]model.Transaction{
		{Id: "1", Book_id: "1", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}, dto.Paging{Page: 1, Size: 10, TotalRows: 1, TotalPages: 1}, nil)

	router := setupRouter()
	taskController := NewTaskController(mockUseCase, router.Group("/"))
	taskController.Routing()

	req, _ := http.NewRequest("GET", "/transaction?page=1&size=10", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	mockUseCase.AssertExpectations(t)
}

func TestGetTransactionByIdHandler(t *testing.T) {
	mockUseCase := new(MockTaskUseCase)
	mockBook := model.Transaction{Id: "1", Book_id: "1", CreatedAt: time.Now(), UpdatedAt: time.Now()}
	mockUseCase.On("FindTransactionById", "1").Return(mockBook, nil)

	router := setupRouter()
	taskController := NewTaskController(mockUseCase, router.Group("/"))
	taskController.Routing()

	req, _ := http.NewRequest("GET", "/transaction/1", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	mockUseCase.AssertExpectations(t)
}

func TestPostTransactionHandler(t *testing.T) {
	mockUseCase := new(MockTaskUseCase)
	mockBook := model.Transaction{Status: "New Book"}
	mockUseCase.On("PostNewTransaction", mockBook).Return(mockBook, nil)

	router := setupRouter()
	taskController := NewTaskController(mockUseCase, router.Group("/"))
	taskController.Routing()

	body, _ := json.Marshal(mockBook)
	req, _ := http.NewRequest("POST", "/transaction", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)
	mockUseCase.AssertExpectations(t)
}

func TestUpdateTransactionHandler(t *testing.T) {
	mockUseCase := new(MockTaskUseCase)
	mockBook := model.Transaction{Status: "Updated Book"}
	mockUseCase.On("UpdateTransaction", "1", mockBook).Return(mockBook, nil)

	router := setupRouter()
	taskController := NewTaskController(mockUseCase, router.Group("/"))
	taskController.Routing()

	body, _ := json.Marshal(mockBook)
	req, _ := http.NewRequest("PUT", "/transaction/1", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	mockUseCase.AssertExpectations(t)
}

func TestDeleteTransactionHandler(t *testing.T) {
	mockUseCase := new(MockTaskUseCase)
	mockUseCase.On("DeleteTransaction", "1").Return(nil)

	router := setupRouter()
	taskController := NewTaskController(mockUseCase, router.Group("/"))
	taskController.Routing()

	req, _ := http.NewRequest("DELETE", "/transaction/1", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	mockUseCase.AssertExpectations(t)
}

// Add similar tests for Borrowers, Admins, and Transactions...
