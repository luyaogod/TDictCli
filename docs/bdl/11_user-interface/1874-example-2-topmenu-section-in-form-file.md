---
title: "Example 2: Topmenu section in form file"
source: "fgl-topics/c_fgl_topmenus_013.html"
breadcrumb: "User interface > Form definitions > Topmenus > Examples > Example 2: Topmenu section in form file"
type: "concept"
description: "TOPMENU GROUP form (TEXT=\"Form\", STYLE=\"mystyle\" ) COMMAND help (TEXT=\"Help\", IMAGE=\"quest\") COMMAND quit (TEXT=\"Quit\", ACCELERATOR=ALT-F4) END GROUP edit (TEXT=\"Edit\") COMMAND accept ..."
---

# Example 2: Topmenu section in form file

```
TOPMENU
  GROUP form (TEXT="Form", STYLE="mystyle" )
     COMMAND help (TEXT="Help", IMAGE="quest")
     COMMAND quit (TEXT="Quit", ACCELERATOR=ALT-F4)
  END
  GROUP edit (TEXT="Edit")
     COMMAND accept (TEXT="Validate", IMAGE="ok")
     COMMAND cancel (TEXT="Cancel", IMAGE="cancel")
  END
  GROUP records (TEXT="Records")
     COMMAND append (TEXT="Add", IMAGE="add")
     COMMAND delete (TEXT="Remove", IMAGE="del")
     COMMAND update (TEXT="Modify", IMAGE="change")
     SEPARATOR
     COMMAND search (TEXT="Search", IMAGE="find")
  END
END
```
