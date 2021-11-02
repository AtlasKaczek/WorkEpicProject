package epictest

import (
	"stankryj/WorkEpicProject/epic"
	"testing"
)

func TestSendSlackMessage(t *testing.T) {
	body, err := epic.SendSlackMessege("testing", "https://hooks.slack.com/services/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")
	if err != nil {
		t.Fatal(err)
	}
	if body == nil {
		t.Errorf("body is nil")
	}
}
