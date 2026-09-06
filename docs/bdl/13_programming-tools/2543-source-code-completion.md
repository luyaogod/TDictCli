---
title: "Source code completion"
source: "fgl-topics/c_fgl_CodeEditing_002.html"
breadcrumb: "Programming tools > Source code edition > Source code completion"
type: "concept"
description: "Purpose of source code completer The fglcomp compiler has a build-in feature to make source code completion. For example, if you start to type FUNC then press the TAB key, the code will be completed ..."
---

# Source code completion

## Purpose of source code completer

The fglcomp compiler has a build-in feature to make source code completion.
For example, if you start to type `FUNC` then press the TAB
key, the code will be completed with the `FUNCTION` keyword.

Source code completion can be enabled in the VIM editor and in VS Code: The editor must be
configured to call fglcomp with options with shortcuts and get code completion
proposals. For more details about VIM configuration, see [Configure VIM for Genero BDL](2544-configure-vim-for-genero-bdl.md)
and [Visual Studio Code extension](2545-visual-studio-code-extension.md).

## Complete column and variable names in static SQL

When editing static SQL statements, code completion proposes available symbols for program
variables, SQL tables and columns (from the [schema
file](../09_advanced-features/0791-database-schema.md "Defines database table structures with column type information to be reused in program variable definitions.")).

Column proposals after `customer` table name:

```
SCHEMA stores
MAIN
    DEFINE cust_id INTEGER
    DEFINE rec RECORD LIKE customer.*
    SELECT * INTO rec.* FROM customer
     WHERE customer.address1
                  | address1     c VARCHAR(20)   |
                  | address2     c VARCHAR(20)   |
                  | city         c VARCHAR(15)   |
                  | company      c VARCHAR(20)   |
                  | customer_num c INTEGER       |
                   ...
```

Variable proposals after `$`
sign:

```
SCHEMA stores
MAIN
    DEFINE cust_id INTEGER
    DEFINE rec RECORD LIKE customer.*
    SELECT * INTO rec.* FROM customer
     WHERE customer.customer_num = $<TAB>
                                  | cust_id   v INTEGER              |
                                  | int_flag  v INTEGER              |
                                  | quit_flag v INTEGER              |
                                  | rec       v like:stores.customer | 
                                    ...
```

## Related links

**Related concepts**  

[Configure VIM for Genero BDL](2544-configure-vim-for-genero-bdl.md "Configure VIM for Genero BDL")
