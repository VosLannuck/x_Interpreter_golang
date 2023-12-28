package token;

type TokenType string

type Token struct{
  Type TokenType
  Literal string
}

// Define the tokens
const (
  ILLEGAL = "ILLEGAL"
  EOF = "EOF"

  // Identifiers + literals
  IDENT = "IDENT" // x , y , zoo, foo
  INT = "INT" // 1, 2, 3

  //Opertors
  PLUS = "+"
  ASSIGN = "="

  // Delimeters
  COMMA = ","
  SEMICOLON = ";"

  LPAREN = "("
  RPAREN = ")"
  LBRACE = "{"
  RBRACE = "}"

  // KEYWORDS
  FUNCTION = "FUNCTION"
  LET = "LET"

)
