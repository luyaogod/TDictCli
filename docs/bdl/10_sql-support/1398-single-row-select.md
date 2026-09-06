---
title: "Single row SELECT"
source: "fgl-topics/c_fgl_odiagora_031.html"
breadcrumb: "SQL support > SQL database guides > Oracle® Database > Data manipulation > Single row SELECT"
type: "concept"
description: "Informix® With Informix, you must use the system table with a condition on the table id: SELECT USER FROM systables WHERE tabid=1 ORACLE Oracle provides the DUAL table to produce a single row result ..."
---

# Single row SELECT

## Informix®

With Informix, you must use the system table with a
condition on the table id:

```
SELECT USER FROM systables WHERE tabid=1
```

## ORACLE

Oracle provides the `DUAL` table to produce a single row result set:

```
SELECT USER FROM DUAL
```

Starting with Oracle 23c, it is possible to omit the `FROM`
clause:

```
SELECT USER
```

## Solution

Check the BDL sources for "`FROM systables WHERE tabid=1`" and use dynamic SQL to
resolve this problem.

Consider writing a `FUNCTION` which produces the `FROM` and
`WHERE` part, depending on the target database type.
