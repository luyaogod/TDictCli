---
title: "Report engine configuration"
source: "fgl-topics/c_fgl_reports_005.html"
breadcrumb: "Reports > Report engine configuration"
type: "concept"
---

# Report engine configuration

> Report engine behavior can be controlled with FGLPROFILE settings.

By default, aggregate instructions such as `SUM()` return a `NULL`
value if all input record values are `NULL`.

You can force the report engine to return a zero decimal value with the following FGLPROFILE
setting:

```
Report.aggregateZero = {true|false}
```

When this entry is set to true, the `SUM()`, `AVG()`,
`MIN()`, or `MAX()` aggregate functions return zero when all values
are `NULL`.

Default value of the configuration parameter is false (this means aggregate functions evaluate to
null if all items are null).

When using `GROUP` aggregates with this entry is set to true, the aggregate
instruction will still return `NULL` in the first `AFTER GROUP OF` output of the report.
Zero values will be returned starting from second group output. This behavior is expected, for
backward compatibility with older versions.

It is not recommended to use the `Report.aggregateZero` entry if you don't need
that specific behavior.

## Related links

**Related concepts**  

[The FGLPROFILE file(s)](../07_configuration/0483-the-fglprofile-file-s.md "FGLPROFILE environment variable defines Genero BDL configuration files")

[NULL](../08_language-basics/0572-null.md "The NULL constant defines a non-value.")
