# BUILD
FROM golang:1.11.4-alpine as build-stage
ARG version
ARG revision
RUN touch /${version} && touch /${revision}
COPY . /go/src/hrkt.io/command-as-a-service
RUN go install -ldflags="-s -w -X \"main.Version=${version}\" -X \"main.Revision=${revision}\" -extldflags \"-static\"" hrkt.io/command-as-a-service

# Production
FROM alpine as production-stage
COPY --from=build-stage /go/bin/command-as-a-service .
COPY app-settings.json .
ENV PORT 8080
ENV GIN_MODE=release
CMD ["./command-as-a-service"]

