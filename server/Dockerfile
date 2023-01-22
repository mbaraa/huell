FROM alpine:latest as build

RUN apk add go git

WORKDIR /app
COPY . .

RUN go get
RUN go build -ldflags="-w -s"

FROM alpine:latest as run

RUN apk add cairo

WORKDIR /app

COPY --from=build /app/huell ./run

EXPOSE 9090

CMD ["./run"]
