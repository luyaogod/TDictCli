---
title: "Using [package.]module prefix"
source: "fgl-topics/c_fgl_programs_IMPORT_FGL_prefix.html"
breadcrumb: "Advanced features > Importing modules > Importing FGL modules > Using [package.]module prefix"
type: "concept"
---

# Using [package.]module prefix

> Resolve symbol name conflicts with module prefix.

If a symbol is defined twice with the same name in two different modules, the symbol must be
qualified by the name of the module or the package name, when the module belongs to a package.

This feature overcomes the traditional 4GL limitation, requiring unique function names within a
program.

In the following example, both imported modules define the same "`init()`"
function, but this can be resolved, by adding the module name followed by a dot before the function
names:

```
IMPORT FGL orders 
IMPORT FGL customers 
MAIN
  CALL orders.init()
  CALL customers.init()
  ...
END MAIN
```

If a symbol is defined twice with the same name in the current and the imported module, an
unqualified symbol will reference the current module symbol.

The following example calls the "`init()`" function with and without a module
qualifier. The second call will reference the local function:

```
IMPORT FGL orders 
MAIN
  CALL orders.init()  -- orders module function 
  CALL init()  -- local function 
  ...
END MAIN
FUNCTION init()
  ...
END FUNCTION
```

When using packages, the prefix to a module symbol can be the full package path, or a sub-path of
the package path:

```
IMPORT FGL myshop.* -- imports several modules including "orders"
MAIN
  CALL myshop.orders.init() 
  ...
END MAIN
```

To make the code more readable, check for unqualified symbols with the `-W
unqualified-imports` warning, and consider using the `--qualify-imports`
option of fglcomp to refactor the sources.

The `--qualify-imports` option is part of the code refactoring tools provided by
Genero BDL.

For more details, see [Qualifying imported symbols](../13_programming-tools/2642-qualifying-imported-symbols.md "The fglcomp compiler can automatically qualify imported symbols.").

## Related links

**Related concepts**  

[Defining aliases for imported modules](0824-defining-aliases-for-imported-modules.md "Use aliases when package path or module names are too long.")
