---
title: "Commenting a module"
source: "fgl-topics/c_fgl_AutoDoc_008.html"
breadcrumb: "Programming tools > Source documentation > Adding comments to sources > Commenting a module"
type: "concept"
description: "To add a summary to a module, add #+ lines at the beginning of the source, before module element declarations such as variables, types, constants and functions. The #+ lines of the module source ..."
---

# Commenting a module

To add a summary to a module, add `#+` lines at the beginning of the source,
before module element declarations such as variables, types, constants and functions.

The `#+` lines of the module source comment must appear at the top of the
source.

In order to have fglcomp --build-doc work well,
there must be an initial #+ comment line for the module description, before any other constant,
variable, type or function documentation directives.

## Example

```
#+ This module implements customer information handling
#+
#+ This code uses the 'customer' and 'custdetail' database tables.
#+ Customer input, query and list handling functions are defined here.
#+
DEFINE r_cust RECORD
   cust_id INTEGER,
   cust_name VARCHAR(50),
   cust_address VARCHAR(200)
END RECORD
```
