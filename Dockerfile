FROM golang:1.24-alpine AS builder
RUN apk add --no-cache build-base

WORKDIR /src/app
ADD go.mod .
RUN go mod download

ADD . .
# Copy everything in assets to the public folder
RUN cp internal/assets/* public/

# Downloading the tailwind binary, musl because this is an Alpine image, and
# running the build command
RUN go tool tailo download -v v4.0.6 --musl
RUN go tool tailo --i internal/system/assets/tailwind.css -o internal/system/assets/application.css

# Building the app
RUN go build -tags osusergo,netgo -buildvcs=false -o bin/app ./cmd/app

# Building the migrator
RUN go build -tags osusergo,netgo -buildvcs=false -o bin/migrate ./cmd/migrate

FROM alpine
RUN apk add --no-cache tzdata ca-certificates bash

# Copying binaries
COPY --from=builder /src/app/bin/migrate /bin/
COPY --from=builder /src/app/bin/app /bin/

WORKDIR /usr/

SHELL ["/bin/ash", "-c"]
CMD migrate && app
