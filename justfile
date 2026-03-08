# Run the application
run: all
    go run ./cmd/server

# build all
all: frontend server test_client

# Build the application for the pi
pi: frontend
    mkdir -p output/web
    cp -r web/build output/web/build
    GOOS=linux GOARCH=arm GOARM=5 go build ./cmd/server -o output/nameplate

# Build just the go stuff
go: server test_client

# Build the server
server:
    go build ./cmd/server

# Build the test client
test_client:
    go build ./cmd/test-client

# Build the frontend
frontend:
    (cd ./frontend && npm run build)

# Clean the project
clean:
    rm -f server
    rm -f test-client
    (cd ./frontend && rm -rf build)

