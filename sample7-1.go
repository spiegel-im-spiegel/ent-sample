//go:build run
// +build run

package main

import (
	"context"
	"fmt"
	"os"
	"sample/dbconn"
	"sample/ent"

	"github.com/spiegel-im-spiegel/errs"
	"github.com/spiegel-im-spiegel/gocli/exitcode"
)

func Run() exitcode.ExitCode {
	// get ent context
	entCtx, err := dbconn.NewEnt()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return exitcode.Abnormal
	}
	defer entCtx.Close()

	// search data and delete
	if err := entCtx.Transaction(context.TODO(), func(tx *ent.Tx) error {
		ct, err := tx.BinaryFile.Delete().Exec(context.TODO())
		if err != nil {
			return errs.Wrap(err)
		}
		if ct <= 0 {
			return errs.New("not delete record", errs.WithContext("username", "Bob"))
		}
		return nil
	}); err != nil {
		entCtx.GetLogger().Error().Interface("error", errs.Wrap(err)).Send()
		return exitcode.Abnormal
	}

	return exitcode.Normal
}

func main() {
	Run().Exit()
}
