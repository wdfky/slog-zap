
```package main

import (
	"log/slog"

	"github.com/slogzap"
	"go.uber.org/zap"
)

func main() {
	zapL := zap.Must(zap.NewProduction())
	defer zapL.Sync()
	slog.SetDefault(slog.New(slogzap.NewHandler(zapL.Core(), nil)))

	ctx := slogzap.StartTrace()
	slog.InfoContext(ctx, "hello")
	slog.InfoContext(ctx, "world")
}
```