package main

import (
	"github.com/GabrielLoureiroGomes/basket-collection/cmd"
	"github.com/GabrielLoureiroGomes/basket-collection/logger"
	"go.uber.org/zap/zapcore"
)

var log = logger.GetLogger()

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatal("error to initialize app", zapcore.Field{Type: zapcore.StringType, String: err.Error()})
	}
}
