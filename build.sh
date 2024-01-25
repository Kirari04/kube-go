docker build -t kube-go:v1 . --load
docker save kube-go > kube-go_image.tar
microk8s ctr image import kube-go_image.tar