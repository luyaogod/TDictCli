---
title: "Front-end stored settings"
source: "fgl-topics/c_fgl_form_rendering_stored_settings.html"
breadcrumb: "User interface > Form definitions > Form rendering > Front-end stored settings"
type: "concept"
---

# Front-end stored settings

> Front-ends can store some layout properties of windows and form elements, for subsequent program executions.

## Purpose of stored settings

With a GUI front-end, when a [window/form](1561-windows-and-forms.md "The section describes the concept of windows and forms in the language.") is
closed, the front-end stores widget sizing and positionning properties locally on the platform where
the front-end executes. When the window/form is reopened (in a new or in the same program instance),
these variable layout properties are restored, to display the forms and its content with the same
aspect as when they were closed.

For example, [`TABLE`](1703-table-item-type.md "Defines a list view widget.") columns
can be resized, reordered, and selected to sort rows. Such variable layout properties are saved in
the stored settings and restored when the form is re-displayed.

## Managing stored settings on front-end side

Stored settings can be controlled by the end-user, in the front-end configuration panel.

For example, it is possible to disable stored settings completely, or to reset them in order to
get the original form layouts.

See front-end specific documentation for more details.

## Controlling stored settings with program files

Some window presentation styles attributes such as
[`forceDefaultSettings`](1655-window-style-attributes-basics.md) or
[`position`](1655-window-style-attributes-basics.md)
can be used to control stored settings.

## Related links

**Related concepts**  

[Form specification files](1662-form-specification-files.md "Form specification files are the source files defining the layout and content of application forms.")

[Style attributes reference](1628-style-attributes-reference.md "A presentation style attribute may be a common attribute that can be applied to any graphical element. Most presentation style attributes apply only to a specific graphical element.")
