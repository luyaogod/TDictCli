---
title: "Saving memory by using dynamic arrays"
source: "fgl-topics/c_fgl_optimization_015.html"
breadcrumb: "Advanced features > Optimization > Optimize your programs > Saving memory by using dynamic arrays"
type: "concept"
description: "The language supports both static arrays and dynamic arrays . For compatibility reasons, static arrays must be allocated in their entirety. This can result in huge memory usage when big structures are ..."
---

# Saving memory by using dynamic arrays

The language supports both [static arrays and dynamic
arrays](../08_language-basics/0729-arrays.md "Arrays (static or dynamic) allow you to handle an ordered collection of elements."). For compatibility reasons, static arrays must be allocated in their entirety. This
can result in huge memory usage when big structures are declared, such
as:

```
DEFINE arr_cust ARRAY[5000] OF RECORD
      id INTEGER,
      shipcode CHAR(5),
      name CHAR(50),
      address CHAR(200)
END RECORD
```

If possible, replace such static arrays with dynamic arrays, and consider using
`VARCHAR` for large character string fields:

```
DEFINE arr_cust DYNAMIC ARRAY OF RECORD
      id INTEGER,
      shipcode CHAR(5),
      name VARCHAR(50),
      address VARCHAR(200)
END RECORD
```

However, be aware that dynamic arrays have a slightly different behavior than static
arrays.
