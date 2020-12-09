package common

import (
	"errors"
	"hash/crc32"
	"sort"
	"strconv"
	"sync"
)

// 声明新切片类型
type units []uint32

// 返回切片长度
func (x units) Len() int {
	return len(x)
}

// 比对两个数的大小
func (x units) Less(i, j int) bool {
	return x[i] < x[j]
}

// 切片中连个值交换
func (x units) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

// 当hash环上没有数据时，提示错误
var errEmpty = errors.New("Hash环没有数据")

// 创建结构体保存一致性hash信息
type Constant struct {
	// hash环，key为hash，值存放节点信息
	circle map[uint32]string
	// 已经排序的节点hash切片
	sortedHashes units
	// 虚拟节点个数，用来增加hash平衡性
	VirtualNode int
	// map读写锁
	sync.RWMutex
}

// 创建一致性hash算法结构体，设置默认节点数量
func NewConstant() *Constant {
	return &Constant{
		// 初始化变量
		circle: make(map[uint32]string),
		// 设置虚拟节点个数
		VirtualNode: 20,
	}
}

// 自动生成key值
func (c *Constant) gennerateKey(element string, index int) string {
	return element + strconv.Itoa(index)
}

func (c *Constant) hashKey(key string) uint32 {
	if len(key) < 64 {
		// 声明一个数组长度为64
		var scratch [64]byte
		// 拷贝数据到数组中
		copy(scratch[:], key)
		// 使用IEEE 多项式返回数据的CRC-32校验和
		return crc32.ChecksumIEEE(scratch[:len(key)])
	}
	return crc32.ChecksumIEEE([]byte(key))
}

// 更新排序，便于查找
func (c *Constant) updateSortedHashes() {
	hashes := c.sortedHashes[:0]
	// 判断切片容纳量，是否过大，过大则重置
	if cap(c.sortedHashes)/(c.VirtualNode*4) > len(c.circle) {
		hashes = nil
	}
	// 添加hash
	for k := range c.circle {
		hashes = append(hashes, k)
	}

	// 对所有节点hash值进行排序
	// 方便之后二分查找
	sort.Sort(hashes)
	c.sortedHashes = hashes
}

// 向hash环添加节点
func (c *Constant) Add(element string) {
	c.Lock()
	defer c.Unlock()
	c.add(element)
}

// 删除一个节点
func (c *Constant) Remove(element string) {
	c.Lock()
	defer c.Unlock()
	c.remove(element)
}

// 删除节点
func (c *Constant) remove(element string) {
	for i := 0; i < c.VirtualNode; i++ {
		delete(c.circle, c.hashKey(c.gennerateKey(element, i)))
	}
	c.updateSortedHashes()
}

// 添加节点
func (c *Constant) add(element string) {
	for i := 0; i < c.VirtualNode; i++ {
		c.circle[c.hashKey(c.gennerateKey(element, i))] = element
	}
	// 更新排序
	c.updateSortedHashes()
}

// 顺时针查找最近的节点
func (c *Constant) search(key uint32) int {
	f := func(x int) bool {
		return c.sortedHashes[x] > key
	}
	// 使用二分查找算法来查找指定切片满足条件的最小值
	i := sort.Search(len(c.sortedHashes), f)

	if i >= len(c.sortedHashes) {
		i = 0
	}
	return i
}

// 根据数据表识，获取最近服务器信息
func (c *Constant) Get(name string) (string, error) {
	c.RLock()
	defer c.RUnlock()

	if len(c.circle) == 0 {
		return "", errEmpty
	}
	// 计算hash值
	key := c.hashKey(name)
	i := c.search(key)
	return c.circle[c.sortedHashes[i]], nil
}
