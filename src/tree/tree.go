package main

import "fmt"

type tree struct {
	val   int
	left  *tree
	right *tree
}

func (t *tree) insert(num int) {

	if t.val == 0 {
		t.val = num
	}

	for t != nil {
		if t.val < num {
			if t.right == nil {
				t.right = node(num)
				break
			} else {
				t = t.right
			}
		} else {
			if t.left == nil {
				t.left = node(num)
				break
			} else {
				t = t.left
			}
		}
	}
}

func (t *tree) init(num int) {
	t.val = num
	t.left = nil
	t.right = nil
}

func node(num int) *tree {
	t := new(tree)
	t.val = num
	t.left = nil
	t.right = nil
	return t
}

func (t *tree) trav() {
	if t == nil {
		return
	}
	fmt.Println(t.val)
	t.left.trav()
	t.right.trav()
}

func main() {
	var t tree
	t.init(5)
	t.insert(3)
	t.insert(6)
	t.insert(9)
	t.insert(1)
	t.insert(32)
	t.trav()
}
