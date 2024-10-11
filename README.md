# GO Short Tutorial & Interview Questions

1.  What keyword is used to specify the default case in a switch statement in Go?

        Answer: default 

2. Which of the following is true about variadic functions in Go?
        


        Answer: They can accept a variable number of arguments
    ```go 
    func printNumbers(nums ...int) {
        for _, num := range nums {
            fmt.Println(num)
        }
    }

    printNumbers(1, 2, 3)      // Works with three arguments
    printNumbers(4, 5)         // Works with two arguments
    printNumbers()   // works with no argument


3. What keyword is used to declare a variable in Go?

        Answer: var

4. The _______ function in Go is used to panic, terminating the program immediately with a runtime error message.

        Answer: panic

5. The _______ data type in Go is used to represent decimal numbers with floating-point precision.

            Answer: Float32

6. The _______ data type in Go is used to represent decimal numbers with floating-point precision.

            Answer: ==

7. Which keyword is used to include a package in Go?

        Answer: import

8. Which keyword is used to include a package in Go?

        Answer: a,b= b, a

9. In Go, can variables be declared without being initialized?

        Answer: Yes

10. Which operator in Go is used to concatenate strings?

        Answer: +

11. In Go, a function can return multiple _______.

        Answer: Values

12. What is the size of the 'int' data type in Go on a 64-bit system?

    Answer: 8 bytes

13. You're developing a Go application where you need to declare a constant named 'MaxRetryAttempts' to specify the maximum number of retry attempts. Which data type would you use for this constant?

        Answer: int

14. What is the difference between the '=' and ':=' operators in Go?

        Answer: Both operators are used for assignment

15. What does the '<<' operator do in Go when used with integers?

        Answer: Left shifts the bits of an integer

16. In Go, which loop is used to iterate over elements in an array, slice, string, or map?

        Answer: for loop

17. You're writing a Go program where you need to define the value of Pi as a constant. Which keyword would you use to declare Pi as a constant?

        Answer: const

18. Which data type in Go is used to represent true or false values?

        Answer: bool

19. What is the difference between import . "package" and import _ "package" in Go?

        Answer: import . "package" allows referring to package symbols without package name prefix whereas import _ "package" only imports the package for its side effects wihtout making its symbols accessible

20. The _______ data type in Go is used to represent a sequence of bytes.

        Answer: bytes.Buffer

21. What does the blank identifier (_) represent when used in an import statement in Go?

        Answer: It represents importing the package only for its side effects

22. In Go, what is the purpose of a package?

        Answer: Encapsulate related code

23. The _______ operator in Go is used to perform logical OR.

        Answer: ||

24. What is the scope of a variable declared inside a function in Go?

        Answer: Local to the function

25. In a Go application, you encounter a circular import error. How would you refactor your code to resolve this issue?

        Answer: Extract the common functionality causing the circular import into a new package

26. What does the '==' operator do in Go when used with slices?

        Answer: It is not a valid operator to use with slices in Go

27. Which keyword is used to declare constants in Go?

        Answer: switch

28. How does Go handle cyclic dependencies between packages?

        Answer: Go does not support cyclic dependencies between packages

29. Variables declared with the ':=' syntax in Go are automatically _______ based on the assigned value.

        Answer: typed

30. The _______ data type in Go is used to represent a pointer to any type.

        Answer: interface{}

31. Which control structure in Go is used to execute a block of code repeatedly based on a condition, but at least once, even if the condition is false initially?

        Answer: for

32. What is a higher-order function in Go?

        Answer: A function that operates on other functions, either by taking them as arguments or returning them.
33. Constants in Go are immutable, meaning their values cannot be _______ after declaration.

        Answer: Changed

34. Which operator in Go is used to perform bitwise XOR?

        Answer: ^

35. Suppose you're implementing a command-line tool in Go, and you want to handle different options provided by the user. Which control structure in Go would be the most suitable choice for this scenario?

        Answer: switch

36. Which data type in Go is used to represent a sequence of bytes?

        Answer: slice

37. Suppose you're building a system where you need to represent and manipulate raw binary data. Which data type would be most appropriate for this task?

        Answer: byte

38. In Go, can you declare multiple variables in a single declaration statement?

        Answer: Yes

39. You're designing a security module for your Go application. Which operator would you use to perform constant-time comparison of two byte slices to prevent timing attacks?

        Answer: bytes.Equal()

40. In Go, which loop is typically used when the number of iterations is known?


        Answer: for

41. The '_____' operator in Go is used to shift bits to the right.

        Answer: >>

42. Which data type in Go is used to represent decimal numbers with floating-point precision?

        Answer: float64

43. What happens if there is no match in a switch statement in Go and there is no default case?

        Answer: It will execute the code block of the case immediately following the switch statement

44. You're working on a large Go project where multiple packages need to be imported. How would you organize your import statements for clarity and maintainability?

        Answer: Group related packages together and import them in separate blocks

45. The _______ data type in Go is used to represent a single Unicode character.

        Answer: rune

46. In Go, which data type is used to represent a pointer to any type?

        Answer: *

47. In Go, the _________ function is executed automatically before the main function in the same package.

        Answer: init

48. Which data type in Go is used to represent a sequence of characters?

        Answer: string

49. In Go, the _________ directory is used to define a package that is shared across multiple projects.

        Answer: src

50. What does the 'defer' keyword do in Go?

        Answer: Delays execution of a function until the end of the surrounding function

51. Which control structure in Go is used to execute a block of code repeatedly as long as a condition is true?

        Answer: for

52. What is the data type used to store whole numbers in Go?

        Answer: int

53. The behavior of the _______ statement in Go can be replicated using if-else chains, but switch statements provide a cleaner and more concise way to express such conditions.

        Answer: swtich

54. In Go, which statement is used to import a package that is within the same module?

        Answer: import "./package"

55. The '_____' operator in Go is used to perform bit clear (AND NOT).

        Answer: &^

56. In Go, what is a closure?

        Answer: A closure is a function value that references variable from outside its body
        


```go
        func makeCounter() func() int {
        count := 0
        return func() int {
                count++
                return count
        }
        }

        func main() {
                counter := makeCounter()

                fmt.Println(counter()) // Output: 1
                fmt.Println(counter()) // Output: 2
                fmt.Println(counter()) // Output: 3
        }
```



57. Can constants be of a complex data type, such as slices or structs, in Go?

        Answer: No

58. In Go, which data type is used to store complex numbers?

        Answer: complex

59. In a Go program, how can you ensure that an imported package is only compiled for testing purposes?

        Answer: By naming the test files with a _"test" suffix

 60. Which control structure in Go is used to execute one of many blocks of code based on the value of an expression?

        Answer: switch statement

61. The _______ function in Go is used to make a copy of a slice or map.

        Answer: copy

62. In Go, what is the difference between 'var' and ':=' when declaring variables?

        Answer: Declares a variable and assigns a value without specifying its type

63. Which of the following is not a valid way to call a function in Go?

        Answer: package.functionName(arguments)

64. You are developing a web application in Go, and you need to handle HTTP requests asynchronously. Which feature of Go would you use for this purpose?

        Answer: Goroutines

65. What is the purpose of the init function in Go packages?

        Answer: The init function is called when a package is imported, allowing initialization code to execute before the package is used

66. In a Go program, you're dealing with a large dataset where memory optimization is crucial. Which data type would you choose to represent a set of integer values ranging from -128 to 127?

        Answer: int8

67. The _______ keyword in Go is used to create a closure.

        Answer: func

68. When importing a package in Go, if we want to execute its init functions, we use the _________ form of the import statement.

        Answer: blank

69. The _______ operator in Go is used to perform pointer indirection.

        Answer: *

70. You're developing a Go library intended for use in other projects. What considerations should you keep in mind regarding package structure and imports?

        Answer: Use unique package names to avoid conflicts

71. Which of the following is true about error handling in Go functions?

        Answer: Errors are values and Go programmer are encouraged to return thems as part of the function's return values

72. In Go, what is the term used to describe a function that doesn't return any value?

        Answer: void