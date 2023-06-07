BINARY_NAME=cp-db

build:
	@echo "Building back end..."
	go build -o ${BINARY_NAME} ./cmd/web/
	@echo "Binary built!"

run: build
	@echo "Starting back end..."
	@env ./${BINARY_NAME} &
	@echo "Back end started!"

clean:
	@echo "Cleaning..."
	@go clean
	@rm ${BINARY_NAME}
	@echo "Cleaned!"

start: run

stop:
	@echo "Stopping back end..."
	@-pkill -SIGTERM -f "./${BINARY_NAME}"
	@echo "Stopped back end!"

restart: stop start