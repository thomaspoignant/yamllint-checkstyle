FROM golang:1.15 as build
WORKDIR /go/src/app
COPY . /go/src/app
RUN go mod vendor && go build -mod=vendor .

FROM gcr.io/distroless/base-debian10:latest
COPY --from=build /go/src/app /
CMD ["/yamllint-checkstyle"]