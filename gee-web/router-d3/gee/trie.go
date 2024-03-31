package gee

import "strings"

type node struct {
	pattern  string  // 待匹配路由 例如： /p/:lang   冒号表示一部分是变量
	part     string  //路由中的一部分 例如： :lang
	children []*node // 子节点
	isWild   bool    // 是否精准匹配，是否叶子结点
}

func (n *node) matchChlid(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

func (n *node) matchAllChlid(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
	}
	part := parts[height]
	child := n.matchChlid(part)
	if child == nil {
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}
	part := parts[height]
	children := n.matchAllChlid(part)

	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
		}
		return result
	}
	return nil
}

func (n *node) travel(list *([]*node)) {
	if n.pattern != "" {
		*list = append(*list, n)
	}
	for _, child := range n.children {
		child.travel(list)
	}
}
