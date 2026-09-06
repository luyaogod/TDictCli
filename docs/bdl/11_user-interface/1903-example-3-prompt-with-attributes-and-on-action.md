---
title: "Example 3: PROMPT with ATTRIBUTES and ON ACTION"
source: "fgl-topics/c_fgl_prompt_015.html"
breadcrumb: "User interface > Dialog instructions > Prompt for values (PROMPT) > Examples > Example 3: PROMPT with ATTRIBUTES and ON ACTION"
type: "concept"
description: "MAIN DEFINE birth DATE LET birth = TODAY PROMPT \"Please enter your birthday: \" FOR birth ATTRIBUTES(WITHOUT DEFAULTS) ON ACTION action1 DISPLAY \"Action 1\" END PROMPT DISPLAY \"Your birthday is \" || ..."
---

# Example 3: PROMPT with ATTRIBUTES and ON ACTION

```
MAIN
  DEFINE birth DATE
  LET birth = TODAY
  PROMPT "Please enter your birthday: " FOR birth 
     ATTRIBUTES(WITHOUT DEFAULTS)
        ON ACTION action1
           DISPLAY "Action 1"
  END PROMPT
  DISPLAY "Your birthday is " || birth 
END MAIN
```
