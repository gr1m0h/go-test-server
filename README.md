# test-server

## Usage

### local

Run test-server

```sh
go run app/main.go
```

#### Docker

Build Docker image

```sh
make docker-build
```

#### Docker Compose

Run test-server with PostgreSQL database and Adminer interface

```sh
make docker-run
```

This will start:
- test-server on http://localhost:9090
- Adminer (database management interface) on http://localhost:8080
- PostgreSQL database on port 5432

Stop the services:

```sh
make docker-stop
```

#### Kubernetes

Prepare Docker Image.

```sh
make docker-build
```

Create K8s Cluster using [kind](https://github.com/kubernetes-sigs/kind) and load local Docker Image.

```sh
kind create cluster
kubectl cluster-info --context kind-kind
kind load docker-image test-server
```

##### Kubernetes Manifest(Raw)

Apply K8s manifest and connect to K8s service.

```sh
kubectl apply -f kubernetes/raw
kubectl port-forward service/test-server 8080:8080 -n test-server
```

##### Kubernetes Manifest(HelmChart)

##### Kubernetes Manifest(Helmfile)
