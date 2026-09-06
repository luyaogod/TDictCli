---
title: "Example 1: Defining a record with explicit member types"
source: "fgl-topics/c_fgl_records_004.html"
breadcrumb: "Language basics > Records > Examples > Example 1: Defining a record with explicit member types"
type: "concept"
---

# Example 1: Defining a record with explicit member types

> This example shows a simple record definition with built-in types.

```
MAIN
  DEFINE rec RECORD
              id INTEGER,
              name VARCHAR(100),
              birth DATE
            END RECORD
  LET rec.id = 50
  LET rec.name = 'Scott'
  LET rec.birth = TODAY
  DISPLAY rec.*
END MAIN
```
