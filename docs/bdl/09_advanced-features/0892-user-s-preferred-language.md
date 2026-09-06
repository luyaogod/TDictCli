---
title: "User's preferred language"
source: "fgl-topics/c_fgl_localization_fe_lang.html"
breadcrumb: "Advanced features > Localization > Application locale > User's preferred language"
type: "concept"
---

# User's preferred language

> An application can get the user's preferred language and territory as configured on the front-end platform.

The user preferred language can, for example, be useful in selecting appropriate content based on
language preferences, and in starting other programs by [setting the expected application locale](0879-defining-the-application-locale.md "This section describes the settings defining the application locale, changing the behavior of the compilers and runtime system.").

To get the user preferred language as defined in the front-end, perform a
`standard.feInfo` front call with the `userPreferredLang`
option:

```
PRIVATE DEFINE fe_lang STRING

FUNCTION get_fe_lang()
  IF fe_lang IS NULL THEN
     CALL ui.Interface.frontCall( "standard", "feInfo", ["userPreferredLang"], [fe_lang] )
  END IF
  RETURN fe_lang
END FUNCTION
```

The front-end locale configuration depends on the type of front-end:

- With Genero Application Server (GAS), the front-end locale is defined in the web browser
  preferences.
- With Genero Desktop Client (GDC), by default the front-end locale is defined by the operating
  system language settings. It can be configured with a GDC option.
- With Genero Mobile for Android (GMA) and Genero Mobile for iOS (GMI), the front-end locale is
  defined by the device language settings.

The format of the returned value
is:

```
language_territory
```

For example, when
running the GDC front-end on a Linux®
platform with LC\_ALL defined as `en_US.utf8`, the front call will
return:

```
en_US
```

## Related links

**Related concepts**  

[standard.feInfo](../15_library-reference/3399-standard-feinfo.md "Queries general front-end properties.")
