---
title: "Compilers command line options"
source: "fgl-topics/c_fgl_preprocessor_003.html"
breadcrumb: "Programming tools > Source preprocessor > Compilers command line options"
type: "concept"
---

# Compilers command line options

> Preprocessor options can be used with fglcomp and fglform compilers.

## File inclusion path

The `-I` option defines a single path used to search files included by
the `&include` directives:

```
 -I path
```

To specify multiple include directories, repeat the `-I path`
option:

```
fglcomp -I /usr/app1/src/common -I /usr/app1/src/stock ...
```

## Macro definition

The `-D` option defines a macro with a default value of 1, so that it
can be used in conditional directives like `&ifdef`:

```
 -D identifier
```

The `-D` option can also define a macro with a value:

```
 -D identifier=value
```

If you want to define a macro with a string value containing spaces, you need to provide the
identifier and value as a single command line parameter. For example on UNIX®: `-D "VSTR=\"Version 1.32\""`

The `-U` option undefines a macro. The macro will not be defined, even
if it is defined with the `-D` option later in the command line, or when it is
defined in the code with a `&define`
directive:

```
 -U identifier
```

However, predefined macros such as `__LINE__` cannot be undefined with
the `-U` option.

## Preprocessing only

```
 -E
```

By using the `-E` option, only the preprocessing phase is done by the compilers.
Result is dumped in standard output.

## Preprocessing style option

```
 -p { nopp | noln | fglpp | auto }
```

When using option `-p nopp`, it disables the preprocessor phase.

By using option `-p noln` with the `-E` preprocessing-only option,
you can remove line number information and unnecessary empty lines.

By default, the preprocessor expects an ampersand '`&`' as preprocessor symbol
for macros. The option `-p fglpp` enables the old syntax, using the hash
'`#`' as preprocessor symbol. The hash '`#` ' syntax is not compatible
with [single-line comments](../08_language-basics/0550-source-comments.md "The --, # and { } characters can be used to add source comments.").

When using `-p auto`, the compiler detects automatically the
'`&`' ampersand or '`#`' hash usage for preprocessor commands in
the source code. The first preprocessor instruction found in the source will define the
preprocessing style.

## Examples

```
fglcomp -D DEBUG -I /usr/sources/headers program.4gl
fglcomp -p fglpp program.4gl
fglcomp -p auto module.4gl
```

## Related links

**Related concepts**  

[fglcomp](2516-fglcomp.md "The fglcomp tool compiles .4gl source files into .42m p-code modules, and does various other tasks.")

[fglform](2514-fglform.md "The fglform tool compiles form specification files into XML formatted files used by programs.")
