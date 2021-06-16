package rpcclient

import (
	"fmt"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc"
	"sync"
)

var (
	nlpOnce   sync.Once
	nlpClient pb.NLPClient
)

func GetNLPClient(client *rpc.Client) pb.NLPClient {
	nlpOnce.Do(func() {
		conn, err := client.Dial("nlp")
		if err != nil {
			fmt.Println(err)
			return
		}
		nlpClient = pb.NewNLPClient(conn)
	})

	return nlpClient
}
