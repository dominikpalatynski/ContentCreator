package content

import (
	"context"
	"fmt"
	"os"

	"github.com/dominikpalatynski/ContentCreator/util"
	openai "github.com/sashabaranov/go-openai"
)

func GenerateContent(prompt string) {
	util.LoadEnv()
	client := openai.NewClient(os.Getenv("OPEN_AI_API_KEY"))
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}
	
	fmt.Println(resp.Choices[0].Message.Content)
	
}