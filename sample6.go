//go:build run
// +build run

package main

import (
	"context"
	"fmt"
	"os"
	"sample/dbconn"
	"sample/ent"
	"sample/ent/binaryfile"
	"sample/ent/user"
	"sample/files"

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

	file4 := "files/file4.txt"
	bin4, err := files.GetBinary(file4)
	if err != nil {
		entCtx.GetLogger().Error().Interface("error", errs.Wrap(err)).Send()
		return exitcode.Abnormal
	}

	// search data and update
	if err := entCtx.Transaction(context.TODO(), func(tx *ent.Tx) error {
		ct, err := tx.BinaryFile.Update().Where(
			binaryfile.HasOwnerWith(user.Username("Bob")),
		).SetFilename(file4).SetBody(bin4).Save(context.TODO())
		if err != nil {
			return errs.Wrap(err)
		}
		if ct <= 0 {
			return errs.New("not change record", errs.WithContext("username", "Bob"))
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
