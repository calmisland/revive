package test

import (
	"testing"

	"github.com/calmisland/go-revive/rule"
)

// BoolLiteral rule.
func TestBoolLiteral(t *testing.T) {
	testRule(t, "bool-literal-in-expr", &rule.BoolLiteralRule{})
}
