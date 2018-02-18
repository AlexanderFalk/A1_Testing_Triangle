# Getting Started Testing - Assignment 1

This program states whether a triangle is either: scalene, equalateral or isosceles, based on three parameters that you insert.

To use the program you can clone the repository into your chosen destination. When you got the folder associated with the files, you can try it out. A prerequisites to compile the program is to have [GOLANG](https://golang.org/dl/) installed on your machine. When it is installed, you can jump into your terminal or command-prompt and type: 

``` go build triangle.go traingle_test.go```

This will compile and build the program to use. You'll on Windows get an .exe file, where on MacOS/LINUX you'll receive an executable file.

For Windows you in your command-prompt type the following example:

```bash
triangle.exe find --values "1 2 2"
```

Which will return:

```bash
Your triangle is of type: isosceles triangle!
```

On MacOS/Linux you can do the same, just with a little adjustment:

```bash
./triangle find --values "2 2 2"
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
=== RUN   TestEquilateralTriangle
Your triangle is of type: equi lateral triangle!
--- PASS: TestEquilateralTriangle (0.00s)
=== RUN   TestIsoscelesTriangle
Your triangle is of type: isosceles triangle!
--- PASS: TestIsoscelesTriangle (0.00s)
=== RUN   TestScaleneTriangle
Your triangle is of type: scalene triangle!
--- PASS: TestScaleneTriangle (0.00s)
=== RUN   TestInvalidTriangle
That is not a triangle my dear friend...
--- PASS: TestInvalidTriangle (0.00s)
=== RUN   TestTriangleOne
That is not a triangle my dear friend...
--- PASS: TestTriangleOne (0.00s)
=== RUN   TestTriangleTwo
Your triangle is of type: isosceles triangle!
--- PASS: TestTriangleTwo (0.00s)
=== RUN   TestTriangleThree
That is not a triangle my dear friend...
--- PASS: TestTriangleThree (0.00s)
=== RUN   TestCheckZeroLengthValue
A triangle can't be 0 in one of its length or it can't be any minus value
--- PASS: TestCheckZeroLengthValue (0.00s)
=== RUN   TestCheckMinusValue
A triangle can't be 0 in one of its length or it can't be any minus value
--- PASS: TestCheckMinusValue (0.00s)
PASS
ok      testcourse/a1_triangle  0.078s
```

