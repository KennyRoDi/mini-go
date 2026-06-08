package main

import (
	"fmt"
)

// MiniGo Type System
type Type interface {
	String() string
	Equals(other Type) bool
}

type BasicType int

const (
	T_INT BasicType = iota
	T_CHAR
	T_BOOL
	T_STRING
	T_VOID
	T_UNKNOWN
)

func (t BasicType) String() string {
	switch t {
	case T_INT: return "int"
	case T_CHAR: return "char"
	case T_BOOL: return "bool"
	case T_STRING: return "string"
	case T_VOID: return "void"
	default: return "unknown"
	}
}

func (t BasicType) Equals(other Type) bool {
	if b, ok := other.(BasicType); ok {
		return t == b
	}
	return false
}

type ArrayType struct {
	Elem Type
	Size int
}

func (t ArrayType) String() string {
	return fmt.Sprintf("[%d]%s", t.Size, t.Elem.String())
}

func (t ArrayType) Equals(other Type) bool {
	if a, ok := other.(ArrayType); ok {
		return t.Size == a.Size && t.Elem.Equals(a.Elem)
	}
	return false
}

type SliceType struct {
	Elem Type
}

func (t SliceType) String() string {
	return fmt.Sprintf("[]%s", t.Elem.String())
}

func (t SliceType) Equals(other Type) bool {
	if s, ok := other.(SliceType); ok {
		return t.Elem.Equals(s.Elem)
	}
	return false
}

type StructType struct {
	Fields map[string]Type
	OrderedFields []string
}

func (t StructType) String() string {
	return "struct{...}"
}

func (t StructType) Equals(other Type) bool {
	if s, ok := other.(StructType); ok {
		if len(t.Fields) != len(s.Fields) { return false }
		for k, v := range t.Fields {
			if v2, ok := s.Fields[k]; !ok || !v.Equals(v2) { return false }
		}
		return true
	}
	return false
}

type FuncType struct {
	Params []Type
	Return Type
}

func (t FuncType) String() string {
	return "func(...)"
}

func (t FuncType) Equals(other Type) bool {
	if f, ok := other.(FuncType); ok {
		if len(t.Params) != len(f.Params) { return false }
		for i, p := range t.Params {
			if !p.Equals(f.Params[i]) { return false }
		}
		return t.Return.Equals(f.Return)
	}
	return false
}

// Ident represents a symbol in the symbol table.
type Ident struct {
	Nombre string
	Nivel  int
	Tipo   Type
}

// TablaSimbolos implements the mandated linear sequential symbol table.
type TablaSimbolos struct {
	Tabla       []*Ident
	NivelActual int
}

func NewTablaSimbolos() *TablaSimbolos {
	return &TablaSimbolos{
		Tabla:       []*Ident{},
		NivelActual: 0,
	}
}

func GetBasicType(typeName string) Type {
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
func (ts *TablaSimbolos) Insert(nombre string, tipo Type) error {
	for _, ident := range ts.Tabla {
		if ident.Nombre == nombre && ident.Nivel == ts.NivelActual {
			return fmt.Errorf("el identificador '%s' ya ha sido declarado en este nivel", nombre)
		}
	}
	ts.Tabla = append(ts.Tabla, &Ident{Nombre: nombre, Nivel: ts.NivelActual, Tipo: tipo})
	return nil
}

func (ts *TablaSimbolos) Lookup(nombre string) (*Ident, bool) {
	for i := len(ts.Tabla) - 1; i >= 0; i-- {
		if ts.Tabla[i].Nombre == nombre {
			return ts.Tabla[i], true
		}
	}
	return nil, false
}

// Update changes the type of an existing identifier in the closest scope.
func (ts *TablaSimbolos) Update(nombre string, nuevoTipo Type) error {
	ident, found := ts.Lookup(nombre)
	if found {
		ident.Tipo = nuevoTipo
		return nil
	}
	return fmt.Errorf("identificador '%s' no encontrado para actualizar", nombre)
}

func (ts *TablaSimbolos) OpenScope() {
	ts.NivelActual++
}

func (ts *TablaSimbolos) CloseScope() {
	var nuevaTabla []*Ident
	for _, ident := range ts.Tabla {
		if ident.Nivel != ts.NivelActual {
			nuevaTabla = append(nuevaTabla, ident)
		}
	}
	ts.Tabla = nuevaTabla
	ts.NivelActual--
}
