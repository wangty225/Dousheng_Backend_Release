FROM alpine
LABEL authors="wangty"

WORKDIR /app
ADD ./dist/server/interaction-server .
COPY config/ config/
EXPOSE 22113
CMD ["./interaction-server"]
