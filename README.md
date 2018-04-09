# GO-Tutorials
This tutorials will give you a brief introduction about GO programming langugage. where you can learn the progrmming from scratch. 
the hands-on is performend on 'ripl.it' IDE used for GO language. This is the link : https://repl.it/repls/NavyAverageSystemsoftware
Go programming language is created by Robert Griesemer, Rob Pike, and Ken Thompson at Google in 2007. It is strongly and statically typed (type checking done at compile-time) language, provides inbuilt support for garbage collection, and supports concurrent programming.
A simple Go program consists: package declaration, import packages, functios, variables, statement and expression, comments.

1) Write your first Go program that would print "Hello World".

    //hello.go
    
 let's look at the each line and see how they work.
Defining package name is mandatory for writing Go programs. it is the starting point to run the program. Next line use import "fmt" it      is a preprocessing command that tells the Go compiler to include the files lying in fmt package.
The next line contain main() function where program execution begins. we can use the single line comment(//.....) and multiline comment(/*......*/) simillarly to java.
The next line contain fmt.Println(....) is another function available in Go which contain the message " Hello World" to display on       screen.
In Go if the name start with capital letter it means they are exported Function or variable/constant. such as Pi is exported from math package if it is pi it behaves differently.

## Basics
Go support different data type such as Boolean type, Numeric type, string type, derived type(Array, structure, union, function, slice, function, interface, map and channel type).

### Variables:
In any language variables are used to store values. In Go it is vary easy to define a static type varibles let's do it. In the following example the statement "var a, b" is defining variables of type int. Simillarly, varibale of type float and string are defined. it support wide varity of int (int8 - int64, uint8 - uint 64), float (float 32 - float 64, complex 64 - complex 124).

//variable.go

There exist two type kind of expression in Go:
lvalue => variables are lvalues that appears on the left side of the expression and refers to memory location.
rvalue => numeric literals apperas on right side of the expression that refer to a data value and stored at some address in memory.
Let's take another example that define the dynamic type declration (DTD)/ type interface in Go. In DTD the compiler itself interpret the type of variable based on the value passed to it. 

//mixedVar.go

##Constants
Constants refers to the fixed values in the program that cannot be altered during program execution. The fixed values are also know as Literals. Go support different type of literals such as Interger litereal, floating point literals and string literals. let's take an example for better understanding. here, const keyword is used to declare constants with a specific type.

//constant.go






