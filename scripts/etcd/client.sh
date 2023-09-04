docker run -it --rm --network app-tier --env ALLOW_NONE_AUTHENTICATION=yes   bitnami/etcd:latest etcdctl --endpoints http://123.57.251.188:2379
