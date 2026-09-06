---
title: "SHOW/HIDE OPTION instruction"
source: "fgl-topics/c_fgl_menus_SHOW_HIDE_OPTION.html"
breadcrumb: "User interface > Dialog instructions > Ring menus (MENU) > Using ring menus > MENU control instructions > SHOW/HIDE OPTION instruction"
type: "concept"
description: "Syntax { HIDE | SHOW } OPTION { ALL | option-name [ , ...] } Usage The SHOW OPTION instruction will show/enable action views corresponding to the listed menu options. The default action views are made ..."
---

# SHOW/HIDE OPTION instruction

## Syntax

```
{ HIDE | SHOW } OPTION
    { ALL
    | option-name [,...]
    }
```

## Usage

The `SHOW OPTION` instruction will show/enable action views corresponding to the
listed menu options. The [default action
views](2258-default-action-views.md "A default action view is created to render an action handler, when no explicit action view exists for it.") are made visible and the explicit action views (`BUTTON` in form) are enabled. The `HIDE
OPTION` instruction will hide default action views and disable explicit action views.

Use the `ALL` keyword reference all menu options. In a menu that contains many
options, you typically do a `HIDE OPTIONS ALL` followed by `HIDE
OPTION` to show a subset of the menu options.

```
MENU "Customers"
   BEFORE MENU
      HIDE OPTION ALL
      SHOW OPTION "Add", "Exit"
   ...
```

The `SHOW OPTION` and `HIDE OPTION` instructions are provided for
backward compatibility. To hide and show default action views, use the `DIALOG.setActionHidden()` method
instead. In GUI applications, it is recommended that you disable actions, instead of hiding them
from the end user.
