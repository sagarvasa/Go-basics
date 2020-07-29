# struct - Summary

1. Array, Slices, Map holds data of same types but using struct we can create different datatypes in one logical entity

2. They are copy by value by default

3. They are normally created as type but we can create anonymous struct too

4. There is no Inheritance in struct but we can embed one struct in other hence forming composition

5. Tags can be added to struct field and retrive using `reflect` library to perform variery of operations

6. Way to create named struct
    type structName struct {
        fieldname dataType
        fieldname2 dataType2
        fieldname3 dataType3
    }

7.  initialize the struct
        employee := structName {
                    fieldname:    123,
                    fieldname2:  "xyz",
                    fieldname3: []string{"technology", "product"},
                }

8. anonymous struct creation and initialization
anonymousStruct := struct{ fieldname dataType }{initializationDataField: initializationDataValue}
