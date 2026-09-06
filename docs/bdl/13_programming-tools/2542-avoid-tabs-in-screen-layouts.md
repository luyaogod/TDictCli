---
title: "Avoid Tabs in screen layouts"
source: "fgl-topics/c_fgl_CodeEditing_004.html"
breadcrumb: "Programming tools > Source code edition > Avoid Tabs in screen layouts"
type: "concept"
description: "When editing .per form files , avoid using Tab characters in sources, especially in the LAYOUT or SCREEN sections of forms. Each kind of text/source editors can expand Tab characters differently, ..."
---

# Avoid Tabs in screen layouts

When editing [.per form files](../11_user-interface/1662-form-specification-files.md "Form specification files are the source files defining the layout and content of application forms."), avoid using Tab
characters in sources, especially in the `LAYOUT` or `SCREEN` sections
of forms.

Each kind of text/source editors can expand Tab characters differently, depending on the
configuration settings. As a result, if two programmers are using different Tab expansion settings,
the form layout will display in different ways.

When used in a grid area of a [`LAYOUT` section](../11_user-interface/1715-layout-section.md "The LAYOUT section defines the graphical alignment of the form by using a tree of layout containers."), a Tab character will be interpreted as 8 blanks by [fglform](2514-fglform.md "The fglform tool compiles form specification files into XML formatted files used by programs."). It is legal to use Tab characters in the rest of the
.per file or .4gl sources (for example, to indent the
code).
