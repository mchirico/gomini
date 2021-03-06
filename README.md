[![Build Status](https://dev.azure.com/mchirico/gomini/_apis/build/status/mchirico.gomini?branchName=master)](https://dev.azure.com/mchirico/gomini/_build/latest?definitionId=37&branchName=master)

[![codecov](https://codecov.io/gh/mchirico/gomini/branch/master/graph/badge.svg)](https://codecov.io/gh/mchirico/gomini)




# gomini

```
docker pull gcr.io/pigdevonlyx/gomini:test
docker run -p 3000:3000 -it --rm  gcr.io/pigdevonlyx/gomini:test
curl localhost:3000/data

```


```
minikube tunnel
# or minikube tunnel &> /dev/null &
k expose deployment gomini-pod --type=LoadBalancer --name=gomini-service
k expose deployment gomini-pv-pod --type=LoadBalancer --name=gomini-pv-service


# Dashboard
minikube dashboard

```



## Build with vendor
```
export GO111MODULE=on
go mod init
# Below will put all packages in a vendor folder
go mod vendor



go test -v -mod=vendor ./...

# Don't forget the "." in "./cmd/script" below
go build -v -mod=vendor ./...
```


## Don't forget golint

```
go get -u golang.org/x/lint/golint
golint -set_exit_status $(go list ./... | grep -v /vendor/)

```


