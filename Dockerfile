FROM golang:1.18.2-alpine3.16 as base
RUN apk update
WORKDIR /src/socius
COPY go.mod go.sum ./
COPY . .
RUN go build -o socius ./cmd/api

FROM alpine:3.16 as binary
COPY --from=base /src/socius/socius .
COPY --from=base /src/socius/web ./web
EXPOSE 3000
CMD ["./socius"]