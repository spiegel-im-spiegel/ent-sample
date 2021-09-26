//go:build run
// +build run

package main

import (
	"context"
	"fmt"
	"os"
	"sample/dbconn"
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
	client := entCtx.GetClient()

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

	// create data
	user, err := client.User.Create().SetUsername("Alice").Save(context.TODO())
	if err != nil {
		entCtx.GetLogger().Error().Interface("error", errs.Wrap(err)).Send()
		return exitcode.Abnormal
	}
	if _, err := client.BinaryFile.CreateBulk(
		client.BinaryFile.Create().SetFilename(file1).SetBody(bin1).SetOwner(user),
		client.BinaryFile.Create().SetFilename(file2).SetBody(bin2).SetOwner(user),
	).Save(context.TODO()); err != nil {
		entCtx.GetLogger().Error().Interface("error", errs.Wrap(err)).Send()
		return exitcode.Abnormal
	}

	return exitcode.Normal
}

func main() {
	Run().Exit()
}
