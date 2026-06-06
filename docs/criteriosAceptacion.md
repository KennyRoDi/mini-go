***

> **TL;DR:** A continuación se presenta la lista estricta de criterios de aceptación organizada por componentes del proyecto, basada directamente en los rubros con puntaje máximo de la rúbrica provista.

### 🖥️ 1. Interfaz Gráfica de Usuario (GUI)

- [ ] **Editor de Código:** El entorno incluye un editor de texto funcional que resalta la sintaxis y permite la edición fluida de código fuente escrito en Mini-GO.
- [ ] **Mapeo de Errores:** Cualquier error detectado (léxico, sintáctico o semántico) se despliega indicando de forma exacta la **línea y la columna** dentro del editor.
- [ ] **Manejo de Archivos:** Permite de manera correcta la creación, apertura, modificación y guardado de archivos de código fuente.
- [ ] **Opciones de Compilación:** La interfaz provee botones o comandos accesibles y funcionales para disparar el proceso de análisis y compilación.

---

### 🔍 2. Analizador Léxico (Scanner)

- [ ] **Tokens Simples:** Reconocimiento infalible de palabras reservadas, operadores básicos y delimitadores explícitos en la gramática.
- [ ] **Tokens Complejos:** Identificación exacta mediante expresiones regulares complejas para:
  - [ ] `IDENTIFIER`, `INTLITERAL`, `FLOATLITERAL`.
  - [ ] Literales de texto bajo el estándar estricto de GoLang: `RUNELITERAL`, `RAWSTRINGLITERAL` (`` `...` ``) e `INTERPRETEDSTRINGLITERAL` (`"..."`).
  - [ ] Comentarios de una sola línea (`//`) y multilínea (`/* ... */`).
- [ ] **Recuperación de Errores Léxicos:** El scanner no se detiene ante el primer error; reporta de forma eficiente y detallada todos los caracteres ilegibles detectados, asociándolos a su posición (línea y columna).

---

### 🌳 3. Analizador Sintáctico (Parser)

- [ ] **Declaraciones:** Reconocimiento exacto del árbol de derivación para declaraciones de paquetes, variables (`var`), tipos complejos (`type`) y funciones (`func`).
- [ ] **Instrucciones (Statements):** Validación estricta de estructuras como `print`, `println`, asignaciones normales/compuestas, retornos y bifurcaciones.
- [ ] **Expresiones:** Procesamiento preciso de árboles sintácticos respetando las reglas de precedencia para operaciones aritméticas, lógicas, unarias y relacionales.
- [ ] **Recuperación de Errores Sintácticos:** Despliegue claro de fallas de sintaxis (p. ej., falta de punto y coma, paréntesis desbalanceados) con su ubicación exacta, intentando continuar el análisis en la medida de lo posible.

---

### 🧠 4. Analizador Semántico (Contextual/Scope)

- [ ] **Validaciones del Árbol:** Comprobación rigurosa de las reglas del lenguaje sobre el árbol sintáctico generado.
- [ ] **Manejo de Ámbitos (Scopes):** El compilador valida correctamente la visibilidad de identificadores (variables, funciones, tipos) en ámbitos globales, locales y bloques anidados.
- [ ] **Tabla de Símbolos:** Construcción y mantenimiento exitoso de una tabla de símbolos para el registro de tipos, firmas de funciones y variables.
- [ ] **Errores Semánticos:** Detiene el avance y reporta con línea/columna errores de tipo (ej: asignación de tipos incompatibles), variables no declaradas o funciones duplicadas.

---

### ⚙️ 5. Generación de Código (Backend)

- [ ] **Estructuras de Datos Finitas:** Generación de código ejecutable u objeto completamente funcional para la asignación y manipulación de arreglos (`array`), rebanadas (`slice`) y estructuras (`struct`).
- [ ] **Variables y Parámetros Locales:** El código generado gestiona de forma adecuada la pila de ejecución, el paso de parámetros a funciones y la mutación de variables locales en múltiples escenarios de prueba.
- [ ] **Instrucciones de Control:** Traducción y ejecución sin fallas de bifurcaciones condicionales complejas (`if-else` anidados) y estructuras de selección múltiple (`switch`).
- [ ] **Métodos Predefinidos:** Soporte operativo total en el código final para las funciones internas incorporadas como `print()`, `println()`, `cap()` y `len()`.
- [ ] **Expresiones Generales:** Correcta resolución y cálculo en tiempo de ejecución de expresiones aritmético-lógicas compuestas utilizando tipos de datos simples permitidos.

---

### 📄 6. Documentación del Proyecto

- [ ] **Calidad del Contenido:** Entrega de un documento escrito con un enfoque altamente descriptivo, claro y conciso que abarca la arquitectura del compilador, las decisiones de diseño y los manuales técnicos solicitados.
- [ ] **Calidad Gramatical:** El documento no presenta errores ortográficos, mantiene una redacción técnica impecable y cuenta con una estructura legible de principio a fin.
