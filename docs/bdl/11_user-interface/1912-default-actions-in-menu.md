---
title: "Default actions in MENU"
source: "fgl-topics/c_fgl_menus_009.html"
breadcrumb: "User interface > Dialog instructions > Ring menus (MENU) > Using ring menus > Default actions in MENU"
type: "concept"
description: "When an MENU instruction executes, the runtime system creates a set of default actions. Table 1. Default actions created for the MENU instruction Default action Control Block execution order close ..."
---

# Default actions in MENU

When an `MENU` instruction executes, the runtime system creates
a set of default actions.

| Default action | Control Block execution order |
| --- | --- |
| `close` | Created to execute `COMMAND KEY(INTERRUPT)` if used (can be overwritten with `ON ACTION close`)Default action view is hidden. See [Implementing the close action](2289-implementing-the-close-action.md "The close action is a predefined action dedicated to close graphical windows (for example, with the X cross button)."). |
| `help` | Shows the help topic defined by the `HELP` clause.Default action view is hidden. |

Window close events can be trapped with `COMMAND KEY(INTERRUPT)`
clause.

## Related links

**Related concepts**  

[The model-view-controller paradigm](2219-the-model-view-controller-paradigm.md "The dynamic user interface architecture is based on the Model-View-Controller (MVC) paradigm.")
