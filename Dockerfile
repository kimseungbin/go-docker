FROM golang:1.19.2-alpine AS build
WORKDIR /usr/app/src

# copy necessary module scripts
COPY go.mod go.sum ./

# copy source code
COPY ./src ./

# install dependencies (i.e. modules)
RUN go mod download

# compile
RUN CGO_ENABLED=0 go build -o /go/bin/app main.go

FROM gcr.io/distroless/base-debian11
WORKDIR /
COPY --from=build /go/bin/app /
USER nonroot:nonroot
ENTRYPOINT ["/app"]