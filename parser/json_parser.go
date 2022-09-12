// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // JSON

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type JSONParser struct {
	*antlr.BaseParser
}

var jsonParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func jsonParserInit() {
	staticData := &jsonParserStaticData
	staticData.literalNames = []string{
		"", "'{'", "'}'", "'['", "']'", "','", "':'", "'true'", "'false'", "'null'",
		"'\"'", "", "", "'-'",
	}
	staticData.symbolicNames = []string{
		"", "LEFT_BRACKET", "RIGHT_BRACKET", "LEFT_SQUARE_BRACKET", "RIGHT_SQUARE_BRACKET",
		"COMMA", "COLORN", "TRUE", "FALSE", "NULL", "DOUBLE_QUOTE", "WHITE_SPACE",
		"LETTER", "MINUS", "NUMBER",
	}
	staticData.ruleNames = []string{
		"json", "value", "object", "members", "member", "array", "elements",
		"str", "num",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 14, 76, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 1, 0, 1, 0, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 28, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 3, 2, 36, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 43, 8,
		3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 55,
		8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 62, 8, 6, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 3, 7, 69, 8, 7, 1, 8, 1, 8, 1, 8, 3, 8, 74, 8, 8, 1, 8, 0, 0,
		9, 0, 2, 4, 6, 8, 10, 12, 14, 16, 0, 0, 78, 0, 18, 1, 0, 0, 0, 2, 27, 1,
		0, 0, 0, 4, 35, 1, 0, 0, 0, 6, 42, 1, 0, 0, 0, 8, 44, 1, 0, 0, 0, 10, 54,
		1, 0, 0, 0, 12, 61, 1, 0, 0, 0, 14, 68, 1, 0, 0, 0, 16, 73, 1, 0, 0, 0,
		18, 19, 3, 2, 1, 0, 19, 1, 1, 0, 0, 0, 20, 28, 3, 4, 2, 0, 21, 28, 3, 10,
		5, 0, 22, 28, 3, 14, 7, 0, 23, 28, 3, 16, 8, 0, 24, 28, 5, 7, 0, 0, 25,
		28, 5, 8, 0, 0, 26, 28, 5, 9, 0, 0, 27, 20, 1, 0, 0, 0, 27, 21, 1, 0, 0,
		0, 27, 22, 1, 0, 0, 0, 27, 23, 1, 0, 0, 0, 27, 24, 1, 0, 0, 0, 27, 25,
		1, 0, 0, 0, 27, 26, 1, 0, 0, 0, 28, 3, 1, 0, 0, 0, 29, 30, 5, 1, 0, 0,
		30, 36, 5, 2, 0, 0, 31, 32, 5, 1, 0, 0, 32, 33, 3, 6, 3, 0, 33, 34, 5,
		2, 0, 0, 34, 36, 1, 0, 0, 0, 35, 29, 1, 0, 0, 0, 35, 31, 1, 0, 0, 0, 36,
		5, 1, 0, 0, 0, 37, 43, 3, 8, 4, 0, 38, 39, 3, 8, 4, 0, 39, 40, 5, 5, 0,
		0, 40, 41, 3, 8, 4, 0, 41, 43, 1, 0, 0, 0, 42, 37, 1, 0, 0, 0, 42, 38,
		1, 0, 0, 0, 43, 7, 1, 0, 0, 0, 44, 45, 3, 14, 7, 0, 45, 46, 5, 6, 0, 0,
		46, 47, 3, 2, 1, 0, 47, 9, 1, 0, 0, 0, 48, 49, 5, 3, 0, 0, 49, 55, 5, 4,
		0, 0, 50, 51, 5, 3, 0, 0, 51, 52, 3, 12, 6, 0, 52, 53, 5, 4, 0, 0, 53,
		55, 1, 0, 0, 0, 54, 48, 1, 0, 0, 0, 54, 50, 1, 0, 0, 0, 55, 11, 1, 0, 0,
		0, 56, 62, 3, 2, 1, 0, 57, 58, 3, 2, 1, 0, 58, 59, 5, 5, 0, 0, 59, 60,
		3, 2, 1, 0, 60, 62, 1, 0, 0, 0, 61, 56, 1, 0, 0, 0, 61, 57, 1, 0, 0, 0,
		62, 13, 1, 0, 0, 0, 63, 64, 5, 10, 0, 0, 64, 65, 5, 12, 0, 0, 65, 69, 5,
		10, 0, 0, 66, 67, 5, 10, 0, 0, 67, 69, 5, 10, 0, 0, 68, 63, 1, 0, 0, 0,
		68, 66, 1, 0, 0, 0, 69, 15, 1, 0, 0, 0, 70, 74, 5, 14, 0, 0, 71, 72, 5,
		13, 0, 0, 72, 74, 5, 14, 0, 0, 73, 70, 1, 0, 0, 0, 73, 71, 1, 0, 0, 0,
		74, 17, 1, 0, 0, 0, 7, 27, 35, 42, 54, 61, 68, 73,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// JSONParserInit initializes any static state used to implement JSONParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewJSONParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func JSONParserInit() {
	staticData := &jsonParserStaticData
	staticData.once.Do(jsonParserInit)
}

// NewJSONParser produces a new parser instance for the optional input antlr.TokenStream.
func NewJSONParser(input antlr.TokenStream) *JSONParser {
	JSONParserInit()
	this := new(JSONParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &jsonParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "java-escape"

	return this
}

// JSONParser tokens.
const (
	JSONParserEOF                  = antlr.TokenEOF
	JSONParserLEFT_BRACKET         = 1
	JSONParserRIGHT_BRACKET        = 2
	JSONParserLEFT_SQUARE_BRACKET  = 3
	JSONParserRIGHT_SQUARE_BRACKET = 4
	JSONParserCOMMA                = 5
	JSONParserCOLORN               = 6
	JSONParserTRUE                 = 7
	JSONParserFALSE                = 8
	JSONParserNULL                 = 9
	JSONParserDOUBLE_QUOTE         = 10
	JSONParserWHITE_SPACE          = 11
	JSONParserLETTER               = 12
	JSONParserMINUS                = 13
	JSONParserNUMBER               = 14
)

// JSONParser rules.
const (
	JSONParserRULE_json     = 0
	JSONParserRULE_value    = 1
	JSONParserRULE_object   = 2
	JSONParserRULE_members  = 3
	JSONParserRULE_member   = 4
	JSONParserRULE_array    = 5
	JSONParserRULE_elements = 6
	JSONParserRULE_str      = 7
	JSONParserRULE_num      = 8
)

// IJsonContext is an interface to support dynamic dispatch.
type IJsonContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsJsonContext differentiates from other interfaces.
	IsJsonContext()
}

type JsonContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJsonContext() *JsonContext {
	var p = new(JsonContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JSONParserRULE_json
	return p
}

func (*JsonContext) IsJsonContext() {}

func NewJsonContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JsonContext {
	var p = new(JsonContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JSONParserRULE_json

	return p
}

func (s *JsonContext) GetParser() antlr.Parser { return s.parser }

func (s *JsonContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *JsonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JsonContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JsonContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JSONVisitor:
		return t.VisitJson(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JSONParser) Json() (localctx IJsonContext) {
	this := p
	_ = this

	localctx = NewJsonContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, JSONParserRULE_json)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(18)
		p.Value()
	}

	return localctx
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JSONParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JSONParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) Object() IObjectContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObjectContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IObjectContext)
}

func (s *ValueContext) Array() IArrayContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayContext)
}

func (s *ValueContext) Str() IStrContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStrContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStrContext)
}

func (s *ValueContext) Num() INumContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumContext)
}

func (s *ValueContext) TRUE() antlr.TerminalNode {
	return s.GetToken(JSONParserTRUE, 0)
}

func (s *ValueContext) FALSE() antlr.TerminalNode {
	return s.GetToken(JSONParserFALSE, 0)
}

func (s *ValueContext) NULL() antlr.TerminalNode {
	return s.GetToken(JSONParserNULL, 0)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JSONVisitor:
		return t.VisitValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JSONParser) Value() (localctx IValueContext) {
	this := p
	_ = this

	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, JSONParserRULE_value)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(27)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case JSONParserLEFT_BRACKET:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(20)
			p.Object()
		}

	case JSONParserLEFT_SQUARE_BRACKET:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(21)
			p.Array()
		}

	case JSONParserDOUBLE_QUOTE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(22)
			p.Str()
		}

	case JSONParserMINUS, JSONParserNUMBER:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(23)
			p.Num()
		}

	case JSONParserTRUE:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(24)
			p.Match(JSONParserTRUE)
		}

	case JSONParserFALSE:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(25)
			p.Match(JSONParserFALSE)
		}

	case JSONParserNULL:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(26)
			p.Match(JSONParserNULL)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IObjectContext is an interface to support dynamic dispatch.
type IObjectContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObjectContext differentiates from other interfaces.
	IsObjectContext()
}

type ObjectContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjectContext() *ObjectContext {
	var p = new(ObjectContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JSONParserRULE_object
	return p
}

func (*ObjectContext) IsObjectContext() {}

func NewObjectContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectContext {
	var p = new(ObjectContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JSONParserRULE_object

	return p
}

func (s *ObjectContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectContext) LEFT_BRACKET() antlr.TerminalNode {
	return s.GetToken(JSONParserLEFT_BRACKET, 0)
}

func (s *ObjectContext) RIGHT_BRACKET() antlr.TerminalNode {
	return s.GetToken(JSONParserRIGHT_BRACKET, 0)
}

func (s *ObjectContext) Members() IMembersContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMembersContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMembersContext)
}

func (s *ObjectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjectContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JSONVisitor:
		return t.VisitObject(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JSONParser) Object() (localctx IObjectContext) {
	this := p
	_ = this

	localctx = NewObjectContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, JSONParserRULE_object)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(35)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(29)
			p.Match(JSONParserLEFT_BRACKET)
		}
		{
			p.SetState(30)
			p.Match(JSONParserRIGHT_BRACKET)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(31)
			p.Match(JSONParserLEFT_BRACKET)
		}
		{
			p.SetState(32)
			p.Members()
		}
		{
			p.SetState(33)
			p.Match(JSONParserRIGHT_BRACKET)
		}

	}

	return localctx
}

// IMembersContext is an interface to support dynamic dispatch.
type IMembersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMembersContext differentiates from other interfaces.
	IsMembersContext()
}

type MembersContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMembersContext() *MembersContext {
	var p = new(MembersContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JSONParserRULE_members
	return p
}

func (*MembersContext) IsMembersContext() {}

func NewMembersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MembersContext {
	var p = new(MembersContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JSONParserRULE_members

	return p
}

func (s *MembersContext) GetParser() antlr.Parser { return s.parser }

func (s *MembersContext) AllMember() []IMemberContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMemberContext); ok {
			len++
		}
	}

	tst := make([]IMemberContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMemberContext); ok {
			tst[i] = t.(IMemberContext)
			i++
		}
	}

	return tst
}

func (s *MembersContext) Member(i int) IMemberContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMemberContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMemberContext)
}

func (s *MembersContext) COMMA() antlr.TerminalNode {
	return s.GetToken(JSONParserCOMMA, 0)
}

func (s *MembersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MembersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MembersContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JSONVisitor:
		return t.VisitMembers(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JSONParser) Members() (localctx IMembersContext) {
	this := p
	_ = this

	localctx = NewMembersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, JSONParserRULE_members)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(42)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(37)
			p.Member()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(38)
			p.Member()
		}
		{
			p.SetState(39)
			p.Match(JSONParserCOMMA)
		}
		{
			p.SetState(40)
			p.Member()
		}

	}

	return localctx
}

// IMemberContext is an interface to support dynamic dispatch.
type IMemberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMemberContext differentiates from other interfaces.
	IsMemberContext()
}

type MemberContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMemberContext() *MemberContext {
	var p = new(MemberContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JSONParserRULE_member
	return p
}

func (*MemberContext) IsMemberContext() {}

func NewMemberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MemberContext {
	var p = new(MemberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JSONParserRULE_member

	return p
}

func (s *MemberContext) GetParser() antlr.Parser { return s.parser }

func (s *MemberContext) Str() IStrContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStrContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStrContext)
}

func (s *MemberContext) COLORN() antlr.TerminalNode {
	return s.GetToken(JSONParserCOLORN, 0)
}

func (s *MemberContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *MemberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MemberContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JSONVisitor:
		return t.VisitMember(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JSONParser) Member() (localctx IMemberContext) {
	this := p
	_ = this

	localctx = NewMemberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, JSONParserRULE_member)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(44)
		p.Str()
	}
	{
		p.SetState(45)
		p.Match(JSONParserCOLORN)
	}
	{
		p.SetState(46)
		p.Value()
	}

	return localctx
}

// IArrayContext is an interface to support dynamic dispatch.
type IArrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayContext differentiates from other interfaces.
	IsArrayContext()
}

type ArrayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayContext() *ArrayContext {
	var p = new(ArrayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JSONParserRULE_array
	return p
}

func (*ArrayContext) IsArrayContext() {}

func NewArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayContext {
	var p = new(ArrayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JSONParserRULE_array

	return p
}

func (s *ArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayContext) LEFT_SQUARE_BRACKET() antlr.TerminalNode {
	return s.GetToken(JSONParserLEFT_SQUARE_BRACKET, 0)
}

func (s *ArrayContext) RIGHT_SQUARE_BRACKET() antlr.TerminalNode {
	return s.GetToken(JSONParserRIGHT_SQUARE_BRACKET, 0)
}

func (s *ArrayContext) Elements() IElementsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElementsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElementsContext)
}

func (s *ArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JSONVisitor:
		return t.VisitArray(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JSONParser) Array() (localctx IArrayContext) {
	this := p
	_ = this

	localctx = NewArrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, JSONParserRULE_array)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(48)
			p.Match(JSONParserLEFT_SQUARE_BRACKET)
		}
		{
			p.SetState(49)
			p.Match(JSONParserRIGHT_SQUARE_BRACKET)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(50)
			p.Match(JSONParserLEFT_SQUARE_BRACKET)
		}
		{
			p.SetState(51)
			p.Elements()
		}
		{
			p.SetState(52)
			p.Match(JSONParserRIGHT_SQUARE_BRACKET)
		}

	}

	return localctx
}

// IElementsContext is an interface to support dynamic dispatch.
type IElementsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElementsContext differentiates from other interfaces.
	IsElementsContext()
}

type ElementsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElementsContext() *ElementsContext {
	var p = new(ElementsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JSONParserRULE_elements
	return p
}

func (*ElementsContext) IsElementsContext() {}

func NewElementsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElementsContext {
	var p = new(ElementsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JSONParserRULE_elements

	return p
}

func (s *ElementsContext) GetParser() antlr.Parser { return s.parser }

func (s *ElementsContext) AllValue() []IValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IValueContext); ok {
			len++
		}
	}

	tst := make([]IValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IValueContext); ok {
			tst[i] = t.(IValueContext)
			i++
		}
	}

	return tst
}

func (s *ElementsContext) Value(i int) IValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ElementsContext) COMMA() antlr.TerminalNode {
	return s.GetToken(JSONParserCOMMA, 0)
}

func (s *ElementsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElementsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElementsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JSONVisitor:
		return t.VisitElements(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JSONParser) Elements() (localctx IElementsContext) {
	this := p
	_ = this

	localctx = NewElementsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, JSONParserRULE_elements)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(56)
			p.Value()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(57)
			p.Value()
		}
		{
			p.SetState(58)
			p.Match(JSONParserCOMMA)
		}
		{
			p.SetState(59)
			p.Value()
		}

	}

	return localctx
}

// IStrContext is an interface to support dynamic dispatch.
type IStrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStrContext differentiates from other interfaces.
	IsStrContext()
}

type StrContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStrContext() *StrContext {
	var p = new(StrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JSONParserRULE_str
	return p
}

func (*StrContext) IsStrContext() {}

func NewStrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StrContext {
	var p = new(StrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JSONParserRULE_str

	return p
}

func (s *StrContext) GetParser() antlr.Parser { return s.parser }

func (s *StrContext) AllDOUBLE_QUOTE() []antlr.TerminalNode {
	return s.GetTokens(JSONParserDOUBLE_QUOTE)
}

func (s *StrContext) DOUBLE_QUOTE(i int) antlr.TerminalNode {
	return s.GetToken(JSONParserDOUBLE_QUOTE, i)
}

func (s *StrContext) LETTER() antlr.TerminalNode {
	return s.GetToken(JSONParserLETTER, 0)
}

func (s *StrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JSONVisitor:
		return t.VisitStr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JSONParser) Str() (localctx IStrContext) {
	this := p
	_ = this

	localctx = NewStrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, JSONParserRULE_str)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(68)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(63)
			p.Match(JSONParserDOUBLE_QUOTE)
		}
		{
			p.SetState(64)
			p.Match(JSONParserLETTER)
		}
		{
			p.SetState(65)
			p.Match(JSONParserDOUBLE_QUOTE)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(66)
			p.Match(JSONParserDOUBLE_QUOTE)
		}
		{
			p.SetState(67)
			p.Match(JSONParserDOUBLE_QUOTE)
		}

	}

	return localctx
}

// INumContext is an interface to support dynamic dispatch.
type INumContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumContext differentiates from other interfaces.
	IsNumContext()
}

type NumContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumContext() *NumContext {
	var p = new(NumContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JSONParserRULE_num
	return p
}

func (*NumContext) IsNumContext() {}

func NewNumContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumContext {
	var p = new(NumContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JSONParserRULE_num

	return p
}

func (s *NumContext) GetParser() antlr.Parser { return s.parser }

func (s *NumContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(JSONParserNUMBER, 0)
}

func (s *NumContext) MINUS() antlr.TerminalNode {
	return s.GetToken(JSONParserMINUS, 0)
}

func (s *NumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JSONVisitor:
		return t.VisitNum(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JSONParser) Num() (localctx INumContext) {
	this := p
	_ = this

	localctx = NewNumContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, JSONParserRULE_num)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(73)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case JSONParserNUMBER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(70)
			p.Match(JSONParserNUMBER)
		}

	case JSONParserMINUS:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(71)
			p.Match(JSONParserMINUS)
		}
		{
			p.SetState(72)
			p.Match(JSONParserNUMBER)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
