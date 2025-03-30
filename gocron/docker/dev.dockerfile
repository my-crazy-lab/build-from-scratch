ARG V_DOCKER=27.3.1
ARG V_GOLANG=1.24
ARG V_RCLONE=1
ARG V_RESTIC=0.17.3
FROM docker:${V_DOCKER}-cli AS docker
FROM rclone/rclone:${V_RCLONE} AS rclone
FROM restic/restic:${V_RESTIC} AS restic
FROM golang:${V_GOLANG}-alpine
RUN apk add --update --no-cache \
    su-exec dumb-init \
    zip tzdata borgbackup rsync curl rdiff-backup

# docker
COPY --from=docker --chmod=0755 \
    /usr/local/bin/docker \
    /usr/local/bin/docker

# docker compose
COPY --from=docker --chmod=0755 \
    /usr/local/libexec/docker/cli-plugins/docker-compose \
    /usr/local/libexec/docker/cli-plugins/docker-compose

# rclone
COPY --from=rclone --chmod=0755 \
    /usr/local/bin/rclone /usr/bin/rclone

# restic
COPY --from=restic --chmod=0755 \
    /usr/bin/restic /usr/bin/restic

# air
RUN go install github.com/air-verse/air@latest

WORKDIR /app

COPY ./go.mod ./go.sum ./
RUN go mod download
