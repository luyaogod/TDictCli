---
title: "Review application ergonomics"
source: "fgl-topics/c_fgl_MigI4GL_020.html"
breadcrumb: "Upgrading > Migrating from IBM® Informix® 4GL to Genero BDL > User interface topics > Review application ergonomics"
type: "concept"
---

# Review application ergonomics

> Genero BDL no longer limits programs to executing a single interactive instruction, and provides additional GUI concepts such as drag-and-drop and tree views.

With IBM® Informix® 4GL, programs can only execute
a single `MENU`, `INPUT`, `CONSTRUCT`,
`DISPLAY ARRAY` or `INPUT ARRAY` instruction
at a time. This may be sufficient for dumb-terminal applications,
but is not adapted for a graphical user interface.

Genero Business Development Language (BDL) introduces the concept
of multi-dialog, where multiple interactive instructions control
several form areas at the same time. Typical GUI concepts such
as Drag and Drop and Tree Views are supported as well. You may wish
to review your code to take advantage of these features.
