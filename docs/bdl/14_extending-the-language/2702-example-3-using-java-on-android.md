---
title: "Example 3: Using Java on Android"
source: "fgl-topics/c_fgl_JavaBridge_example_3.html"
breadcrumb: "Extending the language > The Java interface > Examples > Example 3: Using Java on Android™"
type: "concept"
description: "This example shows how to access Android™ components through Java, it includes: Access to the JDK API to get the number of cores on your device. Access to Android APIs to get the screen dimension, the ..."
---

# Example 3: Using Java on Android

This example shows how to access Android™ components
through Java, it includes:

- Access to the JDK API to get the number of cores on your device.
- Access to Android APIs to get the screen dimension, the
  device manufacturer and model (with no need for any additional authorization)
- Access to the Bluetooth stack to list the paired devices.
  > **Note:**
  >
  > In your GM project, you need to
  > ask for BLUETOOTH authorization.

Form file
formJavaStandard.per:

```
LAYOUT (TEXT="Access to Android API")
GROUP group1(TEXT="Using standard JDK API...")
GRID grid1
{
[l1                        |f1                                  ]    
}
END
END

ATTRIBUTES
LABEL l1 : label1, TEXT="Number of processors available";
LABEL f1 = FORMONLY.nb_proc;
END
```

Form file formAndroidSimple.per:

```
LAYOUT (TEXT="Access to Android API")
GROUP group1(TEXT="Using simple Android API...")
GRID grid1
{
[l1                        |f1                                  ]
[l2                        |f2                                  ]
[l3                        |f3                                  ]
[l4                        |f4                                  ]
}
END
END

ATTRIBUTES
LABEL l1 : label1, TEXT="Device manufacturer";
LABEL f1 = FORMONLY.manufacturer;
LABEL l2 : label2, TEXT="Device model";
LABEL f2 = FORMONLY.model;
LABEL l3 : label3, TEXT="Device serial number";
LABEL f3 = FORMONLY.serial;
LABEL l4 : label4, TEXT="Device screen dimension";
LABEL f4 = FORMONLY.diagonal;
END
```

Form file formAndroidBluetooth.per:

```
LAYOUT (TEXT="Access to Android API")
GROUP group1(TEXT="Using Bluetooth Android API...")
GRID grid1
{
[l1                            |f1                                       ]
<TABLE t                                                                 >
[c1                            |c2                                       ]
[c1                            |c2                                       ]
[c1                            |c2                                       ]
<                                                                        >
}
END
END

ATTRIBUTES
LABEL l1 : label1, TEXT="Bluetooth adapter name";
LABEL f1 = FORMONLY.ba_name;
LABEL c1 = FORMONLY.name;
LABEL c2 = FORMONLY.comment;
END

INSTRUCTIONS
SCREEN RECORD list(FORMONLY.name, FORMONLY.comment);
END
```

Program file:

```
IMPORT util

IMPORT JAVA java.lang.Runtime
IMPORT JAVA java.util.Iterator
IMPORT JAVA java.lang.Class
IMPORT JAVA java.lang.Math

IMPORT JAVA android.bluetooth.BluetoothAdapter
IMPORT JAVA android.bluetooth.BluetoothDevice
IMPORT JAVA android.content.Context
IMPORT JAVA android.os.Build
IMPORT JAVA android.util.DisplayMetrics
IMPORT JAVA android.view.WindowManager

IMPORT JAVA com.fourjs.gma.vm.FglRun

MAIN
	MENU "Samples"
		COMMAND "Android API access"
			CALL androidApiAccess()
		COMMAND "Quit"
			EXIT MENU
		ON ACTION close
			EXIT MENU
	END MENU
END MAIN

FUNCTION androidApiAccess()

	MENU "Android API access"
		COMMAND "Accessing Java standard API"
			CALL androidApiAccess_java_standard()
		COMMAND "Accessing simple android information"
			CALL androidApiAccess_android_simple()
		COMMAND "Accessing sophisticated APIs : bluetooth"
			CALL androidApiAccess_bluetooth()
		ON ACTION CANCEL
			EXIT MENU
	END MENU
END FUNCTION

FUNCTION androidApiAccess_java_standard()
	DEFINE r Runtime

	OPEN WINDOW w WITH FORM "formJavaStandard"

	LET r = java.lang.Runtime.getRuntime()
	DISPLAY r.availableProcessors() TO nb_proc

	MENU
		ON ACTION QUIT
			EXIT MENU
		ON ACTION close
			EXIT MENU
	END MENU

	CLOSE WINDOW w
END FUNCTION

FUNCTION androidApiAccess_android_simple()
	DEFINE s STRING
	DEFINE dm DisplayMetrics
	DEFINE c Context
	DEFINE width, height, dens, wi, hi, x, y FLOAT
	DEFINE screenInches FLOAT
	DEFINE wm android.view.WindowManager

	OPEN WINDOW w WITH FORM "formAndroidSimple"

	LET s = android.os.Build.MANUFACTURER
	DISPLAY s TO manufacturer
	LET s = android.os.Build.MODEL
	DISPLAY s TO model
	LET s = android.os.Build.SERIAL
	DISPLAY s TO serial

	# Get the FglRun Context
	LET c = com.fourjs.gma.vm.FglRun.getContext()

	# Compute display dimension (diagonal)
	LET dm = android.util.DisplayMetrics.create()
	LET wm = CAST ( c.getSystemService("window") AS android.view.WindowManager )
	CALL wm.getDefaultDisplay().getMetrics(dm)
	LET width = dm.widthPixels
	LET height = dm.heightPixels
	LET dens = dm.densityDpi
	LET wi = width/dens
	LET hi = height/dens
	LET x = util.Math.pow(wi,2)
	LET y = util.Math.pow(hi,2);
	LET screenInches = util.Math.sqrt(x+y);

	DISPLAY screenInches TO diagonal
	MENU
		ON ACTION QUIT
			EXIT MENU
		ON ACTION close
			EXIT MENU
	END MENU

	CLOSE WINDOW w
END FUNCTION

FUNCTION androidApiAccess_bluetooth()
	DEFINE ba  BluetoothAdapter
	DEFINE sbd Iterator
	DEFINE bd  BluetoothDevice
	DEFINE bds DYNAMIC ARRAY OF RECORD
		name STRING,
		comment STRING
	END RECORD
	DEFINE i INTEGER
	DEFINE s STRING

	OPEN WINDOW w WITH FORM "formAndroidBluetooth"

	LET ba = android.bluetooth.BluetoothAdapter.getDefaultAdapter()
	LET s = ba.getName()
	DISPLAY s TO ba_name

	LET sbd = ba.getBondedDevices().iterator()
	LET i = 0
	WHILE sbd.hasNext()
		LET bd = CAST(sbd.next() AS BluetoothDevice)
		LET i = i + 1
		LET bds[i].name = bd.getName()
		LET bds[i].comment = bd.getBluetoothClass().toString()
	END WHILE

	DISPLAY ARRAY bds TO list.*
		ON ACTION QUIT
			EXIT DISPLAY
		ON ACTION close
			EXIT DISPLAY
	END DISPLAY

	CLOSE WINDOW w
END FUNCTION
```
