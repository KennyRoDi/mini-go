package test;

var x int = 10;
var x int = 20; // ERROR: redeclaration

func main() {
    var y int = 5;
    print(x, y);
    print(z); // ERROR: undeclared
}

func other() {
    print(y); // ERROR: undeclared (y is local to main)
}
;
