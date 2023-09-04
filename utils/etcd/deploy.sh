docker pull bitnami/etcd:latest
docker run -d --name etcd-server     --network app-tier     --publish 2379:2379     --publish 2380:2380     --env ALLOW_NONE_AUTHENTICATION=yes     --env ETCD_ADVERTISE_CLIENT_URLS=http://123.57.251.188
docker run -d --name etcd-server     --network app-tier     --publish 2379:2379     --publish 2380:2380     --env ALLOW_NONE_AUTHENTICATION=yes     --env ETCD_ADVERTISE_CLIENT_URLS=http://123.57.251.188 bitnami/etcd:latest
docker rm /etcd-server
docker ps -a
docker run -d --name etcd-server     --network app-tier     --publish 2379:2379     --publish 2380:2380     --env ALLOW_NONE_AUTHENTICATION=yes     --env ETCD_ADVERTISE_CLIENT_URLS=http://123.57.251.188 bitnami/etcd:latest
docker ps -a
docker logs etcd-server
docker ps -a
docker rm 9ae
docker run -d --name etcd-server     --network app-tier     --publish 2379:2379     --publish 2380:2380     --env ALLOW_NONE_AUTHENTICATION=yes     --env ETCD_ADVERTISE_CLIENT_URLS=http://123.57.251.188:2379 bitnami/etcd:latest
docker ps -a
