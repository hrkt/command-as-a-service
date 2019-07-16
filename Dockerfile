# BUILD
FROM golang:1.11.4-alpine as build-stage
ARG version
ARG revision
RUN touch /${version} && touch /${revision}
COPY . /go/src/ehw2018//cmd-exec-server
RUN go install -ldflags="-s -w -X \"main.Version=${version}\" -X \"main.Revision=${revision}\" -extldflags \"-static\"" ehw2018/cmd-exec-server

# Production
FROM alpine as production-stage
COPY --from=build-stage /go/bin/cmd-exec-server .
ENV PORT 8080
ENV GIN_MODE=release
CMD ["./cmd-exec-server"]

