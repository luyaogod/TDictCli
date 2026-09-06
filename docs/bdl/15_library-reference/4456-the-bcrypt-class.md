---
title: "The BCrypt class"
source: "fgl-topics/c_gws_SecurityBCrypt.html"
breadcrumb: "Library reference > Extension packages > The security package > The BCrypt class"
type: "concept"
---

# The BCrypt class

> The security.BCrypt class lets you save passwords as BCrypt results instead of clear text.

This class is provided in the `security` [C-Extension](../14_extending-the-language/2703-c-extensions.md "With C-Extensions, you can bind your own C libraries in the runtime system, to call C function from the application code.") library; To use this class, import the
`security` package with:

```
IMPORT security
```

If you need to store passwords on a database for instance, you can save them as BCrypt results
instead of clear text. This makes them difficult to hack, as the time to generate one is
expensive.

## Child topics

- [security.BCrypt methods](4457-bcrypt-methods.md): Methods of the security.BCrypt class.
- [Example: Using security.BCrypt methods](4461-bcrypt-example.md): This example creates (and checks) a hash password as BCrypt results.
