---
title: "Usage of RUN IN FORM MODE"
source: "fgl-topics/c_fgl_Migrate_to_200_run_in_form_mode.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.00 upgrade guide > Usage of RUN IN FORM MODE"
type: "concept"
---

# Usage of RUN IN FORM MODE

> RUN ... IN LINE MODE is recommended to run interactive applications.

Before version 2.00, `RUN ... IN FORM MODE` was recommended to run
interactive applications.

Starting with version 2.00, the `RUN` command must be used as follows (in both GUI
and TUI mode):

1. When starting an interactive program, use `RUN ... IN LINE MODE`. If the
   default run mode is `LINE`, you can use the `RUN` instruction without
   any option.
2. When starting a (silent) batch program that does not display any message, it is recommended that
   you use `RUN ... IN FORM MODE`.

For more details, read the [`RUN`](../09_advanced-features/0830-run.md) instruction reference topic.

## Related links

**Related concepts**  

[Program execution](../09_advanced-features/0828-program-execution.md "This section describes program execution and language instructions related to program execution.")
