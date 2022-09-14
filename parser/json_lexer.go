// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type JSONLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var jsonlexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func jsonlexerLexerInit() {
	staticData := &jsonlexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "'{'", "'}'", "'['", "']'", "','", "':'", "'true'", "'false'", "'null'",
		"'\"'",
	}
	staticData.symbolicNames = []string{
		"", "LEFT_BRACKET", "RIGHT_BRACKET", "LEFT_SQUARE_BRACKET", "RIGHT_SQUARE_BRACKET",
		"COMMA", "COLORN", "TRUE", "FALSE", "NULL", "DOUBLE_QUOTE", "WHITE_SPACE",
		"STRING", "INTEGER",
	}
	staticData.ruleNames = []string{
		"LEFT_BRACKET", "RIGHT_BRACKET", "LEFT_SQUARE_BRACKET", "RIGHT_SQUARE_BRACKET",
		"COMMA", "COLORN", "TRUE", "FALSE", "NULL", "DOUBLE_QUOTE", "WHITE_SPACE",
		"STRING", "DIGIT", "INTEGER",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 13, 95, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 1, 0, 1, 0, 1, 1,
		1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 9, 1, 9, 1, 10, 4, 10, 61, 8, 10, 11, 10, 12, 10, 62, 1, 10, 1,
		10, 1, 11, 1, 11, 5, 11, 69, 8, 11, 10, 11, 12, 11, 72, 9, 11, 1, 11, 1,
		11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 5, 13, 81, 8, 13, 10, 13, 12, 13,
		84, 9, 13, 1, 13, 1, 13, 1, 13, 5, 13, 89, 8, 13, 10, 13, 12, 13, 92, 9,
		13, 3, 13, 94, 8, 13, 0, 0, 14, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13,
		7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 0, 27, 13, 1, 0, 2, 3, 0,
		9, 10, 13, 13, 32, 32, 3, 0, 48, 57, 65, 90, 97, 122, 99, 0, 1, 1, 0, 0,
		0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0,
		0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0,
		0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 27, 1,
		0, 0, 0, 1, 29, 1, 0, 0, 0, 3, 31, 1, 0, 0, 0, 5, 33, 1, 0, 0, 0, 7, 35,
		1, 0, 0, 0, 9, 37, 1, 0, 0, 0, 11, 39, 1, 0, 0, 0, 13, 41, 1, 0, 0, 0,
		15, 46, 1, 0, 0, 0, 17, 52, 1, 0, 0, 0, 19, 57, 1, 0, 0, 0, 21, 60, 1,
		0, 0, 0, 23, 66, 1, 0, 0, 0, 25, 75, 1, 0, 0, 0, 27, 93, 1, 0, 0, 0, 29,
		30, 5, 123, 0, 0, 30, 2, 1, 0, 0, 0, 31, 32, 5, 125, 0, 0, 32, 4, 1, 0,
		0, 0, 33, 34, 5, 91, 0, 0, 34, 6, 1, 0, 0, 0, 35, 36, 5, 93, 0, 0, 36,
		8, 1, 0, 0, 0, 37, 38, 5, 44, 0, 0, 38, 10, 1, 0, 0, 0, 39, 40, 5, 58,
		0, 0, 40, 12, 1, 0, 0, 0, 41, 42, 5, 116, 0, 0, 42, 43, 5, 114, 0, 0, 43,
		44, 5, 117, 0, 0, 44, 45, 5, 101, 0, 0, 45, 14, 1, 0, 0, 0, 46, 47, 5,
		102, 0, 0, 47, 48, 5, 97, 0, 0, 48, 49, 5, 108, 0, 0, 49, 50, 5, 115, 0,
		0, 50, 51, 5, 101, 0, 0, 51, 16, 1, 0, 0, 0, 52, 53, 5, 110, 0, 0, 53,
		54, 5, 117, 0, 0, 54, 55, 5, 108, 0, 0, 55, 56, 5, 108, 0, 0, 56, 18, 1,
		0, 0, 0, 57, 58, 5, 34, 0, 0, 58, 20, 1, 0, 0, 0, 59, 61, 7, 0, 0, 0, 60,
		59, 1, 0, 0, 0, 61, 62, 1, 0, 0, 0, 62, 60, 1, 0, 0, 0, 62, 63, 1, 0, 0,
		0, 63, 64, 1, 0, 0, 0, 64, 65, 6, 10, 0, 0, 65, 22, 1, 0, 0, 0, 66, 70,
		5, 34, 0, 0, 67, 69, 7, 1, 0, 0, 68, 67, 1, 0, 0, 0, 69, 72, 1, 0, 0, 0,
		70, 68, 1, 0, 0, 0, 70, 71, 1, 0, 0, 0, 71, 73, 1, 0, 0, 0, 72, 70, 1,
		0, 0, 0, 73, 74, 5, 34, 0, 0, 74, 24, 1, 0, 0, 0, 75, 76, 2, 48, 57, 0,
		76, 26, 1, 0, 0, 0, 77, 94, 5, 48, 0, 0, 78, 82, 2, 49, 57, 0, 79, 81,
		3, 25, 12, 0, 80, 79, 1, 0, 0, 0, 81, 84, 1, 0, 0, 0, 82, 80, 1, 0, 0,
		0, 82, 83, 1, 0, 0, 0, 83, 94, 1, 0, 0, 0, 84, 82, 1, 0, 0, 0, 85, 86,
		5, 45, 0, 0, 86, 90, 2, 49, 57, 0, 87, 89, 3, 25, 12, 0, 88, 87, 1, 0,
		0, 0, 89, 92, 1, 0, 0, 0, 90, 88, 1, 0, 0, 0, 90, 91, 1, 0, 0, 0, 91, 94,
		1, 0, 0, 0, 92, 90, 1, 0, 0, 0, 93, 77, 1, 0, 0, 0, 93, 78, 1, 0, 0, 0,
		93, 85, 1, 0, 0, 0, 94, 28, 1, 0, 0, 0, 6, 0, 62, 70, 82, 90, 93, 1, 6,
		0, 0,
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

// JSONLexerInit initializes any static state used to implement JSONLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewJSONLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func JSONLexerInit() {
	staticData := &jsonlexerLexerStaticData
	staticData.once.Do(jsonlexerLexerInit)
}

// NewJSONLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewJSONLexer(input antlr.CharStream) *JSONLexer {
	JSONLexerInit()
	l := new(JSONLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &jsonlexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "JSON.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// JSONLexer tokens.
const (
	JSONLexerLEFT_BRACKET         = 1
	JSONLexerRIGHT_BRACKET        = 2
	JSONLexerLEFT_SQUARE_BRACKET  = 3
	JSONLexerRIGHT_SQUARE_BRACKET = 4
	JSONLexerCOMMA                = 5
	JSONLexerCOLORN               = 6
	JSONLexerTRUE                 = 7
	JSONLexerFALSE                = 8
	JSONLexerNULL                 = 9
	JSONLexerDOUBLE_QUOTE         = 10
	JSONLexerWHITE_SPACE          = 11
	JSONLexerSTRING               = 12
	JSONLexerINTEGER              = 13
)
