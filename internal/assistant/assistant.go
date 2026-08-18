package assistant

type Assistant struct{}

func (a *Assistant) Respond(message string) string {
	return "You said: " + message
}
