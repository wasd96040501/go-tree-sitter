package markdown

import (
	"context"

	sitter "github.com/smacker/go-tree-sitter"
)

type MarkdownTree struct {
	blockTree     *sitter.Tree
	inlineTrees   []*sitter.Tree
	inlineIndices map[uintptr]int
}

func (t *MarkdownTree) Edit(edit sitter.EditInput) {
	t.blockTree.Edit(edit)
	for _, tree := range t.inlineTrees {
		tree.Edit(edit)
	}
}

func (t *MarkdownTree) BlockTree() *sitter.Tree {
	return t.blockTree
}

func (t *MarkdownTree) InlineTree(parent *sitter.Node) *sitter.Tree {
	if parent == nil {
		return nil
	}

	index, ok := t.inlineIndices[parent.ID()]
	if ok {
		return t.inlineTrees[index]
	}

	return nil
}

func (t *MarkdownTree) InlineRootNode(parent *sitter.Node) *sitter.Node {
	tree := t.InlineTree(parent)
	if tree == nil {
		return nil
	}

	return tree.RootNode()
}

func (t *MarkdownTree) InlineTrees() []*sitter.Tree {
	return t.inlineTrees
}

func (t *MarkdownTree) Iter(f func(node *Node) bool) {
	root := t.blockTree.RootNode()
	t.iter(&Node{root, t.InlineRootNode(root)}, f)
}

func (t *MarkdownTree) iter(node *Node, f func(node *Node) bool) {
	goNext := f(node)
	if !goNext {
		return
	}

	childCount := node.ChildCount()
	for i := 0; i < int(childCount); i++ {
		child := node.Child(i)

		t.iter(&Node{Node: child, Inline: t.InlineRootNode(child)}, f)
	}
}

type Node struct {
	*sitter.Node
	Inline *sitter.Node
}

func ParseCtx(ctx context.Context, oldTree *MarkdownTree, content []byte) *MarkdownTree {
	return nil
}
