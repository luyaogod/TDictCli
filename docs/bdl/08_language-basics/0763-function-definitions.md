---
title: "FUNCTION definitions"
source: "fgl-topics/c_fgl_Functions_syntax.html"
breadcrumb: "Language basics > Functions > FUNCTION definitions"
type: "concept"
---

# FUNCTION definitions

> A FUNCTION definition defines a named procedure with a set of statements.

## Syntax 1 (legacy function syntax):

```
[PUBLIC|PRIVATE] FUNCTION function-name ( [ parameter-name [,...] ] )
    [ DEFINE parameter-name type-specification [,...] ]
    [ local-declaration [...] ]
    [ instruction 
    | [ RETURN expression [,...] ]
        [...] 
    ]
END FUNCTION
```

1. function-name is the function identifier.
2. parameter-name is the name of a formal argument of the function.
3. type-specification can be one of:
   - A [primitive type](0690-primitive-type-specification.md "Type definitions using a primitive data type define a primitive type.")
   - A [record definition](0717-record.md "The RECORD keyword defines a structured type or variable.")
   - An [array definition](0731-array.md "An array defines a vector variable with a list of elements.")
   - A [dictionary definition](0744-dictionary.md "A dictionary defines an associative array (hash-map) of elements.")
   - A [function type definition](0770-function-references.md "Function can be referenced and invoked dynamically in a CALL instruction, or in an expression.")
   - The name of a [user defined type](0753-type.md "Types define a synonym for a base or structured data type.")
   - The name of a [built-in class](../15_library-reference/2908-built-in-packages.md "These topics cover the built-in classes provided by the Genero Business Development Language.")
   - The name of an [imported extension
     class](../15_library-reference/3480-extension-packages.md "Several utility classes and functions are provided in additional packages.")
   - The name of an [imported Java class](../14_extending-the-language/2655-the-java-interface.md "The Java interface allows you to import Java classes and instantiate Java objects in your programs.")
4. local-declaration is a `DEFINE`, `CONSTANT` or
   `TYPE` instruction.
5. instruction is a language statement to be executed when the function is
   invoked.
6. expression is a value to be returned by the function.

## Syntax 2 (fully typed function):

```
[PUBLIC|PRIVATE] FUNCTION function-name (
     { parameter-name type-specification [ attributes-list ]
     | record-name record-type INOUT
     }
     [,...]
  )
 [ attributes-list ]
 [ RETURNS returns-specification ]
    [ local-declaration [...] ]
    [ instruction 
    | [ RETURN expression [,...] ]
        [...] 
    ]
END FUNCTION
```

where returns-specification
is:

```
{ type-specification [ attributes-list ]
| ( type-specification [ attributes-list ] [,...] )
| ( )
}
```

1. function-name is the identifier of the function, that must follow the
   convention for [identifiers](0551-identifiers.md "A Genero BDL identifier is a sequence of characters used to identify a program entity.").
2. parameter-name is the name of a formal argument of the function.
3. type-specification can be one of:
   - A [primitive type](0690-primitive-type-specification.md "Type definitions using a primitive data type define a primitive type.")
   - A [record definition](0717-record.md "The RECORD keyword defines a structured type or variable.")
   - An [array definition](0731-array.md "An array defines a vector variable with a list of elements.")
   - A [dictionary definition](0744-dictionary.md "A dictionary defines an associative array (hash-map) of elements.")
   - A [function type definition](0770-function-references.md "Function can be referenced and invoked dynamically in a CALL instruction, or in an expression.")
   - The name of a [user defined type](0753-type.md "Types define a synonym for a base or structured data type.")
   - The name of a [built-in class](../15_library-reference/2908-built-in-packages.md "These topics cover the built-in classes provided by the Genero Business Development Language.")
   - The name of an [imported extension
     class](../15_library-reference/3480-extension-packages.md "Several utility classes and functions are provided in additional packages.")
   - The name of an [imported Java class](../14_extending-the-language/2655-the-java-interface.md "The Java interface allows you to import Java classes and instantiate Java objects in your programs.")
4. record-name is the name of a parameter defined as
   record-type.
5. record-type is a the name of a user-defined [`TYPE`](0751-types.md "Types can be defined by the programmer to centralize the definition of complex/structured variables.") declared as a [`RECORD`](0717-record.md "The RECORD keyword defines a structured type or variable.") structure.
6. attributes-list is a comma-separated list of name = value
   pairs or name attributes, and defines [attributes for the function](0769-function-attributes.md "Function attributes can be used to add definition information about the function, its parameters and its return values.").
7. local-declaration is a `DEFINE`, `CONSTANT` or
   `TYPE` instruction.
8. instruction is a language statement to be executed when the function is
   invoked.
9. expression is a value to be returned by the function.

## Syntax 3 (methods):

```
[PUBLIC|PRIVATE] FUNCTION ( receiver-name receiver-type ) method-name (
     { parameter-name type-specification
     | record-name record-type INOUT
     }
     [,...]
  )
 [ RETURNS returns-specification ]
    [ local-declaration [...] ]
    [ instruction 
    | [ RETURN expression [,...] ]
        [...] 
    ]
END FUNCTION
```

where returns-specification
is:

```
{ type-specification
| ( type-specification [,...] )
| ( )
}
```

1. A `(receiver-name receiver-type)` clause
   defines a function as a method for a user-defined type.
2. receiver-name is the identifier of the receiver referenced in the method
   body.
3. receiver-type is the user-defined type to be the target of this method.
4. method-name is the identifier of the method.
5. parameter-name is the name of a formal argument of the method.
6. type-specification can be one of:
   - A [primitive type](0690-primitive-type-specification.md "Type definitions using a primitive data type define a primitive type.")
   - A [record definition](0717-record.md "The RECORD keyword defines a structured type or variable.")
   - An [array definition](0731-array.md "An array defines a vector variable with a list of elements.")
   - A [dictionary definition](0744-dictionary.md "A dictionary defines an associative array (hash-map) of elements.")
   - A [function type definition](0770-function-references.md "Function can be referenced and invoked dynamically in a CALL instruction, or in an expression.")
   - The name of a [user defined type](0753-type.md "Types define a synonym for a base or structured data type.")
   - The name of a [built-in class](../15_library-reference/2908-built-in-packages.md "These topics cover the built-in classes provided by the Genero Business Development Language.")
   - The name of an [imported extension
     class](../15_library-reference/3480-extension-packages.md "Several utility classes and functions are provided in additional packages.")
   - The name of an [imported Java class](../14_extending-the-language/2655-the-java-interface.md "The Java interface allows you to import Java classes and instantiate Java objects in your programs.")
7. record-name is the name of a parameter defined with a record
   `TYPE`.
8. record-type is a the name of a user-defined [`TYPE`](0751-types.md "Types can be defined by the programmer to centralize the definition of complex/structured variables.") declared as a `RECORD`
   structure.
9. local-declaration is a `DEFINE`, `CONSTANT` or
   `TYPE` instruction.
10. instruction is a language statement to be executed when the method is
    invoked.
11. expression is a value to be returned by the method.

## Function names

Like other identifiers, function names are case-insensitive. However, always consider using the
same naming convention when defining and invoking functions.

If the function name is also the name of a built-in function, an error occurs at link time, even
if the program does not reference the built-in function.

## Function definition using legacy syntax

The following example shows a function definition using the legacy syntax, with parameter
definition in the function body:

```
FUNCTION split(str, len)
    DEFINE str STRING, len INT
    RETURN str.subString(1, len),
           str.subString(len+1, str.getLength())
END FUNCTION
```

This syntax does not define a function with a complete signature.

## Function definition with complete function type

By specifying data types in the parameter list, you define a function with a complete function
type.

This syntax allows better compilation
checking:

```
FUNCTION split(str STRING, len INT) RETURNS (STRING, STRING)
    RETURN str.subString(1, len),
           str.subString(len+1, str.getLength())
END FUNCTION
```

Braces after the `RETURNS` clause are not required, when the function returns a
single
value:

```
FUNCTION count_items(sid INT) RETURNS INT
    DEFINE cnt INT
    SELECT COUNT(*) INTO cnt FROM stock WHERE stock_id = sid
    RETURN cnt
END FUNCTION
```

When the function returns a single value/type, it is also possible to enclose the type in
parentheses:

```
FUNCTION count_items(sid INT) RETURNS (INT)
    ...
```

In the next example, the function has no parameters and does not return
values:

```
FUNCTION clean_debug_log() RETURNS ()
    DELETE FROM debug_log
END FUNCTION
```

## Methods

A `FUNCTION` definition using a receiver identifier and type defines a method for
this type. Specify the receiver identifier and type before the method
name:

```
TYPE Rectangle RECORD
    height, width DOUBLE PRECISION
END RECORD

FUNCTION (r Rectangle) area() RETURNS DOUBLE PRECISION
    RETURN r.height * r.width
END FUNCTION
```

For more details see [Methods](0772-methods.md "A function declared with a receiver type defines a method for this type.").

## Function with record passed as reference

The `INOUT` keyword can be used to pass records by reference:

```
SCHEMA custdemo
TYPE CustRec RECORD LIKE customer.*
FUNCTION init_cust(r CustRec INOUT)
    INITIALIZE r.* TO NULL
    LET r.cust_num = 0
    LET r.cust_name = "<undefined>"
END FUNCTION
```

For more details see [Function parameters](0767-function-parameters.md "Functions can take parameters, to specialize their behavior.").

## Function with attributes

Function attributes can be used to complete the definition of the function itself, for its
parameters and return values:

```
FUNCTION add(
    p1 INTEGER ATTRIBUTES(json_name = "parameter 1"),
    p2 INTEGER ATTRIBUTES(json_name = "parameter 2")
    )
    ATTRIBUTES(json_name = "function add")
    RETURNS INTEGER ATTRIBUTES(json_name = "return value 1")
        DEFINE var1 INTEGER
        LET var1 = (p1 + p2)
        RETURN var1
END FUNCTION
```

Function attributes are especially useful to implement [RESTful
Web Services functions](../16_web-services/4854-restful-web-services-low-level-apis.md "With Genero Business Development Language, you can code a RESTful Web Services client or server application using low-level APIs.").

For more details see [Function attributes](0769-function-attributes.md "Function attributes can be used to add definition information about the function, its parameters and its return values.").

## Related links

**Related concepts**  

[Variables](0686-variables.md "Explains how to define program variables.")

[Expressions](0592-expressions.md "Shows the possible expressions supported in the language.")

[Commenting a function](../13_programming-tools/2556-commenting-a-function.md "Commenting a function")
