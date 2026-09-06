---
title: "fgldialog: Common dialog functions"
source: "fgl-topics/r_fgl_utility_functions_fgldialog.html"
breadcrumb: "Library reference > Utility modules > fgldialog: Common dialog functions"
type: "reference"
description: "Table 1. Common dialog functions (fgldialog.4gl) Function Description FUNCTION fgl_winbutton ( title STRING, text STRING, default STRING, buttons STRING, icon STRING, danger SMALLINT ) RETURNS STRING ..."
---

# fgldialog: Common dialog functions

| Function | Description |
| --- | --- |
| FUNCTION fgl_winbutton( title STRING, text STRING, default STRING, buttons STRING, icon STRING, danger SMALLINT ) RETURNS STRING | Displays an interactive message box containing multiple choices, in a pop-up window. |
| FUNCTION fgl_winmessage( title STRING, text STRING, icon STRING ) | Displays an interactive message box containing text and OK button. |
| FUNCTION fgl_winprompt( x INTEGER, y INTEGER, text STRING, default STRING, length INTEGER, type INTEGER ) RETURNS STRING | Displays a dialog box containing a field that accepts a value. |
| FUNCTION fgl_winquestion( title STRING, text STRING, default STRING, buttons STRING, icon STRING, danger SMALLINT ) RETURNS STRING | Displays an interactive message box with configurable Ok / Yes / No / Cancel / Ignore / Abort / Retry buttons. |
| FUNCTION fgl_winwait( text STRING ) | Displays an interactive message box and waits for user validation. |
