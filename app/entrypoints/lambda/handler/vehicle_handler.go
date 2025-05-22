package handler

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/juanbautista0/go-hexagonal-archetype/app/domain/command"
	"github.com/juanbautista0/go-hexagonal-archetype/app/entrypoints/model"
	"github.com/juanbautista0/go-hexagonal-archetype/app/libraries"
)

// VehicleLambdaHandler adapta el CommandHandler al formato de Lambda
type VehicleLambdaHandler struct {
	cmdHandler libraries.CommandHandler
}

// NewVehicleLambdaHandler crea un nuevo handler de Lambda
func NewVehicleLambdaHandler(cmdHandler libraries.CommandHandler) libraries.LambdaHandler {
	return &VehicleLambdaHandler{cmdHandler: cmdHandler}
}

// Handler implementa la interfaz LambdaHandler
func (h *VehicleLambdaHandler) Handler(ctx context.Context, event interface{}) (interface{}, error) {
	apiEvent := event.(events.APIGatewayV2HTTPRequest)
	var req model.CreateVehicleRequest

	err := json.Unmarshal([]byte(apiEvent.Body), &req)
	if err != nil {
		return events.APIGatewayV2HTTPResponse{
			StatusCode: 400,
			Body:       "Invalid request body",
		}, nil
	}

	cmd := command.CreateVehicleCommand{
		Vehicle: req.Vehicle,
	}

	result, err := h.cmdHandler.Handle(ctx, cmd)
	if err != nil {
		return events.APIGatewayV2HTTPResponse{
			StatusCode: 500,
			Body:       "Error processing request",
		}, nil
	}

	success := result.(bool)
	responseBody, _ := json.Marshal(map[string]bool{"success": success})

	return events.APIGatewayV2HTTPResponse{
		StatusCode: 200,
		Body:       string(responseBody),
		Headers:    map[string]string{"Content-Type": "application/json"},
	}, nil
}
