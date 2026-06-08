package main

import (
	"fmt"
	"os"

	"github.com/antlr4-go/antlr/v4"
	"minigo-backend/parser"
)

func runEncoder(encoder *MiniGoEncoder, tree antlr.ParseTree) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
		}
	}()
	tree.Accept(encoder)
	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: minigo <archivo.go>")
		os.Exit(1)
	}

	fileName := os.Args[1]
	input, err := antlr.NewFileStream(fileName)
	if err != nil {
		fmt.Printf("Error al leer el archivo: %v\n", err)
		os.Exit(1)
	}

	// Lexer
	lexer := parser.NewminigoLexer(input)
	lexerErrors := NewCustomErrorListener()
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(lexerErrors)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Parser
	p := parser.NewminigoParser(stream)
	parserErrors := NewCustomErrorListener()
	p.RemoveErrorListeners()
	p.AddErrorListener(parserErrors)

	// Iniciar el parsing desde la regla "root"
	tree := p.Root()

	// Reportar errores de sintaxis
	if lexerErrors.HasErrors() || parserErrors.HasErrors() {
		lexerErrors.PrintErrors()
		parserErrors.PrintErrors()
		os.Exit(1)
	}

	// Fase de Tipos (Semantic Analysis)
	symbols := NewTablaSimbolos()
	typeVisitor := NewTypeVisitor(symbols)
	tree.Accept(typeVisitor)

	if len(typeVisitor.Errors) > 0 {
		for _, err := range typeVisitor.Errors {
			fmt.Println(err)
		}
		os.Exit(1)
	}

	// Fase de Generacion de Codigo (LLVM IR)
	encoder := NewMiniGoEncoder(symbols)
	encoder.NodeTypes = typeVisitor.NodeTypes
	encoder.SymbolMap = typeVisitor.SymbolMap
	if err := runEncoder(encoder, tree); err != nil {
		fmt.Printf("Error interno en generacion de codigo: %v\n", err)
		os.Exit(1)
	}

	outPath := "output"
	err = encoder.Emit(outPath)
	if err != nil {
		fmt.Printf("Error durante la generacion de codigo: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Compilacion exitosa. Ejecutable generado en: %s\n", outPath)
}
