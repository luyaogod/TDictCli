---
title: "Saving memory by using STRING variables"
source: "fgl-topics/c_fgl_optimization_014.html"
breadcrumb: "Advanced features > Optimization > Optimize your programs > Saving memory by using STRING variables"
type: "concept"
description: "The CHAR and VARCHAR data types are provided to hold string data from a database column. When you define a CHAR or VARCHAR variable with a length of 1000, the runtime system must allocate the entire ..."
---

# Saving memory by using STRING variables

The [`CHAR`](../08_language-basics/0557-char-size.md "The CHAR data type is a fixed-length character string data type.") and
[`VARCHAR`](../08_language-basics/0570-varchar-size.md "The VARCHAR data type is a variable-length character string data type, with a maximum size.") data
types are provided to hold string data from a database column.
When you define a `CHAR` or `VARCHAR` variable
with a length of 1000, the runtime system must allocate the entire
size, to be able to fetch SQL data directly into the internal string
buffer.

For character string data that is not stored in the database, consider
using the [`STRING`](../08_language-basics/0567-string.md "The STRING data type is a variable-length, dynamically allocated character string data type, without limitation.") data
type. The `STRING` type is similar to `VARCHAR`,
except that you don't need to specify a maximum length and the
internal string buffer is allocated dynamically as needed. Thus,
by default, a `STRING` variable initially requires
just a bunch of bytes, and grows during the program life time,
with a limitation of 65534 bytes.

Use of a `STRING` variable is typically recommended when building SQL statements
dynamically, for example from a [`CONSTRUCT`](../11_user-interface/2046-query-by-example-construct.md "The CONSTRUCT instruction implements database query criteria input in an application form.") instruction. You may also use the
`STRING` type for utility function parameters, to hold file names for
example.

After a large `STRING` variable is used, it should
be cleared with a [`LET`](../08_language-basics/0701-let.md "The LET statement assigns values to variables.") or
a [`INITIALIZE
TO NULL`](../08_language-basics/0698-initialize.md "The INITIALIZE instruction initializes program variables with NULL or default values.") instruction. However, this is only needed
for `STRING` variables declared as global or module
variables. The variables defined in functions will be automatically
destroyed when the program returns from the function.

Use of the [`base.StringBuffer`](../15_library-reference/3045-the-stringbuffer-class.md "The base.StringBuffer class is a built-in class designed to manipulate character strings.")
build-in class is recommended when a great deal of string manipulation and modification is
required. String data is not copied on the stack when an object of this class is passed to
a function, or when the string is modified with class methods. This can have a big impact
on performance when very large strings are processed.
