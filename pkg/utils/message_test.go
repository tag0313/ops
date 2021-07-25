package utils

import (
	"testing"

	"github.com/asim/go-micro/v3/logger"
)

//test msg type 9
func TestMessage9(t *testing.T) {
	uid := "需要被通知的uid"
	mentioning := "谁@了这个uid"
	err := SendMentionMsg(uid, mentioning, "oid", "content", "at_users,another_at_users")
	if err != nil {
		logger.Error("进入消息队列失败")
		return
	}
}
