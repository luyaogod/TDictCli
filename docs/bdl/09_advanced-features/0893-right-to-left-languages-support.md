---
title: "Right-to-left languages support"
source: "fgl-topics/c_fgl_localization_semitic.html"
breadcrumb: "Advanced features > Localization > Application locale > Right-to-left languages support"
type: "concept"
---

# Right-to-left languages support

> Genero supports right-to-left languages, such as Arabic and Hebrew.

## Right-to-left mode

For specific front-end clients, Genero supports right-to-left languages with the reverse mode.
With reverse mode, all forms are mirrored and the text direction changes to right-to-left, for
display and input.

The reverse mode is enabled at runtime. The same form files are used to display in the default
left-to-right and right-to-left mode.

Right-to-left display is implicit on mobile devices. It is enabled depending on language settings
on the device.

## Application locale

Reverse mode is used with an application locale that defines a language written from right to
left.

For example, for Arabic support on a Linux® platform, you can use the following LC\_ALL value when using the ISO-8859-6
codeset:

```
$ export LC_ALL=ar_DZ.iso88596
```

## Reverse mode configuration

To enable reverse mode, set the `reverse` style attribute for the
`UserInterface` class to `"yes"` in your .4st
file:

```
<StyleList>
  ...
  <Style name="UserInterface">
     <StyleAttribute name="reverse" value="yes" />
  </Style>
  ...
</StyleList>
```

## Related links

**Related reference**  

[UserInterface style attributes](../11_user-interface/1652-userinterface-style-attributes.md "UserInterface presentation style attributes define general options related to the application user interface.")
