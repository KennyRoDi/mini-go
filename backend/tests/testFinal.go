package main;

// === Variables globales ===
var globalInt int = 100;
var globalBool bool = true;

// === Tipos struct ===
type Punto struct {
	x int;
	y int;
};

// === Funciones con parametros y retorno ===
func suma(a int, b int) int {
	return a + b;
};

func esMayor(a int, b int) bool {
	if a > b {
		return true;
	} else {
		return false;
	};
};

func calcularDistancia(p Punto) int {
	return p.x + p.y;
};

// === Test de arreglos ===
func testArrays() {
	var arr [5]int;
	arr[0] = 10;
	arr[1] = 20;
	arr[2] = 30;
	arr[3] = 40;
	arr[4] = 50;
	println(arr[0] + arr[1]);
	println(len(arr));
};

// === Test de estructuras ===
func testStructs() {
	var p Punto;
	p.x = 3;
	p.y = 4;
	println(p.x);
	println(p.y);
	println(calcularDistancia(p));
};

// === Test de control de flujo ===
func testControl() {
	var x int = 10;
	if x > 5 {
		if x > 8 {
			println(1);
		} else {
			println(2);
		};
	} else {
		println(3);
	};

	var i int = 0;
	for i < 3 {
		println(i);
		i = i + 1;
	};

	for j := 0; j < 3; j++ {
		println(j);
	};
};

// === Test de switch ===
func testSwitch(n int) {
	switch n {
	case 1:
		println(1);
	case 2:
		println(2);
	case 3:
		println(3);
	default:
		println(0);
	};
};

// === Test de expresiones ===
func testExprs() {
	var a int = 2 + 3*4 - 1;
	println(a);
	var b bool = a > 10;
	println(b);
	var c bool = (a > 5) && (a < 20);
	println(c);
	var d int = (a * 2 + 10) % 7;
	println(d);
};

// === Test de booleanos ===
func testBool() {
	println(true);
	println(false);
	var flag bool = true;
	if flag {
		println(1);
	};
	var otro bool = esMayor(10, 5);
	println(otro);
};

func main() {
	println(`=== Arreglos ===`);
	testArrays();

	println(`=== Structs ===`);
	testStructs();

	println(`=== Control de flujo ===`);
	testControl();

	println(`=== Switch ===`);
	testSwitch(2);
	testSwitch(3);
	testSwitch(99);

	println(`=== Expresiones ===`);
	testExprs();

	println(`=== Booleanos ===`);
	testBool();

	println(`=== Variables globales ===`);
	println(globalInt);
	println(globalBool);

	println(`=== Funciones ===`);
	println(suma(40, 60));
	println(suma(globalInt, 50));

	println(`=== Fin ===`);
};
