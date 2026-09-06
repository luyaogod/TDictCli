---
title: "Keyboard access"
source: "fgl-topics/c_fgl_Accessibility_002.html"
breadcrumb: "User interface > Form definitions > Accessibility guidelines > Keyboard access"
type: "concept"
---

# Keyboard access

> How to implement keyboard usage to follow accessibility standard?

## Defining keyboard accelerators for every action

Since a mouse or other pointing devices may not be used
by people with reduced vision, an accessible application
must be usable with the keyboard alone. Therefore, all the
possible actions that could be triggered by a user must have
a keyboard shortcut.

We strongly suggest that you define
consistent keyboard shortcuts for all actions through
the use of action defaults. Developers can avoid overriding the system
default shortcuts by checking the target platform guidelines,
especially for system shortcuts that trigger accessible
actions (for example, Ctrl-Shift-Enter, which triggers spoken information
about the currently selected item). Overriding system
shortcuts is generally a bad practice, even for non-accessible
applications, although overriding may be unavoidable due
to compatibility issues.

## Keyboard focus and action views

Generally, keyboard navigation in an application may be easier
if you keep the `MENU` actions
in the menu frame; the actions can have the keyboard
focus and the user can navigate through them using the up and down
arrows.

Avoid using [toolbars](1855-toolbars.md "Toolbars define a bar of buttons that appears at the top of application forms.") alone in accessible
applications, because toolbars by default are not accessible using the keyboard. Toolbars cannot
have the keyboard focus, and there is no way to navigate through all toolbar items or to activate
one of them using the keyboard. If you do use toolbars, provide keyboard shortcuts and duplicate
them in a topmenu.

## Related links

**Related concepts**  

[Configuring actions](2259-configuring-actions.md "Action attributes related to decoration, keyboard shortcuts, and behavior can be defined with action attributes.")
