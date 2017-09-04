# pitc-go-challenge-1-solution

## How to deploy on Openshift
1. `oc login https://ose3-lab-master.puzzle.ch:8443`
2. (optional) `oc delete project pitc-go-workshop-$(whoami)-moviequote`
3. `oc new-project pitc-go-workshop-$(whoami)-moviequote`
4. `oc new-app https://github.com/KeeTraxx/pitc-go-challenge-1-solution.git --name backend`
5. `oc expose svc/backend`