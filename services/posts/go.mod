module voidspace/posts

go 1.24.5

require (
	github.com/georgysavva/scany/v2 v2.1.4
	github.com/jackc/pgerrcode v0.0.0-20250907135507-afb5586c32a6
	github.com/jackc/pgx/v5 v5.8.0
	github.com/joho/godotenv v1.5.1
	github.com/vhysxl/voidspace/shared v0.0.0-00010101000000-000000000000
	go.uber.org/zap v1.27.1
	google.golang.org/grpc v1.74.2
	google.golang.org/protobuf v1.36.7
)

require (
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	github.com/lib/pq v1.10.9 // indirect
	go.opentelemetry.io/otel v1.37.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.37.0 // indirect
)

require (
	go.uber.org/multierr v1.10.0 // indirect
	golang.org/x/net v0.43.0 // indirect
	golang.org/x/sync v0.17.0 // indirect
	golang.org/x/sys v0.35.0 // indirect
	golang.org/x/text v0.29.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250811230008-5f3141c8851a // indirect
)

replace github.com/vhysxl/voidspace/shared => ../../shared
