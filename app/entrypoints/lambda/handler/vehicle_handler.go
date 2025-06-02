package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/juanbautista0/go-hexagonal-archetype/app/domain/command"
	"github.com/juanbautista0/go-hexagonal-archetype/app/domain/entity"
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
	var req model.CreateVehicleRequest
	var apiEvent events.APIGatewayV2HTTPRequest
	var ok bool

	apiEvent, ok = event.(events.APIGatewayV2HTTPRequest)
	if !ok {
		h.logger(http.StatusBadRequest).
			SetCode("HANDLER").
			SetDetail("Invalid event: not APIGatewayV2HTTPRequest").
			SetMessage(http.StatusText(http.StatusBadRequest)).
			SetMetadata(map[string]interface{}{"event": apiEvent}).
			Write()

		return h.Response(http.StatusBadRequest, nil), nil
	}

	err := json.Unmarshal([]byte(apiEvent.Body), &req)
	if err != nil {
		h.logger(http.StatusBadRequest).
			SetCode("HANDLER").
			SetDetail("Invalid request body").
			SetMessage(http.StatusText(http.StatusBadRequest)).
			SetMetadata(map[string]interface{}{"error": err.Error(), "request_body": apiEvent.Body}).
			Write()

		return h.Response(http.StatusBadRequest, nil), nil
	}

	vehicle, err := entity.NewVehicleFromPrimitives(req.Id.String(), req.Brand, req.Model, req.Year)
	if err != nil {
		h.logger(http.StatusBadRequest).
			SetCode("HANDLER").
			SetDetail("Invalid request body").
			SetMessage(http.StatusText(http.StatusBadRequest)).
			SetMetadata(map[string]interface{}{"error": err.Error()}).
			Write()

		return h.Response(http.StatusBadRequest, nil), nil
	}

	cmd := &command.CreateVehicleCommand{Vehicle: vehicle}

	result, err := h.cmdHandler.Handle(ctx, cmd)
	if err != nil {
		statusCode, detail := MapErrorToHTTPResponse(err)
		h.logger(statusCode).
			SetCode("HANDLER").
			SetDetail("Exception").
			SetMessage(http.StatusText(statusCode)).
			SetMetadata(map[string]interface{}{"error": detail}).
			Write()

		return h.Response(statusCode, nil), nil

	}

	h.logger(http.StatusOK).
		SetCode("HANDLER").
		SetDetail("Successful operation").
		SetMessage(http.StatusText(http.StatusOK)).
		SetMetadata(map[string]interface{}{"result": result}).
		Write()

	return h.Response(http.StatusOK, result), nil
}

func (h *VehicleLambdaHandler) Response(httpStatusCode int, body interface{}) events.APIGatewayV2HTTPResponse {

	responseDto := libraries.LambdaDtoResponse{
		Message:    strings.ReplaceAll(strings.ToUpper(http.StatusText(httpStatusCode)), " ", "_"),
		StatusCode: httpStatusCode,
		Data:       body,
	}
	responseBody, _ := json.Marshal(responseDto)

	return events.APIGatewayV2HTTPResponse{
		Headers:    map[string]string{"Content-Type": "application/json; charset=utf-8"},
		StatusCode: httpStatusCode,
		Body:       string(responseBody),
	}

}

func (h *VehicleLambdaHandler) logger(httpStatusCode int) *libraries.LoggerfyBase {
	var log *libraries.Loggerfy = libraries.NewLoggerfy()

	switch true {
	case httpStatusCode >= 200 && httpStatusCode < 300:
		return log.Info()
	case httpStatusCode >= 400 && httpStatusCode < 500:
		return log.Warn()
	default:
		return log.Error()
	}

}
