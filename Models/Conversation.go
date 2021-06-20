package Models

type Conversation struct {
	Senderid string
	Name     string
	Message  string
}

type Chatroom struct {
	Participant1 string
	Participant2 string
	Conversation []Conversation
}
