# Online Car Auction - another clean architecture

# Inspection Service instructions

# Generate gRPC code
protoc -I=services/inspection/internal/transport/gRPC/proto \
  -I=$(go list -m -f '{{.Dir}}' github.com/grpc-ecosystem/grpc-gateway/v2) \
  -I=$(go list -m -f '{{.Dir}}' github.com/googleapis/googleapis) \
  --go_out=paths=source_relative:services/inspection/internal/transport/gRPC/proto \
  --go-grpc_out=paths=source_relative:services/inspection/internal/transport/gRPC/proto \
  --grpc-gateway_out=paths=source_relative:services/inspection/internal/transport/gRPC/proto \
  --openapiv2_out=services/inspection/internal/transport/gRPC/proto \
  services/inspection/internal/transport/gRPC/proto/inspection.proto

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