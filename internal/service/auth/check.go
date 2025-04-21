package auth

import (
	"context"
	"github.com/biryanim/avito-tech-pvz/internal/model"
	"github.com/pkg/errors"
)

var accessibleRoles map[string]map[model.Role]struct{}

func (s *serv) accessibleRoles(ctx context.Context) (map[string]map[model.Role]struct{}, error) {
	if accessibleRoles == nil {
		accessibleRoles = make(map[string]map[model.Role]struct{})

		accessInfo, err := s.accessRepository.GetList(ctx)
		if err != nil {
			return nil, err
		}

		for _, info := range accessInfo {
			key := info.Method + ":" + info.EndpointAddress
			if _, ok := accessibleRoles[key]; !ok {
				accessibleRoles[key] = make(map[model.Role]struct{})
			}
			accessibleRoles[key][info.Role] = struct{}{}
		}
	}

	return accessibleRoles, nil
}

func (s *serv) Check(ctx context.Context, token, method, endpointAddress string) (bool, error) {
	claims, err := s.verifyToken(token, s.jwtConfig.TokenSecret())
	if err != nil {
		return false, err
	}

	accessMap, err := s.accessibleRoles(ctx)
	if err != nil {
		return false, errors.Wrap(err, "could not check access roles for endpoint")
	}

	role, ok := accessMap[method+":"+endpointAddress]
	if !ok {
		return true, nil
	}

	if _, ok = role[claims.Role]; ok {
		return true, nil
	}

	return false, errors.New("access denied")
}
