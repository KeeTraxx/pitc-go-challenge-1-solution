# pitc-go-challenge-1-solution

## How to run locally
1. `go get -v ./...`
2. `go build .`
3. `./pitc-go-challenge-1-solution`

Oneliner:
`go get -v ./... && go build . && ./$(basename $(pwd))`

## How to run in a docker container
`docker run -ti -w=/go/src/app -p 1323:1323 -v $(pwd):/go/src/app golang:1.8 go get -v ./... && go build . && ./app`

## How to deploy on Openshift
1. `oc login https://ose3-lab-master.puzzle.ch:8443`
2. (optional) `oc delete project pitc-go-workshop-$(whoami)-moviequote`
3. `oc new-project pitc-go-workshop-$(whoami)-moviequote`
4. `oc new-app https://github.com/KeeTraxx/pitc-go-challenge-1-solution.git --name backend`
5. `oc expose svc/backend`
6. `oc get routes` - shows the auto generated URL
