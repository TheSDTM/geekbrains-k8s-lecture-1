FROM golang:alpine AS build

WORKDIR /go/src/app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o app .

FROM scratch AS bin
COPY --from=build /go/src/app/app /app
CMD ["/app"]