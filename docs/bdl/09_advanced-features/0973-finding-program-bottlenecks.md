---
title: "Finding program bottlenecks"
source: "fgl-topics/c_fgl_optimization_009.html"
breadcrumb: "Advanced features > Optimization > Optimize your programs > Finding program bottlenecks"
type: "concept"
description: "The best way to find out why a program is slow (and also, to optimize an already fast-running program), is to use the profiler . This tool is included in the runtime system, and generates a report ..."
---

# Finding program bottlenecks

The best way to find out why a program is slow (and also, to optimize an already fast-running
program), is to use the [profiler](../13_programming-tools/2622-program-profiler.md "Find out what function is causing the bottleneck in your program.").

This tool is included in the runtime system, and generates a report that shows what
function in your program is the most time-consuming.

Additionally, you might want to identify part of code of your programs that are never executed,
or executed very often, by using the [coverage tool](../13_programming-tools/2628-source-code-coverage.md "Collect information about used source lines").
