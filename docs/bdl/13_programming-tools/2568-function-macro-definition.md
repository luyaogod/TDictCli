---
title: "Function macro definition"
source: "fgl-topics/c_fgl_preprocessor_006.html"
breadcrumb: "Programming tools > Source preprocessor > Function macro definition"
type: "concept"
---

# Function macro definition

> Function macros are preprocessor macros which can take arguments.

## Syntax

```
&define identifier( arglist ) body
```

1. identifier is the name of the macro. Any valid identifier can be used.
2. body is any sequence of tokens until the end of the line.
3. arglist is a list of identifiers separated with commas and optionally white space.
4. There must be no space or comment between the macro name and the opening parenthesis
   `(`. Otherwise the macro is not a function macro, but a simple macro.

A preprocessor directive must start at the beginning of the line. The `&`
ampersand character can be preceded by blanks.

## Usage

Function macros take arguments that are replaced in the body by the preprocessor.

Source: File
A

```
&define function_macro(a,b) a + b 
&define simple_macro (a,b) a + b 
function_macro( 4 , 5 )
simple_macro (1,2)
```

Result (fglcomp -E
output):

```
& 1 "A"

4 + 5
(a,b) a + b (1,2)
```

A function macro can have an empty argument list. In this case, parentheses `()`
are required for the macro to be expanded. As we can see in the following example, the third line is
not expanded because there is no parentheses after `foo`. The function macro cannot
be applied even if it has no arguments.

Source: File A

```
&define foo() yes 
foo()
foo
```

Result (fglcomp -E
output):

```
& 1 "A"

yes 
foo
```

The comma separates arguments. Macro parameters containing a comma can be used with parentheses.
In this example, the second line has been substituted, but the third line produced an error, because
the number of parameters is incorrect.

Source: File
A

```
&define one_parameter(a) a 
one_parameter((a,b))
one_parameter(a,b)
```

`fglcomp -M
output`

```
source.4gl:3:1:3:1:error:(-8039) Invalid number of parameters
 for macro one_parameter.
```

Macro arguments are completely
expanded and substituted before the function macro expansion.

A
macro argument can be left empty.

Source: File
A

```
&define two_args(a,b) a b 
two_args(,b)
two_args(,)
two_args() 
two_args(,,)
```

`fglcomp -M
output`

```
source.4gl:4:1:4:1:error:(-8039) Invalid number of parameters
 for macro two_args. 
source.4gl:5:1:5:1:error:(-8039) Invalid number of parameters
 for macro two_args.
```

Macro arguments
appearing inside strings are not expanded.

Source: File A

```
&define foo(x) "x"
foo(toto)
```

Result (fglcomp -E
output):

```
& 1 "A"

"x"
```

## Related links

**Related concepts**  

[Simple macro definition](2567-simple-macro-definition.md "A simple macro is identified by its name and body.")
