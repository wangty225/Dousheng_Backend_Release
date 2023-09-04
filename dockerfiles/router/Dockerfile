FROM alpine
LABEL authors="wangty"

WORKDIR /app
ADD ./dist/router/hertz-router .
COPY config/ config/
EXPOSE 8888
CMD ["./hertz-router"]
