---
title: "Presentation styles changes"
source: "fgl-topics/c_fgl_Migrate_to_320_presentation_styles.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 3.20 upgrade guide > Presentation styles changes"
type: "concept"
---

# Presentation styles changes

> Modifications to consider when using presentation styles.

## `common.imageCache` (GDC specific) style attribute is desupported

Starting with GDC 3.20.02, the `imageCache` style attribute is no longer
supported.

The `imageCache` style attribute was only available with the GDC front-end, to
indicate if the image resource of a form element had to be cached by the front-end. By default,
images for image fields (typically changing at runtime) were not cached, while images of static form
items (Button, TopMenu item, Toolbar item) were cached.

> **Note:**
>
> The front-end file cache management has been improved in Genero 3.20. For more details about
> image resource handling, see [The resource file cache of the front-end](../11_user-interface/1586-providing-the-image-resource.md).

## Action views rendering in GBC chromebar

Starting with GBC 1.00.51, the value `"chrome"` can be used for the following
Window style attributes:

- `actionPanelPosition`
- `ringMenuPosition`
- `toolBarPosition`

When setting one of these style attributes to the value `"chrome"`, the action
views of the corresponding panel are rendered in the GBC chromebar.

On mobile devices, this will follow Material Design specification.

For more details, see [Action views in chromebar](../11_user-interface/2287-action-views-in-chromebar.md "Default action views and toolbar action views can be displayed in the chromebar, to save space on small screens.").

## Collapsible groups with GBC

Since GBC 1.00.52, group elements in grid-based or stack-based form layouts can be defined as
collapsible with new presentation style attributes:

- `collapsible`: When set to `"yes"`, makes the group element
  collapsible by the end user.
- `initiallyCollapsed`: When set to `"yes"`, forces the group to be
  collapsed when the form is displayed the first time.

  Starting with GBC 1.00.54, value of
  `initiallyCollapsed` can be `"never"` to never collapse a group when
  it displays, or `"always"`, to always collapse a group when it displays (ignoring
  stored settings or previous display state).
- Starting with GBC 1.00.53: `collapserPosition`: Can be set to
  `"left"` or `"right"` to define the position of the collapser
  icon.

For more details, see [Collapsible groups](../11_user-interface/1693-group-item-type.md).

## New `Folder.collapserPosition` (GBC)

Starting with GBC 1.00.53, the `collapserPosition` style attribute can be set to
`"left"` or `"right"` to define the position of the collapser icon,
when the `position` style attribute is defined to "accordion" for the folder.

For more details, see [FOLDER rendering](../11_user-interface/1691-folder-item-type.md).

## New `HBox.splitViewRendering` (GBC)

Starting with GBC 1.00.52, `HBOX` containers can be rendered as split views in
GBC, when the width of the webview is lower than a pre-defined number of pixels.

For more details, see [c\_fgl\_FormSpecFiles\_HBOX.html#c\_fgl\_FormSpecFiles\_HBOX\_\_section\_hbox\_splitview](../11_user-interface/1694-hbox-item-type.md).

## `Button.alignment` style attribute (GMA)

The `alignment` style attribute for `BUTTON` form items is now
supported by:

- GMA 1.40.03

For the possible values of this attribute, see [`Button.alignment`](../11_user-interface/1631-button-style-attributes.md).

## New `ScrollGrid.itemsAlignment` style attribute (GBC)

Starting with GBC 1.00.55, it is possible to control the alignment of elements inside a
`SCROLLGRID` container with the `itemsAlignment` style attribute.

For more details, see [Controlling element alignment inside a scrollgrid](../11_user-interface/2339-controlling-scrollgrid-rendering.md).

## New `DateTimeEdit.enableCalendar` style attribute (GDC)

Since GDC version 3.20.09, `DATETIMEEDIT` fields can be rendered with a date/time
picker, when defining the `enableCalendar` style attribute to
`"yes"`.

The date/time picker reacts like with `DATEEDIT` fields: ENTER key will close it
and validate the date/time selection, ESC key will close the widget and cancel the date/time
selection.

When using the date/time picker mode of `DATETIMEEDIT`, the content of the field
can be cleared, and it makes the `CONSTRUCT` mode available.

When used with an `INPUT`, `INPUT ARRAY` (or in simple display
mode), the format of the date part is defined by the DBDATE environment variable, and the format for
the time part is defined by the GDC monitor language settings. In `CONSTRUCT` mode,
the display format is set to ISO, when the variable associated to the field is a
`DATETIME`.

## New `ComboBox.qtStyle` style attribute (GDC)

Since GDC version 3.20.11, `COMBOBOX` fields can be rendered with a specific Qt
widget style, by defining the `qtStyle` attribute. In order to get a background color
defined by the `backgroundColor` style attribute, the `qtStyle`
attribute must be set to `"Windows"`.

> **Important:**
>
> This feature is desupported. It may still be
> available, but can be removed in a next version.

This feature is desupported with GDC version 4.00, using now [Universal Rendering](0134-universal-rendering-as-standard.md "All front-ends (GDC, GAS, GMA, GMI) use now the GBC as rendering engine.").

## New `sanitize` style attribute (GBC)

By default, to avoid "Stored XSS" attacks, the GBC (version 1.00.58) cleans the HTML sent to form
elements, to ensure no malicious script can be executed. This security control prevents for example
to create a label with HTML content such as `"<a href='mailto: …"`

Starting with GBC 1.00.59, it is possible to disable the checking of HTML content send to form
elements, by setting the `sanitize` style attribute to `"no"`. This
attribute is available for form items that accept HTML content: [Label](../11_user-interface/1642-label-style-attributes.md "Label presentation style attributes apply to LABEL elements."), [Message](../11_user-interface/1644-message-style-attributes.md "Message presentation style attributes apply to an ERROR or MESSAGE instruction."), and [TextEdit](../11_user-interface/1650-textedit-style-attributes.md "Textedit presentation style attributes apply to the TEXTEDIT element.").

## New `Folder.lateRendering` style attribute (GBC)

Starting with GBC 1.00.60, the layout computation of `FOLDER` containers can be
optimized with the `lateRendering` style attribute set to `"yes"`, to
compute only the layout of the current displayed page.

For more details, see [Folder](../11_user-interface/1638-folder-style-attributes.md) style attribute reference pages.
