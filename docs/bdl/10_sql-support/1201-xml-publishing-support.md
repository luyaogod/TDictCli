---
title: "XML publishing support"
source: "fgl-topics/c_fgl_odiagifx_014.html"
breadcrumb: "SQL support > SQL database guides > IBM® Informix® > Partially supported IBM® Informix® SQL features > XML publishing support"
type: "concept"
description: "IBM® Informix® IDS 11.10 introduced a set of XML built-in functions when the idsxmlvp virtual processor is turned on. Built-in XML functions are of two types: those returning LVARCHAR values, and ..."
---

# XML publishing support

IBM®
Informix® IDS 11.10 introduced a set of XML built-in
functions when the **idsxmlvp** virtual processor is turned on. Built-in XML functions are of two
types: those returning LVARCHAR values, and those returning CLOB values. For example,
`genxml()` returns an LVARCHAR(32739), while `genxmlclob()` returns a
CLOB. XML data is typically stored in LVARCHAR or CLOB columns.

Genero BDL partially supports XML functions:

- Because Genero BDL does not support [BLOB/CLOB](1206-clob-and-blob-data-types.md) types,
  functions returning CLOB values cannot be used. You can however use
  the XML functions returning LVARCHAR values, and fetch the result
  into a [VARCHAR](../08_language-basics/0553-primitive-data-types.md "Selecting the correct data type assists you in the input, storage, and display of your data.") variable
  of the appropriate size.
- Some of the XML functions such as `genxml()` take ROW() values as parameters.
  Because literal unnamed ROW() expressions are like regular function calls, you can use XML functions
  in static SQL statements.

Example:

```
FUNCTION get_cust_data(id)
   DEFINE id INT, v VARCHAR(5000)
   SELECT genxml(ROW(cust_name, cust_address), "custdata") INTO v
      FROM customers WHERE cust_id = id
   RETURN v
END FUNCTION
```

## Related links

**Related concepts**  

[Static SQL statements](1115-static-sql-statements.md "Describes static SQL statements supported in the language.")
