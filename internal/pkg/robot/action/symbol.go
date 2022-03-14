package action

import (
	"fmt"
	"github.com/tsundata/assistant/internal/pkg/util/collection"
	"strings"
)

type Symbol interface{}

type BuiltinTypeSymbol struct {
	Name       string
	Type       Symbol
	ScopeLevel int
}

func NewBuiltinTypeSymbol(name string) *BuiltinTypeSymbol {
	return &BuiltinTypeSymbol{Name: name, ScopeLevel: 0}
}

func (s *BuiltinTypeSymbol) String() string {
	return fmt.Sprintf("<BuiltinTypeSymbol(name=%s)>", s.Name)
}

type OpcodeSymbol struct {
	Name       string
	ScopeLevel int
}

func NewOpcodeSymbol(name string) *OpcodeSymbol {
	return &OpcodeSymbol{Name: name, ScopeLevel: 0}
}

func (s *OpcodeSymbol) String() string {
	return fmt.Sprintf("<OpcodeSymbol(name=%s)>", s.Name)
}

type WebhookSymbol struct {
	Flag   string
	Secret string
}

func NewWebhookSymbol(flag string, secret string) *WebhookSymbol {
	return &WebhookSymbol{Flag: flag, Secret: secret}
}

func (s *WebhookSymbol) String() string {
	return fmt.Sprintf("<WebhookSymbol(flag=%s)>", s.Flag)
}

type CronSymbol struct {
	When string
}

func NewCronSymbol(when string) *CronSymbol {
	return &CronSymbol{When: when}
}

func (s *CronSymbol) String() string {
	return fmt.Sprintf("<CronSymbol(when=%s)>", s.When)
}

type ScopedSymbolTable struct {
	symbols        *collection.OrderedDict
	ScopeName      string
	ScopeLevel     int
	EnclosingScope *ScopedSymbolTable
}

func NewScopedSymbolTable(scopeName string, scopeLevel int, enclosingScope *ScopedSymbolTable) *ScopedSymbolTable {
	table := &ScopedSymbolTable{
		symbols:        collection.NewOrderedDict(),
		ScopeName:      scopeName,
		ScopeLevel:     scopeLevel,
		EnclosingScope: enclosingScope,
	}
	table.Insert(NewBuiltinTypeSymbol("INTEGER"))
	table.Insert(NewBuiltinTypeSymbol("FLOAT"))
	table.Insert(NewBuiltinTypeSymbol("BOOL"))
	table.Insert(NewBuiltinTypeSymbol("STRING"))
	table.Insert(NewBuiltinTypeSymbol("MESSAGE"))
	return table
}

func (t *ScopedSymbolTable) String() string {
	if t == nil {
		return ""
	}

	var lines []string

	lines = append(lines, fmt.Sprintf("Scope name : %s", t.ScopeName))
	lines = append(lines, fmt.Sprintf("Scope level : %d", t.ScopeLevel))

	if t.EnclosingScope != nil {
		lines = append(lines, fmt.Sprintf("Enclosing scope : %s", t.EnclosingScope.ScopeName))
	}

	lines = append(lines, "------------------------------------")
	lines = append(lines, "Scope (Scoped symbol table) contents")

	i := 0
	for v := range t.symbols.Iterate() {
		i++
		lines = append(lines, fmt.Sprintf("%6d: %v", i, v))
	}

	return fmt.Sprintf("\nSCOPE (SCOPED SYMBOL TABLE)\n===========================\n%s\n", strings.Join(lines, "\n"))
}

func (t *ScopedSymbolTable) Insert(symbol Symbol) {
	debugLog(fmt.Sprintf("Insert: %s\n", symbol))

	var name string
	if s, ok := symbol.(*BuiltinTypeSymbol); ok {
		name = s.Name
		s.ScopeLevel = t.ScopeLevel
		t.symbols.Set(name, s)
		return
	}
	if s, ok := symbol.(*OpcodeSymbol); ok {
		name = s.Name
		s.ScopeLevel = t.ScopeLevel
		t.symbols.Set(name, s)
		return
	}
}

func (t *ScopedSymbolTable) Lookup(name string, currentScopeOnly bool) Symbol {
	debugLog(fmt.Sprintf("Lookup: %s. (Scope name: %s)\n", name, t.ScopeName))

	s := t.symbols.Get(name)
	if s != nil {
		return s.(Symbol)
	}
	if currentScopeOnly {
		return nil
	}

	if t.EnclosingScope != nil {
		return t.EnclosingScope.Lookup(name, false)
	}
	return nil
}

type SemanticAnalyzer struct {
	CurrentScope *ScopedSymbolTable
	Webhook      *WebhookSymbol
	Cron         *CronSymbol
}

func NewSemanticAnalyzer() *SemanticAnalyzer {
	return &SemanticAnalyzer{CurrentScope: nil}
}

func (b *SemanticAnalyzer) error(errorCode ErrorCode, token *Token) error {
	return Error{
		ErrorCode: errorCode,
		Token:     token,
		Message:   fmt.Sprintf("%s -> %v", errorCode, token),
		Type:      SemanticErrorType,
	}
}

func (b *SemanticAnalyzer) Visit(node Ast) error {
	if n, ok := node.(*Program); ok {
		return b.VisitProgram(n)
	}
	if n, ok := node.(*Opcode); ok {
		return b.VisitOpcode(n)
	}
	if n, ok := node.(*IntegerConst); ok {
		return b.VisitIntegerConst(n)
	}
	if n, ok := node.(*StringConst); ok {
		return b.VisitStringConst(n)
	}
	if n, ok := node.(*BooleanConst); ok {
		return b.VisitBooleanConst(n)
	}
	if n, ok := node.(*MessageConst); ok {
		return b.VisitMessageConst(n)
	}
	if n, ok := node.(*Var); ok {
		return b.VisitVar(n)
	}
	if n, ok := node.(*NoOp); ok {
		return b.VisitNoOp(n)
	}
	return nil
}

func (b *SemanticAnalyzer) VisitProgram(node *Program) error {
	debugLog("ENTER scope: global")
	globalScope := NewScopedSymbolTable("global", 1, b.CurrentScope)
	b.CurrentScope = globalScope

	for _, item := range node.Statements {
		err := b.Visit(item)
		if err != nil {
			return err
		}
	}

	debugLog(globalScope.String())

	b.CurrentScope = b.CurrentScope.EnclosingScope
	debugLog("LEAVE scope: global")

	return nil
}

func (b *SemanticAnalyzer) VisitOpcode(node *Opcode) error {
	name := node.ID.(*Token).Value.(string)

	// Async opcode
	if name == "webhook" || name == "cron" {
		s := b.CurrentScope.Lookup(name, true)
		if s != nil {
			return b.error(RepeatOpcode, node.Token)
		}

		var args []string
		for _, item := range node.Expressions {
			if s, ok := item.(*StringConst); ok {
				args = append(args, s.Value)
			}
		}

		if name == "webhook" {
			if len(args) >= 2 {
				b.Webhook = NewWebhookSymbol(args[0], args[1])
			} else if len(args) >= 1 {
				b.Webhook = NewWebhookSymbol(args[0], "")
			} else {
				return b.error(ParameterType, node.Token)
			}
		}
		if name == "cron" {
			if len(args) >= 1 {
				b.Cron = NewCronSymbol(args[0])
			} else {
				return b.error(ParameterType, node.Token)
			}
		}
	}

	nodeSymbol := NewOpcodeSymbol(name)
	b.CurrentScope.Insert(nodeSymbol)
	return nil
}

func (b *SemanticAnalyzer) VisitIntegerConst(_ *IntegerConst) error {
	// pass
	return nil
}

func (b *SemanticAnalyzer) VisitStringConst(_ *StringConst) error {
	// pass
	return nil
}
func (b *SemanticAnalyzer) VisitMessageConst(_ *MessageConst) error {
	// pass
	return nil
}

func (b *SemanticAnalyzer) VisitBooleanConst(_ *BooleanConst) error {
	// pass
	return nil
}

func (b *SemanticAnalyzer) VisitVar(node *Var) error {
	varName := node.Value.(string)
	varSymbol := b.CurrentScope.Lookup(varName, false)

	if varSymbol == nil {
		return b.error(IdNotFound, node.Token)
	}
	return nil
}

func (b *SemanticAnalyzer) VisitNoOp(_ *NoOp) error {
	// pass
	return nil
}
