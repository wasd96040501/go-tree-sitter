package markdown_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/smacker/go-tree-sitter/markdown"
)

func TestMarkdown(t *testing.T) {
	content := "# Hello\n- This is a image: ![image](https://example.com/image.jpg \"a image\")"
	tree, err := markdown.ParseCtx(context.Background(), nil, []byte(content))
	if err != nil {
		t.Fatalf("parse failed. err=%v", err)
	}

	fmt.Printf("%+v\n", tree.BlockTree().RootNode().String())
	for _, t := range tree.InlineTrees() {
		fmt.Printf("%+v\n", t.RootNode().String())
	}
}
