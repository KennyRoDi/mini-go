package main;

// 1. DECLARACIONES DE TIPOS
type Punto struct {
    x int;
    y int;
};

type Datos struct {
    id int;
    valor int;
};

// 2. VARIABLES GLOBALES
var globalX int = 100;
var globalY int = 200;

// 3. FUNCIONES CON PARÁMETROS Y RETORNO
func calcular(a int, b int) int {
    var resultado int = 0;
    
    // Operaciones aritméticas con precedencia
    resultado = a * 2 + b / 2 - 5;
    
    // Lógica condicional
    if resultado > 10 {
        resultado = resultado + globalX;
    } else {
        resultado = resultado + 1;
    };
    
    return resultado;
};

func main() {
    // 4. VARIABLES LOCALES Y ESTRUCTURAS
    var p Punto;
    var d Datos;
    var i int = 0;
    var total int = 0;
    var calculo int;

    // 5. ASIGNACIONES (Soportadas correctamente tras el parche)
    p.x = 10;
    p.y = 20;
    d.id = 1;
    d.valor = 50;

    // 6. BUCLE FOR TRADICIONAL (Sin bucle infinito gracias al incremento corregido)
    print("Iniciando bucle: ");
    for i = 0; i < 3; i++ {
        total = total + i;
        print(i);
        print(" ");
    };
    println("");

    // 7. LLAMADA A FUNCIONES
    calculo = calcular(p.x, d.valor);
    
    print("Resultado del calculo: ");
    println(calculo);

    // 8. SENTENCIA SWITCH
    print("Estado del sistema: ");
    switch p.x {
    case 10:
        println("Punto en origen X");
    case 20:
        println("Punto desplazado");
    default:
        println("Desconocido");
    };

    println("Prueba V1 finalizada exitosamente.");
};
