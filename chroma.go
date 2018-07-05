package carrera

import (
	"log"

	"github.com/alecthomas/chroma"
)

var (
	emptyStyle []chroma.TokenType = nil

	allStrings = []chroma.TokenType{
		chroma.String,
		chroma.StringAffix,
		chroma.StringBacktick,
		chroma.StringChar,
		chroma.StringDelimiter,
		chroma.StringDoc,
		chroma.StringDouble,
		chroma.StringEscape,
		chroma.StringHeredoc,
		chroma.StringInterpol,
		chroma.StringOther,
		chroma.StringRegex,
		chroma.StringSingle,
		chroma.StringSymbol,
	}
)

type Style struct {
	Color  Color
	Tokens []chroma.TokenType
}

func NewStyle(color Color, tokens ...chroma.TokenType) Style {
	return Style{Color: color, Tokens: tokens}
}

func (s Style) Match(tokenType chroma.TokenType) bool {
	for _, candidate := range s.Tokens {
		if tokenType == candidate {
			return true
		}
	}
	return false
}

type StyleArray []Style

func (s StyleArray) Color(tokenType chroma.TokenType) Color {
	for _, style := range s {
		if style.Match(tokenType) {
			return style.Color
		}
	}
	return White
}

var DefaultStyles = StyleArray([]Style{
	NewStyle(Green,
		chroma.String,
		chroma.StringAffix,
		chroma.StringBacktick,
		chroma.StringChar,
		chroma.StringDelimiter,
		chroma.StringDoc,
		chroma.StringDouble,
		chroma.StringEscape,
		chroma.StringHeredoc,
		chroma.StringInterpol,
		chroma.StringOther,
		chroma.StringRegex,
		chroma.StringSingle,
		chroma.StringSymbol,
	),
	NewStyle(Red,
		chroma.Number,
		chroma.NumberBin,
		chroma.NumberFloat,
		chroma.NumberHex,
		chroma.NumberInteger,
		chroma.NumberIntegerLong,
		chroma.NumberOct,
	),
	NewStyle(Blue,
		chroma.Keyword,
		chroma.KeywordConstant,
		chroma.KeywordDeclaration,
		chroma.KeywordNamespace,
		chroma.KeywordPseudo,
		chroma.KeywordReserved,
		chroma.KeywordType,
		chroma.Operator,
		chroma.OperatorWord,
	),
	NewStyle(Cyan,
		chroma.NameFunction,
		chroma.NameFunctionMagic,
	),
	NewStyle(BrightCyan,
		chroma.NameBuiltin,
		chroma.NameBuiltinPseudo,
	),
	NewStyle(BrightBlack,
		chroma.Comment,
		chroma.CommentHashbang,
		chroma.CommentMultiline,
		chroma.CommentPreproc,
		chroma.CommentPreprocFile,
		chroma.CommentSingle,
		chroma.CommentSpecial,
	),
})

type Syntax struct {
	lexer  chroma.Lexer
	styles StyleArray
}

func NewSyntax(lexer chroma.Lexer) Syntax {
	return NewSyntaxWithStyles(lexer, DefaultStyles)
}

func NewSyntaxWithStyles(lexer chroma.Lexer, styles []Style) Syntax {
	return Syntax{lexer: lexer, styles: styles}
}

func (s Syntax) Highlight(rs []rune) []Cell {
	it, err := s.lexer.Tokenise(nil, string(rs))
	if err != nil {
		log.Fatal(err)
	}
	cells := make([]Cell, 0, len(rs))
	for token := it(); token != nil; token = it() {
		color := s.styles.Color(token.Type)
		for _, r := range []rune(token.Value) {
			cells = append(cells, Cell{Ch: r, Fg: color, Bg: Default})
		}
	}
	return cells
}
