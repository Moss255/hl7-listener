FROM golang:1.17-alpine
WORKDIR /app
COPY . .
RUN go install
RUN go build
RUN ['./app/hl7-listener']