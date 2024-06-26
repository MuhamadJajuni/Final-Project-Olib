package controller

import (
	"final-project-olib/middleware"
	"final-project-olib/model"
	customdto "final-project-olib/model/dto/common_response"
	"final-project-olib/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type taskController struct {
	taskUseCase    usecase.TaskUseCase
	engine         *gin.RouterGroup
	authMiddleware middleware.AuthMiddleware
}

func (a *taskController) listHandler(ctx *gin.Context) {
	page, er := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	size, er2 := strconv.Atoi(ctx.DefaultQuery("size", "10"))

	if er != nil {
		customdto.SendErrorResponse(ctx, http.StatusBadRequest, er.Error())
	}
	if er2 != nil {
		customdto.SendErrorResponse(ctx, http.StatusBadRequest, er2.Error())
	}
	listData, paging, err := a.taskUseCase.FindAllBook(page, size)

	if err != nil {
		customdto.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}
	var data []interface{}
	for _, b := range listData {
		data = append(data, b)
	}
	customdto.SendManyResponse(ctx, data, paging, "ok")
}

func (a *taskController) getByIdHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	data, err := a.taskUseCase.FindBookById(id)

	if err != nil {
		panic(err)
	}
	customdto.SendSingleResponse(ctx, data, "ok")
}

func (a *taskController) postHandler(ctx *gin.Context) {
	var newTask model.Book
	err := ctx.ShouldBind(&newTask)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, err := a.taskUseCase.PostNewBook(newTask)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	customdto.SendSingleResponse(ctx, data, "created")
}

func (a *taskController) updateHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	var updatedTask model.Book
	err := ctx.ShouldBind(&updatedTask)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, err := a.taskUseCase.UpdateBook(id, updatedTask)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	customdto.SendSingleResponse(ctx, data, "updated")
}

func (a *taskController) deleteHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	err := a.taskUseCase.DeleteBook(id)
	if err != nil {
		customdto.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	data := model.Book{}
	customdto.SendSingleResponse(ctx, data, "deleted")
}

func (a *taskController) listBorrowerHandler(ctx *gin.Context) {
	page, er := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	size, er2 := strconv.Atoi(ctx.DefaultQuery("size", "10"))

	if er != nil {
		customdto.SendErrorResponse(ctx, http.StatusBadRequest, er.Error())
	}
	if er2 != nil {
		customdto.SendErrorResponse(ctx, http.StatusBadRequest, er2.Error())
	}
	listData, paging, err := a.taskUseCase.FindAllBorrower(page, size)

	if err != nil {
		customdto.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}
	var data []interface{}
	for _, b := range listData {
		data = append(data, b)
	}
	customdto.SendManyResponse(ctx, data, paging, "ok")
}

func (a *taskController) getBorrowerByIdHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	data, err := a.taskUseCase.FindBorrowerById(id)

	if err != nil {
		panic(err)
	}
	customdto.SendSingleResponse(ctx, data, "ok")
}

func (a *taskController) registerBorrowerHandler(ctx *gin.Context) {
	var newTask model.Borrower
	err := ctx.ShouldBind(&newTask)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, err := a.taskUseCase.RegisterBorrower(newTask)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	customdto.SendSingleResponse(ctx, data, "created")
}

func (a *taskController) updateBorrowerHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	var updatedTask model.Borrower
	err := ctx.ShouldBind(&updatedTask)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, err := a.taskUseCase.UpdateBorrower(id, updatedTask)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	customdto.SendSingleResponse(ctx, data, "updated")
}

func (a *taskController) deleteBorrowerHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	err := a.taskUseCase.DeleteBorrower(id)
	if err != nil {
		customdto.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	data := model.Borrower{}
	customdto.SendSingleResponse(ctx, data, "deleted")
}

func (a *taskController) listAdminHandler(ctx *gin.Context) {
	page, er := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	size, er2 := strconv.Atoi(ctx.DefaultQuery("size", "10"))

	if er != nil {
		customdto.SendErrorResponse(ctx, http.StatusBadRequest, er.Error())
	}
	if er2 != nil {
		customdto.SendErrorResponse(ctx, http.StatusBadRequest, er2.Error())
	}
	listData, paging, err := a.taskUseCase.FindAllAdmin(page, size)

	if err != nil {
		customdto.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}
	var data []interface{}
	for _, b := range listData {
		data = append(data, b)
	}
	customdto.SendManyResponse(ctx, data, paging, "ok")
}

func (a *taskController) getAdminByIdHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	data, err := a.taskUseCase.FindAdminById(id)

	if err != nil {
		panic(err)
	}
	customdto.SendSingleResponse(ctx, data, "ok")
}

func (a *taskController) registerAdminHandler(ctx *gin.Context) {
	var newTask model.Admin
	err := ctx.ShouldBind(&newTask)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, err := a.taskUseCase.RegisterAdmin(newTask)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	customdto.SendSingleResponse(ctx, data, "created")
}

func (a *taskController) updateAdminHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	var updatedTask model.Admin
	err := ctx.ShouldBind(&updatedTask)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, err := a.taskUseCase.UpdateAdmin(id, updatedTask)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	customdto.SendSingleResponse(ctx, data, "updated")
}

func (a *taskController) deleteAdminHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	err := a.taskUseCase.DeleteAdmin(id)
	if err != nil {
		customdto.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	data := model.Admin{}
	customdto.SendSingleResponse(ctx, data, "deleted")
}

func (a *taskController) listTransactionHandler(ctx *gin.Context) {
	page, er := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	size, er2 := strconv.Atoi(ctx.DefaultQuery("size", "10"))

	if er != nil {
		customdto.SendErrorResponse(ctx, http.StatusBadRequest, er.Error())
	}
	if er2 != nil {
		customdto.SendErrorResponse(ctx, http.StatusBadRequest, er2.Error())
	}
	listData, paging, err := a.taskUseCase.FindAllTransaction(page, size)

	if err != nil {
		customdto.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}
	var data []interface{}
	for _, b := range listData {
		data = append(data, b)
	}
	customdto.SendManyResponse(ctx, data, paging, "ok")
}

func (a *taskController) getTransactionByIdHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	data, err := a.taskUseCase.FindTransactionById(id)

	if err != nil {
		panic(err)
	}
	customdto.SendSingleResponse(ctx, data, "ok")
}

func (a *taskController) postTransactionHandler(ctx *gin.Context) {
	var newTask model.Transaction
	err := ctx.ShouldBind(&newTask)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, err := a.taskUseCase.PostNewTransaction(newTask)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	customdto.SendSingleResponse(ctx, data, "created")
}

func (a *taskController) updateTransactionHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	var updatedTask model.Transaction
	err := ctx.ShouldBind(&updatedTask)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, err := a.taskUseCase.UpdateTransaction(id, updatedTask)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	customdto.SendSingleResponse(ctx, data, "updated")
}

func (a *taskController) deleteTransactionHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	err := a.taskUseCase.DeleteTransaction(id)
	if err != nil {
		customdto.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	data := model.Book{}
	customdto.SendSingleResponse(ctx, data, "deleted")
}

func (a *taskController) Routing() {
	a.engine.GET("/books", a.authMiddleware.CheckToken(), a.listHandler)
	a.engine.GET("/books/:id", a.authMiddleware.CheckToken(), a.getByIdHandler)
	a.engine.POST("/books", a.authMiddleware.CheckToken(), a.postHandler)
	a.engine.PUT("/books/:id", a.authMiddleware.CheckToken(), a.updateHandler)
	a.engine.DELETE("/books/:id", a.authMiddleware.CheckToken(), a.deleteHandler)

	a.engine.GET("/borrower", a.authMiddleware.CheckToken(), a.listBorrowerHandler)
	a.engine.GET("/borrower/:id", a.authMiddleware.CheckToken(), a.getBorrowerByIdHandler)
	a.engine.POST("/borrower", a.registerBorrowerHandler)
	a.engine.PUT("/borrower/:id", a.authMiddleware.CheckToken(), a.updateBorrowerHandler)
	a.engine.DELETE("/borrower/:id", a.authMiddleware.CheckToken(), a.deleteBorrowerHandler)

	a.engine.GET("/admin", a.authMiddleware.CheckToken(), a.listAdminHandler)
	a.engine.GET("/admin/:id", a.authMiddleware.CheckToken(), a.getAdminByIdHandler)
	a.engine.POST("/admin", a.registerAdminHandler)
	a.engine.PUT("/admin/:id", a.authMiddleware.CheckToken(), a.updateAdminHandler)
	a.engine.DELETE("/admin/:id", a.authMiddleware.CheckToken(), a.deleteAdminHandler)

	a.engine.GET("/transaction", a.authMiddleware.CheckToken(), a.listTransactionHandler)
	a.engine.GET("/transaction/:id", a.authMiddleware.CheckToken(), a.getTransactionByIdHandler)
	a.engine.POST("/transaction", a.authMiddleware.CheckToken(), a.postTransactionHandler)
	a.engine.PUT("/transaction/:id", a.authMiddleware.CheckToken(), a.updateTransactionHandler)
	a.engine.DELETE("/transaction/:id", a.authMiddleware.CheckToken(), a.deleteTransactionHandler)
}

func NewTaskController(taskUc usecase.TaskUseCase, rg *gin.RouterGroup, authMiddle middleware.AuthMiddleware) *taskController {
	return &taskController{
		taskUseCase:    taskUc,
		authMiddleware: authMiddle,
		engine:         rg,
	}
}
