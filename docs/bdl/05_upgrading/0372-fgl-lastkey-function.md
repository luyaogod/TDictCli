---
title: "fgl_lastkey() function"
source: "fgl-topics/c_fgl_MigI4GL_059.html"
breadcrumb: "Upgrading > Migrating from IBM® Informix® 4GL to Genero BDL > User interface topics > fgl_lastkey() function"
type: "concept"
---

# fgl_lastkey() function

> In a given context and for a given key press, FGL fgl_lastkey() API does not return the same key numbers than I4GL fgl_lastkey().

With IBM® Informix® 4GL, the `fgl_lastkey()` always returns the exact last key pressed by
the user, even when the key triggers an operation in a dialog instruction, such as when inserting a
new row during an `INPUT ARRAY`.

With Genero BDL, an operation in a dialog instruction is triggered by an action: To insert a new
row in an `INPUT ARRAY`, the "insert" action is fired. Actions are an abstract
concept in Genero, and can be configured with accelerator keys. When pressing a key, only the
corresponding action is known by FGL. Therefore, with the Genero BDL design, it is impossible to get
the exact same key numbers than I4GL `fgl_lastkey()`.

For example, when defining the `CONTROL-T` key for the `INSERT ROW`
option for an `INPUT ARRAY`:

```
DATABASE formonly
SCREEN
{
[f001]
[f001]
[f001]
[f001]
[f001]
}
ATTRIBUTES
f001 = formonly.col1;
END
INSTRUCTIONS
SCREEN RECORD sr[5](formonly.*);
END
```

```
MAIN
    DEFINE p_arr ARRAY [10] OF RECORD
             col1 CHAR(10)
           END RECORD
    DEFINE k1, k2 INT

    OPTIONS INSERT KEY CONTROL-T

    LET p_arr[1].col1 = "AAAA"
    LET p_arr[2].col1 = "BBBB"

    OPEN FORM f_test1 FROM "getkey"
    DISPLAY FORM f_test1

    INPUT ARRAY p_arr WITHOUT DEFAULTS FROM sr.* ATTRIBUTE(COUNT = 2)
    BEFORE INSERT
           LET k1 = fgl_lastkey()
           DISPLAY "BEFORE INSERT fgl_lastkey:", k1 AT 12,5
           LET k1 = fgl_getkey()
           LET k2 = fgl_lastkey()
           DISPLAY "CALL fgl_getkey:", k1, " fgl_lastkey:", k2 AT 13,5
    END INPUT

END MAIN
```

When executing the code, if you press the CONTROL-T key twice, I4GL will
display:

```
    BEFORE INSERT fgl_lastkey:2014
    CALL fgl_getkey:20 fgl_lastkey:20
```

While FGL will display the following, showing the
key numberb 2014 for fgl\_lastkey:

```
    BEFORE INSERT fgl_lastkey:2014
    CALL fgl_getkey:20 fgl_lastkey:2014
```

## Related links

**Related concepts**  

[fgl\_lastkey()](../15_library-reference/2772-fgl-lastkey.md "Returns the key code corresponding to the logical key that the user most recently typed in the form.")
