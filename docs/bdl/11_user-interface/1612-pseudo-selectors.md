---
title: "Pseudo selectors"
source: "fgl-topics/c_fgl_presentation_styles_006.html"
breadcrumb: "User interface > Form definitions > Presentation styles > Using presentation styles > Pseudo selectors"
type: "concept"
---

# Pseudo selectors

> Pseudo selectors can be used to apply style only when some conditions are fulfilled.

Pseudo selectors are preceded by a colon and can be combined:

```
<Style name="Table:even:input" >
<Style name="Edit:focus" >
<Style name="Edit.important:focus" >
```

When combining several pseudo selectors, the style will be applied if all pseudo selector
conditions are fulfilled.

Depending on the type of the front-end, some pseudo selectors are meaningless, or unsupported.
See the table below to check which pseudo selectors are supported on your front-end platform.

Pseudo selectors have different priorities; the style with the most important pseudo selector
will be used when several styles match.

| Priority | Pseudo selector | Condition |
| --- | --- | --- |
| 1 | `focus` | The widget has the focus |
| 2 | `query` | The widget is in construct mode |
| 3 | `display` | The widget is in a display array |
| 4 | `input` | The widget is in an input array, input or construct |
| 5 | `even` | This widget is on an even row if in a list (Table or Tree) |
| 6 | `odd` | This widget is on an odd row if in a list (Table or Tree) |
| 7 | `inactive` | The widget is inactive |
| 8 | `active` | The widget is active |
| 9 | `message` | Applies only to text displayed with the [`MESSAGE`](1885-message.md "The MESSAGE instruction displays a message to the user.") instruction |
| 10 | `error` | Applies only to text displayed with the [`ERROR`](1886-error.md "The ERROR instruction displays an error message to the user.") instruction |
| 11 | `summaryLine` | Applies only to text displayed in [`AGGREGATE`](1674-aggregate-fields.md "An AGGREGATE field defines a screen-record field to display summary information for a TABLE column.") fields of tables |

## Related links

**Related concepts**  

[Summary lines in tables](2329-summary-lines-in-tables.md "Table views can display a summary line, to show aggregate values for columns.")

**Related reference**  

[Message style attributes](1644-message-style-attributes.md "Message presentation style attributes apply to an ERROR or MESSAGE instruction.")
