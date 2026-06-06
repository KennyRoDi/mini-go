# Mini-GO Compiler & Desktop IDE

---

## Overview

**Mini-GO** is designed to provide a high-fidelity development experience. The system is divided into two primary components:
1.  **Backend (Go):** A high-performance compiler built with **ANTLR4** for lexical/syntactic analysis and **llir/llvm** for code generation.
2.  **Frontend (Electron + TypeScript):** A professional, dark-themed IDE with a focused workspace and integrated terminal.

## Technical Architecture

### 1. Compiler Backend
*   **Parser/Lexer:** Generated via ANTLR4 (Go target) using `minigo.g4`.
*   **Symbol Table:** Implemented as a linear sequential list (`O(n)` lookups) with scope-based purging, as per architectural requirements.
*   **Type System:** Uses hardcoded magic numbers (`0=int`, `1=char`, `2=bool`, `3=string`, `4=void`).
*   **Semantic Passes:** 
    *   **Scope Pass:** Handles identifier declaration and visibility.
    *   **Type Pass:** Performs strict type validation with branch isolation (panic/recover).
*   **Code Gen:** Produces **LLVM IR**, which is then linked via **Clang** to create native executables.

### 2. Desktop IDE
*   **Stack:** Electron, React, TypeScript, HeroUI, Tailwind CSS.
*   **Icons:** Lucide React icons for a professional look.
*   **Editor:** Monaco Editor (el motor de VS Code) via `@monaco-editor/react`.
*   **IPC:** Spawns the Go compiler as a child process and parses real-time diagnostics.

## Onboarding & Requirements

To run this project, you need the following tools installed:

1.  **Go (Golang):** Version 1.20+
2.  **Node.js & pnpm:** For the modern React frontend.
3.  **Clang:** Required by the backend to link LLVM IR into executables.
4.  **Java (JRE):** Only if you need to regenerate the ANTLR4 parser from the `.g4` file.

## Execution Steps

### 1. Build the Compiler Backend
Navigate to the backend directory and compile the `minigo` binary:
```bash
cd mini-go/backend
go mod tidy
go build -o minigo .
```

### 2. Build and Launch the IDE
Navigate to the ide directory, install dependencies, and start the application:
```bash
cd ../ide
pnpm install
pnpm start
```

## Mini-GO Language Features
The current version supports:
*   **Variable Declarations:** `var x int = 10;`
*   **Arithmetic:** `+`, `-`, `*`, `/` (Integer only for now).
*   **Control Flow:** `if` statements with boolean conditions.
*   **Standard I/O:** `print()` and `println()` (mapped to native `printf`).
*   **Main Entry:** Automatic synthesis of `func main()` for global statements.

## Future Improvements
*   **Full Struct Support:** Implementation of struct methods and attribute access.
*   **Array & Slices:** Expansion of the type checker to handle complex collections.
*   **Interactive Terminal:** Integration of `xterm.js` for a true console experience.
*   **Project Workspace:** Support for multi-file projects and folders.
