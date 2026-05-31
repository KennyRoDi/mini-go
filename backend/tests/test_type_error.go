package test;

var x int = 10;
var s string = "hello";

func main() {
    var b bool = (x == 10);
    var y int = s; // ERROR: string to int
    
    if (x + s) { // ERROR: int + string, and condition not bool
        print(x);
    };
};
