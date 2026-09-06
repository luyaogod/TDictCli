---
title: "Theme front calls"
source: "fgl-topics/c_fgl_frontcalls_theme.html"
breadcrumb: "Library reference > Built-in front calls > Theme front calls"
type: "concept"
---

# Theme front calls

> This section describes theme handling front calls.

This table shows the functions implemented in the "`theme`" module, to control GBC
theming.

These front calls were made available starting with GBC 1.00.47

| Function Name | Description |
| --- | --- |
| ui.Interface.frontCall("theme", "setTheme", [name], []) | Activates a specific theme. |
| ui.Interface.frontCall("theme", "getCurrentTheme", [], [result]) | Gets the active theme. |
| ui.Interface.frontCall("theme", "listThemes", [], [result]) | Lists all available themes. |

## Related links

**Related concepts**  

[Front calls](../09_advanced-features/0962-front-calls.md "Front call functions execute on the platform where the front-end is installed.")

## Child topics

- [theme.setTheme](3429-theme-settheme.md): Activates a specific theme.
- [theme.getCurrentTheme](3430-theme-getcurrenttheme.md): Gets the active theme.
- [theme.listThemes](3431-theme-listthemes.md): Lists all available themes.
