package gemini

import (
	"context"
	"fmt"
	"log"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

const ImageTemperture = 0.8
const ChatTemperture = 0.3

const (
	GEMINI_TEXT_MODEL_NAME  = "gemini-pro"
	GEMINI_IMAGE_MODEL_NAME = "gemini-pro-vision"
)

// Gemini Image: Input an image and get the response string.
func GeminiImage(imgData []byte) (string, error) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(geminiKey))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel(GEMINI_IMAGE_MODEL_NAME)
	value := float32(ImageTemperture)
	model.Temperature = &value
	prompt := []genai.Part{
		genai.ImageData("png", imgData),
		genai.Text("Describe this image with scientific detail, reply in zh-TW:"),
	}
	log.Println("Begin processing image...")
	resp, err := model.GenerateContent(ctx, prompt...)
	log.Println("Finished processing image...", resp)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return PrintResponse(resp), nil
}

// Print response
func PrintResponse(resp *genai.GenerateContentResponse) string {
	var ret string
	for _, cand := range resp.Candidates {
		for _, part := range cand.Content.Parts {
			ret = ret + fmt.Sprintf("%v", part)
			log.Println(part)
		}
	}
	return ret
}

// send: Send a message to the chat session
func Send(cs *genai.ChatSession, msg string) *genai.GenerateContentResponse {
	if cs == nil {
		cs = StartNewChatSession()
	}

	ctx := context.Background()
	log.Printf("== Me: %s\n== Model:\n", msg)
	res, err := cs.SendMessage(ctx, genai.Text(msg))
	if err != nil {
		log.Fatal(err)
	}
	return res
}

// startNewChatSession	: Start a new chat session
func StartNewChatSession() *genai.ChatSession {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(geminiKey))
	if err != nil {
		log.Fatal(err)
	}
	model := client.GenerativeModel(GEMINI_TEXT_MODEL_NAME)
	value := float32(ChatTemperture)
	model.Temperature = &value
	cs := model.StartChat()
	return cs
}
