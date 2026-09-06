---
title: "Setting default screen modes for sub-programs"
source: "fgl-topics/c_fgl_programs_014.html"
breadcrumb: "Advanced features > Configuration options > OPTIONS (Runtime) > Setting default screen modes for sub-programs"
type: "concept"
---

# Setting default screen modes for sub-programs

> The OPTIONS RUN IN instruction defines the TTY mode to run sub-programs.

## Syntax

```
OPTIONS RUN IN {FORM|LINE} MODE
```

## Usage

This instruction defines the default mode to run child applications with the `RUN`
command.

Details about the *line mode* and *form mode* are available in the reference topic of
the [RUN](0830-run.md "The RUN instruction executes the command passed as argument.") instruction.
