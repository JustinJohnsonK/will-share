package app

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

type ServiceDependencies struct {
	Db *pgxpool.Pool
}

func InitServiceDependencies() *ServiceDependencies {
	db := intiDb()

	return &ServiceDependencies{
		Db: db,
	}
}

func (deps *ServiceDependencies) Close() {
	deps.Db.Close()
}
