---
title: "Why do I need to care about the locale and character set?"
source: "fgl-topics/c_fgl_localization_005.html"
breadcrumb: "Advanced features > Localization > Application locale > Locale and character set basics > Why do I need to care about the locale and character set?"
type: "concept"
description: "If you don't know what you are doing with character sets, the end user might get strange characters displayed on the screen, and will probably not be able to input non-ASCII characters. In the worst ..."
---

# Why do I need to care about the locale and character set?

If you don't know what you are doing with character sets, the end user might get strange
characters displayed on the screen, and will probably not be able to input non-ASCII characters.

In the worst case, as character set conversion can be symmetric for single-byte character sets,
the end user might see correct characters on the workstation, but on the back-end you can get
invalid characters in the database files.

By upgrading to a newer OS, Genero Business Development Language runtime or database system, or
if a character set mapping utility was used somewhere in the chain, you can even get mixed character
encoding in the database files.

## Related links

**Related concepts**  

[Understanding locale settings](0865-understanding-locale-settings.md "This is an introduction to application locale definition.")
