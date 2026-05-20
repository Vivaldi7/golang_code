package main

import (
	"context"
	"log"

	"github.com/vivaldi7/golang_code/mic_curce/week_3/internal/app"
)

func main() {

	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("failed to init app: %v", err.Error())
	}

	err = a.Run()
	if err != nil {
		log.Fatalf("failed to run app: %v", err.Error())
	}

}
