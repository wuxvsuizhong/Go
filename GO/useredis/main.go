package main

import (
	useredis "useredis/useredis1"
)

func main() {
	useredis.AddKeyVal()
	useredis.AddHash()

	useredis.UseRedisPool()
}
