FROM golang:1.19-bullseye

COPY go.mod go.sum /vender_api/
WORKDIR /vender_api/

RUN go mod download
COPY . .

RUN go build -o /vender-api .
RUN chmod +x /vender-api

ENTRYPOINT ["/vender-api"]
