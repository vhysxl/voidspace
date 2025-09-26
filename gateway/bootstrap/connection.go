package bootstrap

import (
	"crypto/tls"
	"crypto/x509"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	grpcInsecure "google.golang.org/grpc/credentials/insecure"
)

// NewConn creates a new gRPC connection.
// host should be of the form domain:port, e.g., example.com:443
func NewConn(host string, insecure bool) (*grpc.ClientConn, error) {

	var opts []grpc.DialOption

	if insecure {
		opts = append(opts, grpc.WithTransportCredentials(grpcInsecure.NewCredentials()))
	} else {
		systemRoots, err := x509.SystemCertPool()
		if err != nil {
			return nil, err
		}
		cred := credentials.NewTLS(&tls.Config{RootCAs: systemRoots})
		opts = append(opts, grpc.WithTransportCredentials(cred))
	}

	return grpc.NewClient(host, opts...)
}
