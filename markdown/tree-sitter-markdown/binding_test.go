package tree_sitter_markdown_test

import (
	"context"
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
	tree_sitter_markdown "github.com/smacker/go-tree-sitter/markdown/tree-sitter-markdown"
	"github.com/stretchr/testify/assert"
)

func TestGrammar(t *testing.T) {
	assert := assert.New(t)

	n, err := sitter.ParseCtx(context.Background(), []byte("<foo />"), tree_sitter_markdown.GetLanguage())
	assert.NoError(err)
	assert.Equal(
		"(program (expression_statement (jsx_self_closing_element name: (identifier))))",
		n.String(),
	)
}
