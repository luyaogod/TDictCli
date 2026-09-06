---
title: "Presentation styles changes"
source: "fgl-topics/c_fgl_Migrate_to_401_presentation_styles.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 4.01 upgrade guide > Presentation styles changes"
type: "concept"
---

# Presentation styles changes

> Modifications to consider when using presentation styles.

## `scaleIcon` supports `nnpx`

Starting with GBC 4.01.00, the `scaleIcon` attribute value can be defined with
`"px"` unit, to specify a number of pixels.

For more details, see [scaleIcon](../11_user-interface/1630-action-menuaction-style-attributes.md).

## New `Window.topmenuDesktopRendering` and `Window.topmenuMobileRendering` style attributes

Starting with GBC 4.01.05, the `topmenuDesktopRendering` and
`topmenuMobileRendering` style attributes for Window objects can be used to control
the rendering of topmenus.

For more details, see [topmenuDesktopRendering](../11_user-interface/1655-window-style-attributes-basics.md) and [topmenuMobileRendering](../11_user-interface/1655-window-style-attributes-basics.md) style attributes reference page, and [Understanding topmenus](../11_user-interface/1867-understanding-topmenus.md "This is an introduction to topmenu definitions.").

## `defaultTTFColor` applies to any type of form element

Starting with GBC 4.01.06, the `defaultTTFColor` style attribute becomes a
common style attribute that can be specified for any kind of form elements. Before GBC
4.01.06, the `defaultTTFColor` style attribute could only be specified for the Window
element.

For more details, see [defaultTTFColor](../11_user-interface/1629-style-attributes-common-to-all-elements.md).

## New defaults for `navigationDots` and `navigationArrows`

Starting with GBC 4.01.08, the defaults for `navigationDots` /
`navigationArrows` style attributes for `HBOX` and
`VBOX` no longer depend on the front-end platform type (desktop vs mobile), nor do
they depend on the use of `NOSWIPE` attribute.

For more details, see [HBox style attributes](../11_user-interface/1640-hbox-style-attributes.md "HBox presentation style attributes apply to an HBox element."), [VBox style attributes](../11_user-interface/1653-vbox-style-attributes.md "VBox presentation style attributes apply to an VBox element.").

## `resizeFillsEmptySpace` attribute is again supported

Starting with GBC 4.01.10, the `resizeFillsEmptySpace` attribute – which had been
desupported in [version
4.00](0143-presentation-styles-changes.md "Modifications to consider when using presentation styles.") – is once again supported for `TABLE` and
`TREE`.

See [table.resizeFillsEmptySpace](../11_user-interface/1648-table-style-attributes.md), [tree.resizeFillsEmptySpace](../11_user-interface/1649-tree-style-attributes.md).

## `packed` attribute for HBOX and VBOX

Starting with GBC 4.01.14, the new `packed` style attribute is available for
`HBOX` and `VBOX` containers, to pack the children elements to the
origin of the box.

See [packed](../11_user-interface/1640-hbox-style-attributes.md).

## New `windowType` value `modalOnLargeScreen`

Starting with GBC 4.01.15, the window style attribute `windowType` accepts a new
value `"modalOnLargeScreen"`, to indicate that the window must be displayed as
`"modal"` (aka popup) on screens with a size that matches the `@LARGE`
size modifier of form definition attributes (.per), or as
`"normal"` for smaller window sizes.

For more details, refer to [`windowType`](../11_user-interface/1655-window-style-attributes-basics.md).

## New `rowAspect` attribute

Starting with GBC 4.01.21, a new style attribute named `rowAspect` can be defined
to `"list"`, in order to get Material Design list rendering for:

- A `TABLE` with `FLIPPED` attribute in action.
- A `SCROLLGRID` with `WANTFIXEDPAGESIZE=NO` (applies also when
  `customWidget` style attribute is set to `"pagedScrollGrid"`)

For more details, see [`Table.rowAspect`](../11_user-interface/1648-table-style-attributes.md), [`ScrollGrid.rowAspect`](../11_user-interface/1646-scrollgrid-style-attributes.md).

## New `position` attribute for ToolBars

Starting with GBC 4.01.21, the new `position` style attribute can be used to
display toolbars individually at the top or bottom of the window, in the chromebar, or to hide the
toolbar completely.

This new `position` style attribute of ToolBar elements overwrites the Window
style attribute `toolbarPosition`.

For more details, see [position](../11_user-interface/1651-toolbar-style-attributes.md).

## Deprecated style attributes

The following style attributes are deprecated:

- No deprecated style attributes in this version.

## Desupported style attributes

The following style attributes are desupported:

- No desupported style attributes in this version.

## Changes in earlier versions

Make sure to check the upgrade notes of earlier versions, to not miss changes introduced in
maintenance releases. For more details, see [Presentation style changes in BDL
4.00](0143-presentation-styles-changes.md "Modifications to consider when using presentation styles.").

Notable changes introduced in maintenance releases:

- No particular change to consider.
