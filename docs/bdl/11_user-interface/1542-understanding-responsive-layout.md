---
title: "Understanding Responsive Layout"
source: "fgl-topics/c_fgl_form_responsive_intro.html"
breadcrumb: "User interface > Form definitions > Form rendering > Responsive Layout > Understanding Responsive Layout"
type: "concept"
---

# Understanding Responsive Layout

> Forms can be designed to adapt to the front-end screen possibilities.

Screens of mobile devices (smart phones and tablets), as well as web browser windows on a desktop
computer, can have various display area sizes.

Modern applications should adapt automatically to the device type, display area size and screen
orientation.

On a small screen or browser window, only the main information should be visible, and when a
larger display area is available, more detailed information should be visible. On small screens, the
end user must have options to access the detailed information, typically through touchscreen
gestures.

In responsive layout, the concept of screen size refers to the available width of the [mobile device screen](../17_mobile-applications/5081-mobile-applications.md "These topics cover programming subjects about mobile applications") or [desktop window container](1563-containers-for-program-windows.md "Program windows are displayed in window containers by the front-end."), that can be used to
render application forms. The height of the screen/window is secondary.

For example, a given form can be rendered on a large screen as follows, showing all sub-groupd of
fields since there is enough room:

![Diagram of a responsive form on a large screen](../_images/rl_concept_large_screen.jpg)

*Responsive form on a large screen*

The same form would automatically adapt for smaller screens, by hiding secondary field groups
and reorganizing others vertically:

![Diagram of a responsive form on a small screen](../_images/rl_concept_small_screen.jpg)

*Responsive form on a small screen*

In Genero BDL, elements of form files can use layout attributes that will define the rendering of
the elements, depending on the screen/window size, by using the
`@screen-size` modifier. The possible screen sizes are referenced
by abstract names: `SMALL` and `MEDIUM` for smartphones and tablets,
`LARGE` for desktop computer screens. This abstract size can also reflect the screen
orientation on any device type.

For example, to define that a `GROUP` must be hidden on small screens,
use:

```
GROUP g1: group1, ... HIDDEN@SMALL ;
```

Form elements can also adapt to the width of the screen/window, by stretching to the available
width:

```
EDIT f01 = customer.cust_name, ... STRETCH=X ;
```
