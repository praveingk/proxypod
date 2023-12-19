# Proxy pod

## Build
  make build
## Run
  proxypod-k8s --name testapp --port 5000 --target testservice:5000

  This will create a deployment named testapp-xxxx which will be listening to port 5000, and redirect the traffic to an imported testservice:5000.

  
  
