FROM golang:1 as builder

COPY . src/github.com/31z4/harvest
RUN cd src/github.com/31z4/harvest && make release


FROM scratch
LABEL maintainer="Elisey Zanko <elisey.zanko@gmail.com>"

ENV COLUMNS=80
ENTRYPOINT ["./harvest"]
COPY --from=builder /go/src/github.com/31z4/harvest/release/harvest .
