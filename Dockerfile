#Docker multi-stage builds

# ------------------------------------------------------------------------------
# Development image
# ------------------------------------------------------------------------------
FROM golang:1.19-alpine AS development

# Force the go compiler to use modules
ENV GO111MODULE=on
#ENV GOPROXY=http://172.17.0.1:3333

# Update OS package and install Git
RUN apk update && \
    apk add git && \
    apk add build-base && \
    apk add openssh && \
    apk add --no-cache bash && \
    apk add git mercurial && \
    apk add linux-headers

# Set working directory
WORKDIR /go/src/github.com/Kafakk/gin-gorm-clean-code

# Copy Go dependency file
COPY go.mod go.mod
COPY go.sum go.sum

# Install Fresh for local development
RUN go get github.com/pilu/fresh

# Download dependency
RUN go mod download

# Copy src
COPY app app

# Use CMD instead of RUN to allow command overwritability
CMD cd app && fresh
