---
title: "Commenting a report"
source: "fgl-topics/c_fgl_AutoDoc_report.html"
breadcrumb: "Programming tools > Source documentation > Adding comments to sources > Commenting a report"
type: "concept"
description: "To document a report, add some lines starting with #+ , before the REPORT declaration. The comment body is composed of paragraphs separated by blank lines. The first paragraph of the comment is a ..."
---

# Commenting a report

To document a report, add some lines starting with `#+`, before the
`REPORT` declaration.

The comment body is composed of paragraphs separated by blank lines. The first paragraph of the
comment is a short description of the report. This description will be placed in the function
summary table. The next paragraph is long text describing the report in detail. Other paragraphs
must start with a tag to identify the type of paragraph; a tag starts with the @ "at" sign.

In order to have fglcomp --build-doc work well,
there must be an initial #+ comment line for the module description, before any other constant,
variable, type or function documentation directives.

| Tag | Description |
| --- | --- |
| `@code` | Indicates that the next lines show a code example using the report. |
| `@param name description` | Defines a report parameter identified by *name*, explained by a *description*.*name* must match the parameter name in the report declaration. |
