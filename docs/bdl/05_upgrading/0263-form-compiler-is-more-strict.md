---
title: "Form compiler is more strict"
source: "fgl-topics/c_fgl_Migrate_to_230_strict_fglform.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.30 upgrade guide > Form compiler is more strict"
type: "concept"
---

# Form compiler is more strict

> The .per grammar parser has been reviewed to deny invalid code.

In version 2.30, the internals of fglform have been reviewed to simplify the
extension of the form syntax with new item types and attributes. This code review has removed
some inconsistencies in the grammar parser; as a result, the form compiler is more strict
regarding invalid syntaxes. Thus, you may experience compilation errors with forms that compiled
with prior versions. Simply fix the invalid syntax in your forms and recompile.

## Related links

**Related concepts**  

[fglform](../13_programming-tools/2514-fglform.md "The fglform tool compiles form specification files into XML formatted files used by programs.")
