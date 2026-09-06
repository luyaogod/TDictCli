---
title: "MONEY type and CLIENT_LOCALE"
source: "fgl-topics/c_fgl_MigI4GL_042.html"
breadcrumb: "Upgrading > Migrating from IBM® Informix® 4GL to Genero BDL > 4GL programming topics > MONEY type and CLIENT_LOCALE"
type: "concept"
---

# MONEY type and CLIENT_LOCALE

> I4GL adapts the scale of MONEY / MONEY(p) types from the client locale setting, Genero BDL always uses the same default scale.

With IBM® Informix® 4GL, when using the `MONEY` or
`MONEY(p)` data type, the number of decimal digits (scale) adapts
the to current locale settings defined by the CLIENT\_LOCALE environment variable.

For example, with `CLIENT_LOCALE=ja_jp.utf8`, a `MONEY(10)` will be
interpreted as a `MONEY(10,0)`.

With Genero BDL, the default scale of a `MONEY` type is always 2.

For example:

```
MAIN
    DEFINE m MONEY(10)
    DEFINE d DECIMAL(10,2)
    LET m = 123.45
    DISPLAY "m = ", m
    LET d = m
    DISPLAY "d = ", d
END MAIN
```

With I4GL, the `MONEY(10)` type is adapted at compile time
to `MONEY(10,2)` in the English/US locale, or to `MONEY(10,0)` in the
Japanese locale:

```
$ export CLIENT_LOCALE="en_us.utf8"
$ c4gl -o money.bin money.4gl
$ ./money.bin
m =       $123.45
d =       123.45

$ export CLIENT_LOCALE="ja_jp.utf8"
$ ./money.bin
m =          ¥123
d =       123.45

$ c4gl -o money.bin money.4gl
$ ./money.bin
m =          ¥123
d =       123.00
```

With FGL, `MONEY(10)` always converts to
`MONEY(10,2)`:

```
$ export CLIENT_LOCALE="en_us.utf8"
$ fglcomp money.4gl
$ fglrun money.42m
m =       $123.45
d =       123.45

$ export CLIENT_LOCALE="ja_jp.utf8"
$ fglrun money.42m
m =       $123.45
d =       123.45
```

## Related links

**Related concepts**  

[MONEY(p,s)](../08_language-basics/0564-money-p-s.md "The MONEY data type is provided to store currency amounts with exact decimal storage.")
