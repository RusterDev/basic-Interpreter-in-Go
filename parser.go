  package main

  import (
    "fmt"
  )

  type Node struct {
    Type     string
    Value    string
    Children []Token
  }

  type AST struct {
    Nodes []Node
  }

func get_next_token(token Token, tokens []Token) Token {
  if token.Column+1 < uint8(len(tokens)) {
    return tokens[token.Column]
  } else {
    return tokens[token.Column+1] 
  }
}

func parse(tokens *[]Token) AST {
  var ast AST
  for _, token := range *tokens {

    switch token.Type {
    case IDENTIFIER:

      tok := get_next_token(token, *tokens)
      *tokens = append(*tokens, tok)

      if tok.Type == ASSIGN {
        node := Node{
          Type:     "assignment",
          Children: []Token{token, tok},
        }

        next := get_next_token(tok, *tokens)
        node.Value = next.Lexeme
        node.Children = append(node.Children, next)

        ast.Nodes = append(ast.Nodes, node)
      }
    case FUNCPRINT:
      lparen := get_next_token(token, *tokens)
      rparen := get_next_token(lparen, *tokens)

      var contentTokens []Token
      for i := lparen.Column + 1; i < rparen.Column; i++ {
        contentTokens = append(contentTokens, (*tokens)[i])
      }

      contentNode := Node{
        Type:     "content",
        Children: contentTokens,
      }

      fmt.Println("content ", contentNode)
      ast.Nodes = append(ast.Nodes, contentNode)
    }
  }

  return ast
}
