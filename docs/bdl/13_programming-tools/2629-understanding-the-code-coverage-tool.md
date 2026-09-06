---
title: "Understanding the code coverage tool"
source: "fgl-topics/c_fgl_coverage_basics.html"
breadcrumb: "Programming tools > Source code coverage > Understanding the code coverage tool"
type: "concept"
---

# Understanding the code coverage tool

> This is an introduction to the code coverage tool.

The code coverage tool built in the [fglrun runtime system](2513-fglrun.md "The fglrun tool is the runtime system program that executes p-code programs.") allows you to collect information about source code
execution.

With the coverage tool, you can:

- Identify dead code that is never executed, or make sure that important code is really
  executed.
- Find what lines of code are executed many times, to optimize your programs.

For each line of code, the tool reports the number of times it has been executed, if it has never
been executed, or if it is not reachable at all.

> **Important:**
>
> The output format of code coverage tool is for debug purpose only and can
> change in future product releases.
