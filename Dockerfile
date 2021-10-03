FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod go.sum ./
COPY . ./

RUN go mod tidy
RUN go build -v -o /testing

EXPOSE 3000

CMD ["/testing"]