FROM golang:1.22

# set working directory
WORKDIR /app

# Copy the source code
COPY . . 

# Download and install the dependencies
# RUN go get -d -v ./...
COPY go.mod go.sum ./
RUN go mod download

# # Build the Go app
# RUN go build -o api .

# #EXPOSE the port
# EXPOSE 8000

# # Run the executable
# CMD ["./api"]