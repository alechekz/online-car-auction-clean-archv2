# Online Car Auction - another clean architecture


# Inspection Service instructions


# Generate Inspection gRPC code
protoc -I=services/inspection/api/proto/ \
  -I=$(go list -m -f '{{.Dir}}' github.com/grpc-ecosystem/grpc-gateway/v2) \
  -I=$(go list -m -f '{{.Dir}}' github.com/googleapis/googleapis) \
  --go_out=paths=source_relative:services/inspection/internal/transport/gRPC/proto \
  --go-grpc_out=paths=source_relative:services/inspection/internal/transport/gRPC/proto \
  --grpc-gateway_out=paths=source_relative:services/inspection/internal/transport/gRPC/proto \
  --openapiv2_out=services/inspection/internal/transport/gRPC/proto \
  services/inspection/api/proto/inspection.proto

# gRPC calls
grpcurl -plaintext -d '{
  "vin": "5YJSA1E26MF168123",
  "year": 2020,
  "odometer": 15000
}' localhost:7073 inspection.InspectionService/InspectVehicle

grpcurl -plaintext -d '{"vin":"5YJSA1E26MF168123"}' localhost:7073 inspection.InspectionService/GetBuildData

# HTTP calls
curl -X POST -k http://localhost:7072/inspection/inspect -d '{
  "vin": "5YJSA1E26MF168123",
  "year": 2020,
  "odometer": 15000
}' -H "Content-Type: application/json"

curl -i http://localhost:7072/inspection/get-build-data/5YJSA1E26MF168123


# Pricing Service instructions

# Generate Light Inspection gRPC client for Pricing Service
protoc -I=services/inspection/api/proto \
  --go_out=paths=source_relative:services/pricing/internal/provider/inspectionclient \
  --go-grpc_out=paths=source_relative:services/pricing/internal/provider/inspectionclient \
  services/inspection/api/proto/inspection_get_build_data.proto

# Generate Pricing gRPC code
protoc -I=services/pricing/api/proto/ \
  -I=$(go list -m -f '{{.Dir}}' github.com/grpc-ecosystem/grpc-gateway/v2) \
  -I=$(go list -m -f '{{.Dir}}' github.com/googleapis/googleapis) \
  --go_out=paths=source_relative:services/pricing/internal/transport/gRPC/proto \
  --go-grpc_out=paths=source_relative:services/pricing/internal/transport/gRPC/proto \
  --grpc-gateway_out=paths=source_relative:services/pricing/internal/transport/gRPC/proto \
  --openapiv2_out=services/pricing/internal/transport/gRPC/proto \
  services/pricing/api/proto/pricing.proto