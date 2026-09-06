---
title: "Single row SELECT"
source: "fgl-topics/c_fgl_odiagdmg_026.html"
breadcrumb: "SQL support > SQL database guides > Dameng® database server > Data manipulation > Single row SELECT"
type: "concept"
description: "Informix® With Informix, you must use the system table with a condition on the table id: SELECT USER FROM systables WHERE tabid=1 Dameng® To retrieve a single row with Dameng use the following syntax: ..."
---

# Single row SELECT

## Informix®

With Informix, you must use the system table with a
condition on the table id:

```
SELECT USER FROM systables WHERE tabid=1
```

## Dameng®

To retrieve a single row with Dameng use the following
syntax:

```
SELECT USER FROM dual
```

## Solution

Check the BDL sources for "`FROM systables WHERE tabid=1`" and use dynamic SQL to
resolve this problem.

Consider writing a `FUNCTION` which produces the `FROM` and
`WHERE` part, depending on the target database type.
