package solutions

import (
//"reflect"
)

// all O(1)
type CntNode struct {
	Cnt     int
	NodeMap map[string]bool
	Pre     *CntNode
	Next    *CntNode
}

type NodeData struct {
	CntPtr *CntNode
	Key    string
	Cnt    int
}

type AllOne struct {
	IdxMap map[string]*NodeData
	Head   *CntNode
	Tail   *CntNode
}

/** Initialize your data structure here. */
func ConstructorAllOne() AllOne {
	ret := AllOne{}
	ret.IdxMap = make(map[string]*NodeData)
	ret.Head = &CntNode{}
	ret.Tail = &CntNode{}
	ret.Head.Next = ret.Tail
	ret.Tail.Pre = ret.Head

	return ret
}

/** Inserts a new key <Key> with value 1. Or increments an existing key by 1. */
func (this *AllOne) Inc(key string) {
	knode, ok := this.IdxMap[key]
	if ok {
		knode.Cnt++
		cntNode := knode.CntPtr
		stepCntNode := cntNode.Pre
		if (cntNode.Cnt + 1) != stepCntNode.Cnt {
			// insert
			newCntNode := &CntNode{
				Cnt: cntNode.Cnt + 1,
			}
			newCntNode.NodeMap = make(map[string]bool)
			newCntNode.NodeMap[key] = true

			newCntNode.Pre = cntNode.Pre
			newCntNode.Next = cntNode
			cntNode.Pre.Next = newCntNode
			cntNode.Pre = newCntNode
			knode.CntPtr = newCntNode
		} else {
			// add
			stepCntNode.NodeMap[key] = true
			knode.CntPtr = stepCntNode
		}
		delete(cntNode.NodeMap, key)
		if len(cntNode.NodeMap) == 0 {
			cntNode.Pre.Next = cntNode.Next
			cntNode.Next.Pre = cntNode.Pre
		}

		return
	}

	nd := &NodeData{
		Key: key,
		Cnt: 1,
	}
	// insert to tail
	stepCntNode := this.Tail.Pre
	if stepCntNode.Cnt == 1 {
		stepCntNode.NodeMap[key] = true
		nd.CntPtr = stepCntNode
	} else {
		newCntNode := &CntNode{
			Cnt: 1,
		}
		newCntNode.NodeMap = make(map[string]bool)
		newCntNode.NodeMap[key] = true

		newCntNode.Pre = stepCntNode
		newCntNode.Next = this.Tail
		stepCntNode.Next = newCntNode
		this.Tail.Pre = newCntNode

		nd.CntPtr = newCntNode
	}

	this.IdxMap[key] = nd
}

/** Decrements an existing key by 1. If Key's value is 1, remove it from the data structure. */
func (this *AllOne) Dec(key string) {
	knode, ok := this.IdxMap[key]
	if ok {
		knode.Cnt--
		cntNode := knode.CntPtr

		if knode.Cnt == 0 {
			delete(this.IdxMap, key)
		}

		downCntNode := cntNode.Next
		if cntNode.Cnt > 0 {
			if (cntNode.Cnt - 1) != downCntNode.Cnt {
				// insert
				newCntNode := &CntNode{
					Cnt: cntNode.Cnt - 1,
				}
				newCntNode.NodeMap = make(map[string]bool)
				newCntNode.NodeMap[key] = true

				newCntNode.Pre = cntNode
				newCntNode.Next = cntNode.Next
				cntNode.Next.Pre = newCntNode
				cntNode.Next = newCntNode
				knode.CntPtr = newCntNode
			} else if downCntNode != this.Tail {
				// add
				downCntNode.NodeMap[key] = true
				knode.CntPtr = downCntNode
			}
		}

		delete(cntNode.NodeMap, key)
		if len(cntNode.NodeMap) == 0 {
			cntNode.Pre.Next = cntNode.Next
			cntNode.Next.Pre = cntNode.Pre
		}
	}
}

/** Returns one of the keys with maximal value. */
func (this *AllOne) GetMaxKey() string {
	if this.Head.Next != this.Tail {
		nmap := this.Head.Next.NodeMap
		ret := ""
		for k, _ := range nmap {
			ret = k
			break
		}
		return ret
	}

	return ""
}

/** Returns one of the keys with Minimal value. */
func (this *AllOne) GetMinKey() string {
	if this.Tail.Pre != this.Head {
		nmap := this.Tail.Pre.NodeMap
		ret := ""
		for k, _ := range nmap {
			ret = k
			break
		}
		return ret
	}

	return ""
}

func init() {
	desc := `
实现一个数据结构支持以下操作：

Inc(key) - 插入一个新的值为 1 的 key。或者使一个存在的 key 增加一，保证 key 不为空字符串。
Dec(key) - 如果这个 key 的值是 1，那么把他从数据结构中移除掉。否者使一个存在的 key 值减一。如果这个 key 不存在，这个函数不做任何事情。key 保证不为空字符串。
GetMaxKey() - 返回 key 中值最大的任意一个。如果没有元素存在，返回一个空字符串""。
GetMinKey() - 返回 key 中值最小的任意一个。如果没有元素存在，返回一个空字符串""。
挑战：以 O(1) 的时间复杂度实现所有操作。
	`
	sol := Solution{
		Title: "全 O(1) 的数据结构",
		Desc:  desc,
		//Method: reflect.ValueOf(twoSum),
		Tests: make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{[]int{2, 7, 11, 15}, 9}
	a.Output = []interface{}{[]int{0, 1}}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0432"] = sol
}
