---
title: "Containers for program windows"
source: "fgl-topics/c_fgl_windows_and_forms_containers.html"
breadcrumb: "User interface > Form definitions > Windows and forms > Containers for program windows"
type: "concept"
---

# Containers for program windows

> Program windows are displayed in window containers by the front-end.

## Program windows and window container

A Genero program can create multiple [windows](1564-the-window-concept.md "Windows are containers for .42f forms.") with the `OPEN WINDOW` instruction.

The front-end uses a window container to display the program windows.

- On desktop (GDC), the window container is a frame of the OS window manager.
- In a web browser (GAS), the window container is a browser tab.
- On mobile devices (GMA/GMI), the window container is the device screen.

All windows created by a given program are displayed in the same window container.
The window visible to the user will be the [current window](1564-the-window-concept.md "Windows are containers for .42f forms."). If the current window is a modal
window, it will appear in a frame on top of the last displayed frame of a normal
window.

When multiple programs display their windows/forms in the same window container, on
the left of the container, the application list can be used to navigate between these
programs.

The chromebar is the top bar with a blue background. It can contain action views
(buttons) to trigger actions and system buttons when GBC debug mode is active. By default, the
chromebar is not visible when using the GDC desktop front-end, but it can be activated by a GBC
theme variable. The chromebar is visible by default when using a browser (GAS) and with mobile
devices (GMA/GMI).

To switch between normal windows created by a given program, use the window list,
displayed on the top left in the chromebar.

Note that by default, with the GDC front-end, no chromebar is visible: It needs to be enabled
with a GBC theme variable (refer to the GBC manual for more details).

![Window container screenshot](../_images/desktop_window_container.jpg)

*Desktop window container and related areas*

## Shared or dedicated window containers for programs

With both GDC/desktop and GAS/web browser, when starting a new child program with a
`RUN` command not using the `WITHOUT WAITING` clause, the parent
program is frozen and is waiting for the child program to finish. In such case, the child program
windows are always displayed in the same window container.

When using the GDC on desktop, by default, new programs started with `RUN ... WITHOUT
WAITING` will display in the same window container as the parent program. To
force each program to display its windows in a dedicated desktop window container,
define the [`desktopMultiWindow`](1652-userinterface-style-attributes.md) style attribute to `"yes"` at the
`UserInterface` node level:

```
<StyleList>
  <Style name="UserInterface">
     <StyleAttribute name="desktopMultiWindow" value="yes"/>
  </Style>
  ...
</StyleList>
```

When using the GAS with a web browser, by default, new programs started with `RUN ...
WITHOUT WAITING` will display in the same web browser tab as the parent program. To force
each child program to display its windows in a dedicated web browser tab, define the [`browserMultiPage`](1652-userinterface-style-attributes.md) style attribute to `"yes"` at the
`UserInterface` node level:

```
<StyleList>
  <Style name="UserInterface">
     <StyleAttribute name="browserMultiPage" value="yes"/>
  </Style>
  ...
</StyleList>
```

## The application list

The application list shows the set of programs running concurrently and displaying
in the front-end window container, these can be child programs started with `RUN
WITHOUT WAITING`.

![Application list screenshot](../_images/gbc_application_list.jpg)

*Desktop window container with application list*

With GAS and a web browser, the browser tabs display titles of regular windows, using the
`windowType="normal"` (default) style attribute: Titles of windows defines as modal
with `windowType="modal"` style attribute are not displayed in the browser tab
caption.

By default, the application list is visible to the end user in the
SideBarRail. It can be controlled with the following `UserInterface`
style attribute [applicationListVisible](1652-userinterface-style-attributes.md) and [applicationListPosition](1652-userinterface-style-attributes.md):

```
<StyleList>
  <Style name="UserInterface">
     <StyleAttribute name="applicationListPosition" value="top"/>
  </Style>
  ...
</StyleList>
```

![Application list on top screenshot](../_images/gbc_application_list_top.jpg)

*Desktop window container with application list on top*

The application list can also be shown in the topmenu and in the toolbar of the
current window/form, by using the `AUTOCOMMANDS` and `AUTOITEMS`
placeholders, with the `CONTENT=PROGRAMS` attribute. These placeholders can be
located anywhere in the topmenu and toolbar. For more details, see [Automatic action views](2298-automatic-action-views.md "Action views can be rendered automatically in some form elements.").

## The window list

The window list is defined by the set of windows opened by the current program
displayed in the window container.

When the parent application launches multiple child applications and waits for them to complete
(`RUN` not using `WITHOUT WAITING`), the window list shows also
windows from those child applications.

The window list can be accessed from the chromebar (no available by default on GDC
desktop):

![Window list screenshot](../_images/gbc_window_list.jpg)

*Desktop window container with window list open*

By default, the window list is displayed on the top left in the chromebar. It can be
hidden by using the following UserInterface style attribute:

```
<StyleList>
  <Style name="UserInterface">
     <StyleAttribute name="windowListVisible" value="no"/>
  </Style>
  ...
</StyleList>
```

For reference, see [windowListVisible](1652-userinterface-style-attributes.md).

The window list can also be shown in the topmenu and in the toolbar of the current
window/form, by using the `AUTOCOMMANDS` and `AUTOITEMS` placeholders,
with the `CONTENT=WINDOWS` attribute. These placeholders can be located anywhere in
the topmenu and toolbar. For more details, see [Automatic action views](2298-automatic-action-views.md "Action views can be rendered automatically in some form elements.").

## RUN command usage and effects

The [`RUN`](../09_advanced-features/0830-run.md "The RUN instruction executes the command passed as argument.") command can be used to chain
application (program) execution.

In order to handle the application/program chain properly (to display the correct application
list hierarchy), the front-end needs to be aware of the child program execution, before executing
other sub-programs that display GUI. When executing a child program that is interactive, the parent
program must make sure to be connected to the front-end, before executing the `RUN`
command to start another interactive program.

For example, program `A` is started and is connected to the front-end. This root
program starts a sub-program `B`, which does not do any user-interaction, before
starting the third sub-program `C`, which opens a form and executes a dialog. Since
`B` did not register to the front-end, it will consider `C` as if it
was started individually and miss the `A->B->C` chain.

The solution in such case is to force the intermediate program to connect to the front-end:

```
-- This is prog-B
MAIN
    CALL ui.Interface.refresh()
    RUN "fglrun prog-C"
END MAIN
```

## Initial window size (GDC/desktop)

When using the GDC front-end (on a desktop screen), the initial size of the window container can
be defined by the [`ui.Interface.setSize()`](../15_library-reference/3121-ui-interface-setsize.md "Specify the initial size of the window container.") method, for example as a % of the display
screen:

```
MAIN
    CALL ui.Interface.setSize("30%","80%")
    OPEN FORM f1 FROM "main_form"
    DISPLAY FORM f1
    ...
END MAIN
```

This method is only used to specify the initial size of the window container: When
stored settings apply, the size is taken from the stored settings.

## Tabbed containers (deprecated)

> **Important:**
>
> `tabbedContainer` is deprecated and replaced by the new user interface elements to
> control application windows. See [Containers for program windows](1563-containers-for-program-windows.md "Program windows are displayed in window containers by the front-end.") for more
> details.

Instead of using the `tabbedContainer=yes` style attribute, use [`applicationListPosition=top`](1652-userinterface-style-attributes.md).

For reference, see [`tabbedContainer`](1655-window-style-attributes-basics.md) in the `Window` style attributes topic.

## Related links

**Related concepts**  

[Position and size of a window](1566-position-and-size-of-a-window.md "Window objects can be created with a position and size for the TUI mode.")

[Presentation styles](1607-presentation-styles.md "Use presentation styles to specify decoration attributes for window and form elements.")
