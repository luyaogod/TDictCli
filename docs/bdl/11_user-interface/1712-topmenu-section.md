---
title: "TOPMENU section"
source: "fgl-topics/c_fgl_FormSpecFiles_TOPMENU_section.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form file structure > TOPMENU section"
type: "concept"
---

# TOPMENU section

> The TOPMENU section defines a pull-down menu with options that are bound to actions.

## Syntax

```
TOPMENU [topmenu-identifier] [ ( topmenu-attribute [,...] ) ]
  group
  [...]
END
```

1. topmenu-identifier defines the name of the top menu (optional).
2. topmenu-attribute can be: `STYLE`, `TAG`.

where topmenu-group is:

```
GROUP group-identifier [ ( group-attribute [,...] ) ]
  { command
  | placeholder
  | group
  | separator
  } [...]
END
```

1. group-identifier defines the name of the group, it is mandatory.
2. group-attribute is one of: `STYLE`, `TEXT`,
   `IMAGE`, `COMMENT`, `TAG`,
   `HIDDEN`.

where command is:

```
COMMAND command-identifier [ ( command-attribute [,...] ) ]
```

1. command-identifier defines the name of the action to bind to, it is
   mandatory.
2. command-attribute is one of: `STYLE`, `TEXT`,
   `IMAGE`, `COMMENT`, `TAG`, `HIDDEN`,
   `ACCELERATOR`, `AUTOHIDE`.

and placeholder
is:

```
AUTOCOMMANDS ( CONTENT = { ACTIONS | PROGRAMS | WINDOWS } )
```

1. `ACTIONS` stands for action views to be rendered for all default action views
   shown in the action panel.
2. `PROGRAMS` stands for the list of current programs displayed by the
   front-end.
3. `WINDOWS` stands for the list of windows created by the current program.

and separator
is:

```
SEPARATOR [separator-identifier] [ ( separator-attribute [,...] ) ]
```

1. separator-identifier defines the name of the separator (optional).
2. separator-attribute is one of: `STYLE`, `TAG`,
   `HIDDEN`.

## Form attributes

[`ACCELERATOR`](1754-accelerator-attribute.md "The ACCELERATOR is an action attribute defining the primary accelerator key for an action."), [`AUTOHIDE`](1761-autohide-attribute.md "The AUTOHIDE attribute hides automatically the form element when the related action gets inactive."), [`COMMENT`](1769-comment-attribute.md "The COMMENT attribute defines a hint for the user about the form element."), [`HIDDEN`](1783-hidden-attribute.md "The HIDDEN attribute indicates if an element must be visible to the user."), [`IMAGE`](1785-image-attribute.md "The IMAGE attribute defines the image resource to be displayed for the form item."), [`STYLE`](1825-style-attribute.md "The STYLE attribute specifies a presentation style for a form element."), [`TEXT`](1828-text-attribute.md "The TEXT attribute defines the label associated with a form item."), [`TAG`](1827-tag-attribute.md "The TAG attribute can be used to identify the form item with a specific string.").

## Style attributes

Common: [`backgroundColor`](1629-style-attributes-common-to-all-elements.md), [`border`](1629-style-attributes-common-to-all-elements.md), [`fontFamily`](1629-style-attributes-common-to-all-elements.md), [`fontSize`](1629-style-attributes-common-to-all-elements.md), [`fontStyle`](1629-style-attributes-common-to-all-elements.md), [`fontWeight`](1629-style-attributes-common-to-all-elements.md), [`textColor`](1629-style-attributes-common-to-all-elements.md), [`textDecoration`](1629-style-attributes-common-to-all-elements.md).

Class-specific: none.

## Usage

The `TOPMENU` section is used to define a pull-down menu in a form.

The `TOPMENU` section must appear in the sequence described in [form file structure](1709-form-file-structure.md "A form specification file is defined by a set of sections.").

The `TOPMENU` section is optional.

In a `TOPMENU` section, you build a tree of `GROUP` elements
to design the pull-down menu. A `GROUP` can contain `COMMAND`,
`SEPARATOR` or `GROUP` children. A `COMMAND`
defines a pull-down menu option that triggers an action when it is selected. In the topmenu
specification, command-identifier defines which action a menu option is bound
to. For example, if you define a topmenu option as "`COMMAND zoom`", it can be
controlled by an "`ON ACTION zoom`" clause in an interactive instruction.

The topmenu commands are enabled depending on the actions defined by the current interactive
instruction. For example, you can define a topmenu option with the action name "cancel" to bind the
pull-down item to this predefined dialog action. Topmenu commands can be automatically hidden when
the corresponding action is disabled, by using the `AUTOHIDE` attribute for the
command.

An accelerator name can be defined for a topmenu command; this accelerator name will be
used for display in the command item. You must define the same accelerator in the action defaults
section for the action name of the topmenu command.

`TOPMENU` elements can include `AUTOCOMMANDS` place holders, with a
mandatory `CONTENT` attribute, to define if the auto-commands must show action views
for the current actions not bound to an explicit action view `(CONTENT=ACTIONS`), to
show the list of current running programs (`CONTENT=PROGRAMS`), or to show the list
of opened windows of the current program (`CONTENT=WINDOWS`).

`TOPMENU` elements can get a `STYLE` attribute in order to use a
specific rendering and decoration, based on presentation style definitions.

## Example

```
TOPMENU tm ( STYLE="mystyle" )
  GROUP form (TEXT="Form")
    COMMAND help (TEXT="Help", IMAGE="quest")
    COMMAND quit (TEXT="Quit")
  END
  GROUP edit (TEXT="Edit")
    COMMAND accept (TEXT="Validate", IMAGE="ok", TAG="acceptMenu")
    COMMAND cancel (TEXT="Cancel", IMAGE="cancel")
  END
  GROUP records (TEXT="Records")
    COMMAND append (TEXT="Add", IMAGE="plus")
    COMMAND delete (TEXT="Remove", IMAGE="minus")
    COMMAND update (TEXT="Modify", IMAGE="accept")
    SEPARATOR (TAG="lastSeparator")
    COMMAND search (TEXT="Search", IMAGE="find")
  END
  GROUP windows (TEXT="Windows")
    AUTOCOMMANDS (CONTENT=WINDOWS)
  END
END
```

## Related links

**Related concepts**  

[Binding action views to action handlers](2280-binding-action-views-to-action-handlers.md "How are action views of the forms bound to action handlers in the program code?")

[Topmenus](1866-topmenus.md "Topmenus define typical pull-down menus that appear at the top of application forms.")

[ACTION DEFAULTS section](1711-action-defaults-section.md "The ACTION DEFAULTS section defines local action view default attributes for the form elements.")
