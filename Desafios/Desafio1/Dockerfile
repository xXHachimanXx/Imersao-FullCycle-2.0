FROM golang:1.15-alpine3.13 AS build

WORKDIR /go/src
COPY . .
ENV PATH="/go/bin:${PATH}"

EXPOSE 8000

RUN GOOS=linux go build main.go

CMD [ "./main" ]




