---
title: "Character set mapping"
source: "fgl-topics/c_fgl_JavaBridge_033.html"
breadcrumb: "Extending the language > The Java interface > Advanced programming > Character set mapping"
type: "concept"
description: "Application programs use a given locale and character set , while Java uses its own charset internally for the char Java type (16- bit UNICODE). When passing character strings to/from Java methods or ..."
---

# Character set mapping

Application programs use a given [locale and character set](../09_advanced-features/0879-defining-the-application-locale.md "This section describes the settings defining the application locale, changing the behavior of the compilers and runtime system."),
while Java uses its own charset internally for the `char`
Java type (16- bit UNICODE).

When passing character strings to/from Java methods or when
assigning program strings to `java.lang.String`,
the runtime system handles character set conversion.
