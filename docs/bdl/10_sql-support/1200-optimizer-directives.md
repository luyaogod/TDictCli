---
title: "Optimizer directives"
source: "fgl-topics/c_fgl_odiagifx_013.html"
breadcrumb: "SQL support > SQL database guides > IBM® Informix® > Partially supported IBM® Informix® SQL features > Optimizer directives"
type: "concept"
description: "IBM® Informix® SQL allows you to specify query optimization directives to force the query optimizer to use a different path than the implicit plan. With IBM Informix, optimizer directives are ..."
---

# Optimizer directives

IBM® Informix®
SQL allows you to specify query optimization directives to force the query optimizer to use a
different path than the implicit plan. With IBM Informix, optimizer directives are specified with the
following SQL comment markers followed by a plus sign:

```
/*+ optimizer-directives */
{+  optimizer-directives }
--+ optimizer-directives
```

Genero BDL partially supports optimizer directives:

- The static SQL syntax does not allow the C-style optimizer syntax.
- The curly-brace and hyphen-hyphen optimizer directive syntaxes cannot be used in static SQL
  statements, because these correspond to the [4GL language comments](../08_language-basics/0544-syntax-features.md "Genero BDL is an English-like programming language, easy to write and read.").
- However, you can execute queries with optimization directives with [Dynamic SQL](1140-dynamic-sql-management.md "Explains how to execute and manage SQL statements at runtime.").

> **Tip:**
>
> Optimization directives are not portable. If you plan to use different types
> of database servers, it is recommended that you avoid the usage of query plan hints.

## Related links

**Related concepts**  

[Static SQL statements](1115-static-sql-statements.md "Describes static SQL statements supported in the language.")
