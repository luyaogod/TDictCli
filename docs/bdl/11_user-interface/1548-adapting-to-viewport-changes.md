---
title: "Adapting to viewport changes"
source: "fgl-topics/c_fgl_form_rendering_responsive_layout.html"
breadcrumb: "User interface > Form definitions > Form rendering > Responsive Layout > Adapting to viewport changes"
type: "concept"
---

# Adapting to viewport changes

> Application forms and functions can be adapted to the front-end viewport size or mobile device orientation.

> **Important:**
>
> This feature is deprecated, its use is
> discouraged although not prohibited.

> **Important:**
>
> The `windowresized` predefined action has been introduced to
> let programs adapt the forms to the screen size or mobile screen orientation. Starting with Genero
> V4, to adapt dynamically the forms to the screen, consider defining forms with [Responsive Layout](1542-understanding-responsive-layout.md "Forms can be designed to adapt to the front-end screen possibilities.") instead. The
> `windowresized` predefined action is deprecated.

## Detecting viewport size / orientation changes

When the mobile device orientation changes, or when the current window is resized on
desktop/web front-ends, the [`windowresized`](2257-list-of-predefined-actions.md) specific predefined action will be sent, if
an `ON ACTION` handler is defined by the current dialog for this
action.

> **Note:**
>
> The `windowresized` action is used only to hide/show items on the current form
> using the standard user interface API ([`ui.Form.setElementHidden()`](../15_library-reference/3156-ui-form-setelementhidden.md "Show or hide form elements.")) and it is not recommended for reloading forms on
> the fly.

This predefined action can be used to detect viewport geometry changes and adapt the
application form to the new size:

```
  ON ACTION windowresized
     -- Code to adapt to the new viewport size
```

> **Note:**
>
> In dialogs allowing field input (`INPUT` / `INPUT ARRAY` or
> `CONSTRUCT`), take care of the current field input: The `windowresized`
> action can force the field validation. Therefore, it is not recommended to use this special action
> in these dialogs. The action can be safely used in `DISPLAY ARRAY` and `MENU`
> dialogs.

To control action view rendering defaults and current field validation behavior when the
`windowresized` action is used, consider setting action default attributes for
this action in your [.4ad file](2259-configuring-actions.md "Action attributes related to decoration, keyboard shortcuts, and behavior can be defined with action attributes.") as
follows:

```
<ActionDefaultList>
  ...
  <ActionDefault name="windowresized" validate="no" defaultView="no" contextMenu="no"/>
  ...
</ActionDefaultList>
```

Another option is to define these action defaults attributes in the `ON ACTION`
handler:

```
ON ACTION windowresized ATTRIBUTES(VALIDATE=NO, DEFAULTVIEW=NO)
   ...
```

## Querying the geometry of the viewport

Use the [feInfo/windowSize front call](../15_library-reference/3399-standard-feinfo.md "Queries general front-end properties.") to
query the actual size of the front-end viewport (desktop window container, web browser tab, or
mobile screen size):

```
DEFINE size STRING
CALL ui.Interface.frontCall("standard","feInfo",["windowSize"],[size])
IF size == "1200x1824" THEN
   ...
END IF
```

## Related links

**Related concepts**  

[List of predefined actions](2257-list-of-predefined-actions.md "List of predefined actions")
