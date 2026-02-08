FROM tomyj/my-env:go
WORKDIR /app
RUN apt update -y && apt full-upgrade -y
RUN apt install -y python3 python3-pip
RUN pip3 install yt-dlp
RUN apt install -y ffmpeg
COPY . .
RUN go build -o bin/backend .
CMD [ "./bin/backend" ]