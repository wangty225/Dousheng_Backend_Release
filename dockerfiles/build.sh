docker build -f dockerfiles/router/Dockerfile -t dousheng_router:v0.1 ./

docker build -f dockerfiles/user/Dockerfile -t dousheng_user-server:v0.1 ./

docker build -f dockerfiles/video/Dockerfile -t dousheng_video-server:v0.1 ./

docker build -f dockerfiles/interaction/Dockerfile -t dousheng_interaction-server:v0.1 ./
