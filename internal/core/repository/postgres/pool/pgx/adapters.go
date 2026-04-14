package core_pgx_pool

import (
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)


type pgxRows struct {
	pgx.Rows
}


type pgxRow struct {
	pgx.Row
}



type pgxComandTag struct {
	pgconn.CommandTag
}

