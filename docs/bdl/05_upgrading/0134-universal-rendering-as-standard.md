---
title: "Universal Rendering as standard"
source: "fgl-topics/c_fgl_Migrate_to_400_universal_rendering.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 4.00 upgrade guide > Universal Rendering as standard"
type: "concept"
---

# Universal Rendering as standard

> All front-ends (GDC, GAS, GMA, GMI) use now the GBC as rendering engine.

## GBC as unique rendering engine

Genero version 3.20 introduced the Universal Rendering concept, which could be used
as an option to display application forms on the GDC, GMA and GMI front-ends, using the GBC
front-end as rendering engine.

Starting with version 4.00, the GBC front-end becomes the only available AUI tree rendering
engine, to get an identical user interface on all type of front-ends (GDC, GMA, GMI, GAS).
> **Note:**
>
> The
> main new GUI feature of Genero 4.00 is [Responsive
> Layout](../11_user-interface/1541-responsive-layout.md "Forms can be designed to adapt to the front-end screen possibilities."), implemented by the new V4 GBC. This feature available for all type of
> front-ends.

## Feature changes

List of feature still supported, with slight change in behavior:

- On desktop (GDC), program WINDOWs display now in the same window container. See
  [Containers for program windows](../11_user-interface/1563-containers-for-program-windows.md "Program windows are displayed in window containers by the front-end.").
- On a `LAYOUT` element (form/window), the [`MINWIDTH`](../11_user-interface/1800-minwidth-attribute.md "The MINWIDTH attribute defines the minimum width of a form.")/[`MINHEIGHT`](../11_user-interface/1799-minheight-attribute.md "The MINHEIGHT attribute defines the minimum height of a form.") attributes apply only to
  windows using the `windowType="modal"` style attribute.
- With GDC 3.20 native mode, `ERROR` and `MESSAGE` notifications
  were displayed in the status bar. These are now rendered by default in a floating window by
  GBC.
- When a `TEXT` (label) of an action view
  (`BUTTON` elements, `TOPMENU` commands, `TOOLBAR`
  buttons, and so on) contains an & ampersand character, it will be hidden by the GBC. In GDC
  native rendering mode, the & ampersand could be used to underline the next character and get an
  automatic `ALT-letter` accelerator; this behavior is not
  implemented with GBC. However, to simplify migration from GDC native rendering, single &
  characters are not displayed in labels by GBC. To display a single ampersand, you must double it
  (&&).

## Deprecated UI features

List of deprecated features:

- [`ui.Interface.getUniversalClientName()`](../15_library-reference/3110-ui-interface-getuniversalclientname.md "Returns the name of the front-end used for Universal Rendering.") method.
- The [`windowresized`](../11_user-interface/1548-adapting-to-viewport-changes.md "Application forms and functions can be adapted to the front-end viewport size or mobile device orientation.") special predefined action.

## Desupported UI features

List of desupported features:

- `gui.rendering` FGLPROFILE parameter. This parameter could
  be used to enable Universal Rendering with older front-ends. It is not longer required with Genero
  V4 front-ends, since Universal Rendering is the only option.
- The `STACK`-based layout. Replacement is [Responsive Layout](../11_user-interface/1541-responsive-layout.md "Forms can be designed to adapt to the front-end screen possibilities.").
- Mobile platforms (GMA/GMI):
  - Mobile native rendering of listviews (AKA "full listviews", with two-column display and
    "embedded listviews"). As replacement, consider using a [`TABLE`](../11_user-interface/2315-table-views.md "Describes how to implement table/list views.") with the [`FLIPPED`](../11_user-interface/1780-flipped-attribute.md "The FLIPPED attribute flips TABLE columns into rows.") attribute, in conjunction with `rowAspect` style attribute set to `"list"`.
  - Splitviews controlled with parallel dialogs, parallel dialogs. Replacement: Use `HBOX`
    with `SPLIT` attribute. For more details, see [Removal of parallel dialogs / splitviews](0136-removal-of-parallel-dialogs-splitviews.md "Parallel dialog and splitview features for mobile are no longer supported.").
  - iOS specific row configuration attributes `ACCESSORYTYPE`, `DETAILACTION`.
  - The iOS specific `DISCLOSUREINDICATOR` attribute.
- The Window Container Interface (WCI) as MDI-style. See [Removal of Windows Container Interface (WCI)](0135-removal-of-windows-container-interface-wci.md "The WCI is no longer supported."). Replacement for WCI are new style attributes for
  [application windows containers](../11_user-interface/1563-containers-for-program-windows.md "Program windows are displayed in window containers by the front-end.").
- The `SPACING` attribute for [`LAYOUT`](../11_user-interface/1715-layout-section.md "The LAYOUT section defines the graphical alignment of the form by using a tree of layout containers.") section in form
  definition files.
- BDS legacy "RIP Widgets" defined in form definition files with the
  `WIDGET="type"` attribute.
- Toolbar [`BUTTONTEXTHIDDEN`](../11_user-interface/1764-buttontexthidden-attribute.md "The BUTTONTEXTHIDDEN attribute indicates that the button labels for an element are not to be displayed.") attribute: Use Toolbar/[`"aspect"`](../11_user-interface/1651-toolbar-style-attributes.md) style attribute instead.
- There is no more status bar displayed as in GDC 3.20 native mode: The [statusBarType](0143-presentation-styles-changes.md "Modifications to consider when using presentation styles.") style attribute is
  desupported.
- Desupported presentation style attributes are listed in [Presentation styles changes](0143-presentation-styles-changes.md "Modifications to consider when using presentation styles.").
- Desupported front calls are listed in [Front calls changes](0144-front-calls-changes.md "Modifications to consider when using front calls.").
- GDC native rendering specific "local actions" (with no VM-side action handler):
  `editcut`, `editcopy`, `editpaste`,
  `nextfield`, `prevfield`, `nextpage`,
  `prevpage`, `nexttab`, `prevtab`. Note that the locale
  actions are marked as [deprecated since Genero
  3.10](0188-gdc-local-actions.md "The concept of \"Local Actions\" is now deprecated.").

## Related links

**Related concepts**  

[Graphical mode rendering (GUI mode)](../11_user-interface/1518-graphical-mode-rendering-gui-mode.md "Graphical mode rendering (GUI mode)")
