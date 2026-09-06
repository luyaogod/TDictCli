---
title: "PROMPT programming steps"
source: "fgl-topics/t_fgl_prompt_005.html"
breadcrumb: "User interface > Dialog instructions > Prompt for values (PROMPT) > PROMPT programming steps"
type: "task"
description: "To use the PROMPT statement, you must: Declare a program variable with the DEFINE statement. Set the int_flag variable to FALSE . Define the PROMPT dialog , with dialog control blocks to control the ..."
---

# PROMPT programming steps

To use the `PROMPT` statement, you must:

1. Declare a [program variable](../08_language-basics/0686-variables.md "Explains how to define program variables.") with the
   `DEFINE` statement.
2. Set the `int_flag` variable to
   `FALSE`.
3. Define the [`PROMPT` dialog](1890-syntax-of-prompt-instruction.md "The PROMPT statement assigns a user-supplied value to a variable."), with
   dialog control blocks to control the instruction. Use the `FOR CHAR` clause if a
   single character is to be entered.
4. After executing the `PROMPT`, check the `int_flag` variable to
   determine whether the input was validated or canceled by the user.

## Related links

**Related concepts**  

[Example 2: PROMPT with validation checking](1902-example-2-prompt-with-validation-checking.md "Example 2: PROMPT with validation checking")
