package main;

// === Test de reportes de errores semanticos ===
// Este archivo contiene errores intencionales para verificar
// que el compilador los detecta con linea y columna correctas.

// Error 1 (linea 13): asignar string a variable int
var x int = "hola";

// Error 2 (linea 16): tipo invalido en operacion aritmetica (bool + bool)
func testAritmetica() {
	var b bool = true;
	var c int = b + b;
	println(c);
};

// Error 3 (linea 22): condicion de if no es bool (se pasa int)
func testCondicion() {
	var n int = 5;
	if n {
		println(1);
	};
};

// Error 4 (linea 29): && con operandos no booleanos
func testLogica() {
	var a int = 1;
	var b int = 2;
	var c bool = a && b;
	println(c);
};

// Error 5 (linea 36): acceder a campo de tipo no-struct
func testCampo() {
	var n int = 10;
	println(n.x);
};

// Error 6 (linea 42): indexar tipo no-arreglo
func testIndex() {
	var n int = 5;
	println(n[0]);
};

func main() {
	testAritmetica();
	testCondicion();
	testLogica();
	testCampo();
	testIndex();
};
