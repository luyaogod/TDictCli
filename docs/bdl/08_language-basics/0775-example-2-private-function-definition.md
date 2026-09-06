---
title: "Example 2: Private function definition"
source: "fgl-topics/c_fgl_Functions_example_2.html"
breadcrumb: "Language basics > Functions > Examples > Example 2: Private function definition"
type: "concept"
description: "This example shows how to define a private function. This function will not be visible to other modules because it is declared as PRIVATE for the module: PRIVATE FUNCTION checkLength( name VARCHAR(50) ..."
---

# Example 2: Private function definition

This example shows how to define a private function.

This function will not be visible to other modules because it is declared as
`PRIVATE` for the module:

```
PRIVATE FUNCTION checkLength( name VARCHAR(50) )
  DEFINE ok BOOLEAN
  IF length(name) == 0 THEN
     LET ok = FALSE
  ELSE
     LET ok = TRUE
  END IF
  RETURN ok
END FUNCTION
```
