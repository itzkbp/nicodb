# NicoDB

## Prerequisites

Golang installed, prefeered version 1.19

## Running parser

1. Create main.go file 
```go
package main

import (
	sqlparser "github.com/itzkbp/nicodb/internal/sql"
)

func main() {

	query := `#YOUR SQL COMMAND;`

	lexer := sqlparser.NewLexer(query)
	parser := sqlparser.NewParser(lexer)
	parser.Parse().Execute()
}

```

2. Run the file
```bash
go run main.go
```