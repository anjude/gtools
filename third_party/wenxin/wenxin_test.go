package wenxin

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestGetWenXinReply(t *testing.T) {
	GetWenXinReply("11", fmt.Sprintf("%d", rand.Intn(1000000)))
}
