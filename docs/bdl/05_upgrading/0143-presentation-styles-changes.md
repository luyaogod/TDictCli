---
title: "Presentation styles changes"
source: "fgl-topics/c_fgl_Migrate_to_400_presentation_styles.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 4.00 upgrade guide > Presentation styles changes"
type: "concept"
---

# Presentation styles changes

> Modifications to consider when using presentation styles.

## New `Table.showGrid` style attribute

Starting with GBC 4.00.01, the `TABLE` and `TREE` containers can
configured to display vertical and horizontal lines individually, with the `showGrid`
style attribute.

For more details, see [Table](../11_user-interface/1648-table-style-attributes.md) and [Tree](../11_user-interface/1649-tree-style-attributes.md) style attributes reference pages and [Showing table grid lines](../11_user-interface/2319-controlling-table-rendering.md).

## New `Table.alternateRows` style attribute

Starting with GBC 4.00.01, the `TABLE` and `TREE` containers can
use the `alternateRows` style attribute to render odd/even rows with a different
background color.

For more details, see [Table](../11_user-interface/1648-table-style-attributes.md) and [Tree](../11_user-interface/1649-tree-style-attributes.md) style attributes reference pages and [Alternate background color of odd/even rows](../11_user-interface/2319-controlling-table-rendering.md).

## New `Table.rowHover` style attribute

Starting with GBC 4.00.01, the `rowHover` style attribute for
`TABLE` and `TREE` containers can be used to disable visual effects
when the mouse hovers a given row.

For more details, see [Table](../11_user-interface/1648-table-style-attributes.md) and [Tree](../11_user-interface/1649-tree-style-attributes.md) style attributes reference pages and [Distinguish row on mouse hover](../11_user-interface/2319-controlling-table-rendering.md).

## Changes with `Window.toolbarPosition` style attribute

Starting with GBC 4.00.01, the `toolbarPosition` style attribute for Window
objects supports now the `"bottom"` position.

For more details, see [Window](../11_user-interface/1655-window-style-attributes-basics.md) style attributes reference page.

## New `Toolbar` style attributes

Starting with GBC 4.00.03, Toolbars can be configured with the `aspect`,
`itemsAlignment` and `"size"` style attributes, to define the
rendering and placement of buttons, as well as the arrangements of icons and texts and button
sizes.

For more details, see [Toolbar](../11_user-interface/1651-toolbar-style-attributes.md "ToolBar presentation style attributes apply to the TOOLBAR element.") style attributes reference page.

## Changes with `UserInterface.browserMultiPage` style attribute

Starting with GBC 4.00.02, when using the GBC front-end with a GAS/web browser, the
`browserMultiPage` style attribute for UserInterface defines if the forms of a child
program are displayed in a different folder tab, only if the child program was started with
`RUN ... WITHOUT WAITING`. In prior GBC versions, child programs started with
`RUN` (not using the `WITHOUT WAITING` clause) also displayed in a
separate brower tab. This makes not much sense, as the parent program is frozen and waits for the
child program to terminate. Consequently, `browserMultiPage=yes` has no more effect
with a simple `RUN` instruction and child program forms will display in the same
browser tab displaying parent program forms.

For more details, see [browserMultiPage](../11_user-interface/1652-userinterface-style-attributes.md).

## New `UserInterface.desktopMultiWindow` style attribute

Starting with GBC 4.00.03, when using the GDC front-end, the `desktopMultiWindow`
style attribute for UserInterface defines if program forms are displayed in the same unique window
container, or in a dedicated window container.

For more details, see [desktopMultiWindow](../11_user-interface/1652-userinterface-style-attributes.md).

## New `UserInterface.applicationListVisible` style attribute

Starting with GBC 4.00.03, the `applicationListVisible` style attribute for
UserInterface can be used to define the visibility of the list of programs displaying on the
front-end.

For more details, see [applicationListVisible](../11_user-interface/1652-userinterface-style-attributes.md).

## New `UserInterface.windowListVisible` style attribute

Starting with GBC 4.00.03, the `windowListVisible` style attribute for
UserInterface can be used to define the visibility of the list of windows opened by the current
program.

For more details, see [windowListVisible](../11_user-interface/1652-userinterface-style-attributes.md).

## Deprecated style attributes

The following style attributes are deprecated:

- `HBox`
  - `splitViewRendering` (use [`SPLIT[@screen-size]`](../11_user-interface/1818-split-attribute.md "The SPLIT attribute forces a horizontal box to show only one child container.") form file attribute instead)
- `Table`
  - [`tableType/listView`](../11_user-interface/1648-table-style-attributes.md) (As replacement, consider using a [`TABLE`](../11_user-interface/2315-table-views.md "Describes how to implement table/list views.") with the [`FLIPPED`](../11_user-interface/1780-flipped-attribute.md "The FLIPPED attribute flips TABLE columns into rows.") attribute, in conjunction with `rowAspect` style attribute set to `"list"`.)

- `Window`
  - [`tabbedContainer`](../11_user-interface/1655-window-style-attributes-basics.md) (use [window container configuration](../11_user-interface/1563-containers-for-program-windows.md "Program windows are displayed in window containers by the front-end.") instead)

## Desupported style attributes

The following style attributes are desupported, as a consequence of common Universal Rendering
using GBC on all front-ends (there is no more Native Rendering mode):

- Common style attributes
  - `localAccelerators` common style attribute
  - `showAcceleratorInToolTip` common style attribute
- `ComboBox`
  - `autoSelectionStart`
  - `comboboxCompleter`
  - `completionTimeout`
  - `qtStyle` (QT lib specific feature)
- `Action`/`MenuAction`
  - `androidActionPosition`
  - `androidActionWithIcon`
  - `androidActionWithText`
- `CheckBox`
  - `iosCheckBoxOnTintColor`
- `DateEdit`
  - `showGrid`
- `DateTimeEdit`
  - `enableCalendar`
  - `showGrid`
- `Edit`
  - `dataTypeHint` (use [`KEYBOARDHINT`](../11_user-interface/1798-keyboardhint-attribute.md "The KEYBOARDHINT attribute gives an indication of the kind of data the form field contains, allowing the front-end to adapt the keyboard accordingly.") form field attribute instead)
  - `spellCheck` (with GBC, browser spell check is used)
- `Form`
  - `resetFormSize`
- `Image`
  - `imageContainerType` (was anyway GDC-experimental, use [`WEBCOMPONENT`](../11_user-interface/1708-webcomponent-item-type.md "Defines a specialized form item that holds an external component.") instead)
- `Menu`
  - `position`
- `Message`
  - `position`
- `RadioGroup`
  - `autoSelectionStart`
  - `completionTimeout`
- `Table`
  - `resizeFillsEmptySpace` (use [`STRETCHCOLUMNS`](../11_user-interface/1822-stretchcolumns-attribute.md "The STRETCHCOLUMNS attribute makes all TABLE/TREE columns stretchable.") or [`STRETCH=X`](../11_user-interface/1821-stretch-attribute.md "The STRETCH attribute defines if the form element can grow or has a fixed size.") on individual columns) - Re-introduced [in GBC 4.01.10](0127-presentation-styles-changes.md).
  - `summaryLineAlwaysAtBottom`
- `Tree`
  - `resizeFillsEmptySpace` (use [`STRETCHCOLUMNS`](../11_user-interface/1822-stretchcolumns-attribute.md "The STRETCHCOLUMNS attribute makes all TABLE/TREE columns stretchable.") or [`STRETCH=X`](../11_user-interface/1821-stretch-attribute.md "The STRETCH attribute defines if the form element can grow or has a fixed size.") on individual columns)
- `TextEdit`
  - `integratedSearch`
  - `spellCheck` (with GBC, browser spell check is used)
- `ToolBar`
  - `toolBarTextPosition` (use [aspect](../11_user-interface/1651-toolbar-style-attributes.md) style attribute instead)
  - `iosSeparatorStretch` (use [`Toolbar.itemsAlignment`](../11_user-interface/1651-toolbar-style-attributes.md) instead)
- `Window`
  - Core:
    - `border` (as the desktop window border)
    - `errorMessagePosition`
    - `formScroll`
    - `ignoreMinimizeSetting`
    - `menuPopupPosition`
    - `messagePosition`
    - `statusBarType`
    - `position`: values `"center2"` and
      `"previous"`
    - `tabbedContainerCloseMethod`
    - `toolBarDocking`
    - `windowMenu`
    - `windowOptionMaximize`
    - `windowOptionMinimize`
    - `windowSystemMenu`
  - Action Panel:
    - `actionPanelButtonSize`
    - `actionPanelButtonSpace`
    - `actionPanelButtonTextAlign`
    - `actionPanelButtonTextHidden`
    - `actionPanelDecoration`
    - `actionPanelHAlign`
    - `actionPanelScroll`
    - `actionPanelScrollStep`
  - Ring Menu:
    - `ringMenuButtonSize`
    - `ringMenuButtonSpace`
    - `ringMenuButtonTextAlign`
    - `ringMenuButtonTextHidden`
    - `ringMenuDecoration`
    - `ringMenuHAlign`
    - `ringMenuScroll`
    - `ringMenuScrollStep`
  - Start Menu:
    - `startMenuAccelerator`
    - `startMenuExecShortcut2`
    - `startMenuShortcut`
    - `startMenuSize`
  - Mobile Specific:
    - `materialFABActionList`
    - `materialFABType`
    - `iosRenderSystemActions`
    - `iosTintColor`
    - `iosNavigationBarTextColor`
    - `iosNavigationBarTintColor`
    - `iosToolBarTintColor`
    - `iosTabBarTintColor`
    - `iosTabBarUnselectedColor`

## Changes in earlier versions

Make sure to check the upgrade notes of earlier versions, to not miss changes introduced in
maintenance releases. For more details, see [Presentation style changes in BDL
3.20](0161-presentation-styles-changes.md "Modifications to consider when using presentation styles.").

Notable changes introduced in maintenance releases:

- The [`sanitize`](0161-presentation-styles-changes.md) style attribute, also available since GBC 4.00.01.
- The [`Folder.lateRendering`](0161-presentation-styles-changes.md) style attribute, also available since GBC
  4.00.02.
