docker run -d --name dousheng_router --restart unless-stopped --network my_static_network dousheng_router:v0.1

docker run -d --name dousheng_user-server --restart unless-stopped --network my_static_network --ip 192.168.1.11 -p 22111:22111 dousheng_user-server:v0.1

docker run -d --name dousheng_video-server --restart unless-stopped --network my_static_network --ip 192.168.1.12 -p 22112:22112 dousheng_video-server:v0.1

docker run -d --name dousheng_interaction-server --restart unless-stopped --network my_static_network --ip 192.168.1.13 -p 22113:22113 dousheng_interaction-server:v0.1
