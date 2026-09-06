---
title: "Elements shared by multiple programs"
source: "fgl-topics/c_fgl_optimization_004.html"
breadcrumb: "Advanced features > Optimization > Runtime system basics > Elements shared by multiple programs"
type: "concept"
description: "The ( .42m ) p-code module instructions and other elements such as constants are shared among several programs running on the same machine. Localized string resource files (.42s) are also shared among ..."
---

# Elements shared by multiple programs

The (.42m) p-code module instructions and other elements such as constants
are shared among several [programs](0785-program-structure.md "Explains the organization of a BDL program.") running on the same
machine.

[Localized string resource files](0901-localized-strings.md "Localized strings provide a means of writing applications in which the text of strings can be customized on site.") (.42s)
are also shared among all fglrun processes running on a computer.

These files are loaded with the system memory mapping facility, which allows multiple
processes to access the same unique memory area.
