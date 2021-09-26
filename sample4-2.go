//go:build run
// +build run

package main

import (
	"context"
	"fmt"
	"os"
	"sample/dbconn"
	"sample/ent"
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

	file1 := "files/file1.txt"
	bin1, err := files.GetBinary(file1)
	if err != nil {
		entCtx.GetLogger().Error().Interface("error", errs.Wrap(err)).Send()
		return exitcode.Abnormal
	}
	file2 := "files/file2.txt"
	bin2, err := files.GetBinary(file2)
	if err != nil {
		entCtx.GetLogger().Error().Interface("error", errs.Wrap(err)).Send()
		return exitcode.Abnormal
	}
	file3 := "files/file3.txt"
	bin3, err := files.GetBinary(file3)
	if err != nil {
		entCtx.GetLogger().Error().Interface("error", errs.Wrap(err)).Send()
		return exitcode.Abnormal
	}

	// create data
	if err := entCtx.Transaction(context.TODO(), func(tx *ent.Tx) error {
		user, err := tx.User.Create().SetUsername("Alice").Save(context.TODO())
		if err != nil {
			return errs.Wrap(err, errs.WithContext("username", "Alice"))
		}
		if _, err := tx.BinaryFile.CreateBulk(
			tx.BinaryFile.Create().SetFilename(file1).SetBody(bin1).SetOwner(user),
			tx.BinaryFile.Create().SetFilename(file2).SetBody(bin2).SetOwner(user),
		).Save(context.TODO()); err != nil {
			return errs.Wrap(err, errs.WithContext("files", []string{file1, file2}))
		}
		user, err = tx.User.Create().SetUsername("Bob").Save(context.TODO())
		if err != nil {
			return errs.Wrap(err, errs.WithContext("username", "Bob"))
		}
		if _, err := tx.BinaryFile.CreateBulk(
			tx.BinaryFile.Create().SetFilename(file3).SetBody(bin3).SetOwner(user),
		).Save(context.TODO()); err != nil {
			return errs.Wrap(err, errs.WithContext("files", []string{file3}))
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
