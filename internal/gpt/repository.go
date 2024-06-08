package gpt

import (
	"context"

	"github.com/sashabaranov/go-openai"
)

type GPTRepository struct {
	Client *openai.Client
}

func (g *GPTRepository) ChatWithImage(img string, prompt string) (*string, error) {
	resp, err := g.Client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model: openai.GPT4o,
		Messages: []openai.ChatCompletionMessage{
			{
				Role: openai.ChatMessageRoleUser,
				MultiContent: []openai.ChatMessagePart{
					{
						ImageURL: &openai.ChatMessageImageURL{
							URL:    img,
							Detail: openai.ImageURLDetailAuto,
						},
						Type: openai.ChatMessagePartTypeImageURL,
					},
					{
						Text: prompt,
						Type: openai.ChatMessagePartTypeText,
					},
				},
			},
		},
	})
	if err != nil {
		return nil, err
	}
	result := resp.Choices[0].Message.Content
	return &result, nil
}

func InitGPTRepository() *GPTRepository {
	return &GPTRepository{}
}

func (g GPTRepository) WithClient(client *openai.Client) GPTRepository {
	g.Client = client
	return g
}
