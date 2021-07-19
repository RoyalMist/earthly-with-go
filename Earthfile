ARG GO_VERSION=1.16
ARG NODE_VERSION=16.5
ARG ALPINE_VERSION=3.14
FROM golang:$GO_VERSION-alpine$ALPINE_VERSION
RUN apk add --no-cache build-base
RUN go install github.com/securego/gosec/v2/cmd/gosec@v2.8.1
RUN go install honnef.co/go/tools/cmd/staticcheck@latest

frontend-base:
    FROM node:$NODE_VERSION-alpine$ALPINE_VERSION
    RUN npm install -g pnpm
    WORKDIR /build

frontend-deps:
    FROM +frontend-base
    COPY web/package.json .
    COPY web/pnpm-lock.yaml .
    RUN pnpm i

frontend-build:
    FROM +frontend-deps
    COPY web/src src
    COPY web/declaration.d.ts .
    COPY web/index.html .
    COPY web/postcss.config.js .
    COPY web/tailwind.config.js .
    COPY web/tsconfig.json .
    COPY web/vite.config.ts .
    RUN pnpm build
    SAVE ARTIFACT dist

backend-base:
    WORKDIR /build
    COPY go.mod .
    COPY go.sum .
    COPY main.go .
    COPY backend backend
    RUN mkdir web
    COPY +frontend-build/dist web/dist

backend-quality:
    FROM +backend-base
    RUN gosec ./...
    RUN staticcheck ./...
    WITH DOCKER --pull postgres:13-alpine
        RUN docker run --name postgres -e POSTGRES_PASSWORD=postgres -p 5432:5432 -d postgres:13-alpine && \
            until docker exec postgres pg_isready ; do sleep 5 ; done && \
            go test
    END

backend-build:
    FROM +backend-base
    RUN go build -ldflags "-s -w" -o wiatt
    SAVE ARTIFACT wiatt

docker:
    FROM alpine:$ALPINE_VERSION
    ENV LANG "en_US.UTF-8"
    ENV LANGUAGE "en_US:en"
    ENV LC_ALL "en_US.UTF-8"
    COPY +backend-build/wiatt /opt/wiatt
    RUN apk update && \
        apk upgrade && \
        apk add --no-cache \
        ca-certificates && \
        update-ca-certificates && \
        chown -R nobody:nobody /opt/wiatt && \
        apk del ca-certificates
    USER nobody:nobody
    EXPOSE 4000
    ENTRYPOINT ["/opt/wiatt"]
    SAVE IMAGE royalmist/wiatt:latest

linux:
    FROM +backend-base
    RUN GOOS=linux go build -ldflags "-s -w" -o wiatt
    SAVE ARTIFACT wiatt AS LOCAL ./artifacts/linux/wiatt

windows:
    FROM +backend-base
    RUN GOOS=windows go build -ldflags "-s -w" -o wiatt.exe
    SAVE ARTIFACT wiatt.exe AS LOCAL ./artifacts/windows/wiatt.exe

osx:
    FROM +backend-base
    RUN GOOS=darwin go build -ldflags "-s -w" -o wiatt
    SAVE ARTIFACT wiatt AS LOCAL ./artifacts/osx/wiatt

all:
    BUILD +linux
    BUILD +windows
    BUILD +osx

quality:
    BUILD +backend-quality
