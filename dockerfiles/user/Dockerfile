FROM alpine
LABEL authors="wangty"

WORKDIR /app
ADD ./dist/server/user-server .
COPY config/ config/
EXPOSE 22111
CMD ["./user-server"]
