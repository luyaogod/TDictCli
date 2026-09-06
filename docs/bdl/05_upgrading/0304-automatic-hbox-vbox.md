---
title: "Automatic HBox/VBox"
source: "fgl-topics/c_fgl_Migrate_to_202_auto_hbox_vbox.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.02 upgrade guide > Automatic HBox/VBox"
type: "concept"
---

# Automatic HBox/VBox

> fglform adds automatically HBox/VBox elements when needed.

Starting with version 2.02.01, the form compiler automatically adds HBox and VBox containers with
splitter around stretchable form elements that are placed side-by-side.

When recompiling your forms with this new version of [fglform](../13_programming-tools/2514-fglform.md "The fglform tool compiles form specification files into XML formatted files used by programs."), the generated
.42f can get additional HBox/VBox nodes even if you did not touch the
.per source file.

## Related links

**Related concepts**  

[Form rendering](../11_user-interface/1538-form-rendering.md "The section explains the layout rules to render forms on graphical front-ends.")

[Compiling form specification files (.per)](../13_programming-tools/2531-compiling-form-specification-files-per.md "The .per form definition files must be compiled to .42f XML files, in order to be loaded by the runtime system.")
