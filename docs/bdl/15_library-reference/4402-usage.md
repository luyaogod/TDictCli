---
title: "Usage"
source: "fgl-topics/c_fgl_ext_reflect_usage.html"
breadcrumb: "Library reference > Extension packages > The reflect package > Usage"
type: "concept"
description: "The Genero BDL reflection API allows you to implement generic code, to introspect program elements at runtime. In programming, introspection is the ability to describe elements of a program at ..."
---

# Usage

The Genero BDL reflection API allows you to implement generic code, to introspect program
elements at runtime.

In programming, introspection is the ability to describe elements of a program at
runtime. This feature allows you to implement common generic code to manipulate program elements the
definition of which is not known at compile time.

Program elements such as [variables](../08_language-basics/0686-variables.md "Explains how to define program variables."), [user-defined types](../08_language-basics/0751-types.md "Types can be defined by the programmer to centralize the definition of complex/structured variables."), [anonymous types](../08_language-basics/0755-anonymous-types.md "Anonymous types are created from DEFINE instructions."), [function references](../08_language-basics/0770-function-references.md "Function can be referenced and invoked dynamically in a CALL instruction, or in an expression."),
[methods](../08_language-basics/0772-methods.md "A function declared with a receiver type defines a method for this type.") and [interfaces](../08_language-basics/0778-interfaces.md "An interface groups a set of methods acting on a user-defined type.") can be inspected by the reflection API.

The reflection API can also modify the value of program variables, and add/remove elements or
arrays and dictionaries.

## Importing the `reflect` package

In order to use the reflection API, you need to import the `reflect` package in
your module.

The reflection API is provided as a [C Extension](../14_extending-the-language/2703-c-extensions.md "With C-Extensions, you can bind your own C libraries in the runtime system, to call C function from the application code.")
module named `reflect`, containing a set of classes to manipulale program
elements.

The `reflect` module must be imported at the beginning of your
.4gl module(s) with:

```
IMPORT reflect
```

## Exception handling

If needed, the errors thrown by the reflection API must be handled with a [`TRY / CATCH`](../09_advanced-features/0853-try-catch-block.md "Use TRY / CATCH blocks to trap runtime exceptions in a delimited code block.")
block:

```
DEFINE val reflect.Value
DEFINE rec RECORD ... END RECORD
TRY
    LET val = reflect.Value.valueOf( rec )
    ...
CATCH
    DISPLAY "ERROR: Could not use reflect.Value object for var"
    RETURN -1
END TRY
```

> **Important:**
>
> The exceptions thrown by the
> `reflect.*` API can only be caught with a [`TRY/CATCH`](../09_advanced-features/0853-try-catch-block.md "Use TRY / CATCH blocks to trap runtime exceptions in a delimited code block.") block: If [`WHENEVER
> ERROR CONTINUE`](../09_advanced-features/0850-whenever-directive.md "Use the WHENEVER directive to define how exceptions must be handled for the rest of the module.") is active and the reflection API throws an exception, the program
> stops.

## Manipulating variables: `reflect.Value`

Variables can be inspected and modified with the `reflect.Value` class.

`reflect.Value` objects must first be created with the
`reflect.Value` class methods `copyOf()` or `valueOf()`:

- [`reflect.Value.copyOf(
  expression )`](4356-reflect-value-copyof.md "Creates a new reflect.Value object from a copy of an expression.") creates a `reflect.Value` object that
  references a copy of the expression passed as parameter. This is typically used to create a value
  from a [literal](../08_language-basics/0585-literals.md "Describes the syntax of literals (constant values) to be used in sources."), to be assigned to a
  `reflect.Value` object with the [`set()`](4377-reflect-value-set.md "Assigns the specified value to this value object.") method.
- [`reflect.Value.valueOf(
  variable )`](4357-reflect-value-valueof.md "Creates a new reflect.Value object as a reference to the original variable.") creates a `reflect.Value` object that
  references the variable passed as parameter. This is used to describe and/or modify the underlying
  variable.

Once you have a `reflect.Value` object, you can get its corresponding type object
with the [`getType()`](4371-reflect-value-gettype.md "Returns the reflect.Type object of a reflect.Value object.") method,
to inspect detailed information about the variable type.

Some `reflect.Value` methods are specific to a given [type kind](4388-reflect-type-getkind.md "Returns the kind of a type.") of the value object. It is for example
only possible to call the [`appendArrayElement()`](4358-reflect-value-appendarrayelement.md "Appends a new element to an array.") method with a value object of the kind
`"ARRAY"`.

For a complete list of methods, see [`reflect.Value` methods](4355-reflect-value-methods.md "Methods for the reflect.Value class.").

## Inspecting types: `reflect.Type`

Types can be inspected with the `reflect.Type` class.

`reflect.Type` objects must first be created with the [`reflect.Type.typeOf()`](4381-reflect-type-typeof.md "Creates a new reflect.Type object representing a type.") class method
or from a `reflect.Value` object with the [`getType()`](4371-reflect-value-gettype.md "Returns the reflect.Type object of a reflect.Value object.") method:

1. [`reflect.Type.typeOf()`](4381-reflect-type-typeof.md "Creates a new reflect.Type object representing a type.")
   can create a `reflect.Type` object directly from the expression or variable passed as
   parameter.
2. [`reflect.Value.getType()`](4371-reflect-value-gettype.md "Returns the reflect.Type object of a reflect.Value object.") returns the `reflect.Type` object
   representing the type of a existing `reflect.Value` object.

Each type belongs to a [type kind](4388-reflect-type-getkind.md "Returns the kind of a type."). Some
methods of the `reflect.Type` class apply to certain kind of types, not to all kind
of types.

For a complete list of methods, see [`reflect.Type` methods](4380-reflect-type-methods.md "Methods for the reflect.Type class.").

## Inspecting methods: `reflect.Method`

Methods can be inspected with the `reflect.Method` class.

`reflect.Method` objects must first be created from a
`reflect.Type` object, with the [`getMethod()`](4389-reflect-type-getmethod.md "Returns a reflect.Method object of record with methods or an interface.") method. The type object must represent a record-type with [methods](../08_language-basics/0772-methods.md "A function declared with a receiver type defines a method for this type."), or an [interface](../08_language-basics/0778-interfaces.md "An interface groups a set of methods acting on a user-defined type.").

For a complete list of methods, see [The reflect.Method class](4394-the-reflect-method-class.md "The reflect.Method class is a generic API to inspect methods.").

## From specific to generic code

The key to write generic code is that the "specific" language elements must be passed over to the
code using the reflection API. For program variables, this is done with the [`reflect.Value.valueOf()`](4357-reflect-value-valueof.md "Creates a new reflect.Value object as a reference to the original variable.") class
method.

To write a library with the reflection API, you can for example implement a function or method
that takes a `reflect.Value` object as parameter, that will be registered locally.
The "specific" code using the library can then pass its variables over for processing.

In the next example, the `setVariable()` method registers the variable to be
processed by the `find()` method, of the `t_find` utility type.

The myutils.4gl source:

```
IMPORT reflect
  
PUBLIC TYPE t_find RECORD
         val reflect.Value
     END RECORD

PUBLIC FUNCTION (self t_find) setVariable( v reflect.Value ) RETURNS ()
    IF v.getType().getKind() != "RECORD" THEN
        DISPLAY "ERROR: Can only process records"
        EXIT PROGRAM 1
    END IF
    LET self.val = v
END FUNCTION

PUBLIC FUNCTION (self t_find) find( pattern STRING ) RETURNS (STRING)
    DEFINE m reflect.Value
    DEFINE t reflect.Type
    DEFINE x, fc INTEGER
    LET t = self.val.getType()
    LET fc = t.getFieldCount()
    FOR x=1 TO fc
        LET m = self.val.getField(x)
        IF m.toString() LIKE pattern THEN
            RETURN t.getFieldName(x)
        END IF
    END FOR
    RETURN NULL
END FUNCTION
```

The main.4gl program:

```
IMPORT reflect
IMPORT FGL myutils
MAIN
    DEFINE r_cust RECORD
               pkey INTEGER,
               name VARCHAR(50),
               addr VARCHAR(200)
           END RECORD

    DEFINE f myutils.t_find
    
    LET r_cust.pkey = 111
    LET r_cust.name = "Mike Rean"
    LET r_cust.addr = "145 Sunset St."

    CALL f.setVariable( reflect.Value.valueOf(r_cust) )
    
    DISPLAY "Search 1: ", f.find( "xxx" )
    DISPLAY "Search 2: ", f.find( "Mike%" )
    DISPLAY "Search 3: ", f.find( "%Sunset%" )
    
END MAIN
```

## Manipulating TEXT/BYTE data

`TEXT` and `BYTE` variable have methods to manipulate the LOB data,
such as [`readFile()`](2941-text-readfile.md "Reads a file into a TEXT locator.").

To write generic code for `TEXT/BYTE` variables, use the [`reflect.Type.isAssignableFrom()`](4392-reflect-type-isassignablefrom.md "Checks if this type can be assigned from another type.") method, to check if the corresponding
`reflect.Value` object can be assigned to a `TEXT` or
`BYTE` variable, assign to a local `TEXT` or `BYTE`
variable with [`reflect.Value.assignToVariable()`](4359-reflect-value-assigntovariable.md "Assigns this reflect.Value to a variable."), then call required methods with that
local variable:

```
IMPORT reflect
  
FUNCTION loadLobFromFile(v reflect.Value, fn STRING) RETURNS ()
    DEFINE tx TEXT, bt BYTE
    DEFINE t reflect.Type
    LET t = v.getType()
    CASE
    WHEN t.isAssignableFrom(reflect.Value.valueOf(tx).getType())
       CALL v.assignToVariable(tx)
       CALL tx.readFile(fn)
    WHEN t.isAssignableFrom(reflect.Value.valueOf(bt).getType())
       CALL v.assignToVariable(bt)
       CALL bt.readFile(fn)
    OTHERWISE
       DISPLAY "ERROR: Invalid reflect.Value object"
       EXIT PROGRAM 1
    END CASE
END FUNCTION

FUNCTION main()
    DEFINE tx TEXT
    DEFINE v_tx reflect.Value
    LOCATE tx IN MEMORY
    LET v_tx = reflect.Value.valueOf(tx)
    CALL loadLobFromFile(v_tx,"data.txt")
    DISPLAY "tx = ", tx
END FUNCTION
```

## Manipulating DYNAMIC ARRAY elements

To append a new element to a `DYNAMIC ARRAY`, first append the element, then get
the value object and set its value:

```
IMPORT reflect

MAIN
    DEFINE a DYNAMIC ARRAY OF STRING
    DEFINE va, ve reflect.Value
    LET va = reflect.Value.valueOf(a)
    CALL va.appendArrayElement()
    LET ve = va.getArrayElement(1)
    CALL ve.set(reflect.Value.copyOf("new-value"))
    DISPLAY va.getArrayElement(1).toString()
END MAIN
```

## Manipulating DICTIONARY elements

To append a new element to a `DICTIONARY`, first get an element with the key (this
will automatically create the new element if it does not exist), then set the value of the
element:

```
IMPORT reflect

MAIN
    DEFINE d DICTIONARY OF STRING
    DEFINE v reflect.Value
    LET v = reflect.Value.valueOf(d)
    CALL v.getDictionaryElement("key1").set(reflect.Value.copyOf("new-value"))
    DISPLAY v.getDictionaryElement("key1").toString()
END MAIN
```
