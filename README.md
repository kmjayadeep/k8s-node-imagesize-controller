# Kubernetes Controller example using client-go

A simple Kubernetes controller implemented in golang using client-go

It watches for nodes in the cluster and reports when the container image storage changes.

This project is inspired from the talk : <https://www.youtube.com/watch?v=QIMz4V9WxVc> and the corresponding repo <https://github.com/alena1108/kubecon2017>.

However, this implementation is using a more recent version of client-go and so there are some major changes in the implementation.

# Running outside cluster

```
❯ go run main.go --config ~/.kube/config
INFO[0000] starting controller                           config=/home/jayadeep/.kube/config
INFO[0009] Image size changed for node ip-10-0-102-198.ec2.internal. Old: [0 B], New: [14 GB]
INFO[0009] Image size changed for node ip-10-0-102-198.ec2.internal. Old: [14 GB], New: [15 GB]
```

# Running in cluster

The file `deployment.yaml` contains the yaml spec for running in a cluster. Change the namespace and image if necessary and apply the config


```
kubectl apply -f deployment.yaml
```

# Skaffold

If you are a fan of `Skaffold`, you can use it to build and run the image on the cluster without worrying about dockerfiles


Change the config in skaffold.yaml and run
```
skaffold run
``
