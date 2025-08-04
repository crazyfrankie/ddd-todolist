package main

import (
	"context"

	"github.com/crazyfrankie/ddd-todolist/backend/application"
)

func main() {
	ctx := context.Background()

	if err := application.Init(ctx); err != nil {
		panic("InitializeInfra failed, err=" + err.Error())
	}
}
