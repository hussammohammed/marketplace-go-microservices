# Start from a Debian image with the latest version of Go installed
FROM golang:1.22.3

# Creates an app directory to hold app’s source code
#WORKDIR /github.com/hussammohammed/marketplace-go-microservices/gateway
WORKDIR /app
#COPY gateway/go.mod gateway/go.sum ./
 
# Download Go modules
COPY go.mod go.sum ./
RUN go mod download
COPY . .

#RUN go mod tidy
#RUN go mod vendor
# Copy the source code. Note the slash at the end, as explained in
#COPY . .

# Set the environment variable
ENV ENVNAME=development

# Build
#RUN CGO_ENABLED=0 GOOS=linux go build -o /gateway-api
# Builds  app with optional configuration
RUN go build -o gateway-api .
#ENTRYPOINT ["./gateway-api", "run", "--cfg-name", "$ENVNAME"]
ENTRYPOINT ./gateway-api run --cfg-name $ENVNAME
EXPOSE 3000


# install git
#RUN apk add --update --no-cache git
############ Copy the local package files to the container's workspace.
#WORKDIR /go/src/github.com/hussammohammed/marketplace-go-microservices/gateway
# Copy the local package files to the container's workspace
#COPY . .

#RUN go get -u github.com/golang/dep/cmd/dep
#RUN go mod tidy
#RUN go build
#ENTRYPOINT ./main run --run $ENVNAME
# CMD ["./reppy-api run"]
#EXPOSE 3000