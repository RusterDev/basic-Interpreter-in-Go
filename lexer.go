package main

import (
  "strings"
  "unicode"
  "fmt"
)

const (
  INT = iota
  STRING
  FLOAT
  BOOL
  BYTE
  INT8
  INT16
  INT32
  INT64
  UINT
  UINT8
  UINT16
  UINT32
  UINT64
  PLUS
  MINUS
  MUL
  DIV
  MOD
  EQ
  NEQ
  LPAREN
  RPAREN
  IDENTIFIER
  ASSIGN
  NUMERIC_LITERAL_INT
  NUMERIC_LITERAL_FLOAT
  FUNCPRINT
)

type Token struct {
  Type   uint8
  Value  uint8
  Lexeme string
  Line   uint8
  Column uint8
}



func lex(input string) ([]Token, uint8) {
  var tokens []Token
  var lines, col uint8 = 1, 1
  unformattedLines := strings.Split(input, "\n")
 
  for _, line := range unformattedLines {
    unformattedTokens := strings.Fields(line)

    for _, word := range unformattedTokens {
      switch word {
      case "int":
        var tok Token
        tok.Type = INT
        tok.Lexeme = word
        tok.Line = lines
        tok.Column = col
        tokens = append(tokens, tok)
      case "string":
        var tok Token
        tok.Type = STRING
        tok.Lexeme = word
        tok.Line = lines
        tok.Column = col
        tokens = append(tokens, tok)
      case "float":
        var tok Token
        tok.Type = FLOAT
        tok.Lexeme = word
        tok.Line = lines
        tok.Column = col
        tokens = append(tokens, tok)
      case "bool":
        var tok Token
        tok.Type = BOOL
        tok.Lexeme = word
        tok.Line = lines
        tok.Column = col
        tokens = append(tokens, tok)
      case "print":
        var tok Token
        tok.Type = FUNCPRINT
        tok.Lexeme = word
        tok.Line = lines
        tok.Column = col
        tokens = append(tokens, tok)
    
      default:
        if len(word) == 1 {
          switch word {
          case "+":
            var tok Token
            tok.Type = PLUS
            tok.Lexeme = word
            tok.Line = lines
            tok.Column = col
            tokens = append(tokens, tok)
          case "-":
            var tok Token
            tok.Type = MINUS
            tok.Lexeme = word
            tok.Line = lines
            tok.Column = col
            tokens = append(tokens, tok)
          case "*":
            var tok Token
            tok.Type = MUL
            tok.Lexeme = word
            tok.Line = lines
            tok.Column = col
            tokens = append(tokens, tok)
          case "/":
            var tok Token
            tok.Type = DIV
            tok.Lexeme = word
            tok.Line = lines
            tok.Column = col
            tokens = append(tokens, tok)
          case "%":
            var tok Token
            tok.Type = MOD
            tok.Lexeme = word
            tok.Line = lines
            tok.Column = col
            tokens = append(tokens, tok)
          case "=":
            var tok Token
            tok.Type = ASSIGN
            tok.Lexeme = word
            tok.Line = lines
            tok.Column = col
            tokens = append(tokens, tok)
          case "(":
            var tok Token
            tok.Type = LPAREN
            tok.Lexeme = word
            tok.Line = lines
            tok.Column = col
            tokens = append(tokens, tok)
          case ")":
            var tok Token
            tok.Type = RPAREN
            tok.Lexeme = word
            tok.Line = lines
            tok.Column = col
            tokens = append(tokens, tok)
          default:
            var tok Token
            tok.Type = IDENTIFIER
            tok.Lexeme = word
            tok.Line = lines
            tok.Column = col
            tokens = append(tokens, tok)
          if unicode.IsDigit(rune(word[0])) {
            var tok Token
            tok.Type = NUMERIC_LITERAL_INT
            tok.Lexeme = word
            tok.Line = lines
            tok.Column = col
            fmt.Println("numero")
            tokens = append(tokens, tok)
          } 
          
          }
        }
      }


      col += uint8(len(word)) + 1 
    }


    col = 1


    lines++
  }

  return tokens, lines - 1
}
