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
}' localhost:8083 inspection.v1.InspectionService/InspectVehicle

grpcurl -plaintext -d '{"vin":"5YJSA1E26MF168123"}' localhost:8083 inspection.v1.InspectionService/GetBuildData

# HTTP calls
curl -X POST -k http://localhost:8082/inspection/inspect -d '{
  "vin": "5YJSA1E26MF168123",
  "year": 2020,
  "odometer": 15000
}' -H "Content-Type: application/json"

curl -i http://localhost:8082/inspection/v1/get-build-data/5YJSA1E26MF168123


# Pricing Service instructions

# gRPC calls
grpcurl -plaintext -d '{
  "vin": "5YJSA1E26MF168123",
  "grade": 47,
  "odometer": 30000
}' localhost:8085 pricing.v1.PricingService/GetRecommendedPrice

# HTTP calls
curl -i -X POST http://localhost:8084/pricing/v1/get-recommended-price \
  -H "Content-Type: application/json" \
  -d '{"vin":"5YJSA1E26MF168123","grade":47,"odometer":30000}'


# Vehicle Service instructions

# HTTP calls
curl -i -X POST http://localhost:8081/vehicles \
  -H "Content-Type: application/json" \
  -d '{"vin":"5YJSA1E26MF168123","year":2022,"odometer":12000}'

curl -i -X PUT http://localhost:8081/vehicles/5YJSA1E26MF168123 \
  -H "Content-Type: application/json" \
  -d '{
    "vin":"5YJSA1E26MF168123",
    "year":1999,
    "odometer":125000,
    "exteriorColor":"Red",
    "interiorColor":"Black"
  }'

curl -i http://localhost:8081/vehicles/5YJSA1E26MF168123

curl -i http://localhost:8081/vehicles

curl -i -X DELETE http://localhost:8081/vehicles/5YJSA1E26MF168123

Resume:
 1. По ходу дела поправил названия файлов и методов к ним, следуя твоим рекомендациям, чтобы не было дублирования понятий
 2. Использовал buf для генерации gRPC кода, все получилось. Buf как бы заставляет изменить структуру проекта и завести папку gen.
 3. В inspection и pricing сервисах у меня основным является gRPC, а HTTP - это просто обертка над ним сгенерированная через google/api/anntations.proto. В vehicle сервисе наоборот, основным является HTTP и его я поднял на Gin, просто чтобы поработать с ним, познакомиться
 4. В vehicle еще более структурировал тесты, появилась папка testhelpers
 5. В vehicle попробовал реализовать graceful shutdown, но так и не смог разобраться как его корректно реализовать и не нарушить принципы чистой архитектуры. В частности вопрос как завершить gRPC клиенты, которые создаются в сервере, а используются в сервисе(inspection.Close(), pricing.Close()). Буду благодарен, если подскажешь, как это правильно делается.
 6. Контекст прокинул, где-то общий контекст сервера, где-то контекст запроса
