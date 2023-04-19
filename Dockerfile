FROM golang:1.18.2-alpine3.16 as base
RUN apk update
WORKDIR /src
COPY go.mod go.sum ./
COPY . .
RUN go build -o socius ./

FROM alpine:3.16 as binary
COPY --from=base /src/socius/socius .
EXPOSE 3000
CMD ["./socius"]