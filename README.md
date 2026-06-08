# Mini-GO Compiler & Desktop IDE

---

## English

### Overview

Mini-GO is a complete compiler and IDE system for a Go-like language. The project consists of two main components:

1. **Backend (Go):** A high-performance compiler built with ANTLR4 for lexical and syntactic analysis, and LLVM IR for code generation.
2. **Frontend (Electron + React + TypeScript):** A professional desktop IDE with Monaco Editor integration and real-time diagnostics.

### Architecture

#### Compiler Backend
- **Lexer/Parser:** Generated via ANTLR4 from the grammar file `minigo.g4`
- **Symbol Table:** Scope-aware identifier tracking with hierarchical level management
- **Type System:** Strict type checking with support for int, bool, string, char, and void types
- **Semantic Analysis:** 
  - Detects variable redeclarations
  - Reports undeclared variable usage
  - Validates type compatibility in assignments and operations
  - Enforces boolean conditions in control flow statements
- **Code Generation:** Produces LLVM IR which is linked into native executables

#### Desktop IDE
- **Stack:** Electron, React, TypeScript, Tailwind CSS, Monaco Editor
- **Architecture:** IPC communication with the backend compiler for real-time error diagnostics
- **Features:** Syntax highlighting, error reporting with line/column information, integrated terminal

### System Requirements

The following tools must be installed on your system:

1. **Go (Golang):** Version 1.20 or later
2. **Node.js & pnpm:** For building and running the frontend (Node.js 16+, pnpm 8+)
3. **Clang/LLVM:** Required to link LLVM IR into executables
4. **Java Runtime (JRE):** Only needed if regenerating the ANTLR4 parser from the grammar file

**Windows additional requirement:** LLVM must be in your PATH or explicitly configured for clang linking.

### Build Instructions

#### Step 1: Prepare the Compiler Backend

**Linux/macOS:**
```bash
cd backend
go mod tidy
go build -o minigo .
cd ..
```

**Windows (PowerShell):**
```powershell
cd backend
go mod tidy
go build -o minigo.exe .
cd ..
```

The compiled binary will be named `minigo` (Linux/macOS) or `minigo.exe` (Windows).

#### Step 2: Prepare the IDE Frontend

**Linux/macOS:**
```bash
cd ide
pnpm install
cd ..
```

**Windows (PowerShell):**
```powershell
cd ide
pnpm install
cd ..
```

### Running the Compiler

#### Command-Line Usage

Compile a Mini-GO source file:

**Linux/macOS:**
```bash
./backend/minigo path/to/program.go
```

**Windows (PowerShell):**
```powershell
.\backend\minigo.exe path\to\program.go
```

On successful compilation, an executable named `output` (Linux/macOS) or `output.exe` (Windows) is generated in the current directory.

#### Running the Generated Executable

**Linux/macOS:**
```bash
./output
```

**Windows (PowerShell):**
```powershell
.\output.exe
```

### Running the IDE

Launch the Electron application:

**Linux/macOS:**
```bash
cd ide
pnpm start
```

**Windows (PowerShell):**
```powershell
cd ide
pnpm start
```

The IDE will open with a ready-to-use editor and terminal. Click the "Run" button or press Ctrl+R to compile and execute code directly within the IDE.

### Supported Language Features

- **Variables:** `var x int = 10;`
- **Types:** int, bool, string, char, void
- **Structures:** `type Point struct { x int; y int; }`
- **Arrays:** Fixed-size arrays `[5]int` and slices `[]int`
- **Functions:** Functions with parameters and return types
- **Control Flow:** if/else statements, for loops, switch statements
- **Operators:** Arithmetic, logical, bitwise, comparison, and assignment operators
- **Built-in Functions:** print, println, len, cap, append
- **Error Detection:** Syntax errors, type mismatches, undeclared variables, redeclarations

### Recent Changes

This version includes critical fixes for error detection:

1. **Syntax Error Handling:** The compiler now properly terminates (exit code 1) when syntax errors are encountered, preventing downstream semantic analysis on corrupted ASTs.
2. **Variable Redeclaration Detection:** The compiler now correctly identifies and reports when a variable is declared twice in the same scope.
3. **Undeclared Variable Detection:** Using an identifier that was never declared now produces a clear diagnostic message.

These improvements ensure that the IDE's error display and the command-line compiler behavior are fully synchronized.

---

## Español

### Descripción General

Mini-GO es un sistema completo de compilador e IDE para un lenguaje similar a Go. El proyecto consta de dos componentes principales:

1. **Backend (Go):** Un compilador de alto rendimiento construido con ANTLR4 para análisis léxico y sintáctico, e LLVM IR para generación de código.
2. **Frontend (Electron + React + TypeScript):** Un IDE de escritorio profesional con integración de Monaco Editor y diagnósticos en tiempo real.

### Arquitectura

#### Backend del Compilador
- **Lexer/Parser:** Generado mediante ANTLR4 a partir del archivo de gramática `minigo.g4`
- **Tabla de Símbolos:** Seguimiento de identificadores consciente del alcance con gestión jerárquica de niveles
- **Sistema de Tipos:** Verificación de tipos estricta con soporte para tipos int, bool, string, char y void
- **Análisis Semántico:**
  - Detecta redeclaraciones de variables
  - Reporta uso de variables no declaradas
  - Valida compatibilidad de tipos en asignaciones y operaciones
  - Impone condiciones booleanas en sentencias de control de flujo
- **Generación de Código:** Produce LLVM IR que se vincula en ejecutables nativos

#### IDE de Escritorio
- **Stack:** Electron, React, TypeScript, Tailwind CSS, Monaco Editor
- **Arquitectura:** Comunicación IPC con el compilador backend para diagnósticos de error en tiempo real
- **Características:** Resaltado de sintaxis, reporte de errores con información de línea/columna, terminal integrada

### Requisitos del Sistema

Las siguientes herramientas deben estar instaladas en su sistema:

1. **Go (Golang):** Versión 1.20 o posterior
2. **Node.js & pnpm:** Para construir y ejecutar el frontend (Node.js 16+, pnpm 8+)
3. **Clang/LLVM:** Requerido para vincular LLVM IR en ejecutables
4. **Java Runtime (JRE):** Solo necesario si regenera el analizador ANTLR4 desde el archivo de gramática

**Requisito adicional en Windows:** LLVM debe estar en su PATH o configurado explícitamente para la vinculación con clang.

### Instrucciones de Compilación

#### Paso 1: Preparar el Backend del Compilador

**Linux/macOS:**
```bash
cd backend
go mod tidy
go build -o minigo .
cd ..
```

**Windows (PowerShell):**
```powershell
cd backend
go mod tidy
go build -o minigo.exe .
cd ..
```

El binario compilado se llamará `minigo` (Linux/macOS) o `minigo.exe` (Windows).

#### Paso 2: Preparar el Frontend del IDE

**Linux/macOS:**
```bash
cd ide
pnpm install
cd ..
```

**Windows (PowerShell):**
```powershell
cd ide
pnpm install
cd ..
```

### Ejecutar el Compilador

#### Uso en Línea de Comandos

Compilar un archivo fuente Mini-GO:

**Linux/macOS:**
```bash
./backend/minigo ruta/a/programa.go
```

**Windows (PowerShell):**
```powershell
.\backend\minigo.exe ruta\a\programa.go
```

En caso de compilación exitosa, se genera un ejecutable llamado `output` (Linux/macOS) u `output.exe` (Windows) en el directorio actual.

#### Ejecutar el Ejecutable Generado

**Linux/macOS:**
```bash
./output
```

**Windows (PowerShell):**
```powershell
.\output.exe
```

### Ejecutar el IDE

Inicie la aplicación Electron:

**Linux/macOS:**
```bash
cd ide
pnpm start
```

**Windows (PowerShell):**
```powershell
cd ide
pnpm start
```

El IDE se abrirá con un editor y terminal listos para usar. Haga clic en el botón "Run" o presione Ctrl+R para compilar y ejecutar código directamente desde el IDE.

### Características del Lenguaje Soportadas

- **Variables:** `var x int = 10;`
- **Tipos:** int, bool, string, char, void
- **Estructuras:** `type Point struct { x int; y int; }`
- **Arreglos:** Arreglos de tamaño fijo `[5]int` y slices `[]int`
- **Funciones:** Funciones con parámetros y tipos de retorno
- **Control de Flujo:** Sentencias if/else, bucles for, sentencias switch
- **Operadores:** Operadores aritméticos, lógicos, bitwise, comparación y asignación
- **Funciones Incorporadas:** print, println, len, cap, append
- **Detección de Errores:** Errores de sintaxis, desajustes de tipo, variables no declaradas, redeclaraciones

### Cambios Recientes

Esta versión incluye correcciones críticas para la detección de errores:

1. **Manejo de Errores de Sintaxis:** El compilador ahora termina correctamente (código de salida 1) cuando se encuentran errores de sintaxis, previniendo análisis semántico aguas abajo en ASTs corruptos.
2. **Detección de Redeclaración de Variables:** El compilador ahora identifica y reporta correctamente cuando una variable se declara dos veces en el mismo alcance.
3. **Detección de Variables No Declaradas:** Usar un identificador que nunca fue declarado ahora produce un mensaje de diagnóstico claro.

Estas mejoras aseguran que la visualización de errores del IDE y el comportamiento del compilador en línea de comandos estén completamente sincronizados.

---

## Project Structure

```
mini-go/
├── backend/               # Go compiler source
│   ├── main.go           # Entry point
│   ├── type_checker.go   # Semantic analysis
│   ├── symbol_table.go   # Symbol table management
│   ├── encoder.go        # LLVM IR generation
│   ├── minigo.g4         # ANTLR4 grammar
│   ├── parser/           # Generated ANTLR4 files
│   ├── go.mod            # Go module definition
│   └── tests/            # Test programs
├── ide/                  # Electron IDE source
│   ├── src/              # TypeScript/React source
│   ├── public/           # Static assets
│   ├── package.json      # Node.js dependencies
│   └── tsconfig.json     # TypeScript configuration
└── README.md             # This file
```

## License

This project is part of an academic curriculum at Tecnológico de Costa Rica (TEC).

