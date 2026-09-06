---
title: "ON ACTION block"
source: "fgl-topics/c_fgl_prompt_009.html"
breadcrumb: "User interface > Dialog instructions > Prompt for values (PROMPT) > Using simple prompt inputs > Interaction blocks > ON ACTION block"
type: "concept"
description: "Use ON ACTION blocks in a PROMPT dialog, to execute a sequence of instructions when the user selects the action. This is the preferred solution compared to ON KEY blocks, because ON ACTION blocks use ..."
---

# ON ACTION block

Use `ON ACTION` blocks in a `PROMPT` dialog, to execute a sequence
of instructions when the user selects the action. This is the preferred solution compared to
`ON KEY` blocks, because `ON ACTION` blocks use abstract names to
control user interaction.

The [`PROMPT`](1888-prompt-for-values-prompt.md "The PROMPT instruction provides unique field input in an automatic pop-up window.")
dialog is automatically terminated after `ON
IDLE`, `ON TIMER`,
`ON ACTION`, or `ON KEY` block execution.

## Related links

**Related concepts**  

[Dialog actions](2253-dialog-actions.md "Describes how to program action handling when the end user triggers an action on the front-end.")
