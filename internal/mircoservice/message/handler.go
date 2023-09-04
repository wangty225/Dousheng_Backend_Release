package main

import (
	message "Dousheng_Backend/internal/mircoservice/message/kitex-gen/message"
	"context"
)

// MessageServiceImpl implements the last service interface defined in the IDL.
type MessageServiceImpl struct{}

// ChatMessage implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) ChatMessage(ctx context.Context, req *message.DouyinMessageChatRequest) (resp *message.DouyinMessageChatResponse, err error) {
	// TODO: Your code here...
	return
}

// ActionMessage implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) ActionMessage(ctx context.Context, req *message.DouyinMessageActionRequest) (resp *message.DouyinMessageActionResponse, err error) {
	// TODO: Your code here...
	return
}
