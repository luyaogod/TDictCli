---
title: "Example 1: Calling a JavaScript function"
source: "fgl-topics/c_fgl_webcomponent_gicapi_example_1.html"
breadcrumb: "User interface > User interface programming > Web components > Using a gICAPI web component > Examples > Example 1: Calling a JavaScript function"
type: "concept"
description: "This example shows how to call a JavaScript function with the webcomponent.call front call The form file: form.per LAYOUT GRID { [wc1 ] [ ] [ ] Result: [f1 ] } END END ATTRIBUTES WEBCOMPONENT wc1 = ..."
---

# Example 1: Calling a JavaScript function

This example shows how to call a JavaScript function with the `webcomponent.call` front call

The form file: form.per

```
LAYOUT
GRID
{
[wc1                                    ]
[                                       ]
[                                       ]
Result: [f1                             ]
}
END
END
ATTRIBUTES
WEBCOMPONENT wc1 = FORMONLY.mywebcomp,
    COMPONENTTYPE="mywebcomp",
    STRETCH=BOTH;
EDIT f1 = FORMONLY.result, SCROLL;
END
```

The HTML file: mywebcomp.html

```
<html>

<head>
<title>The title</title>
<script language="JavaScript" type="text/javascript">

var echoString = function(str) {
   document.getElementById("p_info").innerText = str;
   return str;
}

var echoCustFields = function(id,name,crea) {
   var str = "FIELDS: Cust id:" + id + " Name:" + name + " ("+crea+")";
   document.getElementById("p_info").innerText = str;
   return str;
}

var echoCustRecord = function(obj) {
   var str = "RECORD: Cust id:" + obj.id + " Name:" + obj.name + " ("+obj.crea+")";
   document.getElementById("p_info").innerText = str;
   return str;
}

var echoJSONString = function(ostr) {
   var o = JSON.parse(ostr);
   document.getElementById("p_info").innerText = ostr;
   return JSON.stringify(o);
}

</script>
</head>

<body>
<div style="background-color:#EEFFFF">
<p id="p_info">aaaaaa</p>
</div>
</body>

</html>
```

The program file: main.4gl

```
IMPORT util

MAIN
  DEFINE rec RECORD
             mywebcomp STRING,
             result STRING
         END RECORD
  DEFINE cust RECORD
             id INTEGER,
             name STRING,
             crea DATE
         END RECORD

  OPEN FORM f FROM "form"
  DISPLAY FORM f

  LET cust.id = 999
  LET cust.name = "Mike TORN"
  LET cust.crea = MDY(12,24,1998)

  INPUT BY NAME rec.* ATTRIBUTES(UNBUFFERED)
    ON ACTION get_title
       CALL ui.Interface.frontCall("webcomponent","getTitle",
               ["formonly.mywebcomp"],
               [rec.result])
    ON ACTION echo_string
       CALL ui.Interface.frontCall("webcomponent","call",
               ["formonly.mywebcomp","echoString",
                "Hello!"],
               [rec.result])
    ON ACTION echo_fields
       CALL ui.Interface.frontCall("webcomponent","call",
               ["formonly.mywebcomp","echoCustFields",
                cust.id, cust.name, cust.crea ],
               [rec.result])
    ON ACTION echo_record
       CALL ui.Interface.frontCall("webcomponent","call",
               ["formonly.mywebcomp","echoCustRecord", cust ],
               [rec.result])
    ON ACTION echo_json_string
       CALL ui.Interface.frontCall("webcomponent","call",
               ["formonly.mywebcomp","echoJSONString",
                util.JSON.stringify(cust) ],
               [rec.result])
  END INPUT

END MAIN
```
