package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/juanbautista0/go-hexagonal-archetype/app/adapters/persistence/memory"
	"github.com/juanbautista0/go-hexagonal-archetype/app/domain/command"
	"github.com/juanbautista0/go-hexagonal-archetype/app/domain/ports"
	"github.com/juanbautista0/go-hexagonal-archetype/app/entrypoints/lambda/handler"
	"github.com/juanbautista0/go-hexagonal-archetype/app/libraries"
)

// Función adaptadora para convertir el tipo concreto a la interfaz
func createCommandHandler(r ports.VehicleRepository) libraries.CommandHandler {
	return command.NewCreateVehicleCommandHandler(r)
}

func main() {
	repository := memory.NewInMemoryVehicleRepository()
	handler := libraries.CreateHandler(
		createCommandHandler,
		handler.NewVehicleLambdaHandler,
		repository,
	)

	lambda.Start(handler.Handler)
}
