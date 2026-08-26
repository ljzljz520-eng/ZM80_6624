package notify

import (
	"context"
	"fmt"
	"museumlogistics/model"
)

type Message struct{ To, Subject, Body string }

func Build(ctx context.Context, r model.Record, reason string) (Message, error) {
	if err := ctx.Err(); err != nil {
		return Message{}, err
	}
	if r.Status == "rejected" && reason == "" {
		return Message{}, fmt.Errorf("missing rejection reason")
	}
	return Message{To: r.Origin, Subject: "借展状态更新", Body: fmt.Sprintf("%s: %s", r.ID, reason)}, nil
}
func Channel(status string) string {
	switch status {
	case "approved":
		return "email"
	case "rejected":
		return "review-queue"
	default:
		return "dashboard"
	}
}
