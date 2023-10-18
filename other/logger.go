package main

import (
	"log/slog"
	"os"
)

func main() {
	jsonHandler := slog.NewJSONHandler(os.Stdout, nil)

	myLog := slog.New(jsonHandler)
	myLog.Info("fowefowejfo") // {"time":"2023-10-17T17:49:...","level":"INFO","msg":"fowefowejfo"}
	myLog.Error("doqwjodqw")  // {"time":"2023-10-17T18...","level":"ERROR","msg":"doqwjodqw"}
}
