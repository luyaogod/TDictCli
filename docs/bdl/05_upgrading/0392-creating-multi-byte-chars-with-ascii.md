---
title: "Creating multi-byte chars with ASCII"
source: "fgl-topics/c_fgl_MigI4GL_049.html"
breadcrumb: "Upgrading > Migrating from IBM® Informix® 4GL to Genero BDL > 4GL programming topics > Creating multi-byte chars with ASCII"
type: "concept"
---

# Creating multi-byte chars with ASCII

> I4GL allows to create multi-byte characters with the ASCII operator. This is not supported by Genero BDL.

With IBM® Informix® 4GL, when defining the `CLIENT_LOCALE` environment variable for a
multi-byte character set like BIG5, it is possible to create a multi-byte character by concatenating
the bytes produced by the the ASCII operator.

With Genero BDL, the ASCII operator returns only valid characters, in the [current application encoding](0344-localization-support-in-genero.md "I4GL and Genero BDL use different libraries and environment variables in support of localization.").

For example:

```
MAIN
    DEFINE var VARCHAR(10)
    LET var[1] = ASCII(195)
    LET var[2] = ASCII(169)
    LET var[3] = ASCII(195)
    LET var[4] = ASCII(160)
    DISPLAY var
END MAIN
```

With I4GL (using UTF-8 multi-byte locale - make sure your terminal is correctly
configured):

```
$ export CLIENT_LOCALE="en_us.utf8"
$ c4gl -o m.bin m.4gl && ./m.bin
éà
```

With Genero BDL:

```
$ export LC_ALL="en_US.utf8"
$ fglrun -i
Charmap      : UTF-8
Multibyte    : yes
Stateless    : yes
Length Semantics : BYTE
$ fglcomp m.4gl && fglrun m.42m
      (spaces)
```

Creating multi-byte characters by concatenating single bytes is bad practice and should be
reviewed.

## Related links

**Related concepts**  

[Application locale](../09_advanced-features/0864-application-locale.md "The application locale defines the language and codeset for your application.")
