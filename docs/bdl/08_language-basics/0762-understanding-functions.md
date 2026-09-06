---
title: "Understanding functions"
source: "fgl-topics/c_fgl_Functions_002.html"
breadcrumb: "Language basics > Functions > Understanding functions"
type: "concept"
---

# Understanding functions

> This is an introduction to functions.

## Function basics

Functions are named program blocks containing a set of statements to be executed when the
function is invoked with a [`CALL`](0675-call.md "The CALL instruction invokes a specified function or method.")
statement, or when the function is used in an [expression](0592-expressions.md "Shows the possible expressions supported in the language."), or when the function is registered in a callback mechanism like [`WHENEVER ERROR CALL`](../09_advanced-features/0850-whenever-directive.md "Use the WHENEVER directive to define how exceptions must be handled for the rest of the module.").

## Parameters and return values

Functions can get [input parameters](0767-function-parameters.md "Functions can take parameters, to specialize their behavior.") and [return](0768-returning-values.md "A function can return values with the RETURN instruction.") zero, one or several values.

## Scope of functions

A function is defined in a program module, and is by default visible to all modules (it is
`PUBLIC`). A function can be declared as `PRIVATE` to the module where
it is defined. In this case the function is hidden to other modules. This concept is explained in
[Scope of a function](0764-scope-of-a-function.md "A functions can be isolated to control its visibility to other modules.")

## Function references

It is possible to hold a [function reference](0770-function-references.md "Function can be referenced and invoked dynamically in a CALL instruction, or in an expression.")
in a program variable. The variable can then be assigned with any function reference that is
declared with the same function signature as the type used to define the variable.

Function references apply only to regular functions. It is not possible to use a method
reference.

## Functions as methods for a type

When declaring a function with a receiver variable and receiver type enclosed in parentheses
before the function name, it becomes a method acting on the specified type.

Use [methods](0772-methods.md "A function declared with a receiver type defines a method for this type.") in conjunction with [interfaces](0778-interfaces.md "An interface groups a set of methods acting on a user-defined type.") to write clear robust code by using encapsulation and
polymorphism concepts.

## Function attributes

Functions can be declared with [function
attributes](0769-function-attributes.md "Function attributes can be used to add definition information about the function, its parameters and its return values."), to define additional information on the function.
