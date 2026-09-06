---
title: "Understanding the PROMPT instruction"
source: "fgl-topics/c_fgl_prompt_002.html"
breadcrumb: "User interface > Dialog instructions > Prompt for values (PROMPT) > Understanding the PROMPT instruction"
type: "concept"
---

# Understanding the PROMPT instruction

> The PROMPT instruction is used to query for a single value from the user.

`PROMPT` requires the text of the question to be
displayed to the user and the variable that receives the value entered
by the user. The variable can be of any simple data type except
`TEXT` and `BYTE`.

The runtime system displays the question in the prompt area, waits for the user to enter a value,
reads whatever value was entered until the user validates (for example with the Enter key), and
stores this value in a response variable. The prompt dialog remains visible until the user enters a
response.

The [`PROMPT`](1888-prompt-for-values-prompt.md "The PROMPT instruction provides unique field input in an automatic pop-up window.")
dialog is automatically terminated after `ON
IDLE`, `ON TIMER`,
`ON ACTION`, or `ON KEY` block execution.

## Prompt display in TUI mode

In TUI mode, the `PROMPT` question and input field is displayed in the prompt line
of the current window, which is defined by the `OPTIONS PROMPT LINE` instruction or
with the `ATTRIBUTES` clause of `OPEN WINDOW`.

If the prompt line is not as wide as the prompt string and the size required for the variable
input, runtime error [-1146](../15_library-reference/4483-genero-bdl-errors.md)
occurs.

## Prompt display in GUI mode

In GUI mode, the `PROMPT` instruction opens a modal window with an OK and a Cancel
button, and waits for input from the user.

![PROMPT window screenshot](../_images/Prompt01_gbc.jpg)

*PROMPT window*

## Related links

**Related concepts**  

[Defining the position of reserved lines](../09_advanced-features/0925-defining-the-position-of-reserved-lines.md "The OPTIONS element LINE defines position of dedicated screen lines.")

[Dialog programming basics](2218-dialog-programming-basics.md "This section describes basic dialog programming concepts.")
