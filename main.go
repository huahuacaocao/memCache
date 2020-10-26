/*
@Time : 26/10/2020
@Author : GC
@Desc : 
*/

package main

import (
	"memcache/cache"
	"memcache/server"
)

func main() {
	memCache := cache.New()
	server.New(memCache).Listen()
}
