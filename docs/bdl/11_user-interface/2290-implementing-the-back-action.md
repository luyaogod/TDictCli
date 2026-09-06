---
title: "Implementing the back action"
source: "fgl-topics/c_fgl_prog_dialogs_back_action.html"
breadcrumb: "User interface > User interface programming > Dialog actions > Implementing the back action"
type: "concept"
---

# Implementing the back action

> The back action is a predefined action dedicated to move back in the stack of windows/forms.

## Purpose of the `back` action

In graphical applications and especially web applications, it is common to see a back
button with a left arrow, to move back to the previous window/from.

Genero BDL has a predefined action named "`back`" which is dedicated to this
specific need.

When defining an `ON ACTION back` action handler, the front-end will automatically
display a button on the left of the chromebar (if the chromebar is visible):

![Back button rendering on desktop](../_images/BackButton1_gbc.jpg)

*Back button rendering on desktop*

The back button in the chromebar is considered as [default action view](2258-default-action-views.md "A default action view is created to render an action handler, when no explicit action view exists for it."), and consequently, no
button is displayed in the action frame. However, when the chromebar is hidden, the back button for
this action will be created in the action frame.

## Controlling the `back` action

Like regular action views, it is possible to enable/disable the `back`
action:

```
MENU "test"
    ON ACTION enable
       CALL DIALOG.setActionActive("back",TRUE)
    ON ACTION disable
       CALL DIALOG.setActionActive("back",FALSE)
    ON ACTION back
       MESSAGE "GO BACK!"
END MENU
```

Like other default action views displayed in the chromebar, the back button is hidden, when the
action is disabled.

If needed, like regular default action views, you can control the visibility of the back button
with the [`DEFAULTVIEW`](2273-defaultview-action-attribute.md "The DEFAULTVIEW attribute defines if a default view (a button) must be displayed for a given action.")
attribute.

When using the [`UNBUFFERED`](2235-the-buffered-and-unbuffered-modes.md "The buffered and unbuffered mode control the synchronization of program variables and form fields.")
attribute in the input dialog, you can control the field validation with the [`VALIDATE`](2277-validate-action-attribute.md "The VALIDATE action attribute defines the data validation level for a given action.") action attribute. The
`back` action is typically similar to an `accept` action, not
configured with `VALIDATE=NO`.

The button widget created for the `back` action ignores the action attributes
[`TEXT`](2276-text-action-attribute.md "The TEXT attribute defines the label associated to the action."), [`IMAGE`](2274-image-action-attribute.md "The IMAGE attribute defines the image resource to be displayed for the action.") and [`COMMENT`](2271-comment-action-attribute.md "The COMMENT attribute defines hint for the user about the action.").

## Related links

**Related concepts**  

[Predefined actions](2255-predefined-actions.md "Genero predefines some action names for common operations of interactive instructions.")
