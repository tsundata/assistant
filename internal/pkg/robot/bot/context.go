package bot

import "github.com/tsundata/assistant/api/pb"

type Context struct {
	Setting   map[int64][]*pb.KV
	Message   *pb.Message
	FieldItem []FieldItem

	Input  PluginValue
	Output PluginValue
}
