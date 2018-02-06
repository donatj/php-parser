package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
	"github.com/z7zmey/php-parser/walker"
)

// While node
type While struct {
	Token token.Token
	Cond  node.Node
	Stmt  node.Node
}

// NewWhile node constuctor
func NewWhile(Token token.Token, Cond node.Node, Stmt node.Node) *While {
	return &While{
		Token,
		Cond,
		Stmt,
	}
}

// Attributes returns node attributes as map
func (n *While) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *While) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Cond != nil {
		vv := v.GetChildrenVisitor("Cond")
		n.Cond.Walk(vv)
	}

	if n.Stmt != nil {
		vv := v.GetChildrenVisitor("Stmt")
		n.Stmt.Walk(vv)
	}

	v.LeaveNode(n)
}