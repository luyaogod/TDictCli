---
title: "Predefined application strings"
source: "fgl-topics/c_fgl_localized_strings_011.html"
breadcrumb: "Advanced features > Localization > Localized strings > Predefined application strings"
type: "concept"
---

# Predefined application strings

> The runtime system may need to display text to the user.

For example, the runtime system library includes a report viewer, which displays a form. By
default the text is in English, and you may need to localize the text in another language. So the
strings of this component must be 'localizable', as in other application strings.

To
customize the built-in strings, the runtime system uses the mechanism
of localized strings.

All strings used by the runtime system
are centralized in a unique file:

```
$FGLDIR/src/default.str
```

which
is compiled into:

```
$FGLDIR/lib/default.42s
```

This
file is always loaded by the runtime system.

To overwrite the
defaults, you can redefine these strings in your own localized string
files.
