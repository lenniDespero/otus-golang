FROM golang
WORKDIR /app

## Copy go mod and sum files
COPY go.mod ./
COPY go.sum ./

## Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

## Copy the source from the current directory to the Working Directory inside the container
COPY . .

## Build the Go app
RUN go build -o calendar_api cmd/api/api.go

## Expose port 2000 to the outside world
EXPOSE 2000
## Command to run the executable
CMD ["./calendar_api", "--config=config/application.yml"]