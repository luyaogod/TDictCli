---
title: "Example 2: PROMPT with validation checking"
source: "fgl-topics/c_fgl_prompt_014.html"
breadcrumb: "User interface > Dialog instructions > Prompt for values (PROMPT) > Examples > Example 2: PROMPT with validation checking"
type: "concept"
description: "MAIN DEFINE birth DATE LET int_flag = FALSE PROMPT \"Please enter your birthday: \" FOR birth IF int_flag THEN DISPLAY \"Prompt input has been canceled.\" ELSE DISPLAY \"Your birthday is: \" || birth END IF ..."
---

# Example 2: PROMPT with validation checking

```
MAIN
  DEFINE birth DATE
  LET int_flag = FALSE
  PROMPT "Please enter your birthday: " FOR birth 
  IF int_flag THEN
    DISPLAY "Prompt input has been canceled."
  ELSE
    DISPLAY "Your birthday is: " || birth 
  END IF
END MAIN
```
