---
title: "Understanding dictionaries"
source: "fgl-topics/c_fgl_Dictionary_intro.html"
breadcrumb: "Language basics > Dictionaries > Understanding dictionaries"
type: "concept"
---

# Understanding dictionaries

> This is an introduction to dictionaries.

Dictionaries can store an unordered collection of elements, that will be accessed by key string,
rather than by index as in arrays.

Dictionaries are used to implement a hash map (a.k.a. associative array).

The keys must be character
string:

```
DEFINE dict DICTIONARY OF INTEGER
LET dict["abcdef"] = 999
```

Dictionary elements can be of simple built-in types such as [`INTEGER`](0562-integer.md "The INTEGER data type is used for storing large whole numbers."), [`VARCHAR(n)`](0570-varchar-size.md "The VARCHAR data type is a variable-length character string data type, with a maximum size."), [`RECORD`](0715-records.md "Records allow structured program variables definitions.") structured types, or [user-defined types](0751-types.md "Types can be defined by the programmer to centralize the definition of complex/structured variables."). A dictionary can also be defined with
built-in classes, imported module classes (from the Web Services extensions for example), or Java
classes.

## Related links

**Related concepts**  

[Arrays](0729-arrays.md "Arrays (static or dynamic) allow you to handle an ordered collection of elements.")
