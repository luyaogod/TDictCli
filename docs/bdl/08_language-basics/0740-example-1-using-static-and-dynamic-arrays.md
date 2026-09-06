---
title: "Example 1: Using static and dynamic arrays"
source: "fgl-topics/c_fgl_Arrays_008.html"
breadcrumb: "Language basics > Arrays > Examples > Example 1: Using static and dynamic arrays"
type: "concept"
---

# Example 1: Using static and dynamic arrays

> This example illustrates the syntax difference of static and dynamic arrays.

```
MAIN
  DEFINE a1 DYNAMIC ARRAY OF INTEGER
  DEFINE a2 DYNAMIC ARRAY WITH DIMENSION 2 OF INTEGER
  DEFINE a3 ARRAY[10,20] OF RECORD
              id INTEGER,
              name VARCHAR(100),
              birth DATE
           END RECORD
  LET a1[5000] = 12456
  LET a2[5000,300] = 12456
  LET a3[5,1].id = a1[50]
  LET a3[5,1].name = 'Scott'
  LET a3[5,1].birth = TODAY
END MAIN
```
