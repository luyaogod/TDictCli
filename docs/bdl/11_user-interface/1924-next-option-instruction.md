---
title: "NEXT OPTION instruction"
source: "fgl-topics/c_fgl_menus_NEXT_OPTION.html"
breadcrumb: "User interface > Dialog instructions > Ring menus (MENU) > Using ring menus > MENU control instructions > NEXT OPTION instruction"
type: "concept"
description: "Syntax NEXT OPTION option-name Usage The NEXT OPTION instruction selects a menu option to make current. MENU \"Customers\" BEFORE MENU NEXT OPTION \"Modify\" ... The specified menu option will be ..."
---

# NEXT OPTION instruction

## Syntax

```
NEXT OPTION option-name
```

## Usage

The `NEXT OPTION` instruction selects a menu option to make current.

```
MENU "Customers"
   BEFORE MENU
      NEXT OPTION "Modify"
   ...
```

The specified menu option will be highlighted and the user can simply press the RETURN key to
choose that current option.
