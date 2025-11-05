# test-server

## Usage

### local

Run test-server

```sh
go run app/main.go
```

#### Docker

Run test-server

```sh
docker build -t test-server .
docker run -p 8080:8080 test-server
```

#### Kubernetes

Prepare Docker Image.

```
docker build -t test-server .
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
