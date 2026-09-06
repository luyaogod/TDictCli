---
title: "Spacer items in hbox tags"
source: "fgl-topics/c_fgl_form_rendering_hbox_tags_spacer.html"
breadcrumb: "User interface > Form definitions > Form rendering > Grid-based layout > Using hbox tags to align form items > Spacer items in hbox tags"
type: "concept"
description: "HBox tags also introduces the spacer items concept: when a grid hbox is created, the content may be smaller than the container. LAYOUT GRID { [a :b :c ] [ :d :e :f ] [g : :h ] [i: :j: :k : :l ] } END ..."
---

# Spacer items in hbox tags

HBox tags also introduces the spacer items concept: when a grid hbox is created, the content
may be smaller than the container.

```
LAYOUT
GRID
{
[a      :b      :c       ]
[ :d    :e      :f       ]
[g      :     :h         ]
[i: :j: :k    : :l       ]
}
END
END
ATTRIBUTES
EDIT a = FORMONLY.a;
EDIT b = FORMONLY.b;
EDIT c = FORMONLY.c;
EDIT d = FORMONLY.d;
EDIT e = FORMONLY.e;
EDIT f = FORMONLY.f;
EDIT g = FORMONLY.g;
EDIT h = FORMONLY.h;
EDIT i = FORMONLY.i;
EDIT j = FORMONLY.j;
EDIT k = FORMONLY.k;
EDIT l = FORMONLY.l;
END
```

Because of the checkbox, cell 1 is very large, and then the hbox gets larger than the three
fields it contains. A spacer item object is automatically created by the form compiler; the role of
the spacer item is to take all the free space in the container. Then all the widgets are packed to
the left side of the hbox:

*Spacer items*

![Spacer items diagram](../_images/layout16.jpg)

By default, a spacer item is created at the right of the hbox container, but spacers can also be
defined in other places inside the hbox tag:

```
GRID
{
[a      :b      :c       ] -- default spacer on the right
[ :d    :e      :f       ] -- spacer on the left
[g      :     :h         ] -- spacer between g and h
[i: :j: :k    : :l       ] -- multiple spacers (between i and j, j and k, k and l
}
END
```

![Form using hbox tags with spacers screenshot](../_images/layout19_gbc.jpg)

*Form using hbox tags with spacers*
