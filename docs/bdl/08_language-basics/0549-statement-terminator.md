---
title: "Statement terminator"
source: "fgl-topics/c_fgl_language_features_statement_terminator.html"
breadcrumb: "Language basics > Syntax features > Statement terminator"
type: "concept"
---

# Statement terminator

> The semicolon ( ; ) is optional statement terminator in Genero BDL.

Genero BDL is a language that does not require a statement terminator like in C/C++ or Java.

However, you can use the semicolon ( `;` ) as a statement terminator in some
cases.

For example:

```
MAIN
  DISPLAY "Hello, World"  DISPLAY "Hello, World"
  DISPLAY "Hello, World"; DISPLAY "Hello, World"
END MAIN
```

The semicolon statement terminator is sometimes mandatory, when nesting instruction blocks like in
the following
example:

```
INPUT BY NAME ...
    ON ACTION other_input
       INPUT BY NAME ... ;
END INPUT
```
