# relayr-app
A cloud native app run in minikube

[![Go Report Card](https://goreportcard.com/badge/github.com/andrleite/relayr-app)](https://goreportcard.com/report/github.com/andrleite/relayr-app)
[![codecov](https://codecov.io/gh/andrleite/relayr-app/branch/master/graph/badge.svg)](https://codecov.io/gh/andrleite/relayr-app)
[![Build Status](https://travis-ci.org/andrleite/relayr-app.svg?branch=master)](https://travis-ci.org/andrleite/relayr-app)

# Getting Started
## Pre Requisites
- [minikube v1.2+](https://kubernetes.io/docs/tasks/tools/install-minikube/)
- [kubectl v1.12+](https://kubernetes.io/docs/tasks/tools/install-kubectl/)
- [helm v2.13+](https://helm.sh/docs/using_helm/#installing-helm)

## Starting minikube and enable addons
```bash
minikube start
minikube addons enable ingress
minikube addons enable metrics-server
```
## Creating and deploy ssl certificates
```bash
openssl req -x509 -newkey rsa:4096 -sha256 \
            -nodes -keyout hack/certs/tls.key \
            -out hack/certs/tls.crt \
            -subj "/CN=relayr.app" \
            -days 365
```
- Now with ssl certificates created, let's deploy it on minikube
- First of all we'll create namespace.
```bash
kubectl create ns relayr
```
- Now we can create a secret to store tls certificates
```bash
kubectl --namespace relayr create secret tls relayr-tls \
        --cert=hack/certs/tls.crt \
        --key=hack/certs/tls.key
```
## Initializing helm and update dependencies
```bash
helm init
helm dep update helm/relayr-app
```
## Deploy the realyr app
```bash
helm install -n relayr \
  --namespace relayr helm/relayr-app/
```
- setting app domain to etc/hosts
```bash
echo "$(minikube ip) relayr.app" | sudo tee -a /etc/hosts
```
## API
#### /sensors
* `GET` : Get all iot sensors
* `POST` : Create a new iot sensor

#### /sensors/:id
* `GET` : Get a iot sensor
* `PUT` : Update a iot sensor
* `DELETE` : Delete a iot sensor

#### examples:

- create new sensor
```bash
curl  --cacert hack/certs/tls.crt \
      -X POST \
      -H "Content-Type: application/json" \
      -d '{"name": "sensorX", "type": "temperature"}' \
      https://relayr.app/sensors
```
- get all sensors
```bash
curl  --cacert hack/certs/tls.crt \
      https://relayr.app/sensors | jq .
```
- get sensor by id
```bash
curl  --cacert hack/certs/tls.crt \
      https://relayr.app/sensors/1 | jq .
```
- update sensor
```bash
curl  --cacert hack/certs/tls.crt \
      -X PUT \
      -d '{"type": "pressure"}' \
      -H "Content-Type: application/json" \
      https://relayr.app/sensors/1
```
- delete sensor
```bash
curl  --cacert hack/certs/tls.crt \
      -X DELETE \
      -H "Content-Type: application/json" \
      https://relayr.app/sensors/1
```
## Tear Down
```bash
helm del --purge relayr
minikube stop
minikube delete
```