/*
 * @Author: Daiming Liu (xingrufeng)
 */

package meteor

import (
	"errors"
	"strconv"
	"sync"
	"time"
)

// sign		second			nodeid		seq num		rand seq
// 1 bit	29bit			20bit		11bit		3bit

const (
	initSecond int64 = 1607558400                             // start date：2020-12-10
	nodeBits   uint8 = 20                                     // node num
	seqBits    uint8 = 11                                     // seq num
	randBits   uint8 = 3                                      // over seq
	secBits    uint8 = 64 - 1 - nodeBits - seqBits - randBits // seconds num
	maxSecond  int64 = -1 ^ (-1 << secBits)                   // max seconds
	maxNode    int64 = -1 ^ (-1 << nodeBits)                  // max nodeid
	maxSeq     int64 = -1 ^ (-1 << seqBits)                   // max seq num
	maxRand    int64 = -1 ^ (-1 << randBits)                  // max over num
	timeShift  uint8 = nodeBits + seqBits + randBits
	nodeShift  uint8 = seqBits + randBits
	seqShift   uint8 = randBits
)

var curSecond int64 = time.Now().Unix()

// Node contains infomation of generator id
type Node struct {
	sync.Mutex
	second  int64
	nodeID  int64
	seqNum  int64
	randNum int64
	seed    uint16
}

// NewNode create a new node id,if the service is reload,nodeID should increase
func NewNode(nodeID int64) (*Node, error) {
	if nodeID < 0 || nodeID > maxNode {
		return nil, errors.New("Node id must between 0 and " + strconv.FormatInt(maxNode, 10))
	}
	return &Node{
		second:  time.Now().Unix() - initSecond,
		nodeID:  nodeID,
		seqNum:  0,
		randNum: 0,
		seed:    1,
	}, nil
}

// Generate generate a uniq id
func (node *Node) Generate() (int64, error) {
	node.Lock()
	defer node.Unlock()
	node.seqNum = (node.seqNum + 1) & maxSeq
	node.randNum = node.rand(maxRand)
	if node.seqNum == 0 {
		node.second = (node.second + 1) & maxSecond
		if node.second == 0 {
			return 0, errors.New("Seconds over flow. The max seconds is " + strconv.FormatInt(maxSecond, 10))
		}
	}
	return node.second<<timeShift | node.nodeID<<nodeShift | node.seqNum<<seqShift | node.randNum, nil
}

// get rand num; Xorshift: http://www.retroprogramming.com/2017/07/xorshift-pseudorandom-numbers-in-z80.html
func (node *Node) rand(maxNum int64) int64 {
	node.seed ^= node.seed << 7
	node.seed ^= node.seed >> 9
	node.seed ^= node.seed << 8
	return int64(node.seed) % maxNum
}
