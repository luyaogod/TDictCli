---
title: "theme.setTheme"
source: "fgl-topics/c_fgl_frontcall_theme_settheme.html"
breadcrumb: "Library reference > Built-in front calls > Theme front calls > theme.setTheme"
type: "concept"
---

# theme.setTheme

> Activates a specific theme.

## Syntax

```
ui.Interface.frontCall("theme", "setTheme",
   [name], [])
```

1. name - The name of the theme to be activated.

## Usage

The "`setTheme`" front call allows the application to change the GBC theme used to
render the forms of the application.

The list of available themes can be retrieved with the [`themes.listThemes`](3431-theme-listthemes.md "Lists all available themes.") front
call.

For a complete example of theme front call usage, see the [`themes.listThemes` front call code example](3431-theme-listthemes.md).
