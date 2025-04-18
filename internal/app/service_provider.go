package app

import (
	"context"
	"log"

	"github.com/biryanim/avito-tech-pvz/internal/api/auth"
	"github.com/biryanim/avito-tech-pvz/internal/config"
	"github.com/biryanim/avito-tech-pvz/internal/repository"
	accessRepository "github.com/biryanim/avito-tech-pvz/internal/repository/access"
	pvzRepository "github.com/biryanim/avito-tech-pvz/internal/repository/pvz"
	userRepository "github.com/biryanim/avito-tech-pvz/internal/repository/user"
	"github.com/biryanim/avito-tech-pvz/internal/service"
	authService "github.com/biryanim/avito-tech-pvz/internal/service/auth"
	pvzService "github.com/biryanim/avito-tech-pvz/internal/service/pvz"
	"github.com/jackc/pgx/v5/pgxpool"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	jwtConfig  config.JWTConfig
	httpConfig config.HTTPConfig

	pool             *pgxpool.Pool
	userRepository   repository.UserRepository
	pvzRepository    repository.PvzRepository
	accessRepository repository.AccessRepository

	authService service.AuthService
	pvzService  service.PVZService

	authImpl *auth.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := config.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to load pg config: %v", err)
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) JWTConfig() config.JWTConfig {
	if s.jwtConfig == nil {
		cfg, err := config.NewJWTConfig()
		if err != nil {
			log.Fatalf("failed to load jwt config: %v", err)
		}

		s.jwtConfig = cfg
	}

	return s.jwtConfig
}

func (s *serviceProvider) HTTPConfig() config.HTTPConfig {
	if s.httpConfig == nil {
		cfg, err := config.NewHTTPConfig()
		if err != nil {
			log.Fatalf("failed to load http config: %v", err)
		}

		s.httpConfig = cfg
	}

	return s.httpConfig
}

func (s *serviceProvider) Pool(ctx context.Context) *pgxpool.Pool {
	if s.pool == nil {
		conn, err := pgxpool.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to connect to database: %v", err)
		}

		err = conn.Ping(ctx)
		if err != nil {
			log.Fatalf("failed to ping database: %v", err)
		}
		//TODO: добавить Closer
		//closer.Add(conn.Close)
		s.pool = conn
	}

	return s.pool
}

func (s *serviceProvider) UserRepository(ctx context.Context) repository.UserRepository {
	if s.userRepository == nil {
		s.userRepository = userRepository.NewRepository(s.Pool(ctx))
	}

	return s.userRepository
}

func (s *serviceProvider) PvzRepository(ctx context.Context) repository.PvzRepository {
	if s.pvzRepository == nil {
		s.pvzRepository = pvzRepository.NewRepository(s.Pool(ctx))
	}

	return s.pvzRepository
}

func (s *serviceProvider) AccessRepository(ctx context.Context) repository.AccessRepository {
	if s.accessRepository == nil {
		s.accessRepository = accessRepository.NewRepository(s.Pool(ctx))
	}

	return s.accessRepository
}

func (s *serviceProvider) AuthService(ctx context.Context) service.AuthService {
	if s.authService == nil {
		s.authService = authService.NewService(
			s.UserRepository(ctx),
			s.AccessRepository(ctx),
			s.JWTConfig(),
		)
	}

	return s.authService
}

func (s *serviceProvider) PvzService(ctx context.Context) service.PVZService {
	if s.pvzService == nil {
		s.pvzService = pvzService.NewService(s.PvzRepository(ctx))
	}

	return s.pvzService
}

func (s *serviceProvider) AuthImpl(ctx context.Context) *auth.Implementation {
	if s.authImpl == nil {
		s.authImpl = auth.NewImplementation(s.AuthService(ctx))
	}

	return s.authImpl
}
