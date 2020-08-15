FROM golang:alpine

RUN mkdir /app
WORKDIR /app
ADD main.go .
ADD go.mod .
RUN go build -o random .
EXPOSE 8080
RUN adduser -S -D -H -h /app user
USER user

CMD ["./random"]