---
title: "ORACLE and INTERVAL columns"
source: "fgl-topics/c_fgl_Migrate_to_230_oracle_interval.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.30 upgrade guide > ORACLE and INTERVAL columns"
type: "concept"
---

# ORACLE and INTERVAL columns

> INTERVAL storage bug fix needs a review of existing databases in production.

Before 2.30.00 (build 1566), negative (and only negative) INTERVAL values were inserted
incorrectly. This a critical bug.

For example, it was not possible to compare an INTERVAL value inserted by a program with an
INTERVAL
literal:

```
SELECT ... FROM table
  WHERE interval_col = INTERVAL '-55555-11' YEAR(9) TO MONTH
```

The problem concerns database columns with the following interval
types:

```
INTERVAL YEAR(p) TO MONTH
INTERVAL DAY(p) TO FRACTION(n)
```

(Other INTERVAL types are stored in a CHAR(50))

A simple INTERVAL to CHAR to INTERVAL conversion will fix the
values:

```
UPDATE table SET interval_col = TO_CHAR(interval_col)
```

## Related links

**Related concepts**  

[Interval expressions](../08_language-basics/0600-interval-expressions.md "This section covers interval expression evaluation rules.")

[Interval literals](../08_language-basics/0591-interval-literals.md "Interval literals define an interval value in an expression.")
