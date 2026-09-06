---
title: "Text mode rendering (TUI mode)"
source: "fgl-topics/c_fgl_DynamicUI_003.html"
breadcrumb: "User interface > User interface basics > Genero user interface modes > Text mode rendering (TUI mode)"
type: "concept"
description: "What is the text mode (TUI)? The text user interface (TUI) has been designed for character-based terminals. This mode can be used to run your application on a text terminal hardware or in a terminal ..."
---

# Text mode rendering (TUI mode)

## What is the text mode (TUI)?

The text user interface (TUI) has been designed for character-based terminals. This mode can be
used to run your application on a text terminal hardware or in a terminal emulator.

By default, Genero uses the [Graphical mode](1518-graphical-mode-rendering-gui-mode.md). In order
to run a Genero program on text mode, set the [FGLGUI](../07_configuration/0525-fglgui.md "Defines the user interface mode to be used by the program.") environment variable to 0 (zero).

On UNIX™ platforms, you need to configure your terminal
capabilities with the TERM, TERMINFO or TERMCAP environment variables.

In TUI mode, all application forms will display within the current terminal device or emulator as
shown.

![Text mode rendering screenshot](../_images/TextMode1.jpg)

*Text mode console example*

## Form specification files for TUI mode

When designing a .per form file, use a [`SCREEN`](1714-screen-section.md "The SCREEN section defines the form layout for TUI mode forms.") section to define the
layout of the form:

```
SCREEN
{
---------------------------------------------------
                   Customer form
---------------------------------------------------
 Name: [f1                            ] Id:[f2   ]
 Phone:   [f3                         ]
 Address: [f4                                    ]
          [f4                                    ]
 City:    [f5                                    ]
 Zipcode: [f6       ] State: [f7           ]
---------------------------------------------------
}
END
ATTRIBUTES
f1 = FORMONLY.cust_name;
f2 = FORMONLY.cust_id, NOENTRY;
f3 = FORMONLY.cust_phone;
f4 = FORMONLY.cust_addr;
f5 = FORMONLY.cust_city;
f6 = FORMONLY.cust_zipcode;
f7 = FORMONLY.cust_state;
END
```

Using a [`LAYOUT`](1715-layout-section.md "The LAYOUT section defines the graphical alignment of the form by using a tree of layout containers.")
section is for [GUI mode](1518-graphical-mode-rendering-gui-mode.md) only. When displaying a form
defined with `LAYOUT` section in TUI mode (FGLGUI=0), the runtime system will raise
the error [-6315](../15_library-reference/4483-genero-bdl-errors.md).

## Checking for text mode in programs

In the program code, use the [`ui.Interface.getFrontEndName()`](../15_library-reference/3103-ui-interface-getfrontendname.md "Returns the type of the front-end currently in use.") method to query for the front-end type.

When this method returns `"console"`, the program executes in text mode.

## Related links

**Related concepts**  

[Using a text terminal](1531-using-a-text-terminal.md "This section covers topics about text terminal configuration when using the TUI mode.")
