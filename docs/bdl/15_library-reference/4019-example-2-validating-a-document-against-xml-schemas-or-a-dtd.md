---
title: "Example 2 : Validating a document against XML schemas or a DTD"
source: "fgl-topics/c_gws_XML_DomDocument_class_016.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomDocument class > Examples > Example 2 : Validating a document against XML schemas or a DTD"
type: "concept"
description: "This code example loads one or more XML schemas or uses an embedded DTD to validate against a XML document: IMPORT xml MAIN DEFINE location STRING DEFINE xmlfile STRING DEFINE doc xml.DomDocument ..."
---

# Example 2 : Validating a document against XML schemas or a DTD

This code example loads one or more XML schemas or uses an embedded
DTD to validate against a XML document:

```
IMPORT xml
 
MAIN 
  DEFINE location STRING
  DEFINE xmlfile STRING
  DEFINE doc xml.DomDocument
  DEFINE ind INTEGER
 
  IF num_args()<2 THEN
    # Checks the number of arguments
    CALL ExitHelp() 
  ELSE
    LET doc = xml.DomDocument.Create()
    LET xmlfile = arg_val(num_args())
    IF num_args() == 2 AND arg_val(1) == "-dtd" THEN
      # User choosed DTD validation
      CALL doc.setFeature("validation-type", "DTD")
    ELSE 
      # User choosed XML Schema validation
      IF arg_val(1) == "-ns" THEN
        # Handle namespace qualified XML schemas
        IF num_args() MOD 2 != 0 THEN
          CALL ExitHelp()
        END IF
        FOR ind = 2 TO num_args()-1 STEP 2
          IF location IS NULL THEN
            LET location = arg_val(ind) || " " || arg_val(ind+1)
          ELSE
            LET location = location || " " || arg_val(ind) ||
                 " " || arg_val(ind+1)
          END IF
        END FOR
        TRY
          CALL doc.setFeature("external-schemaLocation", location)
        CATCH
          FOR ind = 1 TO doc.getErrorsCount()
            DISPLAY "Schema error ("||ind||") :",doc.getErrorDescription(ind)
          END FOR 
          EXIT PROGRAM (-1)
        END TRY
      ELSE
        # Handle unqualified XML schemas
        FOR ind = 1 TO num_args()-1 
          IF location IS NULL THEN
            LET location = arg_val(ind)
          ELSE
            LET location = location || " " || arg_val(ind)
          END IF
        END FOR
        TRY
          CALL doc.setFeature("external-noNamespaceSchemaLocation", location)
        CATCH
          FOR ind = 1 TO doc.getErrorsCount()
            DISPLAY "Schema error ("||ind||") :",doc.getErrorDescription(ind)
          END FOR 
          EXIT PROGRAM (-1)
        END TRY
      END IF
    END IF
  END IF
  TRY
    # Load XML document from disk
    CALL doc.load(xmlfile)
  CATCH
    # Display errors if loading fails
    IF doc.getErrorsCount()>0 THEN
      FOR ind = 1 TO doc.getErrorsCount()
        DISPLAY "LOADING ERROR #"||ind||" :",doc.getErrorDescription(ind)
      END FOR
      EXIT PROGRAM(-1)
    ELSE
      DISPLAY "Unable to load file :",xmlfile
      EXIT PROGRAM(-1)
    END IF
  END TRY
  TRY
    # Validate loaded document     
    LET ind = doc.validate()
    IF ind == 0 THEN
      # Successful validation
      DISPLAY "OK"
    ELSE
      # Display validation errors
      FOR ind = 1 TO doc.getErrorsCount()
        DISPLAY "VALIDATING ERROR #"||ind||" :",doc.getErrorDescription(ind)
      END FOR
      EXIT PROGRAM(-1)
    END IF 
  CATCH
    DISPLAY "Unable to validate file :",xmlfile
    EXIT PROGRAM(-1)
  END TRY 
END MAIN

# Display help

FUNCTION ExitHelp()
  DISPLAY "Validator < -dtd | -ns [namespace schema]+ | [schema]+ > xmlfile"
  EXIT PROGRAM
END FUNCTION
```

## Example

```
$ fglrun Validator -dtd MyFile.xml
```

Validates XML file using DTD embedded in the XML file.

```
$ fglrun Validator Schema1.xsd Schema2.xsd MyFile.xml
```

Validates unqualified XML file using two unqualified XML schemas.

```
$ fglrun Validator -ns http://tempuri.org/one Schema1.xsd
   http://tempuri.org/two Schema2.xsd MyFile.xml
```

Validates namespace qualified XML file using two namespace qualified XML schemas.
