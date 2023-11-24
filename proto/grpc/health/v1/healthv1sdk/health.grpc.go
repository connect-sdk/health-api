package healthv1sdk

import (
	"context"
	http "net/http"

	connect "connectrpc.com/connect"
	interceptor "github.com/connect-sdk/interceptor"
	middleware "github.com/connect-sdk/middleware"
	chi "github.com/go-chi/chi/v5"

	healthv1 "github.com/connect-sdk/health-api/proto/grpc/health/v1"
	healthv1connect "github.com/connect-sdk/health-api/proto/grpc/health/v1/healthv1connect"
)

var _ healthv1.HealthServiceClient = &HealthServiceClient{}

// HealthServiceClient is a client for the grpc.health.v1.HealthService service.
type HealthServiceClient struct {
	client healthv1connect.HealthServiceClient
}

// NewHealthServiceClient creates a new connect.runtime.v1.healthv1.HealthServiceClient client.
func NewHealthServiceClient(uri string, options ...connect.ClientOption) healthv1.HealthServiceClient {
	// prepare the options
	options = append(options, interceptor.WithContext())
	options = append(options, interceptor.WithTracer())
	options = append(options, interceptor.WithLogger())
	// prepare the client
	client := &HealthServiceClient{
		client: healthv1connect.NewHealthServiceClient(http.DefaultClient, uri, options...),
	}

	return client
}

// Check implements healthv1.HealthServiceClient.
func (x *HealthServiceClient) Check(ctx context.Context, r *healthv1.HealthCheckRequest) (*healthv1.HealthCheckResponse, error) {
	response, err := x.client.Check(ctx, connect.NewRequest(r))
	if err != nil {
		return nil, err
	}

	return response.Msg, nil
}

var _ healthv1connect.HealthServiceHandler = &HealthServiceHandler{}

// HealthServiceHandler represents a controller for grpc.health.v1.HealthServiceHandler handler.
type HealthServiceHandler struct {
	// HealthService contains an instance of grpc.health.v1.HealthService service.
	HealthService healthv1.HealthService
}

// Mount mounts the handler to a given router.
func (x *HealthServiceHandler) Mount(r chi.Router) {
	var options []connect.HandlerOption
	// prepare the options
	options = append(options, interceptor.WithContext())
	options = append(options, interceptor.WithValidator())
	options = append(options, interceptor.WithRecovery())

	r.Group(func(r chi.Router) {
		// mount the middleware
		r.Use(middleware.WithLogger())
		// create the handler
		path, handler := healthv1connect.NewHealthServiceHandler(x, options...)
		// mount the handler
		r.Mount(path, handler)
	})
}

// Check implements HealthServiceHandler.
func (x *HealthServiceHandler) Check(ctx context.Context, r *connect.Request[healthv1.HealthCheckRequest]) (*connect.Response[healthv1.HealthCheckResponse], error) {
	response, err := x.HealthService.Check(ctx, r.Msg)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(response), nil
}

// Watch implements HealthServiceHandler.
func (x *HealthServiceHandler) Watch(ctx context.Context, r *connect.Request[healthv1.HealthCheckRequest], s *connect.ServerStream[healthv1.HealthCheckResponse]) error {
	response, err := x.HealthService.Check(ctx, r.Msg)
	if err != nil {
		return err
	}

	return s.Send(response)
}
