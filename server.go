package main

import (
	"database/sql"
	"final-project-olib/config"
	"final-project-olib/controller"
	"final-project-olib/middleware"
	"final-project-olib/repository"
	"final-project-olib/service"
	"final-project-olib/usecase"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Server struct {
	taskUc     usecase.TaskUseCase
	authUc     usecase.AuthUseCase
	jwtService service.JwtService
	engine     *gin.Engine
}

func (s *Server) initRoute() {
	rg := s.engine.Group("/api/v1")
	authMiddleware := middleware.NewAuthMiddleware(s.jwtService)
	controller.NewAuthController(s.authUc, rg, authMiddleware).Route()
	controller.NewTaskController(s.taskUc, rg, authMiddleware).Routing()
}

func (s *Server) Run() {
	s.initRoute()

	s.engine.Run()
}

func NewServer() *Server {
	c, _ := config.NewConfig()

	urlConnect := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.Host, c.DbPort, c.DbUser, c.DbPassword, c.DbName)

	database, err := sql.Open(c.Driver, urlConnect)
	if err != nil {
		panic("connection Error")
	}
	bookRepo := repository.NewtaskRepo(database)

	jwtService := service.NewJwtService(c.TokenConfig)

	authorUC := usecase.NewAuthUseCase(jwtService, bookRepo)
	taskUc := usecase.NewTaskUseCase(bookRepo)

	return &Server{
		taskUc:     taskUc,
		engine:     gin.Default(),
		jwtService: jwtService,
		authUc:     authorUC,
	}
}
