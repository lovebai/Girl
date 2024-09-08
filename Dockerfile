FROM alpine:latest

WORKDIR /Girl
COPY ./girl ./girl
COPY ./data/ ./data

RUN apk update \
    && apk add --no-cache tzdata \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone \
    && chmod +x ./girl 

EXPOSE 5200
VOLUME [ "/Girl/data"]

ENTRYPOINT ["./girl"]