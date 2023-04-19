package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type ChatCompletionMessage struct {
	Content string `json:"content,omitempty"`
}

func main() {
	// resp, err := http.Get("http://localhost:5000/")
	client := http.Client{}
	message := ChatCompletionMessage{
		Content: "耳机 18",
	}
	jsonMessage, err := json.Marshal(message)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", "http://localhost:8080/api/v1/gpt/2/chat-completion/steam", bytes.NewReader(jsonMessage))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "text/event-stream")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	buf := make([]byte, 10)
	for {
		n, err := resp.Body.Read(buf)
		if err != nil {
			break
		}

		line := string(buf[:n])
		if line == "" {
			continue
		}
		fmt.Print(string(line))
	}

	resp.Body.Close()
}
