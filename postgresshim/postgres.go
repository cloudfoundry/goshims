package postgresshim

import (
	"github.com/lib/pq"
)

//go:generate counterfeiter -o postgres_fake/fake_postgres.go . PostgreSQL
type PostgreSQL interface {
	ParseDSN(dsn string) (*pq.Connector, error)
}
