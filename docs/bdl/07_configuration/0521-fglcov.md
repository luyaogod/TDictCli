---
title: "FGLCOV"
source: "fgl-topics/c_fgl_EnvVariables_FGLCOV.html"
breadcrumb: "Configuration > Environment variables > Genero environment variables > FGLCOV"
type: "concept"
---

# FGLCOV

> Enables code coverage data collection.

The FGLCOV environnement variable controls code coverage data collection while executing
programs.

When the FGLCOV variable is set to a numeric value different from zero, fglrun
produces module.42m.cov files.

FGLCOV=0 disables the code coverage option.

For more details, see [Source code coverage](../13_programming-tools/2628-source-code-coverage.md "Collect information about used source lines").
