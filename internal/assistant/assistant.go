package assistant

import "ultron/internal/llm"

type Assistant struct {
	LLM *llm.Client
}

func (a *Assistant) Respond(messages []llm.Message) (string, error) {
	request := llm.ChatRequest{
		Model:    "openai/gpt-oss-20b",
		Messages: messages,
	}

	return a.LLM.Chat(request)
}