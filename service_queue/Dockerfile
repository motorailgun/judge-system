FROM golang:1.19 as builder
COPY . /build

WORKDIR /build
RUN [ "go", "build", "-a" ]


FROM gcr.io/distroless/base-debian11
COPY --from=builder /build/service_queue /

ENTRYPOINT [ "/service_queue" ]
