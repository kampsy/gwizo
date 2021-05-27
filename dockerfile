FROM golang:alpine AS build-env
WORKDIR /app
ADD . /app
RUN cd /app && go mod tidy
RUN cd /app && go build -o dazwallet

FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apl/*
WORKDIR /app
COPY --from=build-env /app/dazwallet /app

EXPOSE 8080
ENTRYPOINT ./dazwallet
