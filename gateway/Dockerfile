FROM golang:1.22.3
# Creates an app directory to hold app’s source code
WORKDIR /app

# Download Go modules
COPY go.mod ./
RUN go mod tidy
RUN go mod vendor
#RUN go mod download
COPY . .
# Set the environment variable
ENV ENVNAME=development

# Build
# Builds  app with optional configuration
#to debug errors at build step
RUN ls -l 
RUN go build -o gateway-api .
ENTRYPOINT ./gateway-api run --cfg-name $ENVNAME
EXPOSE 3000