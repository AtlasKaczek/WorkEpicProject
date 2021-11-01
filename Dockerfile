FROM golang:1.17.2-alpine

WORKDIR /app

COPY . ./

RUN apk add make

RUN make tidy

CMD make run