package main;

type Estudiante struct {
    id int;
    nota int;
};

func main() {
    var alumno Estudiante;
    // Esto debería causar error semántico y NO pánico en el encoder
    var total int = alumno.nota + 10;
    println(total);
};
