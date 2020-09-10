package tree

import (
	"errors"
	"sort"
)

type Record struct {
	ID     int
	Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

type NodeSlice []*Node

func (n NodeSlice) Len() int           { return len(n) }
func (n NodeSlice) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n NodeSlice) Less(i, j int) bool { return n[i].ID < n[j].ID }

func Build(records []Record) (root *Node, err error) {
	if len(records) == 0 {
		return nil, nil
	}

	r := make(map[int]Record)
	m := make(map[int]*Node)

	// Check for duplicates
	// Check parent ID, shout be the same if node is 0 or
	// should have a lower ID
	for _, item := range records {
		_, prs := m[item.ID]

		if item.Parent >= item.ID && item.ID != 0 {
			err = errors.New("Parent ID")
			break
		}

		if !prs {
			r[item.ID] = item
			m[item.ID] = &Node{ID: item.ID}
		} else {
			err = errors.New("Repeated node")
			break
		}
	}

	if err != nil {
		return
	}
	root, ok := m[0]

	if len(m) == 1 && ok {
		return root, nil
	}

	index := 0

	for index < len(records) {
		for _, v := range m {
			if v.ID != 0 && r[v.ID].Parent == index {
				m[index].Children = append(m[index].Children, v)
				sort.Sort(NodeSlice(m[index].Children))
			}
		}

		index++
	}

	return m[0], nil
}
