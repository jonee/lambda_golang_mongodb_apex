/**
 * Config class
 */

package main

const (
	MONGODB_DATABASE_DEV  = "sample"

	MONGODB_AUTH_DATABASE_DEV  = "admin"

	MONGODB_REGULAR_USER_DEV           = "readWrite"
	MONGODB_REGULAR_USER_PASSWORD_DEV  = "readWritePassword"
)

var (
	MONGODB_SERVERS_DEV  = []string{"cluster0-shard-00-00-xxxxx.mongodb.net:27017", "cluster0-shard-00-01-xxxxx.mongodb.net:27017", "cluster0-shard-00-02-xxxxx.mongodb.net:27017"}
)
