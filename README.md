# Mini-GO Compiler & Desktop IDE

---

## Descripción General

Mini-GO es un sistema completo de compilador e IDE para un lenguaje subconjunto de Go. El proyecto consta de dos componentes principales:

1. **Backend (Go):** Un compilador construido con ANTLR4 para análisis léxico y sintáctico, e LLVM IR para generación de código.
2. **Frontend (Electron + React + TypeScript):** Un entorno integrado de desarrollo con Monaco Editor y diagnósticos en tiempo real.

### Arquitectura

#### Backend del Compilador
- **Lexer/Parser:** Generado mediante ANTLR4 a partir del archivo de gramática `minigo.g4`
- **Tabla de Símbolos:** Seguimiento de identificadores con gestión jerárquica de alcances
- **Sistema de Tipos:** Verificación de tipos estricta con soporte para int, bool, string, char y void
- **Análisis Semántico:**
  - Detección de redeclaraciones de variables
  - Validación de variables declaradas
  - Compatibilidad de tipos en asignaciones y operaciones
  - Verificación de condiciones booleanas en control de flujo
- **Generación de Código:** Producción de LLVM IR vinculado en ejecutables nativos

#### IDE de Escritorio
- **Stack:** Electron, React, TypeScript, Tailwind CSS, Monaco Editor
- **Arquitectura:** Comunicación IPC con el compilador para diagnósticos en tiempo real
- **Características:** Resaltado de sintaxis, reportes de error con posición, terminal integrada

---

## Requisitos del Sistema

Las siguientes herramientas deben estar instaladas:

1. **Go (Golang):** Versión 1.20 o posterior
2. **Node.js & pnpm:** Node.js 16+, pnpm 8+
3. **Clang/LLVM:** Requerido para vinculación de LLVM IR
4. **Java Runtime (JRE):** Opcional, solo si se regenera el parser ANTLR4

**Nota para Windows:** LLVM debe estar en PATH o configurado explícitamente para clang.

---

## Instrucciones de Compilación

### Paso 1: Backend del Compilador

```bash
cd backend
go mod tidy
go build -o minigo .
cd ..
```

**En Windows (PowerShell):**
```powershell
go build -o minigo.exe .
```

### Paso 2: Frontend del IDE

```bash
cd ide
pnpm install
cd ..
```

---

## Uso del Compilador

### Línea de Comandos

Compilar un archivo Mini-GO:

```bash
./backend/minigo ruta/a/programa.go
```

**En Windows (PowerShell):**
```powershell
.\backend\minigo.exe ruta\a\programa.go
```

Genera un ejecutable llamado `output` (o `output.exe` en Windows) en el directorio actual.

### Ejecutar el Programa Compilado

```bash
./output
```

**En Windows:**
```powershell
.\output.exe
```

---

## Ejecución del IDE

```bash
cd ide
pnpm start
```

El IDE se abrirá con editor y terminal listos. Use el botón "Run" o Ctrl+R para compilar y ejecutar código.

---

## Características del Lenguaje Soportadas

- **Variables:** `var x int = 10;`
- **Tipos:** int, bool, string, char, void
- **Estructuras:** `type Point struct { x int; y int; }`
- **Arreglos:** Arreglos fijos `[5]int` y slices `[]int`
- **Funciones:** Con parámetros y tipos de retorno
- **Control de Flujo:** if/else, for loops, switch
- **Operadores:** Aritméticos, lógicos, bitwise, comparación y asignación
- **Funciones Incorporadas:** print, println, len, cap, append
- **Detección de Errores:** Errores sintácticos, desajustes de tipo, variables no declaradas, redeclaraciones

---

## Estructura del Proyecto

```
mini-go/
├── backend/
│   ├── main.go
│   ├── type_checker.go
│   ├── symbol_table.go
│   ├── encoder.go
│   ├── minigo.g4
│   ├── parser/
│   ├── go.mod
│   └── tests/
├── ide/
│   ├── src/
│   ├── public/
│   ├── package.json
│   └── tsconfig.json
└── README.md
```

---