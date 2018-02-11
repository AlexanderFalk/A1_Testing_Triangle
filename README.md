# Getting Started Testing - Assignment 1

This program states whether a triangle is either: scalene, equalateral or isosceles, based on three parameters that you insert.

To use the program you can clone the repository into your chosen destination. When you got the folder associated with the files, you can try it out. A prerequisites to compile the program is to have [GOLANG](https://golang.org/dl/) installed on your machine. When it is installed, you can jump into your terminal or command-prompt and type: 

``` go build triangle.go traingle_test.go```

This will compile and build the program to use. You'll on Windows get an .exe file, where on MacOS/LINUX you'll receive an executable file.

For Windows you in your command-prompt type the following example:

```bash
triangle.exe whatis --values "1 2 2"
```

Which will return:

```bash
Your triangle is of type: isosceles triangle!
```

On MacOS/Linux you can do the same, just with a little adjustment:

```bash
./triangle whatis --values "2 2 2"
```

Which will return:

```bash
Your triangle is of type: equalateral triangle!
```



#### Test

The file  ``` triangle_test.go``` contains the UNIT Testing for the Triangle program. The Unit Testing has been written in GOLANG and you can see if program passes with typing:

```bash
go test -v
```

You've to be located at the folder location to make the test. This will automatically see if their is any testing to be executed. 

Based on the Unit Testing from the file you'll see a respone in the console similiar to below, unless you change the values in the test file:

```bash
=== RUN   Test_equalateral_triangle
Your triangle is of type: equalateral triangle!
--- PASS: Test_equalateral_triangle (0.00s)
=== RUN   Test_isosceles_triangle
Your triangle is of type: isosceles triangle!
--- PASS: Test_isosceles_triangle (0.00s)
=== RUN   Test_scalene_triangle
Your triangle is of type: scalene triangle!
--- PASS: Test_scalene_triangle (0.00s)
PASS
ok
```

