---
title: "Syntax of a toolbar file (.4tb)"
source: "fgl-topics/c_fgl_toolbars_003.html"
breadcrumb: "User interface > Form definitions > Toolbars > Syntax of a toolbar file (.4tb)"
type: "concept"
---

# Syntax of a toolbar file (.4tb)

> A .4tb toolbar file is an XML file that holds a tree of elements defining a toolbar structure.

## Syntax: XML Toolbar

```
<ToolBar [ toolbar-attribute="value" [...]] >
  { <ToolBarItem item-attribute="value" [...] />
  | <ToolBarSeparator separator-attribute="value" [...] />
  | <ToolBarAutoItems autoitems-attribute="value" [...] />
  } [...]
</ToolBar>
```

1. toolbar-attribute defines a property of the toolbar.
2. item-attribute defines a property of a toolbar item.
3. autoitems-attribute defines a property of a toolbar auto-items
   placeholder.
4. separator-attribute defines a property of a toolbar
   separator.
5. value defines the value to be assigned to the attribute.

## Toolbar XML attributes

| Attribute | Type | Description |
| --- | --- | --- |
| `style` | `STRING` | Use to decorate the element with a presentation style. |
| `tag` | `STRING` | User-defined attribute to identify the node. |
| `name` | `STRING` | Identifies the toolbar. |
| `buttonTextHidden` | `INTEGER` | Defines if the text of toolbar buttons must appear by default (0 = visible, 1 = hidden).**Important:**The `buttonTextHidden` attribute is deprecated for `ToolBar` elements: Use Toolbar/[`"aspect"`](1651-toolbar-style-attributes.md) style attribute instead.**Note:**On front-ends where the toolbar button texts can be hidden with an option (context menu), the stored settings take precedence over the `BUTTONTEXTHIDDEN` attribute. |

| Attribute | Type | Description |
| --- | --- | --- |
| `name` | `STRING` | Identifies the action corresponding to the toolbar button.Can be prefixed with the sub-dialog identifier. |
| `style` | `STRING` | Use to decorate the element with a presentation style. |
| `tag` | `STRING` | User-defined attribute to identify the node. |
| `text` | `STRING` | The text to be displayed in the toolbar button. |
| `comment` | `STRING` | The message to be shown as tooltip when the user selects a toolbar button. |
| `hidden` | `INTEGER` | Indicates if the item is hidden (0 = visible, 1 = hidden). |
| `image` | `STRING` | The icon to be used in the toolbar button. |
| `autohide` | `INTEGER` | When set to 1, indicates that the element must be automatically hidden when corresponding action is disabled. |

| Attribute | Type | Description |
| --- | --- | --- |
| `content` | `STRING` | Defines the type of content for this element. Values can be `"actions"`, `"programs"` or `"windows"`. This attribute is mandatory! |
| `name` | `STRING` | Identifies the toolbar separator. |
| `style` | `STRING` | Use to decorate the element with a presentation style. |
| `tag` | `STRING` | User-defined attribute to identify the node. |
| `hidden` | `INTEGER` | Indicates if the auto-item is hidden (0 = visible, 1 = hidden). |

| Attribute | Type | Description |
| --- | --- | --- |
| `name` | `STRING` | Identifies the toolbar separator. |
| `style` | `STRING` | Use to decorate the element with a presentation style. |
| `tag` | `STRING` | User-defined attribute to identify the node. |
| `hidden` | `INTEGER` | Indicates if the separator is hidden (0 = visible, 1 = hidden). |

## Related links

**Related concepts**  

[Configuring actions](2259-configuring-actions.md "Action attributes related to decoration, keyboard shortcuts, and behavior can be defined with action attributes.")

[Presentation styles](1607-presentation-styles.md "Use presentation styles to specify decoration attributes for window and form elements.")

[Binding action views to action handlers](2280-binding-action-views-to-action-handlers.md "How are action views of the forms bound to action handlers in the program code?")
