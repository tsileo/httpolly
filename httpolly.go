package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/polly"
)

type Payload struct {
	Text    string `json:"text"`
	VoiceID string `json:"voice_id"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	p := &Payload{}
	if err := json.NewDecoder(r.Body).Decode(p); err != nil {
		panic(err)
	}

	svc := polly.New(session.New())
	input := &polly.SynthesizeSpeechInput{
		OutputFormat: aws.String("mp3"),
		Text:         aws.String(p.Text),
		TextType:     aws.String("text"),
		VoiceId:      aws.String(p.VoiceID),
	}

	result, err := svc.SynthesizeSpeech(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			fmt.Println(aerr.Code(), aerr.Error())
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "audio/mpeg")
	io.Copy(w, result.AudioStream)
}

func main() {
	// TODO(tsileo): add CLI flags for options
	http.HandleFunc("/", handler)
	http.ListenAndServe("127.0.0.1:8015", nil)
}
