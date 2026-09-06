---
title: "Example 3: PUBLIC module variables"
source: "fgl-topics/c_fgl_variables_example_3.html"
breadcrumb: "Language basics > Variables > Examples > Example 3: PUBLIC module variables"
type: "concept"
description: "This example shows how to declare public and private module variables. The module \" mydebug.4gl \": PUBLIC DEFINE level INTEGER, logfile STRING PRIVATE DEFINE count INTEGER FUNCTION message(m) DEFINE m ..."
---

# Example 3: PUBLIC module variables

This example shows how to declare public and private module variables.

The module
"mydebug.4gl":

```
PUBLIC DEFINE level INTEGER,
              logfile STRING
PRIVATE DEFINE count INTEGER

FUNCTION message(m)
  DEFINE m STRING
  IF level THEN
     -- Write message to debug_logfile
     DISPLAY m
  END IF
  LET count = count + 1
END FUNCTION
```

The main
module:

```
IMPORT FGL mydebug

MAIN
  LET mydebug.level = 4
  LET mydebug.logfile = "myfile.log"
  CALL mydebug.message("Some debug info...")
END MAIN
```
