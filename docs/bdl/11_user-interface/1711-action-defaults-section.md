---
title: "ACTION DEFAULTS section"
source: "fgl-topics/c_fgl_FormSpecFiles_ACTION_DEFAULTS_section.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form file structure > ACTION DEFAULTS section"
type: "concept"
---

# ACTION DEFAULTS section

> The ACTION DEFAULTS section defines local action view default attributes for the form elements.

## Syntax

```
ACTION DEFAULTS
   ACTION action-identifier ( action-attribute [,...] )
   [...]
END
```

1. action-identifier defines the name of the action.
2. action-attribute defines an attribute for the
   action.

## Form attributes

[`ACCELERATOR`](1754-accelerator-attribute.md "The ACCELERATOR is an action attribute defining the primary accelerator key for an action."), [`ACCELERATOR2`](1755-accelerator2-attribute.md "The ACCELERATOR2 is an action attribute defining the secondary accelerator key for an action."), [`ACCELERATOR3`](1756-accelerator3-attribute.md "The ACCELERATOR3 is an action attribute defining the third accelerator key for an action."), [`ACCELERATOR4`](1757-accelerator4-attribute.md "The ACCELERATOR4 is an action attribute defining the fourth accelerator key for an action."), [`DEFAULTVIEW`](1773-defaultview-attribute.md "The DEFAULTVIEW attribute defines if a default view (a button) must be displayed for a given action."), [`COMMENT`](1769-comment-attribute.md "The COMMENT attribute defines a hint for the user about the form element."), [`CONTEXTMENU`](1768-contextmenu-attribute.md "The CONTEXTMENU attribute defines whether a context menu option must be displayed for an action."), [`IMAGE`](1785-image-attribute.md "The IMAGE attribute defines the image resource to be displayed for the form item."), [`TEXT`](1828-text-attribute.md "The TEXT attribute defines the label associated with a form item."), [`VALIDATE`](1839-validate-attribute.md "The VALIDATE action attribute defines the data validation level for a given action.").

## Style attributes

Not applicable.

## Usage

The `ACTION DEFAULTS` section centralizes action view attributes (text, comment,
image, accelerators) at the form level.

The `ACTION DEFAULTS` section must appear in the sequence described in [form file structure](1709-form-file-structure.md "A form specification file is defined by a set of sections.").

The `ACTION DEFAULTS` section is optional.

The section holds a list of `ACTION` elements that specify attributes for each
action. The action is identified by the name following the `ACTION` keyword, and
attributes are specified in a list between parenthesis.

The attributes defined in this section are applied to form action views like buttons, toolbar
buttons, or topmenu options, if the individual action views do not explicitly define their own
attributes.

Action attributes can be defined at different levels, see [action configuration](2259-configuring-actions.md "Action attributes related to decoration, keyboard shortcuts, and behavior can be defined with action attributes.") for more details.

## Example

```
ACTION DEFAULTS
  ACTION accept ( COMMENT="Commit order record changes",
                    CONTEXTMENU=NO )
  ACTION cancel ( TEXT="Stop", IMAGE="stop",
                    ACCELERATOR=SHIFT-F2, VALIDATE=NO )
  ACTION print ( COMMENT="Print order information",
                   ACCELERATOR=CONTROL-P,
                   ACCELERATOR2=F5 )
  ACTION zoom1 ( COMMENT="Open items list", VALIDATE=NO )
  ACTION zoom2 ( COMMENT="Open customers list", VALIDATE=NO )
END
```

## Related links

**Related concepts**  

[Action handling basics](2254-action-handling-basics.md "This topic describes the basic concepts of dialog actions.")

[ON ACTION block](1918-on-action-block.md "ON ACTION block")
