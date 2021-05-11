module web

go 1.15

require (
	github.com/asaskevich/govalidator v0.0.0-20210307081110-f21760c49a8d // indirect
	github.com/eko/gocache v1.1.1
	github.com/go-co-op/gocron v1.3.1
	github.com/go-ozzo/ozzo-validation v3.6.0+incompatible
	github.com/go-redis/redis/v7 v7.4.0
	github.com/golang/protobuf v1.4.3
	github.com/grpc-ecosystem/go-grpc-middleware v1.0.1-0.20190118093823-f849b5445de4
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/labstack/echo/v4 v4.2.1
	github.com/labstack/gommon v0.3.0
	github.com/mitchellh/go-homedir v1.1.0
	github.com/neko-neko/echo-logrus/v2 v2.0.1 // indirect
	github.com/pkg/errors v0.9.1
	github.com/pressly/goose v2.7.0+incompatible
	github.com/prometheus/client_golang v1.9.0
	github.com/spf13/cobra v1.1.3
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.7.0
	go.opentelemetry.io/contrib v0.19.0
	go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho v0.19.0
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.19.0
	go.opentelemetry.io/otel v0.19.0
	go.opentelemetry.io/otel/exporters/trace/jaeger v0.19.0
	go.opentelemetry.io/otel/oteltest v0.19.0
	go.opentelemetry.io/otel/sdk v0.19.0
	go.opentelemetry.io/otel/trace v0.19.0
	google.golang.org/grpc v1.36.0
	gorm.io/driver/mysql v1.0.5
	gorm.io/driver/sqlite v1.1.4
	gorm.io/gorm v1.21.3
)
