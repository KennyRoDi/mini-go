package main

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
)

// CustomErrorListener implements antlr.ErrorListener to capture syntax errors.
type CustomErrorListener struct {
	*antlr.DefaultErrorListener
	Errors []string
}

func NewCustomErrorListener() *CustomErrorListener {
	return &CustomErrorListener{
		DefaultErrorListener: &antlr.DefaultErrorListener{},
		Errors:               []string{},
	}
}

// SyntaxError captures syntax errors and formats them for the UI.
func (c *CustomErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	// The specification requires specific error formats for semantic errors, 
	// but syntax errors should also be clearly reported with line and column.
	errorMsg := fmt.Sprintf("SYNTAX ERROR: %s [linea:%d - columna:%d]", msg, line, column)
	c.Errors = append(c.Errors, errorMsg)
}

func (c *CustomErrorListener) HasErrors() bool {
	return len(c.Errors) > 0
}

func (c *CustomErrorListener) PrintErrors() {
	for _, err := range c.Errors {
		fmt.Println(err)
	}
}
