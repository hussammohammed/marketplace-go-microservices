### Project Description:
this project as an example for building marketplace application using microservice architecture with **Golang, GIN, CLI (cobera), Kafka, gRPC , JWT, Viper, Docker, Kubernetes, Helm Charts, GitOps**

#### system components
[Api gateway] reponsible for communication between UI and all microservices. it's secured via JWT token and midleware on each api call. it takes api request and call proper microservice then return response to UI<br/>

[User service] it's responsible for signup and handling authentication and authorization of users. it's called from api gateway via gRPC protocol <br/>

[order service] it's respobible for creating and updating order at system. it's called from api gateway via kafka messages<br/>


### Api gateway:

http://localhost:3000


For local development:

```
cd /gateway
go run ./gateway-api run --cfg-name=development

// to deploy to kubernetes
cd /gateway/deployments/helm
helm install api-gateway  api-gateway-chart // deploy first time
helm upgrade api-gateway  api-gateway-chart // to redeploy updated files
helm uninstall api-gateway  // to cancel any deployment

```
