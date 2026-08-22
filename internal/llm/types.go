package llm

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type ChatResponse struct {
	Choices []Choice `json:"choices"`
}

type Choice struct {
	Message ResponseMessage `json:"message"`
}

type ResponseMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
