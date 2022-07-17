FROM golang:1.18-alpine AS build-env
WORKDIR /src
ADD . .
RUN apk update && \
    apk add --no-cache make build-base
ENV GOOS=linux
ENV GOARCH=amd64
RUN make build

FROM golang:1.18-alpine AS test-env
ENV PACKAGES build-base wget
RUN apk update && \
    apk add --no-cache $PACKAGES
RUN wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.46.2
WORKDIR /src
COPY --from=build-env /src .
RUN make lint
RUN make test

FROM alpine AS runtime
WORKDIR /app
COPY --from=build-env /src/bin/workerd /app/
COPY --from=build-env /src/migrations /app/migrations
ENTRYPOINT /app/workerd
