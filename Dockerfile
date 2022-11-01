FROM golang:1.19.2-alpine AS build
WORKDIR /usr/app/src

# copy necessary module scripts
COPY go.mod ./
COPY go.sum ./

# copy source code
COPY ./src ./

# install dependencies (i.e. modules)
RUN go mod download

# compile
RUN CGO_ENABLED=0 go build -o /go/bin/app main.go

# run!
CMD ["/go/bin/app"]
