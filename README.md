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
