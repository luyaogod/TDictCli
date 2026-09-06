---
title: "Example 1: URL-based web component using Wikipedia"
source: "fgl-topics/c_fgl_webcomponent_url_example_1.html"
breadcrumb: "User interface > User interface programming > Web components > Using a URL-based web component > Examples > Example 1: URL-based web component using Wikipedia"
type: "concept"
description: "This example shows how to use a WEBCOMPONENT field displaying Wikipedia pages. The form file: webcomp.per LAYOUT GRID { [f1 ] [ ] [ ] [ ] [ ] [ ] [ ] [ ] [f2 ] [f3 ] [ ] } END END ATTRIBUTES ..."
---

# Example 1: URL-based web component using Wikipedia

This example shows how to use a `WEBCOMPONENT` field displaying Wikipedia
pages.

The form file: webcomp.per

```
LAYOUT
GRID
{
[f1                     ]
[                       ]
[                       ]
[                       ]
[                       ]
[                       ]
[                       ]
[                       ]
[f2                     ]
[f3                     ]
[                       ]
}
END
END
ATTRIBUTES
WEBCOMPONENT f1 = FORMONLY.mywebcomp, STRETCH=BOTH;
BUTTONEDIT f2 = FORMONLY.topic, ACTION=load;
TEXTEDIT f3 = FORMONLY.value, STRETCH=X;
END
```

The program file: webcomp.4gl

```
CONSTANT c_wikipedia = "https://en.wikipedia.org/wiki"

DEFINE rec RECORD
           mywebcomp STRING,
           topic STRING,
           value STRING
       END RECORD

MAIN

    OPEN FORM f1 FROM "webcomp"
    DISPLAY FORM f1

    LET rec.topic = "Cat"

    INPUT BY NAME rec.* WITHOUT DEFAULTS
          ATTRIBUTES(UNBUFFERED)
        BEFORE INPUT
           CALL sync_webcomp_value()
        ON ACTION load ATTRIBUTE(ACCELERATOR="F5")
           CALL sync_webcomp_value()
        ON CHANGE topic
           CALL sync_webcomp_value()
        ON CHANGE mywebcomp
           LET rec.value = rec.mywebcomp
           MESSAGE "URL has changed! "||CURRENT HOUR TO FRACTION(3)
    END INPUT

END MAIN

PRIVATE FUNCTION sync_webcomp_value()
    LET rec.mywebcomp = c_wikipedia||"/"||rec.topic
    LET rec.value = rec.mywebcomp
END FUNCTION
```

## Related links

**Related concepts**  

[Defining a URL-based web component in forms](2386-defining-a-url-based-web-component-in-forms.md "A URL-based web component needs to be declared in the form definition.")
