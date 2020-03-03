FROM golang:1.13-alpine
WORKDIR /go/src/app
RUN apk update --no-cache && \
    apk add --no-cache git && \
    rm -rf /var/cache/apk/* && \
    go get -u github.com/gofiber/fiber/... && \
    go get github.com/pilu/fresh
COPY ./app .
EXPOSE 3000
CMD ["fresh"]