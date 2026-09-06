---
title: "Preprocessor changes"
source: "fgl-topics/c_fgl_Migrate_to_300_preprocessor.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 3.00 upgrade guide > Preprocessor changes"
type: "concept"
---

# Preprocessor changes

> Several bugs have been fixed in the preprocessor, that can now result in a compilation error.

## String token expansion

Before version 3.00, the following preprocessor syntax was used to expand a string macro
parameter:

```
&define T(x) DISPLAY "head_"#x"_tail"
-- macro usage:
T(body)
```

This produced the following result (after preprocessing):

```
"head_""body""_tail"
```

It was accepted by the compiler, because it was interpreted as a single string literal.

The new preprocessor now produces (as expected):

```
"head_" "body" "_tail"
```

However, this will now result in a compiler error, because this is not a valid string
literal.

To solve an issue such as this and get the same result string as before version 3.00, use the
`||` concatenation operator in the preprocessor macro and add
(escaped) double quotes before and after the
`#ident` placeholder:

```
&define T(x) DISPLAY "head_""" || #x || """_tail"
```

or, by using single quotes as border strings delimiters:

```
&define T(x) DISPLAY 'head_"' || #x || '"_tail'
```

## Identifier concatenation

Before version 3.00, the following type of macro:

```
&define FOO() foo
-- macro usage:
FOO()bar
```

was producing a single identifier token (accepted by the compiler):

```
foobar
```

But it will now produce two distinct identifier tokens (as expected):

```
foo bar
```

And this will result in a compilation error.

## Backslash in macro parameters

Before version 3.00.00 it was possible to use the backslash to escape a comma in preprocessor
macro parameters. This syntax is no longer allowed by the preprocessor, it is not a
valid usage. To solve this issue, replace parameters by real string literals in the
macro:

```
-- bad coding
&define FOO(p1) DISPLAY #p1
FOO(hello world)       -- expands to: DISPLAY "hello world"
FOO(hello \, world)    -- error 

-- good coding
&define FOO(p1) DISPLAY p1
FOO("hello world")     -- expands to: DISPLAY "hello world"
FOO("hello , world")   -- expands to: DISPLAY "hello , world"
```

## The ## paste operator

Before version 3.00.00, the `##` paste operator was used to construct code with
two elements that did not result in a valid token, for
example:

```
&define FOO(name) rec_ ## [ x ]
FOO(x)
```

produced:

```
rec_[ x ]
```

This kind of preprocessor macro is no longer allowed in version 3.00.00 and will result in a
compiler
error:

```
x.4gl:2:1:2:1:error:(-8042) The operator '##' formed 'rec_[',
 an invalid preprocessing token.
```

The `##` paste operator must be used to join two identifiers, to
create a new identifier:

```
&define REC_PREFIX(name) rec_ ## name
LET REC_PREFIX(customer) = NULL
```

This will produce:

```
LET rec_customer = NULL
```

## File inclusion search path

Before version 3.00, the compiler option `-I` was allowing you to specify a list
of paths separated by the OS path separator.

This behavior was unexpected and wrong: The `-I` option must define a single path
(or directory).

Command lines using the followin (UNIX)
notation:

```
fglcomp -I path1:path2 ...
```

must
be reviewed by specifying each directory in a separate `-I`
option:

```
fglcomp -I path1 -I path2 ...
```

## Related links

**Related concepts**  

[Source preprocessor](../13_programming-tools/2562-source-preprocessor.md "A typical preprocessor like in the C language.")
