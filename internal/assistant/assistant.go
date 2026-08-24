package assistant

import "ultron/internal/llm"

type Assistant struct {
	LLM *llm.Client
}

func (a *Assistant) Respond(message string) (string, error) {
	request := llm.ChatRequest{
		Model: "openai/gpt-oss-20b",
		Messages: []llm.Message{
			{
				Role:    "user",
				Content: message,
			},
		},
	}

	return a.LLM.Chat(request)
}
