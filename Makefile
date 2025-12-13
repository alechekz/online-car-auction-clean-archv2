# Inspection Service
.PHONY: inspection-local-build, inspection-build, inspection-run, inspection-test, inspection-testcover, inspection-lint

inspection-local-build:
	@go build -o inspection-service -v ./services/inspection/cmd/main.go
	@echo "inspection service successfully built"

inspection-build:
	CGO_ENABLED=0 GOOS=linux go build -o /bin/inspection -v ./services/inspection/cmd/main.go
	@echo "inspection service successfully built"

inspection-run:
	@INSPECTION_HTTP=:7072 INSPECTION_GRPC=:7073 ./inspection-service

inspection-test:
	@go test -v ./services/inspection/...

inspection-testcover:
	@go test --cover ./services/inspection/... --coverprofile=testscoverprofile
	@go tool cover -html=testscoverprofile

inspection-lint:
	golangci-lint run ./services/inspection/...

inspection: inspection-lint inspection-test inspection-local-build inspection-run