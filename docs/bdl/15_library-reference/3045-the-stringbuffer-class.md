---
title: "The StringBuffer class"
source: "fgl-topics/c_fgl_ClassStringBuffer.html"
breadcrumb: "Library reference > Built-in packages > The base package > The StringBuffer class"
type: "concept"
---

# The StringBuffer class

> The base.StringBuffer class is a built-in class designed to manipulate character strings.

This class is optimized for string operations such as scanning, replacements, concatenation.

Use the `base.StringBuffer` class instead of [`STRING`](../08_language-basics/0567-string.md "The STRING data type is a variable-length, dynamically allocated character string data type, without limitation.") variables to implement heavy-duty
string manipulations. When you use a `base.StringBuffer` object, you work directly on
the internal string buffer. When you use the `STRING` data type and modify a string,
the runtime system creates a new buffer. While this does not impact the performance of programs with
a user interface or even batch programs doing SQL, it can impact performance when you need to
rapidly process large character strings.

For example, you get much better performance with a `base.StringBuffer` object
than with a `STRING` variable when you need to:

- Process 500 KB of text (such as when you are performing a global search-and-replace of specific
  words).
- Concatenate many strings.

When you pass a `base.StringBuffer` object as a function parameter, the function
receives a variable that references the object. Passing the object by reference is much more efficient
than using a `STRING` that is passed by value, because `STRING` data is
copied on the stack. The function manipulates the original string, not a copy of the string.

> **Important:**
>
> The methods of this class use character positions and string length. When
> using byte length semantics, the length is expressed in bytes. When using [char length semantics](../09_advanced-features/0881-length-semantics-settings.md), the unit is characters. This
> matters when using a multibyte locale such as UTF-8.

## Child topics

- [base.StringBuffer methods](3046-base-stringbuffer-methods.md)
- [Examples](3068-examples.md): base.StringBuffer usage examples.
