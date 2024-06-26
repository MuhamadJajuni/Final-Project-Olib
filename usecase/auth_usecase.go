package usecase

import (
	"final-project-olib/model"
	"final-project-olib/model/dto"
	"final-project-olib/repository"
	"final-project-olib/service"
)

type AuthUseCase interface {
	Login(payload dto.AuthReqDto) (dto.AuthResponDto, error)
	LoginAdmin(payload dto.AuthReqDto) (dto.AuthResponDto, error)
}
type authUseCase struct {
	jwtService service.JwtService
	authorUC   repository.TaskRepo
}

// Login implements AuthUseCase.
func (a *authUseCase) Login(payload dto.AuthReqDto) (dto.AuthResponDto, error) {
	author, err := a.authorUC.FindBorrowerByEmail(payload.Email)
	if err != nil {
		return dto.AuthResponDto{}, err
	}
	authorps := model.Borrower{
		Email:    author.Email,
		Password: author.Password,
	}
	token, err := a.jwtService.CreateToken(authorps)
	if err != nil {
		return dto.AuthResponDto{}, err
	}
	return token, nil
}
func (a *authUseCase) LoginAdmin(payload dto.AuthReqDto) (dto.AuthResponDto, error) {
	author, err := a.authorUC.FindAdminByEmail(payload.Email)
	if err != nil {
		return dto.AuthResponDto{}, err
	}
	authorps := model.Borrower{
		Email:    author.Email,
		Password: author.Password,
	}
	token, err := a.jwtService.CreateToken(authorps)
	if err != nil {
		return dto.AuthResponDto{}, err
	}
	return token, nil
}
func NewAuthUseCase(jwtService service.JwtService, authorUc repository.TaskRepo) AuthUseCase {
	return &authUseCase{jwtService: jwtService, authorUC: authorUc}
}
