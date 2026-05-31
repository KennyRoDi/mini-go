package main

import (
	"fmt"
)

const (
	T_INT     = 0
	T_CHAR    = 1
	T_BOOL    = 2
	T_STRING  = 3
	T_VOID    = 4
	T_UNKNOWN = -1
)

// Ident represents a symbol in the symbol table.
type Ident struct {
	Nombre string
	Nivel  int
	Tipo   int // Magic number according to SPEC.md
}

// TablaSimbolos implements the mandated linear sequential symbol table.
type TablaSimbolos struct {
	Tabla       []Ident
	NivelActual int
}

func NewTablaSimbolos() *TablaSimbolos {
	return &TablaSimbolos{
		Tabla:       []Ident{},
		NivelActual: 0,
	}
}

func GetMagicType(typeName string) int {
	switch typeName {
	case "int":
		return T_INT
	case "char":
		return T_CHAR
	case "bool":
		return T_BOOL
	case "string":
		return T_STRING
	case "void":
		return T_VOID
	default:
		return T_UNKNOWN
	}
}

// Insert adds a new identifier to the current scope.
// It checks for redeclaration in the SAME scope level.
func (ts *TablaSimbolos) Insert(nombre string, tipo int) error {
	// Mandated O(n) search for redeclaration
	for _, ident := range ts.Tabla {
		if ident.Nombre == nombre && ident.Nivel == ts.NivelActual {
			return fmt.Errorf("el identificador '%s' ya ha sido declarado en este nivel", nombre)
		}
	}
	ts.Tabla = append(ts.Tabla, Ident{Nombre: nombre, Nivel: ts.NivelActual, Tipo: tipo})
	return nil
}

// Lookup performs a mandated O(n) linear sequential search for an identifier.
// It searches from the most recent (deepest) to the oldest.
func (ts *TablaSimbolos) Lookup(nombre string) (Ident, bool) {
	for i := len(ts.Tabla) - 1; i >= 0; i-- {
		if ts.Tabla[i].Nombre == nombre {
			return ts.Tabla[i], true
		}
	}
	return Ident{}, false
}

// OpenScope increases the nesting depth.
func (ts *TablaSimbolos) OpenScope() {
	ts.NivelActual++
}

// CloseScope decreases depth and purges symbols from the current level.
// Mandated rule: tabla.removeIf(ident.nivel == nivelActual)
func (ts *TablaSimbolos) CloseScope() {
	var nuevaTabla []Ident
	for _, ident := range ts.Tabla {
		if ident.Nivel != ts.NivelActual {
			nuevaTabla = append(nuevaTabla, ident)
		}
	}
	ts.Tabla = nuevaTabla
	ts.NivelActual--
}
