---
title: "Presentation styles changes"
source: "fgl-topics/c_fgl_Migrate_to_600_presentation_styles.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 6.00 upgrade guide > Presentation styles changes"
type: "concept"
---

# Presentation styles changes

> Modifications to consider when using presentation styles.

## New `customWidget` value `phoneEdit`

Starting with GBC 5.01.06, the `customWidget` style attribute can be set to
`"phoneEdit"`, to get decoration and input helpers for international phone
numbers, with country flags.

See [`customWidget`](../11_user-interface/1637-edit-style-attributes.md) style attribute for `EDIT` fields.

## New `customWidget` value `tagEdit`

Starting with GBC 5.01.06, the `customWidget` style attribute can be set to
`"tagEdit"`, to get multi-value input in `EDIT` and
`BUTTONEDIT` form fields.

See [`customWidget` for `EDIT`](../11_user-interface/1637-edit-style-attributes.md) and [`customWidget` for
`BUTTONEDIT`](../11_user-interface/1632-buttonedit-style-attributes.md).

## New `imageBorderRadius` style attribute

Starting with GBC 6.00.01, the `imageBorderRadius` style attribute can be defined
for `TABLE` (`IMAGECOLUMN`) and `IMAGE` form items
(fields and static image form items).

See [`imageBorderRadius` for `TABLE`](../11_user-interface/1648-table-style-attributes.md "Table presentation style attributes apply to a TABLE container.") and [`imageBorderRadius`
for `IMAGE`](../11_user-interface/1641-image-style-attributes.md "Image style presentation attributes apply to an IMAGE element.").

## New `applicationListPosition` style attribute

Starting with GBC 6.00.02, the `applicationListPosition` style attribute can be
used to set the appication list in the ChromeBar.

See [applicationListPosition](../11_user-interface/1652-userinterface-style-attributes.md) for more
details.

## New `aiWritingAssistant` style attribute

Starting with GBC 6.00.03, the `aiWritingAssistant` style attribute can be used
for `TEXTEDIT` fields, to get help the end user improve the text.

See [aiWritingAssistant](../11_user-interface/1650-textedit-style-attributes.md) for more details.

## New `customWidget` style attribute for `COMBOBOX`

Starting with GBC 6.00.03, the `customWidget` style attribute can be set to
"`comboBoxSimpleConstruct`" for `COMBOBOX` fields, to restrict the
item selection and the input with `QUERYEDITABLE`, during a
`CONSTRUCT` dialog.

See [customWidget](../11_user-interface/1634-combobox-style-attributes.md) for more details.

## Rendering of deprecated `tabbedContainer` style attribute

Starting with GBC version 6.00.03, the rendering for `tabbedContainer` has
changed.

Setting `tabbedContainer=yes` is now the same than
`applicationListPosition=top`: The specific `tabbedContainer=yes`
implementation is no more available.

Noticable changes:

1. There is no more permanent view of the first application: Each application displays its topmenu
   and ring menu action buttons in the main viewport/container, as it is the case today in any normal
   application.
2. There is no more requirement for having a StartMenu when
   `tabbedContainer=yes`.
3. SideBarDrawer is no longer automatically docked.
4. App section of SideBarDrawer remains visible.
5. Closing parent app does no longer close all the children apps.

Reference: [`Window.tabbedContainer` style attribute](../11_user-interface/1655-window-style-attributes-basics.md).

## `showCurrentMonthOnly` for `DATEEDIT`/`DATETIMEEDIT`

Starting with GBC version 6.00.07, the `showCurrentMonthOnly` style attribute can
be used for `DATEEDIT` and `DATETIMEEDIT` fields, to show only days of
the current month.

Reference:

- [`DATEEDIT` style attributes](../11_user-interface/1635-dateedit-style-attributes.md)
- [`DATETIMEEDIT` style attributes](../11_user-interface/1636-datetimeedit-style-attributes.md)

## Deprecated style attributes

The following style attributes are deprecated:

- No deprecated style attributes in this version.

## Desupported style attributes

The following style attributes are desupported:

- Starting with GBC 5.01.05, the `startMenuPosition` style attribute is no longer
  supported: The new UX includes a rework of the StartMenu, which will always be visible in the
  SideBarDrawer as a popup treeview.

## Changes in earlier versions

Make sure to check the upgrade notes of earlier versions, to not miss changes introduced in
maintenance releases. For more details, see [Presentation style changes in BDL
5.01](0105-presentation-styles-changes.md "Modifications to consider when using presentation styles.").

Notable changes introduced in maintenance releases:

- No particular change to consider.
