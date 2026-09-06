---
title: "Setting action key labels"
source: "fgl-topics/c_fgl_DynamicUI_026.html"
breadcrumb: "User interface > User interface programming > Dialog actions > Setting action key labels"
type: "concept"
---

# Setting action key labels

> Labels can be defined to decorate buttons controlled by ON KEY / COMMAND KEY action handlers.

> **Important:**
>
> Key label configuration is provided for backward compatibility. Consider
> using [action configuration](2259-configuring-actions.md "Action attributes related to decoration, keyboard shortcuts, and behavior can be defined with action attributes.") in new
> programs. However, key labels can be used to easily improve the graphical rendering of your
> application, without touching legacy code using `ON KEY` / `COMMAND
> KEY` in dialogs.

## Syntax

Key label configuration can take place at different levels.

- FGLPROFILE definitions

  ```
  key.key-name.text = "label"
  ```
- Program-level key labels

  ```
  CALL fgl_setkeylabel( "key-name", "label" )
  ```
- Form level key labels (in KEYS section)

  ```
  KEYS key-name = [%]"label"
  [...]
  [END]
  ```
- Dialog level key labels

  ```
  CALL fgl_dialog_setkeylabel( "key-name", "label" )
  ```
- Form field level key labels (in field definition)

  ```
  KEY key-name = [%]"label"
  ```

1. *key-name* is the name of the key.
2. *label* is the text to be displayed in the default action view (button).

## Usage

In GUI mode, `ON KEY` and `COMMAND KEY` action handlers in dialogs
can be shown as form buttons when a label text is defined for the key. By defining a label for a
key, the runtime system will automatically show a default button for the key action.

In the example, the function key `F10` is used to show a detail window in this
interactive dialog:

```
INPUT BY NAME myrecord.*
  ON KEY (F10)
    CALL ShowDetail()
END INPUT
```

By default, if you do not specify a label for F10, no default action button is displayed for a
function key or control key. Furthermore, if the text provided for the key label is empty or null, the
default action button will not be displayed.

In order to get a default action view button for F10, define for example the KEY attribute in the
form file, for the corresponding fields where this action can be
fired:

```
ATTRIBUTES
...
f07 = customer.cust_city, KEY F10 = "City list";
f08 = customer.cust_state, KEY F10 ="State list";
...
```

| Key Name | Description |
| --- | --- |
| `f1` to `f255` | Function keys. |
| `control-a` to `control-z` | Control keys. |
| `accept` | Predefined dialog validation action. |
| `interrupt` | Predefined dialog cancellation action. The action name is *cancel*, not *interrupt*. |
| `insert` | Predefined `INPUT ARRAY` dialog row insertion action. |
| `append` | Predefined `INPUT ARRAY` dialog row addition action. |
| `delete` | Predefined `INPUT ARRAY` dialog row deletion action. |
| `help` | Predefined help action. |

Key labels can be defined at different levels. The order of precedence for key label definition
is the following:

1. The label defined with the [KEY](1797-key-attribute.md "The KEY attribute is used to define the labels of keys when the field is made current.") attribute of the
   form field.
2. The label defined for the current dialog, using the [FGL\_DIALOG\_SETKEYLABEL](../15_library-reference/2755-fgl-dialog-setkeylabel.md "Sets the label associated to a key for the current interactive instruction.")
   function.
3. The label defined in the [KEYS section](1752-keys-section.md "The KEYS section can be used to define default key labels for the current form.")
   of the form specification file.
4. The label defined as default for a program, using the [FGL\_SETKEYLABEL](../15_library-reference/2777-fgl-setkeylabel.md "Sets the default label associated to a key.") function.
5. The label defined in the FGLPROFILE configuration file
   (`key.key-name.text` entries).

In Genero, you typically define action labels with action attributes. However, if key labels are
defined, they will overwrite the text defined in action attributes for the corresponding key action.
In BDS 3.xx versions, default key labels are defined in $FGLDIR/etc/fglprofile.
These defaults have been commented out in Genero to have action attribute text applied (In Genero,
by default, `fgl_getkeylabel()` returns `NULL` for all keys). If you
want to get the same default key labels as in BDS 3.xx, uncomment the key.\* lines in
$FGLDIR/etc/fglprofile.

You can query the label defined at the program level with the [FGL\_GETKEYLABEL](../15_library-reference/2764-fgl-getkeylabel.md "Returns the default label associated to a key.") function and, for the
current interactive instruction, with the [FGL\_DIALOG\_GETKEYLABEL](../15_library-reference/2753-fgl-dialog-getkeylabel.md "Returns the label associated to a key for the current interactive instruction.")
function.

## Related links

**Related concepts**  

[Configuring actions](2259-configuring-actions.md "Action attributes related to decoration, keyboard shortcuts, and behavior can be defined with action attributes.")

[ON KEY block](1898-on-key-block.md "ON KEY block")

[COMMAND KEY() block](1917-command-key-block.md "COMMAND KEY() block")

[The FGLPROFILE file(s)](../07_configuration/0483-the-fglprofile-file-s.md "FGLPROFILE environment variable defines Genero BDL configuration files")
