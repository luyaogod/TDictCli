---
title: "Objects private to a program"
source: "fgl-topics/c_fgl_optimization_006.html"
breadcrumb: "Advanced features > Optimization > Runtime system basics > Objects private to a program"
type: "concept"
description: "Program objects such as global variables, module variables as well as resources used by the user interface and SQL connections and cursors, are private to a program. This implies that each of these ..."
---

# Objects private to a program

Program objects such as global variables, module variables as well as resources used by the user
interface and SQL connections and cursors, are private to a program.

This implies that each of these objects requires private memory to be allocated. If
memory is an issue, do not allocate unnecessary resources. For example, don't create
windows / load forms or declare / prepare cursors until these are really needed by the
program. When the resource is not longer needed, consider freeing them (`CLOSE
WINDOW`, `FREE cursor`).
