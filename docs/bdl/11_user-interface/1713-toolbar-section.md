---
title: "TOOLBAR section"
source: "fgl-topics/c_fgl_FormSpecFiles_TOOLBAR_section.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form file structure > TOOLBAR section"
type: "concept"
---

# TOOLBAR section

> The TOOLBAR section defines a toolbar with buttons that are bound to actions.

## Syntax

```
TOOLBAR [toolbar-identifier] [ ( toolbar-attribute [,...] ) ]
  { item
  | placeholder
  | separator
  } [...]
END
```

1. toolbar-identifier defines the name of the toolbar (optional).
2. toolbar-attribute is one of: `STYLE`, `TAG`,
   `BUTTONTEXTHIDDEN`.

where item is:

```
ITEM item-identifier [ ( item-attribute [,...] ) ]
```

1. item-identifier defines the name of the action to bind to, it is
   mandatory.
2. item-attribute is one of: `STYLE`, `TAG`,
   `TEXT`, `IMAGE`, `COMMENT`, `HIDDEN`,
   `AUTOHIDE`.

and placeholder is:

```
AUTOITEMS ( CONTENT = { ACTIONS | PROGRAMS | WINDOWS } )
```

1. `ACTIONS` stands for action views to be rendered for all default action views
   shown in the action panel.
2. `PROGRAMS` stands for the list of current programs displayed by the
   front-end.
3. `WINDOWS` stands for the list of windows created by the current program.

and separator is:

```
SEPARATOR [separator-identifier] [ (separator-attribute [,...] )
```

1. separator-identifier defines the name of the separator (optional).
2. separator-attribute is one of: `STYLE`, `TAG`,
   `HIDDEN`.

## Form attributes

[`AUTOHIDE`](1761-autohide-attribute.md "The AUTOHIDE attribute hides automatically the form element when the related action gets inactive."), [`BUTTONTEXTHIDDEN`](1764-buttontexthidden-attribute.md "The BUTTONTEXTHIDDEN attribute indicates that the button labels for an element are not to be displayed."), [`COMMENT`](1769-comment-attribute.md "The COMMENT attribute defines a hint for the user about the form element."), [`HIDDEN`](1783-hidden-attribute.md "The HIDDEN attribute indicates if an element must be visible to the user."), [`IMAGE`](1785-image-attribute.md "The IMAGE attribute defines the image resource to be displayed for the form item."), [`STYLE`](1825-style-attribute.md "The STYLE attribute specifies a presentation style for a form element."), [`TEXT`](1828-text-attribute.md "The TEXT attribute defines the label associated with a form item."), [`TAG`](1827-tag-attribute.md "The TAG attribute can be used to identify the form item with a specific string.").

## Style attributes

Common: [`backgroundColor`](1629-style-attributes-common-to-all-elements.md), [`border`](1629-style-attributes-common-to-all-elements.md), [`fontFamily`](1629-style-attributes-common-to-all-elements.md), [`fontSize`](1629-style-attributes-common-to-all-elements.md), [`fontStyle`](1629-style-attributes-common-to-all-elements.md), [`fontWeight`](1629-style-attributes-common-to-all-elements.md), [`textColor`](1629-style-attributes-common-to-all-elements.md), [`textDecoration`](1629-style-attributes-common-to-all-elements.md).

Class-specific: [`aspect`](1651-toolbar-style-attributes.md), [`itemsAlignment`](1651-toolbar-style-attributes.md), [`position`](1651-toolbar-style-attributes.md), [`scaleIcon`](1651-toolbar-style-attributes.md), [`size`](1651-toolbar-style-attributes.md).

## Usage

The `TOOLBAR` section defines a toolbar in a form.

The `TOOLBAR` section must appear in the sequence described in [form file structure](1709-form-file-structure.md "A form specification file is defined by a set of sections.").

The `TOOLBAR` section is optional.

A `TOOLBAR` section defines a set of `ITEM` elements that
can be grouped by using a `SEPARATOR` element. Each `ITEM` defines a
toolbar button associated with an action by name. The `SEPARATOR` keyword specifies a
vertical line.

The toolbar buttons are enabled depending on the actions defined by the current interactive
instruction. For example, you can define a toolbar button with the action name "cancel" to bind the
toolbar item to this predefined dialog action. Toolbar items can be automatically hidden when the
corresponding action is disabled, by using the `AUTOHIDE` attribute for the item.

Toolbar button labels are visible by default. The `TOOLBAR` supports the
`BUTTONTEXTHIDDEN` attribute to hide the labels of buttons.

`TOOLBAR` elements can include `AUTOITEMS` place holders, with a
mandatory `CONTENT` attribute, to define if the auto-commands must show action views
for the current actions not bound to an explicit action view `(CONTENT=ACTIONS`), to
show the list of current running programs (`CONTENT=PROGRAMS`), or to show the list
of opened windows of the current program (`CONTENT=WINDOWS`).

When the toolbar is displayed in the [chromebar](2287-action-views-in-chromebar.md "Default action views and toolbar action views can be displayed in the chromebar, to save space on small screens.") (with [`toolBarPosition="chrome"`](1655-window-style-attributes-basics.md) style attribute), `AUTOITEMS` are
not rendered.

`TOOLBAR` elements can get a `STYLE` attribute in order to use a
specific rendering and decoration, based on presentation style definitions.

## Example

```
TOOLBAR tb ( STYLE="mystyle" )
  ITEM accept ( TEXT="Ok", IMAGE="ok" )
  ITEM cancel ( TEXT="Cancel", IMAGE="cancel" )
  SEPARATOR ( TAG="lastSeparator")
  ITEM append ( TEXT="Append", IMAGE="add" )
  ITEM update ( TEXT="Update", IMAGE="modify" )
  ITEM delete ( TEXT="Delete", IMAGE="del" )
  ITEM search ( TEXT="Search", IMAGE="find" )
  SEPARATOR
  AUTOITEMS ( CONTENT=ACTIONS )
END
```

## Related links

**Related concepts**  

[Binding action views to action handlers](2280-binding-action-views-to-action-handlers.md "How are action views of the forms bound to action handlers in the program code?")

[Toolbars](1855-toolbars.md "Toolbars define a bar of buttons that appears at the top of application forms.")

[ACTION DEFAULTS section](1711-action-defaults-section.md "The ACTION DEFAULTS section defines local action view default attributes for the form elements.")
