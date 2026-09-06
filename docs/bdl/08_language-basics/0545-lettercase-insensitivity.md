---
title: "Lettercase insensitivity"
source: "fgl-topics/c_fgl_language_features_lettercase_insensitivity.html"
breadcrumb: "Language basics > Syntax features > Lettercase insensitivity"
type: "concept"
---

# Lettercase insensitivity

> Genero BDL is case insensitive (with some exceptions).

Genero BDL does not make distinction between uppercase and lowercase letters, except within
quoted strings.

You can mix uppercase and lowercase letters in the identifiers that you assign to language
entities, but any uppercase letters in identifiers are automatically shifted to lowercase during
compilation:

```
INPUT BY NAME ...
    ON ACTION PrintReport     -- becomes "printreport" in 42m pcode
```

It is strongly recommended that you define a naming convention
for your projects. For example, you can use underscore notation (`get_user_name`).
If you plan to use the Java notation
(`getUserName`), do not forget that Genero BDL is case
insensitive (`getusername` is the same identifier as `getUserName`).

For better readability, and to be consistent with SQL syntax conventions, consider writing BDL
language keywords in UPPERCASE, and other language elements like identifiers in
lowercase:

```
INPUT BY NAME cust_rec.* ATTRIBUTES(UNBUFFERED)
```

Exceptions regarding case-senstivity:

1. Genero BDL interfaces with [Java](../14_extending-the-language/2655-the-java-interface.md "The Java interface allows you to import Java classes and instantiate Java objects in your programs.") classes and
   objects. Java symbols are case-sensitive.
2. The identifiers (and filenames) used with [`IMPORT FGL`](../09_advanced-features/0815-import-fgl.md "The IMPORT FGL instruction imports module symbols.") are case sensitive.
3. Resource files using XML such as [.4ad](../11_user-interface/1600-action-defaults-files.md "Action defaults files allow to centralize action configuration parameters such as text, icon, accelerators and behavior options in XML format.") files are case-sensitive (regarding action names for
   example).
