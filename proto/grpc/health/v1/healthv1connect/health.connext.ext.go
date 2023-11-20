package healthv1connect

type (
	// HealthServiceClient is a client for the grpc.health.v1.Health service.
	HealthServiceClient HealthClient

	// HealthServiceHandler is an implementation of the grpc.health.v1.Health service.
	HealthServiceHandler HealthHandler
)

var (
	// NewHealthServiceClient constructs a client for the grpc.health.v1.Health service. By default, it uses
	// the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
	// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
	// connect.WithGRPCWeb() options.
	//
	// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
	// http://api.acme.com or https://acme.com/grpc).
	NewHealthServiceClient = NewHealthClient

	// NewHealthServiceHandler builds an HTTP handler from the service implementation. It returns the path on
	// which to mount the handler and the handler itself.
	//
	// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
	// and JSON codecs. They also support gzip compression.
	NewHealthServiceHandler = NewHealthHandler
)
