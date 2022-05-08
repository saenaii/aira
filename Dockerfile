FROM golang:1.18-alpine AS base
RUN apk update && apk upgrade && \
    apk --no-cache add tzdata && \
    apk add git curl gcc cloc the_silver_searcher && \
    rm -rf /var/cache/apk/* && \
    wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.45.2 && \
    go install github.com/fzipp/gocyclo/cmd/gocyclo@latest

FROM base
WORKDIR /app/
COPY . .
COPY .golangci.yaml /etc/.golangci.yaml
RUN go env -w GO111MODULE=on && \
    go env -w GOPROXY=https://goproxy.cn,https://gocenter.io,https://goproxy.io,direct && \
    go mod download && \
    go build -o aira && mv aira /usr/bin/aira
CMD ["aira"]