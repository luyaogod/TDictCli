---
title: "Source documentation structure"
source: "fgl-topics/c_fgl_AutoDoc_005.html"
breadcrumb: "Programming tools > Source documentation > Source documentation structure"
type: "concept"
---

# Source documentation structure

> The source documentation structure is based on the well-known Java-doc technique.

The generated documentation reflects the structure of your sources; in order to have nicely
structured source documentation, you must have a nicely structured source tree.

The source documentation elements are structured as follows (elements in
italic must be created by hand, others are generated files):

- Top/root directory (the root of your project)
  - `overview.4gl` (description of the project)
  - `overview-summary.html`
  - `overview-frame.html`
  - `allclasses-frame.html`
  - `index-all.html`
  - `index.html`
  - `fgldoc.css`
  - sub-directory1 (package)
    - sub-directory11 (package)
    - ...
    - sub-directory12 (package)
      - sub-directory121 (package)
        - `package-info.4gl` (description of
          the package/directory)
        - `package-summary.html`
        - `package-frame.html`
        - source1211.4gl (module)
        - `source1211.html`
        - source1212.4gl (module)
          - `#+ Module 1212 comment` (description of the
            module)
          - `#+ Function 12121 comment` (description of the
            function)
          - function12121
          - `#+ Function 12121 comment`
          - function12122
          - ...
        - `source1212.html`
        - source1213.4gl (module)
        - `source1213.html`
        - ...
  - sub-directory2 (package)
  - ...
  - sub-directory3 (package)
  - ...

First create a file named `overview.4gl` in the
top directory of the project. This file contains the overall description
of the project. In that directory, the documentation generator
creates the files `overview-summary.html`, `overview-frame.html`, `allclasses-frame.html`,
`index-all.html`, `index.html` and `fgldoc.css`
.

The documentation generator can scan sub-directories to build the documentation for a whole
project; each source directory defines a package. For each directory (package), the generator
creates a `package-summary.html` and a `package-frame.html` file. If a
file with the name `package-info.4gl` exists, it will be scanned to complete the
`package-summary.html` file with the package description.

The documentation generator creates a filename.html file
for each .4gl source module, seen as a *class* in the documentation.
