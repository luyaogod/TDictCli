---
title: "Implement Android activities in GMA"
source: "fgl-topics/c_fgl_JavaBridge_gma_activity.html"
breadcrumb: "Extending the language > The Java interface > Advanced programming > Executing Java code with GMA > Implement Android™ activities in GMA"
type: "concept"
description: "Android activities can be bundled with your GMA app and called from the Genero code."
---

# Implement Android activities in GMA

> Android™ activities can be bundled with your GMA app and called from the Genero code.

A Java-based extension that interacts with the end user must be implemented as an Android Activity, by using the [`android.app.Android`](http://developer.android.com/reference/android/app/Activity.html) class.

In order to use your Android Activity
from the program, it must be integrated in the mobile app Android package (.apk), which is created in the Genero
Studio deployment procedure.

This code example implements a simple Android Activity:

```
package com.myextension;

import android.app.Activity;
import android.os.Bundle;
import android.widget.TextView;

public class MyActivity extends Activity {
    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        Button button = new Button(this);
        button.setText("Quit");
        setContentView(button);
        button.setOnClickListener(
             new View.OnClickListener() {
             public void onClick(View v) {
                 int resultCode = RESULT_OK;
                 Intent resultData = new Intent();
                 resultData.putExtra("MyKey", "MyValue");
                 setResult(resultCode, resultData);
                 finish();
             }
        });
    }
}
```

In order to execute this activity from a Genero app, use the [`startActivity`](../15_library-reference/3467-android-startactivity.md "Starts an external Android application (activity), and returns to the GMA application immediately.") front
call:

```
MAIN
    DEFINE data, extras STRING
    MENU
        ON ACTION activity ATTRIBUTES(TEXT="Call bundled activity")
            CALL ui.Interface.frontCall("android", "startActivityForResult",
                 ["android.intent.action.VIEW", NULL, NULL, NULL,
                  "com.myextension.MyActivity"],
                 [ data, extras ])
                 MESSAGE "data=",data," / extras=",extras
        ON ACTION quit
            EXIT MENU
    END MENU
END MAIN
```

The component name (fifth parameter) of the startActivity front call does normally take the APK
package name followed by the Java Activity class name
(`apk-package-name/java-class-name`).
The APK Android package name can
be defined for the application project in Genero Studio. When using an
user-defined activity that is part of the GMA binary archive, do not specify the APK
package in the component parameter, because the Java Activity class will be included in the
current APK package. This is true when using the customized GMA front-end in development
mode, and in the final application that is deployed on the device. For more details about
the component parameter, see [android.startActivity](../15_library-reference/3467-android-startactivity.md "Starts an external Android application (activity), and returns to the GMA application immediately.").

## Related links

**Related concepts**  

[Packaging custom Java extensions for GMA](2697-packaging-custom-java-extensions-for-gma.md "Custom Java extension must be integrated in the GMA to run on Android devices.")
