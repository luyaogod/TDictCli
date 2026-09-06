---
title: "Step 1: Use the I4GL function and the I4GL.4cf configuration file"
source: "fgl-topics/c_gws_i4gl_migration_guide_003.html"
breadcrumb: "Web services > SOAP Web Services > How Do I ... ? > How to migrate I4GL web service to Genero > Migrate an I4GL web service provider to Genero > Step 1: Use the I4GL function and the I4GL.4cf configuration file"
type: "concept"
---

# Step 1: Use the I4GL function and the I4GL.4cf configuration file

> The I4GL .4cf configuration file has all the information you need about the I4GL Web service.

For example, the I4GL zipcode demo has the following .4cf configuration
file:

```
[SERVICE]
	TYPE               = publisher
	INFORMIXDIR        = /dbs/32bits/ifx/11.70.uc2
	DATABASE           = i4glsoa
	CLIENT_LOCALE      = en_US.8859-1
	DB_LOCALE          = en_US.8859-1
	INFORMIXSERVER     = ol_moscou1170uc2
	HOSTNAME           = moscou.strasbourg.4js.com
	PORTNO             = 9876
	I4GLVERSION        = 7.50.xC4
	WSHOME             = /dbs/32bits/ifx/11.70.uc2/AXIS2C
	WSVERSION          = AXIS1.5
	TMPDIR             = /tmp/zipcodedemo
	SERVICENAME        = ws_zipcode
	[FUNCTION]
		NAME           = zipcode_details
		[INPUT]
			[VARIABLE]NAME = pin TYPE = CHAR(10)[END-VARIABLE]
		[END-INPUT]
		[OUTPUT]
			[VARIABLE]NAME = city TYPE = CHAR(100)[END-VARIABLE]
			[VARIABLE]NAME = state TYPE = CHAR(100)[END-VARIABLE]
		[END-OUTPUT]
	[END-FUNCTION]
	[DIRECTORY]
		NAME           = /home/f4gl/fg/i4gl
		FILE           = soademo.4gl,
	[END-DIRECTORY]
[END-SERVICE]
```

Then simply copy your I4GL function without any modification into a new Genero file and add the
Genero `IMPORT com` instruction at the beginning of the file.

For example, the I4GL SOA demo contains the zipcode\_details service
(soademo.4gl)

```
IMPORT com

FUNCTION zipcode_details(pin)
		DEFINE state_rec RECORD
				     pin CHAR(10),
				     city CHAR(100),
				     state CHAR(100)
			    END RECORD,
			    pin CHAR(10),
			    sel_stmt CHAR(512);

		LET sel_stmt= "SELECT * FROM statedetails WHERE pin = ?";
		PREPARE st_id FROM sel_stmt;
		DECLARE cur_id CURSOR FOR st_id;
		OPEN cur_id USING pin;
		FETCH cur_id INTO state_rec.*;
		CLOSE cur_id;
		FREE cur_id;
		FREE st_id;
		RETURN state_rec.city, state_rec.state
		
END FUNCTION
```

You may need to make some minor code modification for compatibility.
