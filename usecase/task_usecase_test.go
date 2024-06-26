package usecase

import (
	"final-project-olib/mocking"
	"final-project-olib/model"
	"final-project-olib/model/dto"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var expectedBook = model.Book{
	Id:           "1",
	Title:        "The Ripper",
	Author:       "Jack",
	Release_Year: 2012,
	CreatedAt:    time.Now(),
	UpdatedAt:    time.Now(),
}
var expectedBorrower = model.Borrower{
	Id:        "1",
	Name:      "Jack",
	Email:     "jack@mail.com",
	Password:  "password",
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
	Address:   "London",
}
var expectedAdmin = model.Admin{
	Id:        "1",
	Name:      "Jack",
	Email:     "jack@mail.com",
	Password:  "password",
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
}
var expectedTransaction = model.Transaction{
	Id:          "1",
	Book_id:     "1",
	Borrower_id: "1",
	Status:      "Delivered",
	CreatedAt:   time.Now(),
	UpdatedAt:   time.Now(),
	Admin_id:    "1",
}

type TaskUseCaseTestSuite struct {
	suite.Suite
	arm *mocking.TaskRepoMock
	tuc TaskUseCase
}

func (suite *TaskUseCaseTestSuite) SetupTest() {
	suite.arm = new(mocking.TaskRepoMock)
	suite.tuc = NewTaskUseCase(suite.arm)
}

func (suite *TaskUseCaseTestSuite) TestDeleteBook() {
	suite.arm.On("DeleteBook", expectedBook.Id).Return(nil)
	err := suite.tuc.DeleteBook(expectedBook.Id)

	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
}

func (suite *TaskUseCaseTestSuite) TestDeleteBorrower() {
	suite.arm.On("DeleteBorrower", expectedBorrower.Id).Return(nil)
	err := suite.tuc.DeleteBorrower(expectedBorrower.Id)

	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
}

func (suite *TaskUseCaseTestSuite) TestDeleteAdmin() {
	suite.arm.On("DeleteAdmin", expectedAdmin.Id).Return(nil)
	err := suite.tuc.DeleteAdmin(expectedAdmin.Id)

	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
}

func (suite *TaskUseCaseTestSuite) TestDeleteTransaction() {
	suite.arm.On("DeleteTransaction", expectedTransaction.Id).Return(nil)
	err := suite.tuc.DeleteTransaction(expectedTransaction.Id)

	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
}
func (suite *TaskUseCaseTestSuite) TestFindBookById() {
	suite.arm.On("FindBookById", expectedBook.Id).Return(expectedBook, nil)
	actual, err := suite.tuc.FindBookById(expectedBook.Id)

	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedBook.Id, actual.Id)
}
func (suite *TaskUseCaseTestSuite) TestFindBorrowerById() {
	suite.arm.On("FindBorrowerById", expectedBorrower.Id).Return(expectedBorrower, nil)
	actual, err := suite.tuc.FindBorrowerById(expectedBorrower.Id)

	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedBorrower.Id, actual.Id)
}
func (suite *TaskUseCaseTestSuite) TestFindAdminById() {
	suite.arm.On("FindAdminById", expectedAdmin.Id).Return(expectedAdmin, nil)
	actual, err := suite.tuc.FindAdminById(expectedAdmin.Id)

	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedAdmin.Id, actual.Id)
}
func (suite *TaskUseCaseTestSuite) TestFindTransactionById() {
	suite.arm.On("FindTransactionById", expectedTransaction.Id).Return(expectedTransaction, nil)
	actual, err := suite.tuc.FindTransactionById(expectedTransaction.Id)

	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedTransaction.Id, actual.Id)
}
func (suite *TaskUseCaseTestSuite) TestFindBorrowerByEmail() {
	suite.arm.On("FindBorrowerByEmail", expectedBorrower.Email).Return(expectedBorrower, nil)
	actual, err := suite.tuc.FindBorrowerByEmail(expectedBorrower.Email)

	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedBorrower.Email, actual.Email)
}
func (suite *TaskUseCaseTestSuite) TestFindAdminByEmail() {
	suite.arm.On("FindAdminByEmail", expectedAdmin.Email).Return(expectedAdmin, nil)
	actual, err := suite.tuc.FindAdminByEmail(expectedAdmin.Email)

	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedAdmin.Email, actual.Email)
}

func (suite *TaskUseCaseTestSuite) TestFindAllBook() {
	mockData := []model.Book{expectedBook}

	mockPaging := dto.Paging{
		Page:       1,
		Size:       1,
		TotalRows:  5,
		TotalPages: 1,
	}
	suite.arm.On("FindAllBook", 1, 5).Return(mockData, mockPaging, nil)

	actual, paging, err := suite.tuc.FindAllBook(1, 5)

	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), actual, 1)
	assert.Equal(suite.T(), mockPaging.Page, paging.Page)
}
func (suite *TaskUseCaseTestSuite) TestFindAllBorrower() {
	mockData := []model.Borrower{expectedBorrower}

	mockPaging := dto.Paging{
		Page:       1,
		Size:       1,
		TotalRows:  5,
		TotalPages: 1,
	}
	suite.arm.On("FindAllBorrower", 1, 5).Return(mockData, mockPaging, nil)

	actual, paging, err := suite.tuc.FindAllBorrower(1, 5)

	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), actual, 1)
	assert.Equal(suite.T(), mockPaging.Page, paging.Page)
}
func (suite *TaskUseCaseTestSuite) TestFindAllAdmin() {
	mockData := []model.Admin{expectedAdmin}

	mockPaging := dto.Paging{
		Page:       1,
		Size:       1,
		TotalRows:  5,
		TotalPages: 1,
	}
	suite.arm.On("FindAllAdmin", 1, 5).Return(mockData, mockPaging, nil)

	actual, paging, err := suite.tuc.FindAllAdmin(1, 5)

	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), actual, 1)
	assert.Equal(suite.T(), mockPaging.Page, paging.Page)
}
func (suite *TaskUseCaseTestSuite) TestFindAllTransaction() {
	mockData := []model.Transaction{expectedTransaction}

	mockPaging := dto.Paging{
		Page:       1,
		Size:       1,
		TotalRows:  5,
		TotalPages: 1,
	}
	suite.arm.On("FindAllTransaction", 1, 5).Return(mockData, mockPaging, nil)

	actual, paging, err := suite.tuc.FindAllTransaction(1, 5)

	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), actual, 1)
	assert.Equal(suite.T(), mockPaging.Page, paging.Page)
}

func TestTaskUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(TaskUseCaseTestSuite))
}
