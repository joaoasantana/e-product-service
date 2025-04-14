package connect

import (
	"fmt"
)

func MongoPattern(host, port string) string {
	return fmt.Sprintf("mongodb://%s:%s", host, port)
}
