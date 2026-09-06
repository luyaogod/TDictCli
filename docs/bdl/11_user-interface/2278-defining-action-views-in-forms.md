---
title: "Defining action views in forms"
source: "fgl-topics/c_fgl_prog_dialogs_action_view.html"
breadcrumb: "User interface > User interface programming > Dialog actions > Defining action views in forms"
type: "concept"
---

# Defining action views in forms

> How to define action views that will fire action events.

Actions views are form items that can be activated to fire an action event. The action event
triggers user code in an `ON ACTION`
block.

We distinguish action views defined explicitly in form files from [default action views](2258-default-action-views.md "A default action view is created to render an action handler, when no explicit action view exists for it."). A default action
view will automatically appear when an action handler is implemented in the current dialog (if no
explicit action view with the same name exists in the form). Default action view creation can be
controlled with the [`DEFAULTVIEW`](2273-defaultview-action-attribute.md "The DEFAULTVIEW attribute defines if a default view (a button) must be displayed for a given action.") action attribute.

To fire user code, action views are bound to [action handlers](2279-implementing-dialog-action-handlers.md "How to execute user code in ON ACTION blocks when an action is fired.") by name.

Action view decoration attributes (`IMAGE` for icons, `TEXT` for
label, `COMMENT` for hint) can be centralized in [action defaults](2259-configuring-actions.md "Action attributes related to decoration, keyboard shortcuts, and behavior can be defined with action attributes.").

Action views can be items of form elements dedicated to action execution, such as [`TOOLBAR`](1713-toolbar-section.md "The TOOLBAR section defines a toolbar with buttons that are bound to actions.") items (toolbar
buttons) or [`TOPMENU`](1712-topmenu-section.md "The TOPMENU section defines a pull-down menu with options that are bound to actions.")
options:

```
TOOLBAR
   ITEM accept
   ITEM cancel
   ...
END
```

Action views can be typical [`BUTTON`](1684-button-item-type.md "Defines a push-button that can trigger an action.") items defined in the form `LAYOUT`:

```
LAYOUT
GRID
{
   [b1     ]
   ...
}
...
ATTRIBUTES
BUTTON b1 : print, IMAGE="printer";
...
```

Action views can be sub-elements of other elements, as when defining a [`BUTTONEDIT`](1685-buttonedit-item-type.md "Defines a line-edit with a push-button that can trigger an action.") with an
`ACTION` attribute:

```
LAYOUT
GRID
{
   [f1              ]
   ...
}
...
ATTRIBUTES
BUTTONEDIT f1 = customer.cust_city, ACTION=choose_city, IMAGE="zoom";
...
```

Action views can also be simple [`IMAGE`](1695-image-item-type.md "Defines an area that can display an image resource.") items, when the `ACTION` attribute is specified:

```
LAYOUT
GRID
{
   [i1              ]
   ...
}
...
ATTRIBUTES
IMAGE i1: image1, ACTION=show_details, IMAGE="mylogo";
...
```

Note that `IMAGE` fields can be defined as `TABLE` columns and
define the `ACTION` attribute to trigger user code:

```
LAYOUT
GRID
{
<TABLE t1          >
[c1  |c2       |c3 ]
[c1  |c2       |c3 ]
[c1  |c2       |c3 ]
...
}
...
ATTRIBUTES
...
IMAGE c3: FORMONLY.image, ACTION=delete;
...
```

For more details about image column actions see [Image columns firing actions](2321-image-columns-firing-actions.md "Columns in tables displaying images can trigger action events, when the user selects the image.").

The row selection in a [`TABLE`](2330-defining-the-action-for-a-row-choice.md "The row choice in a TABLE can be associated with a dedicated action.") (or `TREE`) will be considered an action view when
defining the `DOUBLECLICK` attribute:

```
DISPLAY ARRAY arr TO sr.*
        ATTRIBUTES(UNBUFFERED, DOUBLECLICK=select)
  ...
END DISPLAY
```

Action views can also be graphical elements that are standard action triggers on the front-end
platform, such as the [[x] cross button](2289-implementing-the-close-action.md "The close action is a predefined action dedicated to close graphical windows (for example, with the X cross button).") of
desktop windows, that will automatically bind to a "close" action.

Action views can be automatically created in topmenus with the `AUTOCOMMANDS`
placeholder, and in toolbars with the `AUTOITEMS` placeholders. Use the
`CONTENT=ACTIONS` attribute to specify that a list of actions views must be
rendered:

```
TOPMENU
    ...
    AUTOCOMMANDS (CONTENT=ACTIONS)
    ...
TOOLBAR
    ...
    AUTOITEMS (CONTENT=ACTIONS)
    ...
```

## Related links

**Related concepts**  

[Form specification files](1662-form-specification-files.md "Form specification files are the source files defining the layout and content of application forms.")
