---
title: "theme.listThemes"
source: "fgl-topics/c_fgl_frontcall_theme_listthemes.html"
breadcrumb: "Library reference > Built-in front calls > Theme front calls > theme.listThemes"
type: "concept"
---

# theme.listThemes

> Lists all available themes.

## Syntax

```
ui.Interface.frontCall("theme", "listThemes",
   [], [result])
```

1. result - A string containing a JSON array of the available themes. For each
   theme, the name, title, and conditions are listed. The conditions are contained within a second JSON
   array, and define in which contexts the theme can be used.

## Usage

The `listThemes` front call returns a JSON array of the GBC themes available to
the running application.

The returned value is a JSON formatted array containing JSON objects with name, title and
conditions of each available GBC theme.

Note that the "`listThemes`" front call only returns the themes that are valid in
the current conditions of the application display context. Additional GBC themes may be available,
when other conditions are satisfied.

For example, consider a scenario where application forms are rendered on a desktop machine using
the GDC with Universal Rendering, and where only the "`default`" and
"`highcontrast`" themes were made available for the desktop. In this scenario, the
front call would return the following JSON string (line breaks added for
readability):

```
{{[{"name":"default","title":"Default","conditions":["isUR","isDesktop"]},
 {"name":"highcontrast","title":"High contrast","conditions":["isUR","isDesktop"]}]}}
```

This JSON array can be easily converted with [`util.JSON.parse()`](3582-util-json-parse.md "Parses a JSON string and fills program variables with the values.") to a dynamic array defined with the following
type:

```
TYPE t_themes DYNAMIC ARRAY OF RECORD
    name STRING,
    title STRING,
    conditions DYNAMIC ARRAY OF STRING
END RECORD
...
    DEFINE result STRING
    DEFINE themes t_themes
    CALL ui.Interface.frontcall("theme", "listThemes", [], [result])
    CALL util.JSON.parse(result, themes)
```

The `name` is the key to identify a GBC theme, it can for example be called with
[`theme.setTheme`](3429-theme-settheme.md "Activates a specific theme.").

The `title` is the label of the GBC theme as seen by the end user.

The `conditions` defines a list of conditions where the GBC theme can apply. For
example, the "`isBrowser`" condition indicates that the theme can be used when
application forms are displayed in a browser.

For more details about GBC Themes, see "Theme reference" chapter in the Genero Browser Client User Guide.

## Example

main.4gl

```
IMPORT util

TYPE t_themes DYNAMIC ARRAY OF RECORD
    name STRING,
    title STRING,
    conditions DYNAMIC ARRAY OF STRING
END RECORD

MAIN

    DEFINE themes t_themes
    DEFINE rec RECORD
        curr_theme STRING,
        result STRING
    END RECORD

    OPEN FORM f1 FROM "form"
    DISPLAY FORM f1

    INPUT BY NAME rec.* ATTRIBUTES(UNBUFFERED)
        BEFORE INPUT
            CALL ui.Interface.frontcall("theme", "listThemes",
                                        [], [rec.result])
            CALL util.JSON.parse(rec.result, themes)
            CALL init_combobox_themes(themes)
            CALL ui.Interface.frontcall("theme", "getCurrentTheme",
                                        [], [rec.curr_theme])
        ON CHANGE curr_theme
            CALL ui.Interface.frontcall("theme", "setTheme",
                                        [rec.curr_theme], [])
        ON ACTION get_curr_theme
            CALL ui.Interface.frontcall("theme", "getCurrentTheme",
                                        [], [rec.curr_theme])
    END INPUT

END MAIN

FUNCTION init_combobox_themes(tl t_themes) RETURNS()
    DEFINE x INTEGER
    DEFINE cb ui.ComboBox
    LET cb = ui.ComboBox.forName("formonly.curr_theme")
    CALL cb.clear()
    FOR x = 1 TO tl.getLength()
        CALL cb.addItem(tl[x].name, tl[x].title)
    END FOR
END FUNCTION
```

form.per

```
LAYOUT
GRID
{
Current Theme  [f1                         |b1    ]
Result of theme.listThemes:
[f2                                               ]
[                                                 ]
}
END
END

ATTRIBUTES
COMBOBOX f1 = FORMONLY.curr_theme, NOT NULL,
  COMMENT="Change the current GBC theme";
BUTTON b1: get_curr_theme, TEXT="Get current theme",
  COMMENT="Get the current GBC theme after change in GBC settings";
TEXTEDIT f2 = FORMONLY.result, STRETCH=BOTH;
END
```
