package gmail

import "testing"

func TestLabels(t *testing.T) {
	Labels()
}

func TestDelete(t *testing.T) {
	Delete("INBOX")
	Delete("TRASH")
	Delete("SPAM")
}