package app

import "github.com/bwmarrin/snowflake"

var node *snowflake.Node

func init() {
	n, err := snowflake.NewNode(1)
	if err != nil {
		log.Fatal(err)
	}
	node = n
}
func GetSnowId() string {
	return node.Generate().String()
}
