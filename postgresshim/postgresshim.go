package postgresshim

import (
	"github.com/lib/pq"
)

type PostgreSQLShim struct{}

func (*PostgreSQLShim) ParseDSN(dsn string) (*pq.Connector, error) {
	return pq.NewConnector(dsn)
}
