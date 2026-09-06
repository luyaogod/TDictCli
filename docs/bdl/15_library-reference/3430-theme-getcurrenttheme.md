---
title: "theme.getCurrentTheme"
source: "fgl-topics/c_fgl_frontcall_theme_getcurrenttheme.html"
breadcrumb: "Library reference > Built-in front calls > Theme front calls > theme.getCurrentTheme"
type: "concept"
---

# theme.getCurrentTheme

> Gets the active theme.

## Syntax

```
ui.Interface.frontCall("theme", "getCurrentTheme",
   [], [result])
```

1. result - The name of the active theme.

## Usage

The "`getCurrentTheme`" front call allows the application to identify the GBC
theme that is currently used to display application forms.

For a complete example of theme front call usage, see the [`themes.listThemes` front call code example](3431-theme-listthemes.md).
