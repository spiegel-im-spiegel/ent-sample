package dbconn

import (
	"context"
	"fmt"
	"sample/ent"
	"sample/env"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/rs/zerolog"
	"github.com/spiegel-im-spiegel/errs"
)

type EntContext struct {
	Client *ent.Client
	Logger *zerolog.Logger
}

func NewEnt() (*EntContext, error) {
	pgxCtx, err := NewPgx()
	if err != nil {
		return nil, errs.Wrap(err)
	}
	entCtx := &EntContext{
		Logger: pgxCtx.GetLogger(),
	}
	entCtx.Client = ent.NewClient(
		ent.Driver(
			sql.OpenDB(dialect.Postgres, pgxCtx.GetDb()),
		),
		ent.Log(func(v ...interface{}) {
			entCtx.Logger.Debug().Msg(fmt.Sprint(v...))
		}),
	)
	if env.LogLevel() >= env.LevelDebug {
		entCtx.Client = entCtx.Client.Debug()
	}
	return entCtx, nil
}

func (entCtx *EntContext) GetClient() *ent.Client {
	if entCtx == nil {
		return nil
	}
	return entCtx.Client
}

func (entCtx *EntContext) GetLogger() *zerolog.Logger {
	if entCtx == nil {
		lggr := zerolog.Nop()
		return &lggr
	}
	return entCtx.Logger
}

func (entCtx *EntContext) Close() error {
	if client := entCtx.GetClient(); client != nil {
		return errs.Wrap(client.Close())
	}
	return nil
}

func (entCtx *EntContext) Transaction(ctx context.Context, fn func(tx *ent.Tx) error) error {
	client := entCtx.GetClient()
	if client == nil {
		return errs.New("null reference instance")
	}
	logger := entCtx.GetLogger()

	logger.Info().Msg("begining transaction")
	tx, err := client.Tx(ctx)
	if err != nil {
		return errs.Wrap(err)
	}
	defer func() {
		if v := recover(); v != nil {
			_ = tx.Rollback()
			panic(v)
		}
	}()

	if err := fn(tx); err != nil {
		txErr := errs.Wrap(err)
		if err := tx.Rollback(); err != nil {
			return errs.Wrap(err, errs.WithCause(txErr))
		}
		return txErr
	}

	logger.Info().Msg("committing transaction")
	if err := tx.Commit(); err != nil {
		return errs.Wrap(err)
	}
	return nil
}
