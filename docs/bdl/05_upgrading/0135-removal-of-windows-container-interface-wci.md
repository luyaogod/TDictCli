---
title: "Removal of Windows Container Interface (WCI)"
source: "fgl-topics/c_fgl_Migrate_to_400_wci_desupport.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 4.00 upgrade guide > Removal of Windows Container Interface (WCI)"
type: "concept"
---

# Removal of Windows Container Interface (WCI)

> The WCI is no longer supported.

Starting with version 4.00, the Windows Container Interface (WCI) is desupported. This feature
was only available with the GDC front-end in native rendering mode.

The related `ui.Interface` methods to defined container/child relationship still
exist, but are now deprecated, namely: [`setType()`](../15_library-reference/3123-ui-interface-settype.md "Defines the type of the program for the front-end."), [`getType()`](../15_library-reference/3109-ui-interface-gettype.md "Returns the type of the program."), [`setContainer()`](../15_library-reference/3118-ui-interface-setcontainer.md "Define the parent container for the current program."), [`getContainer()`](../15_library-reference/3101-ui-interface-getcontainer.md "Get the parent container of the current program."), [`getChildCount()`](../15_library-reference/3099-ui-interface-getchildcount.md "Get the number of children in a parent container."), [`getChildInstances()`](../15_library-reference/3100-ui-interface-getchildinstances.md "Get the number of child instances for a given program name."). Programs using these APIs to define a WCI application
can keep the code: The Window Container Interface (aka MDI) rendering is no longer available, but
the `getChildCount()` and `getChildInstance()` methods still return
the child counts, except when `browserMultiPage=yes` (with GAS/GBC) or when
`desktopMultiWindow=yes` (with GDC/UR/GBC): In this case the methods return
zero.

As a replacement for WCI, Genero provides the following user interface options:

- When using the GAS and a web browser, the windows of all applications display by default in the
  same web browser tab, and the sidebar on the left allows you to toggle between applications. This
  can be controlled with the [`UserInterface/browserMultiPage`](../11_user-interface/1652-userinterface-style-attributes.md) style attribute.
- When using the GDC front-end on desktop, the [`UserInterface/desktopMultiWindow`](../11_user-interface/1652-userinterface-style-attributes.md) style attribute can be defined to render
  the windows of all applications in the same desktop window container, or in dedicated window
  containers.

See [Containers for program windows](../11_user-interface/1563-containers-for-program-windows.md "Program windows are displayed in window containers by the front-end.") for more details.

## Related links

**Related concepts**  

[Containers for program windows](../11_user-interface/1563-containers-for-program-windows.md "Program windows are displayed in window containers by the front-end.")
