//go:build run
// +build run

package main

import (
	"context"
	"fmt"
	"os"
	"sample/dbconn"

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

	// output DDL
	// if err := entCtx.GetClient().Schema.WriteTo(context.TODO(), os.Stdout, schema.WithGlobalUniqueID(true)); err != nil {
	if err := entCtx.GetClient().Schema.WriteTo(context.TODO(), os.Stdout); err != nil {
		entCtx.GetLogger().Error().Interface("error", errs.Wrap(err)).Send()
		return exitcode.Abnormal
	}

	return exitcode.Normal
}

func main() {
	Run().Exit()
}
