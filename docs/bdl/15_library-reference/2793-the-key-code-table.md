---
title: "The key code table"
source: "fgl-topics/r_fgl_BuiltInFunctions_key_code_table.html"
breadcrumb: "Library reference > Built-in functions > The key code table"
type: "reference"
description: "This table lists internal key codes. Avoid hard-coding these numbers in your sources; otherwise the source will not be compatible with future versions of the language. Always use the fgl_keyval() ..."
---

# The key code table

This table lists internal key codes. Avoid hard-coding these numbers in your
sources; otherwise the source will not be compatible with future versions of the
language.

Always use the [`fgl_keyval()`](2771-fgl-keyval.md "Returns the key code of a logical or physical key.")
function instead.

| Value | Key name | Description |
| --- | --- | --- |
| `1 to 26` | `CONTROL-letter` | Control key, where letter is the any letter from A to Z. For example, the key code corresponding to CONTROL-B is 2. |
| `13` | `RETURN` | The enter physical key. |
| `27` | `ESC` or `ESCAPE` | The escape physical key. In TUI mode this key is used to accept a dialog. |
| `others < 256` | `ASCII chars` | Other codes correspond to the ASCII characters set. |
| `2000` | `UP` | The up-arrow physical key. |
| `2001` | `DOWN` | The down-arrow physical key. |
| `2002` | `LEFT` | The left-arrow physical key. |
| `2003` | `RIGHT` | The right-arrow physical key. |
| `2005` | `NEXT` or `NEXTPAGE` | The next-page physical key. |
| `2006` | `PREVIOUS` or `PREVPAGE` | The previous-page physical key. |
| `2008` | `HELP` | The help logical key. |
| `2011` | `INTERRUPT` | The interrupt logical key. |
| `2014` | `INSERT` | The insert row logical key. |
| `2015` | `DELETE` | The delete row logical key. |
| `2016` | `ACCEPT` | The accept logical key. |
| `2017` | `BACKSPACE` | The backspace physical key. |
| `2020` | `HOME` | The home physical key. |
| `2021` | `END` | The end physical key. |
| `3000 to 3255` | `Fx` | Function key, where x is the number of the function key. The key code corresponding to a function key Fx is 3000+x-1, for example, 3011 corresponds to F12. |
