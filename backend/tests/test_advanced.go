package main;

// ---------------------------------------------------------
// 1. DECLARACIONES DE TIPOS Y VARIABLES GLOBALES
// ---------------------------------------------------------
type Vector [3] int;

type Estudiante struct {
    id int;
    nota int;
};

var globalX, globalY int = 10, 20;
var miVector Vector;

// ---------------------------------------------------------
// 2. FUNCIONES DE PRUEBA SINTÁCTICA
// ---------------------------------------------------------
func calcularMatematica(a int, b int) int {
    var resultado int = 0;
    
    // Operaciones aritméticas y lógicas complejas de precedencia
    resultado = a * b + globalX / 2 - (b % 3);
    
    if resultado > 100 {
        resultado = resultado >> 1;
    } else {
        resultado = resultado << 2;
    };
    
    return resultado;
};

func procesarEstructuras() int {
    var alumno Estudiante;
    var total int = 0;
    
    alumno.id = 4026;
    alumno.nota = 95;
    
    // Simulación de Switch y comparaciones binarias
    switch alumno.nota {
    case 100:
        total = 10;
    case 95:
        total = 20;
    default:
        total = 0;
    };
    
    return alumno.id + total;
};

// ---------------------------------------------------------
// 3. PUNTO DE ENTRADA PRINCIPAL (MAIN)
// ---------------------------------------------------------
func main() {
    var iterador int = 0;
    var sumaBucles int = 0;
    var operacionCompleja int;
    var structResult int;

    // Inicialización del arreglo global
    miVector[0] = 5;
    miVector[1] = 10;
    miVector[2] = 15;

    // Ejecución de funciones secundarias
    operacionCompleja = calcularMatematica(miVector[1], miVector[2]);
    structResult = procesarEstructuras();

    // Ciclo For complejo tradicional de tres componentes
    for iterador = 0; iterador < 5; iterador++ {
        sumaBucles = sumaBucles + iterador * 2;
    };

    // Impresión estricta usando las palabras reservadas print/println y ';' final
    print("Resultado Operacion Matematica: ");
    println(operacionCompleja);
    
    print("Resultado Datos Estudiante: ");
    println(structResult);
    
    print("Resultado Suma Acumulada del Bucle: ");
    println(sumaBucles);
};
