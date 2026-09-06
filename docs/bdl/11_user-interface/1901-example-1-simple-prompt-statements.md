---
title: "Example 1: Simple PROMPT statements"
source: "fgl-topics/c_fgl_prompt_013.html"
breadcrumb: "User interface > Dialog instructions > Prompt for values (PROMPT) > Examples > Example 1: Simple PROMPT statements"
type: "concept"
description: "MAIN DEFINE birth DATE DEFINE chkey CHAR(1) PROMPT \"Please enter your birthday: \" FOR birth DISPLAY \"Your birthday is: \" || birth PROMPT \"Now press a key... \" FOR CHAR chkey DISPLAY \"You pressed: \" || ..."
---

# Example 1: Simple PROMPT statements

```
MAIN
  DEFINE birth DATE
  DEFINE chkey CHAR(1)
  PROMPT "Please enter your birthday: " FOR birth 
  DISPLAY "Your birthday is: " || birth 
  PROMPT "Now press a key... " FOR CHAR chkey 
  DISPLAY "You pressed: " || chkey 
END MAIN
```
