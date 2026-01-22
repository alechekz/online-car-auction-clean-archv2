# Online Car Auction - another clean architecture


# Generate services gRPC code
buf generate
mv gen/pricing/v1/api/proto/* gen/pricing/v1/
rm -r gen/pricing/v1/api
rm gen/pricing/v1/inspection*
mv gen/inspection/v1/api/proto/* gen/inspection/v1/
rm -r gen/inspection/v1/api
rm gen/inspection/v1/pricing*
goimports -w gen/


# Inspection Service instructions

# gRPC calls
grpcurl -plaintext -d '{
  "vin": "5YJSA1E26MF168123",
  "year": 2020,
  "odometer": 15000
}' localhost:7073 inspection.v1.InspectionService/InspectVehicle

grpcurl -plaintext -d '{"vin":"5YJSA1E26MF168123"}' localhost:7073 inspection.v1.InspectionService/GetBuildData

# HTTP calls
curl -X POST -k http://localhost:7072/inspection/inspect -d '{
  "vin": "5YJSA1E26MF168123",
  "year": 2020,
  "odometer": 15000
}' -H "Content-Type: application/json"

curl -i http://localhost:7072/inspection/v1/get-build-data/5YJSA1E26MF168123


# Pricing Service instructions

# gRPC calls
grpcurl -plaintext -d '{
  "vin": "5YJSA1E26MF168123",
  "grade": 47,
  "odometer": 30000
}' localhost:7075 pricing.v1.PricingService/GetRecommendedPrice

# HTTP calls
curl -i -X POST http://localhost:7074/pricing/v1/get-recommended-price \
  -H "Content-Type: application/json" \
  -d '{"vin":"5YJSA1E26MF168123","grade":47,"odometer":30000}'