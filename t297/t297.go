package t297

import (
	"fmt"
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	sb := strings.Builder{}
	toString(root, &sb)
	return sb.String()
}

func toString(root *TreeNode, sb *strings.Builder) {
	if root == nil {
		sb.WriteString("{}")
		return
	}
	sb.WriteString("{")
	sb.WriteString(strconv.Itoa(root.Val))
	toString(root.Left, sb)
	toString(root.Right, sb)
	sb.WriteString("}")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	return toTree([]rune(data))
}

func toTree(data []rune) *TreeNode {
	if data[1] == '}' {
		return nil
	}
	node := &TreeNode{}
	data = data[1 : len(data)-1] // 去掉首尾括号
	lStart, lEnd, count := 0, 0, 0
	for i, ch := range data {
		if ch == '{' {
			count++
			if lStart == 0 {
				lStart = i
			}
		}
		if ch == '}' {
			count--
			if count == 0 {
				lEnd = i + 1
				break
			}
		}
	}
	fmt.Printf("%d,%d,%d,%v\n", lStart, lEnd, len(data), string(data))
	fmt.Printf("%v,%v,%v\n", string(data[:lStart]), string(data[lStart:lEnd]), string(data[lEnd:len(data)]))

	// 左子树
	node.Left = toTree(data[lStart:lEnd])
	// 右子树
	node.Right = toTree(data[lEnd : len(data)])

	node.Val, _ = strconv.Atoi(string(data[:lStart]))
	return node
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
