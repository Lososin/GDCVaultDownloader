package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
)

var (
	URL string = ""
	FILE_PATH string = ""
	QUEUE_FILE string = ""
)

func DownloadFile(url string, file string) {
	if url == "" || file == "" {
		return
	}

	file += ".mkv"

	var exePath = "./ffmpeg"
	if runtime.GOOS == "windows" {
		exePath += ".exe"
	}

	cmd := exec.Command(exePath, "-y", "-i", url, "-c", "copy", file)
	fmt.Println(cmd)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}

func main() {

	flag.StringVar(&URL, "url", URL, "url to m3u8 file")
	flag.StringVar(&FILE_PATH, "path", FILE_PATH, "Path to save video")
	flag.StringVar(&QUEUE_FILE, "queue", QUEUE_FILE, "file, that contains names and urls")
	flag.Parse()

	if QUEUE_FILE == "" {
		DownloadFile(URL, FILE_PATH)
		return
	}

	file, err := os.Open(QUEUE_FILE)
	if err != nil {
		log.Fatalf("failed to open")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var text []string
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	for i := 0; i < len(text); i+=2 {
		DownloadFile(text[i+1], text[i])
	}
}
