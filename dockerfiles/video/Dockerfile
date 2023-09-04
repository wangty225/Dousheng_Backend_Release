FROM alpine
LABEL authors="wangty"

WORKDIR /app
ADD ./dist/server/video-server .
COPY config/ config/
EXPOSE 22112
CMD ["./video-server"]
