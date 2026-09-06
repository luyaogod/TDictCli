---
title: "Complex grid layout example"
source: "fgl-topics/c_fgl_form_rendering_complex_grid_1.html"
breadcrumb: "User interface > Form definitions > Form rendering > Grid-based layout > Complex grid layout example"
type: "concept"
---

# Complex grid layout example

> Describes how form items align in grid-based front-ends with an example.

These diagrams show the virtual grid of a complex form, with several field item tags.

The source file to produce this rendering is following:

```
LAYOUT
GRID
{
[a       |b           ][c          ]
[d      ][e            |f      |g  ]
[h|i|j|k                      ][l  ]
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

![Grid containing several fields diagram](../_images/layout05.jpg)

*Grid containing several fields*

For each form field, the position and the number of cells is computed by the form compiler.

At runtime, the front-end creates the widgets and sets them on the virtual grid.

![Widgets set on grid diagram](../_images/layout06.jpg)

*Widgets set on grid*

Once widgets are on the grid, their minimum size is computed based on widget size, [`SAMPLE`](1814-sample-attribute.md "The SAMPLE attribute defines the text to be used to compute the width of a form field widget.") and [`SIZEPOLICY`](1817-sizepolicy-attribute.md "The SIZEPOLICY attribute is a sizing directive based on the content of a form item.") attributes.

Then the sizes of the grid cells adapt to the size of the widgets.

![Widget size in layout diagram](../_images/layout07.jpg)

*Widget size computations*

![Widgets in rendered form screenshot](../_images/layout08_gbc.jpg)

*Widgets in rendered form*

In this screenshot, the fields `k` and `c` are much bigger than
expected:

- Field `g` and `l` make columns 33, 34 and 35 bigger than the
  others,
- Field `f` extends columns 25 to 31.
- As field `c` has to fill columns 25 to 35, its size grows; the same for field
  `k`.

Some fields are proportionally bigger than others because some parameters are variable, while
others are fixed.

The width of the widget is the sum of border width, plus the content width (depending on
`SIZEPOLICY` and the `SAMPLE` attributes).

Since the default `SAMPLE` is MMMMMM000..., the graphical width of a field is not
linearly proportional to the width defined in the form file. For example, a field of 1 will be as
wide as 2 borders + 1 'M'. A field of 10 will be as wide as 2 borders + 6 'M' + 4 '0'. This means
that a field of 1 is far from being 10 times smaller than a field of 10.
