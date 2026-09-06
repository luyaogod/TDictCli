---
title: "Understanding arrays"
source: "fgl-topics/c_fgl_Arrays_intro.html"
breadcrumb: "Language basics > Arrays > Understanding arrays"
type: "concept"
---

# Understanding arrays

> This is an introduction to arrays.

Arrays can store a one-, two- or three-dimensional set of elements.

The language supports three kind of array types:

- Static arrays - introduced in early versions of the language.
- Dynamic arrays - to be used in new developments.
- Java arrays - to define a Java array, to interface with Java classes.

For static and dynamic arrays, elements can be of simple built-in types such as [`INTEGER`](0562-integer.md "The INTEGER data type is used for storing large whole numbers."), [`VARCHAR(n)`](0570-varchar-size.md "The VARCHAR data type is a variable-length character string data type, with a maximum size."), [`RECORD`](0715-records.md "Records allow structured program variables definitions.") structured types, or [user-defined types](0751-types.md "Types can be defined by the programmer to centralize the definition of complex/structured variables."). A BDL array can also be defined with
built-in classes, imported module classes (from the Web Services extensions for example), or Java
classes.

The first element in a BDL array is at position 1, for static, dynamic and Java arrays.

Arrays can be used to define a list of records that will be controlled by dialog instructions
such as [`DISPLAY ARRAY`](../11_user-interface/1959-record-list-display-array.md "The DISPLAY ARRAY instruction provides record list navigation in an application form, with optional record modification actions.") and [`INPUT ARRAY`](../11_user-interface/2005-editable-record-list-input-array.md "The INPUT ARRAY instruction provides always-editable record list handling in an application form.").

Dynamic arrays can be manipulated with [methods](0737-array-methods.md "Native BDL arrays and Java arrays can be used to invoke built-in methods."), to
append, insert, delete elements, clear the array, copy one dynamic array to another, search in the
elements.

## Related links

**Related concepts**  

[Dictionaries](0742-dictionaries.md "A dictionary holds an unordered collection of elements accessed by a key.")

[Using Java arrays](../14_extending-the-language/2687-using-java-arrays.md "Using Java arrays")
