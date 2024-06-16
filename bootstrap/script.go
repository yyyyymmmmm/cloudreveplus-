package bootstrap

import (
	"context"

	"github.com/yyyyymmmmm/Test/models/scripts/invoker"
	"github.com/yyyyymmmmm/Test/pkg/util"
)

func RunScript(name string) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if err := invoker.RunDBScript(name, ctx); err != nil {
		util.Log().Error("Failed to execute database script: %s", err)
		return
	}

	util.Log().Info("Finish executing database script %q.", name)
}
