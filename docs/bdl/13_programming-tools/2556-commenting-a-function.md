---
title: "Commenting a function"
source: "fgl-topics/c_fgl_AutoDoc_007.html"
breadcrumb: "Programming tools > Source documentation > Adding comments to sources > Commenting a function"
type: "concept"
description: "To document a function, add some lines starting with #+ , before the FUNCTION declaration. The comment body is composed of paragraphs separated by blank lines. The first paragraph of the comment is a ..."
---

# Commenting a function

To document a function, add some lines starting with `#+`, before the
`FUNCTION` declaration.

The comment body is composed of paragraphs separated by blank lines. The first paragraph of the
comment is a short description of the function. This description will be placed in the function
summary table. The next paragraph is long text describing the function in detail. Other paragraphs
must start with a tag to identify the type of paragraph; a tag starts with the @ "at" sign.

In order to have fglcomp --build-doc work well,
there must be an initial #+ comment line for the module description, before any other constant,
variable, type or function documentation directives.

| Tag | Description |
| --- | --- |
| `@code` | Indicates that the next lines show a code example using the function. |
| `@param name description` | Defines a function parameter identified by *name*, explained by a *description*.*name* must match the parameter name in the function declaration. |
| `@returnType data-type [,...]` | Defines the data type of the value returned by the function.If the function returns several values, write a comma-separated list of types.The `@returnType` tag does not need to be used when the `FUNCTION` is defined with a `RETURNS` clause. |
| `@return description` | Describes the values returned by the function.Several @return comment lines can be written. |

## Example using a FUNCTION definition with signature

```
#+ Compute the amount of the orders for a given customer
#+
#+ This function calculates the total amount of all orders for the
#+ customer identified by the cust_id number passed as parameter.
#+
#+ @code
#+ DEFINE total DECIMAL(10,2)
#+ LET total = ordersTotal(r_customer.cust_id)
#+
#+ @param cid Customer identifier
#+
#+ @return The total amount as DECIMAL(10,2)
#+
FUNCTION ordersTotal(cid INTEGER) RETURNS DECIMAL(10,2)
   DEFINE ordtot DECIMAL(10,2)
   SELECT SUM(ord_amount) INTO ordtot
        FROM orders WHERE orders.cust_id = cid
   RETURN ordtot
END FUNCTION
```

## Example using a FUNCTION definition without signature (needs `@returnType`)

```
#+ Compute the amount of the orders for a given customer
#+
#+ This function calculates the total amount of all orders for the
#+ customer identified by the cust_id number passed as parameter.
#+
#+ @code
#+ DEFINE total DECIMAL(10,2)
#+ LET total = ordersTotal(r_customer.cust_id)
#+
#+ @param cid Customer identifier
#+
#+ @returnType DECIMAL(10,2)
#+ @return The total amount as DECIMAL(10,2)
#+
FUNCTION ordersTotal(cid)
   DEFINE cid INTEGER
   DEFINE ordtot DECIMAL(10,2)
   SELECT SUM(ord_amount) INTO ordtot
        FROM orders WHERE orders.cust_id = cid
   RETURN ordtot
END FUNCTION
```
