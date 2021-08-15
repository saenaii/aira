FROM golang:1.16-alpine AS base
RUN apk update && apk upgrade && \
    apk --no-cache add tzdata && \
    apk add git curl gcc && \
    rm -rf /var/cache/apk/* && \
    curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.41.1 && \
    go get -u github.com/fzipp/gocyclo/cmd/gocyclo

FROM base
WORKDIR /app/
COPY . .
RUN go env -w GO111MODULE=on && \
    go env -w GOPROXY=https://goproxy.cn,https://gocenter.io,https://goproxy.io,direct && \
    go mod download && \
    go build -o aira && mv aira /usr/bin/aira
CMD ["aira"]