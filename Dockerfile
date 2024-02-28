FROM reg.klab.sh/base/go:1.22.0 as builder
WORKDIR /go/src/app
COPY . .
RUN go fmt && CGO_ENABLED=0 go build .

FROM reg.klab.sh/base/abi:edge as runner
COPY --from=builder /go/src/app/klab-site /klab-site
RUN chown appuser:appgroup /klab-site && chmod +x /klab-site
USER appuser
CMD ["/klab-site"]

