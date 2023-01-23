FROM golang:1.19 AS build
COPY . /app

WORKDIR /app
RUN [ "go", "build", "." ]


FROM gcr.io/distroless/base-debian11:nonroot
COPY --from=build --chown=nonroot:nonroot /app/distributor /distributor

USER nonroot

ENTRYPOINT [ "/distributor" ]