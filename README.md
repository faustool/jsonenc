#jsonenc
Go library to encode Json file using a Stream API. The intention when compared to the available Json Encoder is 
to be extremely fast.

##Example 1
The following code:
```golang
	buffer := bytes.NewBufferString("")
	stream := NewJsonStream(buffer)
	stream.WriteStartObject() // start json object
	stream.WriteStartObjectWithName("my-object")
	stream.WriteNameValueString("my-name", "my-value")
	stream.WriteEndObject() // my-object object
	stream.WriteEndObject() // json object
```

Will produce the Json object below:
```json
{  
   "my-object":{  
      "my-name":"my-value"
   }
}
```

##Example 2
Arrays and inner objects can be created as follows:
```golang
	buffer := bytes.NewBufferString("")
	stream := NewJsonStream(buffer)
	stream.WriteStartObject() // start json object
	stream.WriteStartObjectWithName("my-object")
	stream.WriteNameValueString("my-name", "my-value")
	stream.WriteNameValueLiteral("my-int", "10")
	stream.WriteNameValueLiteral("my-bool", "false")
	stream.WriteStartArrayWithName("my-array")
	stream.WriteStringValue("value 1")
	stream.WriteStringValue("value 2")
	stream.WriteStringValue("value 3")
	stream.WriteEndArray() // my-array
	stream.WriteStartObjectWithName("inner-object")
	stream.WriteNameValueString("name", "value")
	stream.WriteEndObject() // inner-object
	stream.WriteEndObject() // my-object object
	stream.WriteEndObject() // json object
```

The code above produces this Json object:
```json
{  
   "my-object":{  
      "my-name":"my-value",
      "my-int":10,
      "my-bool":false,
      "my-array":[  
         "value 1",
         "value 2",
         "value 3"
      ],
      "inner-object":{  
         "name":"value"
      }
   }
}
```