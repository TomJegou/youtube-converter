# YouTube Converter

A small web app that downloads audio from a YouTube URL and returns it as an MP3 file. The UI is a [Next.js](https://nextjs.org/) frontend; conversion runs on a [Go](https://go.dev/) API backed by [yt-dlp](https://github.com/yt-dlp/yt-dlp) and [ffmpeg](https://ffmpeg.org/).

## Architecture

```
Browser  →  Frontend (Next.js, :45450)  →  Backend API (Go, :46460)  →  yt-dlp + ffmpeg
```

| Service   | Stack              | Port  | Role                                      |
|-----------|--------------------|-------|-------------------------------------------|
| Frontend  | Next.js 16, React  | 45450 | Form UI; proxies download requests to API |
| Backend   | Go                 | 46460 | Runs yt-dlp, streams MP3 to the client    |

Both services are built as **multi-stage Docker images** and orchestrated with Docker Compose at the repository root.

## Prerequisites

- [Docker](https://www.docker.com/) and Docker Compose

For local development without Docker:

- **Backend:** Go 1.19+, `yt-dlp`, `ffmpeg` on your `PATH`
- **Frontend:** Node.js 20+ and npm

## Quick start (Docker)

From the project root:

```bash
docker compose up --build
```

Open [http://localhost:45450](http://localhost:45450), paste a YouTube URL, and submit. The browser receives the converted MP3 as a download.

Stop the stack:

```bash
docker compose down
```

## Configuration

| Variable        | Service  | Default (Docker) | Description                          |
|-----------------|----------|------------------|--------------------------------------|
| `API_HOSTNAME`  | Frontend | `api:46460`      | Host and port of the backend API     |
| `PORT`          | Frontend | `45450`          | HTTP port for the Next.js server     |

In Docker Compose, the frontend reaches the backend via the service name `api` on the internal network. Only the frontend port `45450` is published to the host.

## API

### `GET /`

Health-style welcome message.

### `POST /dl`

**Request body (JSON):**

```json
{ "url": "https://www.youtube.com/watch?v=..." }
```

**Response:** MP3 file (`Content-Type: audio/mp3`, `Content-Disposition: attachment`) on success, or `500` on failure.

The frontend exposes this through `POST /api`, which forwards the form field `url` to the backend.

## Local development

### Backend

```bash
cd backend
go run .
```

API listens on [http://localhost:46460](http://localhost:46460).

### Frontend

```bash
cd frontend
npm ci
npm run dev
```

Point the app at the local API (default in code is `localhost:46460` when `API_HOSTNAME` is unset):

```bash
# Windows PowerShell
$env:API_HOSTNAME="localhost:46460"; npm run dev
```

## Project structure

```
youtube-converter/
├── docker-compose.yml    # Runs frontend + backend
├── backend/
│   ├── Dockerfile        # Go builder + Debian runtime (yt-dlp, ffmpeg)
│   ├── main.go
│   ├── handlers/         # HTTP handlers (/dl, /)
│   └── utils/            # yt-dlp wrapper
└── frontend/
    ├── Dockerfile        # Node builder + Node runtime (standalone Next.js)
    ├── app/              # Next.js App Router (page + /api route)
    └── public/           # Static assets (e.g. youtube.png)
```

## Build images only

```bash
docker compose build
```

Images are tagged as `tomyj/my-env:youtube-converter-frontend` and `tomyj/my-env:youtube-converter-backend`.

## Notes

- Conversion quality is set to best audio (`--audio-quality 0`) with embedded metadata and thumbnail.
- Temporary MP3 files are written in the backend working directory and removed after the response is sent.
- Respect YouTube’s terms of service and copyright when downloading content.

## License

No license file is included in this repository. All rights reserved unless stated otherwise by the author.
