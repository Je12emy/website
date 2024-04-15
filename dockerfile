FROM golang:1.22.2
WORKDIR app

# Keeping this commented out until I add dependencies
# COPY go.mod go.sum .
# RUN go mod download

# Copy all the source code
COPY . ./

# Generate a static binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /blog

CMD ["/blog"]
