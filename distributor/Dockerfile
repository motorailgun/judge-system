FROM rust:1.66-slim-bullseye as builder

COPY . /distributor
WORKDIR /distributor

RUN cargo build --release


FROM debian:bullseye-slim

RUN mkdir /app

COPY --from=builder /distributor/target/release/distributor /app/

ENTRYPOINT [ "/app/distributor" ]
