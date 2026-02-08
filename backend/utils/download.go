package utils

import (
	"fmt"
	"os"
	"os/exec"
)

const executable string = "yt-dlp"
const extrAudio string = "-x"
const audioFormat string = "--audio-format"
const audioQuality string = "--audio-quality"
const addMetadata string = "--embed-metadata"
const embedThumbnail string = "--embed-thumbnail"

func Download(url string) error {
	cmd := exec.Command(executable, extrAudio, audioFormat, "mp3", audioQuality, "0", addMetadata, embedThumbnail, url)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("cmd.Run(): %v", err)
	}
	return nil
}
