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
func NewHealthServiceClient(uri string, options ...HealthServiceClientOption) healthv1.HealthServiceClient {
	config := &HealthServiceClientConfig{
		Client:        http.DefaultClient,
		ClientURL:     uri,
		ClientOptions: []connect.ClientOption{},
	}

	// Apply the options
	for _, opt := range options {
		opt.Apply(config)
	}

	var interceptors []connect.Interceptor
	// prepare the interceptors
	interceptors = append(interceptors, interceptor.WithTracer())
	interceptors = append(interceptors, interceptor.WithLogger())
	// prepare the configuration
	config.ClientOptions = append(config.ClientOptions, connect.WithInterceptors(interceptors...))

	client := healthv1connect.NewHealthServiceClient(
		config.Client,
		config.ClientURL,
		config.ClientOptions...)

	return &HealthServiceClient{client: client}
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
	var interceptors []connect.Interceptor
	// prepare the interceptors
	interceptors = append(interceptors, interceptor.WithValidator())

	var options []connect.HandlerOption
	// prepare the options for interceptor collection
	options = append(options, connect.WithInterceptors(interceptors...))
	// prepare the options
	options = append(options, interceptor.WithRecover())

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
