package lexer

import (
	"testing"
  "fmt"
	"github.com/VosLannuck/x_interpeter_golang/token"
)

func TestNextToken(t *testing.T) {
  // var input string = `=+(){},;`
  var input string = `let five = 5;
    let ten = 10;

    let add = fn(x, y) {
      x + y;
    };

    let result = add(five, ten);
    ` 

  type Test struct {
    expectedType token.TokenType
    expectedLiteral string
  }
  /* tests := []Test {
    {token.ASSIGN, "="},
    {token.PLUS, "+"},
    {token.LPAREN,"("},
    {token.RPAREN, ")"},
    {token.LBRACE, "{"},
    {token.RBRACE, "}"},
    {token.COMMA, ","},
    {token.SEMICOLON, ";"},
    {token.EOF, ""},
  } */
      tests := []Test {
      {token.LET, "let"},
      {token.IDENT, "five"},
      {token.ASSIGN, "="},
      {token.INT, "5"},
      {token.SEMICOLON, ";"},
      
      {token.LET, "let"},
      {token.IDENT, "ten"},
      {token.ASSIGN, "="},
      {token.INT, "10"},
      {token.SEMICOLON, ";"},

      {token.LET, "let"},
      {token.IDENT, "add"},
      {token.ASSIGN, "="},
      {token.FUNCTION, "fn"},
      {token.LPAREN, "("},
      {token.IDENT, "x"},
      {token.COMMA, ","},
      {token.IDENT, "y"},
      {token.RPAREN, ")"},
      {token.LPAREN, "{"},
      {token.IDENT, "x"},
      {token.PLUS, "+"},
      {token.IDENT, "y"},
      {token.SEMICOLON, ";"},
      {token.RBRACE, "}"},
      {token.SEMICOLON, ";"},

      {token.LET, "let"},
      {token.IDENT, "result"},
      {token.ASSIGN, "="},
      {token.IDENT, "add"},
      {token.LPAREN, "("},
      {token.IDENT, "five"},
      {token.COMMA, ","},
      {token.IDENT, "ten"},
      {token.RPAREN, ")"},
      {token.SEMICOLON, ";"},
      {token.EOF, ""},
  }

  inp := New(input)
  fmt.Println(inp.ch);
  for indx, t_exp := range tests {
    tkn := inp.NextToken()
    if tkn.Type != t_exp.expectedType {
      // Raise ILLEGAL
      t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
        indx, t_exp.expectedType, tkn.Type)
    }

    if tkn.Literal != t_exp.expectedLiteral {
      // Raise Illegal
      t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
        indx, t_exp.expectedLiteral, tkn.Literal)
    }

  }



}
