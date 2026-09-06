---
title: "Syntax of a topmenu file (.4tm)"
source: "fgl-topics/c_fgl_topmenus_003.html"
breadcrumb: "User interface > Form definitions > Topmenus > Syntax of a topmenu file (.4tm)"
type: "concept"
---

# Syntax of a topmenu file (.4tm)

> A .4tm topmenu file is an XML file that holds a tree of elements defining a topmenu structure.

## Syntax: XML Topmenu

```
<TopMenu [ topmenu-attribute="value" [...] ] >
    group
    [...]
  </TopMenu>
```

where group is:

```
<TopMenuGroup group-attribute="value" [...]>
    { <TopMenuSeparator separator-attribute="value" [...] />
    | <TopMenuCommand command-attribute="value" [...] />
    | <TopMenuAutoCommands autocommands-attribute="value" [...] />
    | group
    } [...]
  </TopMenuGroup>
```

1. topmenu-attribute defines a property of the topmenu.
2. group-attribute defines a property of a topmenu group.
3. command-attribute defines a property of a topmenu command.
4. autocommands-attribute defines a property of a topmenu auto-commands
   placeholder.
5. separator-attribute defines a property of a topmenu separator.
6. value defines the value to be assigned to the attribute.

## Topmenu XML attributes

| Attribute | Type | Description |
| --- | --- | --- |
| `name` | `STRING` | Identifies the topmenu. |
| `style` | `STRING` | Can be used to decorate the element with a presentation style. |
| `tag` | `STRING` | User-defined attribute to identify the node. |

| Attribute | Type | Description |
| --- | --- | --- |
| `name` | `STRING` | Identifies the action corresponding to the topmenu command.Can be prefixed with the sub-dialog identifier. |
| `style` | `STRING` | Can be used to decorate the element with a presentation style. |
| `tag` | `STRING` | User-defined attribute to identify the node. |
| `text` | `STRING` | The text to be displayed in the pull-down menu option. |
| `comment` | `STRING` | The message to be shown for this element. |
| `hidden` | `INTEGER` | Indicates if the command is hidden (0 = visible, 1 = hidden). |
| `image` | `STRING` | The icon to be used in the pull-down menu option. |
| `acceleratorName` | `STRING` | Defines the accelerator name to be display on the left of the menu option text.Note this attribute is only used for decoration (you must also define an action default accelerator). |
| `autohide` | `INTEGER` | When set to 1, indicates that the element must be automatically hidden when corresponding action is disabled. |

| Attribute | Type | Description |
| --- | --- | --- |
| `name` | `STRING` | Identifies the topmenu group. |
| `style` | `STRING` | Can be used to decorate the element with a presentation style. |
| `tag` | `STRING` | User-defined attribute to identify the node. |
| `text` | `STRING` | The text to be displayed in the pull-down menu group. |
| `comment` | `STRING` | The message to be shown for this element. |
| `hidden` | `INTEGER` | Indicates if the group is hidden (0 = visible, 1 = hidden). |
| `image` | `STRING` | The icon to be used in the pull-down menu group. |

| Attribute | Type | Description |
| --- | --- | --- |
| `content` | `STRING` | Defines the type of content for this element. Values can be `"actions"`, `"programs"` or `"windows"`. This attribute is mandatory! |
| `name` | `STRING` | Identifies the topmenu auto-commands element. |
| `style` | `STRING` | Can be used to decorate the element with a presentation style. |
| `tag` | `STRING` | User-defined attribute to identify the node |
| `hidden` | `INTEGER` | Indicates if the element is hidden (0 = visible, 1 = hidden). |

| Attribute | Type | Description |
| --- | --- | --- |
| `name` | `STRING` | Identifies the topmenu separator. |
| `style` | `STRING` | Can be used to decorate the element with a presentation style. |
| `tag` | `STRING` | User-defined attribute to identify the node |
| `hidden` | `INTEGER` | Indicates if the separator is hidden (0 = visible, 1 = hidden). |

## Related links

**Related concepts**  

[Configuring actions](2259-configuring-actions.md "Action attributes related to decoration, keyboard shortcuts, and behavior can be defined with action attributes.")

[Presentation styles](1607-presentation-styles.md "Use presentation styles to specify decoration attributes for window and form elements.")

[Binding action views to action handlers](2280-binding-action-views-to-action-handlers.md "How are action views of the forms bound to action handlers in the program code?")
