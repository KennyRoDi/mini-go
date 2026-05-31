package test;

var x int = 10;

func main() {
    if (x) { // ERROR: condition not bool
        print(x);
    };
};
