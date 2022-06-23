package integration

import (
	"testing"

	"github.com/travisjeffery/proto-go-sql/integration/proto"
)

func TestInitialize(t *testing.T) {
	_ = &proto.MessageNoOption{}
	_ = &proto.MessageWithOption{}
}
