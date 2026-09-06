---
title: "OPTIONS (Compilation)"
source: "fgl-topics/c_fgl_programs_OPTIONS_compiler.html"
breadcrumb: "Advanced features > Configuration options > OPTIONS (Compilation)"
type: "concept"
---

# OPTIONS (Compilation)

> OPTIONS outside program blocks defines semantics of the language for the compiler.

## Syntax

```
OPTIONS
{ SHORT CIRCUIT
} [,...]
```

## Usage

The `OPTIONS` statement
used before any `MAIN`, `FUNCTION` or `REPORT` program
block defines language semantics options, that will
take effect for the current module only. Unlike runtime
options, compiler options cannot be changed during program
execution.

The statement to define compiler options must be placed before the `MAIN` block in
the main module, or before the first `FUNCTION` /
`REPORT` block in other modules.

The OPTIONS compiler directive allows for the control of features [Controlling semantics of AND / OR operators](0923-controlling-semantics-of-and-or-operators.md "The OPTIONS SHORT CIRCUIT defines the semantics of AND/OR operators.")

## Example

```
OPTIONS SHORT CIRCUIT
MAIN
    DISPLAY "Global Options example"
END MAIN
```

## Related links

**Related concepts**  

[OPTIONS (Runtime)](0924-options-runtime.md "The OPTIONS instruction inside program blocks controls program behavior at runtime.")

## Child topics

- [Controlling semantics of AND / OR operators](0923-controlling-semantics-of-and-or-operators.md): The OPTIONS SHORT CIRCUIT defines the semantics of AND/OR operators.
