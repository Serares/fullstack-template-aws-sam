FROM public.ecr.aws/docker/library/golang:1.21-alpine3.18 AS builder
RUN mkdir /homepage
WORKDIR /homepage
COPY go.mod go.sum ./
COPY services/ssr/homepage/ ./
RUN CGO_ENABLED=0 GOOS=linux go build -o homepage -ldflags="-s -w" .
FROM public.ecr.aws/docker/library/alpine:latest
RUN mkdir /app && adduser -h /app -D homepage
WORKDIR /app
COPY --chown=homepage --from=builder /homepage/homepage .
CMD ["/app/homepage"]