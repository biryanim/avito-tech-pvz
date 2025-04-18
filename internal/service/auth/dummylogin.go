package auth

import (
	"context"
	"fmt"
	"github.com/biryanim/avito-tech-pvz/internal/model"
	"github.com/biryanim/avito-tech-pvz/internal/utils"
	"github.com/pkg/errors"
)

func (s *serv) DummyLogin(ctx context.Context, role string) (string, error) {

	fmt.Println("333333333333333333333333")
	userRole := model.Role(role)
	if !userRole.IsValid() {
		//TODO: добавить кастомные ошибки для моделек
		return "", errors.New("invalid role")
	}
	token, err := utils.GenerateToken(userRole, s.jwtConfig.TokenSecret(), s.jwtConfig.TokenExpiration())
	fmt.Println("444444444444444444444444")
	if err != nil {
		return "", err
	}
	return token, nil
}
