.PHONY: test-handlers test-config test-register-admin migrate-server run-server run-frontend build-docker run-docker

test-handlers:
	@echo "Testing handlers package..."
	cd backend && go test -v ./handlers

test-config:
	@echo "Testing config package..."
	cd backend && go test -v ./config

test-register-admin:
	@echo "Testing registerAdmin..."
	cd backend && go test -v ./migration

migrate-server:
	@echo "Creating admin user and initializing the database..."
	cd backend && ADMIN_USERNAME=admin ADMIN_PASSWORD=admin123 go run -mod=mod ./cmd/migrate/main.go

run-server:
	@echo "Starting the server..."
	cd backend && MONGO_CONNECTION_STRING=mongodb://localhost:27017 MONGO_DATABASE_NAME=loginapp go run main.go

run-frontend:
	@echo "Starting the React development server..."
	cd frontend && npm install && npm start

build-docker:
	@echo "Building Docker image..."
	cd backend && docker build -t mongodb .

run-docker:
	@echo "Running MongoDB Docker container..."
	docker run -d --name mongodb -p 27017:27017 mongodb
