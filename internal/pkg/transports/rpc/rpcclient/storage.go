package rpcclient

import (
	"fmt"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/transports/rpc"
	"sync"
)

var (
	storageOnce   sync.Once
	storageClient pb.StorageClient
)

func GetStorageClient(client *rpc.Client) pb.StorageClient {
	storageOnce.Do(func() {
		conn, err := client.Dial("storage")
		if err != nil {
			fmt.Println(err)
			return
		}
		storageClient = pb.NewStorageClient(conn)
	})

	return storageClient
}
