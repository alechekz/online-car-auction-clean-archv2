Online Car Auction - another clean architecture


# Generate gRPC code
protoc -I=services/inspection/internal/transport/gRPC/v1 \
  --go_out=paths=source_relative:services/inspection/internal/transport/gRPC/v1 \
  --go-grpc_out=paths=source_relative:services/inspection/internal/transport/gRPC/v1 \
  services/inspection/internal/transport/gRPC/v1/inspection.proto

protoc -I=services/pricing/internal/transport/gRPC/v1 \
  --go_out=paths=source_relative:services/pricing/internal/transport/gRPC/v1 \
  --go-grpc_out=paths=source_relative:services/pricing/internal/transport/gRPC/v1 \
  services/pricing/internal/transport/gRPC/v1/pricing.proto
