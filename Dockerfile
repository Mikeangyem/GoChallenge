ARG baseImage=golang:1.24-alpine

FROM ${baseImage} AS builder

WORKDIR /app

ADD . /app/

RUN go mod download && go mod verify

RUN go build -v /GoChallenge ./cmd/api/main.go

FROM ${baseImage} AS candidate

ARG NAME=GoChallenge

WORKDIR /app

COPY --from=builder /${NAME} .

EXPOSE 8080

ENTRYPOINT [ "./GoChallenge" ]