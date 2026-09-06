---
title: "SELECT * (asterisk)"
source: "fgl-topics/c_fgl_odiagora_048.html"
breadcrumb: "SQL support > SQL database guides > Oracle® Database > Data manipulation > SELECT * (asterisk)"
type: "concept"
description: "Informix® Informix allows you to use the star character in the select list along with other expressions: SELECT col1, * FROM tab1 ... ORACLE Oracle® does not support the asterisk notation after ..."
---

# SELECT * (asterisk)

## Informix®

Informix allows you to use the star character in the
select list along with other expressions:

```
SELECT col1, * FROM tab1 ...
```

## ORACLE

Oracle® does not support the asterisk
notation after another expression in the `SELECT` list.

Use the table name as a prefix to the star:

```
SELECT col1, tab1.* FROM tab1 ...
```

## Solution

Always use the table name before
the star.
