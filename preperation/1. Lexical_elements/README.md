
- Comments:
  - `//` line comments
  - `/* */` Multiple line comments
  
- Tokens:
  - `Identifiers`- variables and types name. 
    - Must start with a letter(including `_`) and then combination of letter and digits
  - `Keywords`- Reserved keywords, such as, `var`, `func`, `struct`, `if` etc 
  - `Operators and Punctuation`- `+ & += &&` etc
  - `Literals`:
    - Integer literal: `4_2` is equivalent to `42`
        - ```go
            package main
            
            import "fmt"
            
            func main() {
                x := 4_2
                fmt.Println(x + 1)
            }
            // Outputs: 43
          ```
    - Floating-point literal:
        - ```go
            package main
            
            import "fmt"
            
            func main() {
                x := 0.
                y := .25
                fmt.Println(x)
                fmt.Println(y + 1)
            }
                  
            ```      
  - Rune literals: One or more characters enclosed in single quotes, as in 'x' or '\n'
  - String literals: concatenating a sequence of characters.
    - Raw string literals: sequence of chars between ` `` `, any character may appear except back quote "`"   
    - Interpreted string literals: sequence of chars between `" "`
- White space, horizontal tabs, carriage returns, new lines is ignored.
- Newline or EOF may trigger the insertion of a semicolon.
- ";" is used as a terminator.