/*
@Time : 26/10/2020
@Author : GC
@Desc : 
*/

package cache

type Stat struct {
	Count     int
	KeySize   int // 字节数
	ValueSize int // 字节数
}

func (s *Stat) add(k string, v []byte) {
	// TODO 非线程安全的做法, 会有 BUG
	s.Count += 1
	s.KeySize += len(k)
	s.ValueSize += len(v)
}

func (s *Stat) del(k string, v []byte) {
	// TODO 非线程安全
	s.Count--
	// TODO 为了统计ValueSize 代价太大
	s.KeySize -= len(k)
	s.ValueSize -= len(v)
}
