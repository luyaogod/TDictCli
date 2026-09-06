---
title: "Run the documentation generator"
source: "fgl-topics/t_fgl_autodoc_run_the_documentation_generator.html"
breadcrumb: "Programming tools > Source documentation > Run the documentation generator"
type: "task"
---

# Run the documentation generator

> A step-by-step procedure to build the source documentation.

Follow the steps below to produce HTML documentation from your Genero BDL source code:

1. Go to the top directory of your sources.
2. Create a file named overview.4gl, with a `#+` comment
   describing your project.
3. Go to the subdirectories and create files named package-info.4gl with a
   `#+` comment describing the package.
4. Edit the .4gl modules to add `#+` comments to functions
   that must be documented.
5. Go back to the top directory of your sources.
6. Run fglcomp --build-doc overview.4gl

   Use the `-W apidoc` compiler option to get warnings for invalid comment tags.
   For example, when a `@param` tag is missing for a function parameter.
7. To test the result, load the generated index.html file in your preferred
   browser.
