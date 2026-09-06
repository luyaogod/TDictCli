---
title: "Example: Reading a JSON file"
source: "fgl-topics/c_gws_jsonJSONReader_example_1.html"
breadcrumb: "Library reference > Extension packages > The json package > The streaming API for JSON classes > The json.JSONReader class > Examples > Example: Reading a JSON file"
type: "concept"
---

# Example: Reading a JSON file

> This example shows two ways to read a JSON file.

## Example: JSON parsing and serializing directly

This program reads JSON data from customers.json. In this example, it loads
the file into the `reader` object and advances to the first token. The
`JSONToVariable` method is called to serialize the data and fill the program array
`custlist` directly. Then the record content is written to the standard output.

A `TRY/CATCH` block is used to detect potential JSON format errors in the source
file.

```
IMPORT json

DEFINE custlist DYNAMIC ARRAY OF RECORD
            num INTEGER,
            name VARCHAR(40)
         END RECORD

DEFINE reader json.JSONReader
DEFINE ind INTEGER

MAIN
    TRY
        LET reader = json.JSONReader.Create()
        CALL reader.setInputCharset("UTF-8")
        CALL reader.readfrom("customers.json")
        CALL reader.next()
        DISPLAY "Start parsing "
        CALL json.Serializer.JSONToVariable(reader,custlist)
        FOR ind = 1 to custlist.getLength()
            DISPLAY custlist[ind].*
        END FOR
        CALL reader.close()
    CATCH
        DISPLAY "JSON Serializer ERROR: shouldn't raise error :",
        status || " " || sqlca.sqlerrm
        EXIT PROGRAM -1
    END TRY
    DISPLAY "parsing complete"
END MAIN
```

The file customers.json:

```
[
   {
        "num": 823,
        "name": "Mark Renbing"
    }, {
        "num": 234,
        "name": "Clark Gambler"
    }
]
```

The output:

```
Start parsing 
        823Mark Renbing
        234Clark Gambler
parsing complete
```

## Example: event type JSON parsing

This program reads JSON data from client.json. Use this example to control
end-to-end JSON parsing. The program loads the file into the `reader` object. The
main loop will parse the file by checking the [event type](3659-jsonreader-event-types.md "Event types of the json.JSONReader class.") as it reads the file.

A `TRY/CATCH` block is used to detect potential JSON format errors in the source
file.

```
IMPORT json
IMPORT util

DEFINE reader json.JSONReader
DEFINE str STRING

MAIN

    DEFINE startReadClient BOOLEAN
    DEFINE array BOOLEAN
    DEFINE nextElement BOOLEAN

    TRY
        LET reader = json.JSONReader.Create()
        CALL reader.setInputCharset("UTF-8")
        CALL reader.readfrom("client.json")
        DISPLAY "Hello ! Welcome to your client analyzer :) \n"
        WHILE reader.hasnext()
            CASE reader.getEventType()
                WHEN "START_DOCUMENT"
                    DISPLAY "Start reading data.. \n"
                    CALL reader.next()
                WHEN "START_OBJECT"
                    IF NOT startReadClient THEN
                        DISPLAY "\n---------- New client discovered: ----------\n"
                        LET startReadClient = TRUE
                    END IF
                    CALL reader.next()
                WHEN "END_OBJECT"
                    DISPLAY "\n"
                    CALL reader.next()
                WHEN "START_ARRAY"
                    CALL reader.next()
                WHEN "END_ARRAY"
                    CALL reader.next()
                WHEN "PROPERTY"
                    LET str = reader.getProperty()
                    CASE
                        WHEN reader.getProperty().equalsIgnoreCase("name")
                            DISPLAY "Name: "
                        WHEN reader.getProperty().equalsIgnoreCase("firstname")
                            CALL reader.next()
                            DISPLAY "   Firstname: ", reader.getValue()
                        WHEN reader.getProperty().equalsIgnoreCase("lastname")
                            CALL reader.next()
                            DISPLAY "   Lastname: ", reader.getValue()
                        WHEN reader.getProperty().equalsIgnoreCase("sex")
                            CALL reader.next()
                            DISPLAY "Sex: ", reader.getValue()
                            DISPLAY "\n"
                        WHEN reader.getProperty().equalsIgnoreCase("address")
                            DISPLAY "Address: ", reader.getValue()
                        WHEN reader.getProperty().equalsIgnoreCase("streetAddress")
                            CALL reader.next()
                            DISPLAY "   StreetAddress: ", reader.getValue()
                        WHEN reader.getProperty().equalsIgnoreCase("City")
                            CALL reader.next()
                            DISPLAY "   City: ", reader.getValue()
                        WHEN reader.getProperty().equalsIgnoreCase("State")
                            CALL reader.next()
                            DISPLAY "   State: ", reader.getValue()
                        WHEN reader.getProperty().equalsIgnoreCase("phoneNumber")
                            DISPLAY "PhoneNumber: "
                        WHEN reader.getProperty().equalsIgnoreCase("type")
                            CALL reader.next()
                            DISPLAY "   Type: ", reader.getValue()
                            WHEN reader.getProperty().equalsIgnoreCase("number")
                            CALL reader.next()
                            DISPLAY "   number: ", reader.getValue()
                    END CASE
                    CALL reader.next()
                WHEN "STRING"
                    CALL reader.next()
                WHEN "NUMBER"
                    CALL reader.next()
                WHEN "BOOLEAN"
                    CALL reader.next()
                WHEN "NULL"
                    DISPLAY "Null"
                    CALL reader.next()
                WHEN "END_DOCUMENT"
                    DISPLAY "\n..end of document"
                    EXIT WHILE
                OTHERWISE
                    DISPLAY "json data not recognized"
                    EXIT PROGRAM -1
            END CASE
        END WHILE
    CATCH
        DISPLAY "JSON Serializer ERROR: shouldn't raise error :",
            status || " " || sqlca.sqlerrm
        EXIT PROGRAM -1
    END TRY
    CALL reader.close()
END MAIN
```

The file client.json:

```
{
    "name": {
      "firstname": "Pierre",
      "lastname": "Kant"
    },
    "age": 32,
    "sex": "Male",
    "address": {
        "streetAddress": "45 Rue de Londres",
        "City": "Toulouse",
        "State": "France",
        "postalCode":31555
      },
    "phoneNumber": [
        {
            "type": "home",
            "number": "555 555 555"
        },
        {
            "type": "mobile",
            "number": "556 556 556"
        }
    ]
}
```

The output:

```
Hello ! Welcome to your client analyzer :) 

Start reading data..

---------- New client discovered: ----------

Name:
   Firstname: Pierre
   Lastname: Kant

Address:
   StreetAddress: 45 Rue de Londres
   City: Toulouse
   State: France

PhoneNumber:
   Type: home
   number: 555 555 555

   Type: mobile
   number: 556 556 556

..end of document
```
