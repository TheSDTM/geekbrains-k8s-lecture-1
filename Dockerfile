FROM golang:alpine AS build

WORKDIR /go/src/app

COPY . .

RUN go mod download
RUN CGO_ENABLED=0 go build -o app .

FROM scratch AS bin
COPY --from=build /go/src/app/app /app
CMD ["/app"]