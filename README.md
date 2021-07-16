# Kubernetes Controller example using client-go

A simple Kubernetes controller implemented in golang using client-go

It watches for nodes in the cluster and reports when the container image storage changes.

This project is inspired from the talk : <https://www.youtube.com/watch?v=QIMz4V9WxVc> and the corresponding repo <https://github.com/alena1108/kubecon2017>.

However, this implementation is using a more recent version of client-go and so there are some major changes in the implementation.
