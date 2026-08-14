package impl

import (
	"context"
	"database/sql"
	"log"

	"github.com/aarondl/sqlboiler/v4/boil"
	"github.com/tortillaproduction/go-microservices/apps/command-service/infra/sqlboiler/handler"
)

type transaction struct{}

func (inc *transaction) begin(ctx context.Context) (*sql.Tx, error) {
	tran, err := boil.BeginTx(ctx, nil)
	if err != nil {
		return nil, handler.DBErrHandler(err)
	}

	return tran, nil
}

func (ins *transaction) complete(tran *sql.Tx, err error) error {
	if err != nil {
		if e := tran.Rollback(); e != nil {
			return handler.DBErrHandler(err)
		} else {
			log.Println("transaction rolled back")
		}
	} else {
		if e := tran.Commit(); e != nil {
			return handler.DBErrHandler(err)
		} else {
			log.Println("transaction committed")
		}
	}

	return nil
}
