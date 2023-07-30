FROM golang:1.20 AS build_stage

WORKDIR /app

COPY go.sum go.mod ./

RUN go mod download 

COPY . ./

RUN CGO_ENABLED=1 go build -o bin/app ./cmd/server.go

FROM frolvlad/alpine-glibc

WORKDIR /app/

COPY --from=build_stage ./app/bin/app ./

EXPOSE 5000

CMD [ "./app", "-p","5000" ]