---
title: "Reduce multiple text screens"
source: "fgl-topics/c_fgl_MigI4GL_019.html"
breadcrumb: "Upgrading > Migrating from IBM® Informix® 4GL to Genero BDL > User interface topics > Reduce multiple text screens"
type: "concept"
---

# Reduce multiple text screens

> Moving beyond the 80x25 dimensions of a display may require a review of your dumb-terminal-oriented programs.

Applications designed for dumb terminals (TUI mode) use various techniques to ensure that all
display fits in an 80x25 screen. This may mean iterating through a number of dialogs using different
forms, only displaying certain columns in an record list, using abbreviations for labels, and so
on.

With a graphical user interface (GUI mode), windows are wider, re-sizable, and can contain
different layout elements and widgets, displaying much more information than in a simple dumb
terminal. For example, TABLE containers display record lists and have the ability to scroll
horizontally so that you can show more than 78 characters of data.

It is recommended that you review dumb-terminal oriented programs to see how to take advantage of
the GUI possibilities. However, avoid ending up with
over-crowded screens that may be unreadable to the end user.

## Related links

**Related concepts**  

[TABLE container](../11_user-interface/1724-table-container.md "Defines a re-sizable table designed to display a list of records.")
