FROM tomyj/my-env:go AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o bin/backend .

FROM debian:13.3
WORKDIR /app
RUN apt-get update \
    && apt-get install -y --no-install-recommends \
        python3 \
        python3-pip \
        ffmpeg \
        ca-certificates \
    && pip3 install --break-system-packages yt-dlp \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*
COPY --from=builder /app/bin/backend ./bin/backend
EXPOSE 46460
CMD ["./bin/backend"]
