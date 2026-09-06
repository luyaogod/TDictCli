---
title: "Example"
source: "fgl-topics/c_fgl_localized_strings_012.html"
breadcrumb: "Advanced features > Localization > Localized strings > Example"
type: "concept"
---

# Example

> Here is an example using localized strings.

The source string file "common.str" (to be compiled with
fglmkstr):

```
common.accept = "OK"
common.cancel = "Cancel"
common.yes = "Yes!"
common.no = "No!"
```

The source string file "customer.str" (to be compiled with
fglmkstr):

```
customer.mainwindow.title = "Customers"
customer.listwindow.title = "Customer List"
customer.l_custnum = "Number:"
customer.l_custname = "Name:"
customer.c_custname = "The customer name"
customer.q_delete = "Do you want to delete this customer?"
```

The FGLPROFILE configuration file parameters:

```
fglrun.localization.file.count = 1
fglrun.localization.file.1.name = "common.42s"
```

Remark: The 'customer' string file does not have to listed in FGLPROFILE
since it is loaded as it has the same name as the program.

The form specification file "customer.per":

```
ACTION DEFAULTS
  ACTION accept (TEXT=%"common.accept")
  ACTION cancel (TEXT=%"common.cancel")
END
LAYOUT (TEXT=%"customer.mainwindow.title")
GRID
{
[lab1       ] [f01            ]
[lab2       ] [f02                     ]
}
END
END
ATTRIBUTES
  LABEL lab1: TEXT=%"customer.l_custnum";
  EDIT f01 = FORMONLY.custnum;
  LABEL lab2: TEXT=%"customer.l_custname";
  EDIT f02 = FORMONLY.custname, COMMENT=%"customer.c_custname";
END
```

The program "customer.4gl" using the strings file:

```
MAIN
  DEFINE rec RECORD
           custnum INTEGER,
           custname CHAR(20)
         END RECORD
  OPEN FORM f1 FROM "customer"
  DISPLAY FORM f1
  INPUT BY NAME rec.*
    ON ACTION delete
       MENU %"customer.mainwindow.title"
                ATTRIBUTES(STYLE="dialog", COMMENT=%"customer.q_delete")
           COMMAND %"common.yes"
           COMMAND %"common.no"
       END MENU
  END INPUT
END MAIN
```

## Related links

**Related concepts**  

[Steps for application internationalization](0902-steps-for-application-internationalization.md "Follow these steps to internationalize your application.")

[Creating source string files](0903-creating-source-string-files.md "A source string file contains localized string definitions for a given language (or localization context).")

[Compiling string resource files (.str)](0908-compiling-string-resource-files-str.md "The .str source string files must be compiled to .42s binary files, in order to be loaded by the runtime system.")
