/*
 * ZGrab Copyright 2015 Regents of the University of Michigan
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not
 * use this file except in compliance with the License. You may obtain a copy
 * of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
 * implied. See the License for the specific language governing
 * permissions and limitations under the License.
 */

package cachehash

import (
	"fmt"
	"hash/crc32"
)

type ShardedCacheHash struct {
	shards    []CacheHash
	shardsLen int
}

func (c *ShardedCacheHash) Init(maxLen int, shards int) {
	fmt.Println("INIT ")
	c.shardsLen = shards
	shardLen := maxLen / shards
	c.shards = make([]CacheHash, shards)
	for i := 0; i < shards; i++ {
		c.shards[i].Init(shardLen)
	}
}

//This function leave as is?
func (c *ShardedCacheHash) getShardID(k interface{}) int {
	fmt.Println("GET SHARD ID")
	kb := []byte(fmt.Sprintf("%v", k))
	return int(crc32.ChecksumIEEE(kb)) % c.shardsLen
}

//uses the other cachehashe looks like we need to replace both?
func (c *ShardedCacheHash) getShard(k interface{}) *CacheHash {
	fmt.Println("Get Shard ")
	return &c.shards[c.getShardID(k)]
}

func (c *ShardedCacheHash) Add(k interface{}, v interface{}) bool {
	fmt.Println("ADD SHARD")
	return c.getShard(k).Add(k, v)
}

func (c *ShardedCacheHash) Get(k interface{}) (interface{}, bool) {
	fmt.Println("GET SHARD")
	return c.getShard(k).Get(k)
}

func (c *ShardedCacheHash) GetNoMove(k interface{}) (interface{}, bool) {
	fmt.Println("GET NO MOVE SHARD")
	return c.getShard(k).GetNoMove(k)
}

func (c *ShardedCacheHash) Has(k interface{}) bool {
	fmt.Println("HAS SHARD")
	return c.getShard(k).Has(k)
}

func (c *ShardedCacheHash) Delete(k interface{}) (interface{}, bool) {
	fmt.Println("DELETE SHARD")
	return c.getShard(k).Delete(k)
}

func (c *ShardedCacheHash) RegisterCB(newCB func(interface{}, interface{})) {
	fmt.Println("REGISTER CB ")
	for i := 0; i < c.shardsLen; i++ {
		c.shards[i].RegisterCB(newCB)
	}
}

func (c *ShardedCacheHash) Lock(k interface{}) {
	fmt.Println("LOCK SHARD")
	c.getShard(k).Lock()
}

func (c *ShardedCacheHash) Unlock(k interface{}) {
	fmt.Println("UNLOCK SHARD")
	c.getShard(k).Unlock()
}
