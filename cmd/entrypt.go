package cmd

import (
	"context"
	"github.com/mchirico/gomini/point"
)

func EntryPt(ctx context.Context, file string, port string) {
	api := point.NewPointFile(file)
	api.MainListen(ctx, port)
}
