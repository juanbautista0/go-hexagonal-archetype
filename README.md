# Go Hexagonal Architecture for AWS Lambda

This project demonstrates a hexagonal architecture (ports and adapters) implementation for AWS Lambda functions using Go. The architecture ensures a clear separation of concerns, making the codebase more maintainable, testable, and adaptable to changes.

> **Note**: This implementation is inspired by the official AWS guide on hexagonal architectures:
> https://docs.aws.amazon.com/pdfs/prescriptive-guidance/latest/hexagonal-architectures/hexagonal-architectures.pdf

## System Requirements

To work with this project, you'll need:

- **Docker**: Required for local testing and containerization
- **AWS SAM CLI**: For local development and deployment to AWS
- **Go**: Version 1.x or higher
- **Git**: For version control

## Architecture Overview

```
┌─────────────────────────────────────────────────────────────────────────┐
│                                                                         │
│  ┌─────────────┐     ┌───────────────┐     ┌───────────────────────┐    │
│  │             │     │               │     │                       │    │
│  │  API        │     │  Lambda       │     │  Domain               │    │
│  │  Gateway    │────▶│  Handler      │────▶│  Command Handler      │    │
│  │             │     │  (Adapter)    │     │                       │    │
│  └─────────────┘     └───────────────┘     └───────────────┬───────┘    │
│                                                            │            │
│                                                            │            │
│                                                            ▼            │
│                                                  ┌───────────────────┐  │
│                                                  │                   │  │
│                                                  │  Repository       │  │
│                                                  │  (Port)           │  │
│                                                  │                   │  │
│                                                  └─────────┬─────────┘  │
│                                                            │            │
│                                                            │            │
│                                                            ▼            │
│                                                  ┌───────────────────┐  │
│                                                  │                   │  │
│                                                  │  Repository       │  │
│                                                  │  Implementation   │  │
│                                                  │  (Adapter)        │  │
│                                                  │                   │  │
│                                                  └───────────────────┘  │
│                                                                         │
└─────────────────────────────────────────────────────────────────────────┘
```

## Request Flow

1. **API Gateway** receives an HTTP request
2. The request is forwarded to the **Lambda Handler** (adapter)
3. The handler deserializes the request and creates a **Command**
4. The command is passed to the **Command Handler** in the domain
5. The command handler uses the **Repository Port** to interact with data
6. The **Repository Implementation** (adapter) performs the actual data operations
7. The result flows back through the layers
8. The Lambda Handler formats and returns the HTTP response

## Project Structure

```
app/
├── adapters/                 # Implementations of the ports (adapters)
│   └── vehicle_repository/   # Repository implementations
│       └── memory_repository.go
├── domain/                   # Business logic and domain models
│   ├── command/              # Command pattern implementation
│   │   ├── command.go        # Command definitions
│   │   └── command_handler.go # Command handlers
│   ├── model/                # Domain models
│   │   └── vehicle.go
│   └── ports/                # Interface definitions (ports)
│       └── vehicle_repository.go
├── entrypoints/              # Entry points to the application
│   ├── lambda/               # Lambda function handlers
│   │   └── main.go
│   └── model/                # Request/response models
│       └── vehicle_request.go
└── libraries/                # Shared utilities and helpers
    └── lambda_instance_builder.go
```

## Core Components

### Ports (Interfaces)

Ports define the contracts between the domain and the outside world:

```go
// app/domain/ports/vehicle_repository.go
type VehicleRepository interface {
    Save(ctx context.Context, vehicle *model.Vehicle) (*model.Vehicle, error)
}
```

### Domain Models

Domain models represent the business entities:

```go
// app/domain/model/vehicle.go
type Vehicle struct {
    Id    uuid.UUID `json:"id"`
    Brand string    `json:"brand"`
    Model string    `json:"model"`
    Year  int       `json:"year"`
}
```

### Commands

Commands represent actions to be performed:

```go
// app/domain/command/command.go
type CreateVehicleCommand struct {
    model.Vehicle
}
```

### Command Handlers

Command handlers contain the business logic:

```go
// app/domain/command/command_handler.go
type CreateVehicleCommandHandler struct {
    repository ports.VehicleRepository
}

func (h *CreateVehicleCommandHandler) Execute(ctx context.Context, command CreateVehicleCommand) (bool, error) {
    if _, err := h.repository.Save(ctx, &command.Vehicle); err != nil {
        return false, err
    }
    return true, nil
}
```

### Adapters

Adapters implement the ports and connect to external systems:

```go
// app/adapters/vehicle_repository/memory_repository.go
type MemoryRepository struct {
    vehicles map[string]*model.Vehicle
}

func (r *MemoryRepository) Save(ctx context.Context, vehicle *model.Vehicle) (*model.Vehicle, error) {
    r.vehicles[vehicle.Id.String()] = vehicle
    return vehicle, nil
}
```

### Lambda Handler

The Lambda handler adapts between AWS Lambda and the domain:

```go
// app/entrypoints/lambda/main.go
type VehicleLambdaHandler struct {
    cmdHandler libraries.CommandHandler
}

func (h *VehicleLambdaHandler) Handler(ctx context.Context, event interface{}) (interface{}, error) {
    // Convert API Gateway event to domain command
    // Execute command
    // Format response
}
```

## Implementation Guide

### 1. Define Domain Models

Create your domain models in `app/domain/model/`.

### 2. Define Ports (Interfaces)

Create interfaces in `app/domain/ports/` that define how the domain interacts with external systems.

### 3. Create Commands

Define commands in `app/domain/command/command.go` that represent actions in your domain.

### 4. Implement Command Handlers

Create command handlers in `app/domain/command/command_handler.go` that contain your business logic.

### 5. Implement Adapters

Create adapters in `app/adapters/` that implement the ports defined in step 2.

### 6. Create Lambda Handlers

Implement Lambda handlers in `app/entrypoints/lambda/` that adapt between AWS Lambda and your domain.

### 7. Wire Everything Together

Use dependency injection to connect all components in your `main.go` file.

## Deployment

This project uses AWS SAM for deployment:

```bash
# Build the application
make build

# Test locally
make local-invoke

# Deploy to AWS
make deploy
```

## Benefits of This Architecture

1. **Separation of Concerns**: Business logic is isolated from external dependencies
2. **Testability**: Components can be tested in isolation
3. **Flexibility**: Implementations can be swapped without changing the domain logic
4. **Maintainability**: Clear boundaries make the codebase easier to understand
5. **Adaptability**: New interfaces can be added without disrupting existing code

## License

MIT