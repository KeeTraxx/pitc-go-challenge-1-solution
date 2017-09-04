# pitc-go-challenge-1-solution

## How to deploy on Openshift

1. `oc login http://ose3-lab-master.puzzle.ch`
2. `oc new project pitc-go-workshop-$(whoami)-moviequote`
3. `oc new-app https://github.com/KeeTraxx/pitc-go-challenge-1-solution.git --name backend`
4. `oc expose svc/backend`