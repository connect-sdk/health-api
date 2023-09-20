package healthv1sdk

import (
	context "context"
	slog "log/slog"
	http "net/http"

	connect "connectrpc.com/connect"
	idtoken "google.golang.org/api/idtoken"
)

// HealthServiceClientConfig represents the config for google.health.v1.HealthServiceClient client.
type HealthServiceClientConfig struct {
	Client        *http.Client
	ClientURL     string
	ClientOptions []connect.ClientOption
}

// HealthServiceClientOption represents an option for google.health.v1.HealthServiceClient client.
type HealthServiceClientOption interface {
	// Apply applies the option.
	Apply(config *HealthServiceClientConfig)
}

var _ HealthServiceClientOption = HealthServiceClientOptionFunc(nil)

// HealthServiceClientOptionFunc represent a function that implementes google.health.v1.HealthServiceClientOption option.
type HealthServiceClientOptionFunc func(*HealthServiceClientConfig)

// Apply applies the option.
func (fn HealthServiceClientOptionFunc) Apply(config *HealthServiceClientConfig) {
	fn(config)
}

// HealthServiceClientWithAuthorization enables an oidc authorization.
func HealthServiceClientWithAuthorization() HealthServiceClientOption {
	fn := func(config *HealthServiceClientConfig) {
		// client uri
		uri := config.ClientURL
		// prepare the options
		options := []idtoken.ClientOption{}
		options = append(options, idtoken.WithHTTPClient(config.Client))
		// prepare the client
		client, err := idtoken.NewClient(context.Background(), uri, options...)
		if err != nil {
			slog.Error("unable to create an id token", err)
		}
		// set the client
		config.Client = client
	}

	return HealthServiceClientOptionFunc(fn)
}

// HealthServiceClientWithProtocol enables a given protocol.
func HealthServiceClientWithProtocol(name string) HealthServiceClientOption {
	fn := func(config *HealthServiceClientConfig) {
		// prepare the protocol
		switch name {
		case "grpc":
			config.ClientOptions = append(config.ClientOptions, connect.WithGRPC())
		case "grpc+web":
			config.ClientOptions = append(config.ClientOptions, connect.WithGRPCWeb())
		}
	}

	return HealthServiceClientOptionFunc(fn)
}

// HealthServiceClientWithCodec enables a given codec.
func HealthServiceClientWithCodec(name string) HealthServiceClientOption {
	fn := func(config *HealthServiceClientConfig) {
		// prepare the protocol
		switch name {
		case "json":
			config.ClientOptions = append(config.ClientOptions, connect.WithProtoJSON())
		}
	}

	return HealthServiceClientOptionFunc(fn)
}
