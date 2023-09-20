package healthv1

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
)

//go:generate counterfeiter -generate

//counterfeiter:generate -o ./healthv1fake . HealthService

// HealthService is an implementation of the grpc.health.v1.HealthService service.
type HealthService interface {
	// If the requested service is unknown, the call will fail with status
	// NOT_FOUND.
	Check(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error)
}

var _ HealthService = &HealthServiceRegistry{}

// HealthServiceRegistry represents a map of grpc.health.v1.HealthService service.
type HealthServiceRegistry map[string]HealthService

// Check checks the health of a given service.
func (x *HealthServiceRegistry) Check(ctx context.Context, r *HealthCheckRequest) (*HealthCheckResponse, error) {
	if service, ok := (*x)[r.Service]; ok {
		return service.Check(ctx, r)
	}

	return nil, connect.NewError(connect.CodeNotFound, fmt.Errorf("not found a registered instance of grpc.health.v1.HealthService for %v", r.Service))
}

//counterfeiter:generate -o ./healthv1fake . HealthServiceClient

// HealthServiceClient is an implementation of the grpc.health.v1.HealthServiceClient client.
type HealthServiceClient interface {
	// If the requested service is unknown, the call will fail with status
	// NOT_FOUND.
	Check(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error)
}

var _ HealthServiceClient = &NopHealthServiceClient{}

// NopHealthServiceClient is a no-op implementation of the grpc.health.v1.HealthServiceClient client.
type NopHealthServiceClient struct{}

// Check implements HealthServiceClient.
func (*NopHealthServiceClient) Check(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error) {
	return &HealthCheckResponse{}, nil
}
