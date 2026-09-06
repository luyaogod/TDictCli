---
title: "Manipulating character strings"
source: "fgl-topics/c_fgl_optimization_017.html"
breadcrumb: "Advanced features > Optimization > Optimize your programs > Manipulating character strings"
type: "concept"
description: "When parsing, concatenating or accessing parts of a character string contained in CHAR , VARCHAR and STRING variables, the runtime system must follow the BDL language semantics, which can result in ..."
---

# Manipulating character strings

When parsing, concatenating or accessing parts of a character string contained in [`CHAR`](../08_language-basics/0557-char-size.md "The CHAR data type is a fixed-length character string data type."), [`VARCHAR`](../08_language-basics/0570-varchar-size.md "The VARCHAR data type is a variable-length character string data type, with a maximum size.") and [`STRING`](../08_language-basics/0567-string.md "The STRING data type is a variable-length, dynamically allocated character string data type, without limitation.") variables, the runtime system must
follow the BDL language semantics, which can result in performance issues with very large character
strings.

When heavy string manipulation is to be done, consider using the [`base.StringBuffer`](../15_library-reference/3045-the-stringbuffer-class.md "The base.StringBuffer class is a built-in class designed to manipulate character strings.") class instead of
`CHAR`/`VARCHAR`/`STRING` variables: The
`base.StringBuffer` class is designed for character string manipulation.

Furthermore, execution time of character iterations (to find a specific sub-string for example or
get the position of a character in a string), can be affected when using the UTF-8 [locale](0864-application-locale.md "The application locale defines the language and codeset for your application.") and [CHAR
length semantics usage](0881-length-semantics-settings.md): The runtime needs to compute the number of bytes representing a
character when iterating through the string.

## Related links

**Related concepts**  

[Passing CHAR parameters to functions](0975-passing-char-parameters-to-functions.md "Passing CHAR parameters to functions")

[Substring ([s,e])](../08_language-basics/0633-substring-s-e.md "The [] (square brackets) operator extracts a substring from a variable.")
