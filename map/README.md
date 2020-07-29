# map - Summary

1. Multiple way to use map
    1. Using built in make function (recommended)
    e := make(map[keyDataType]valueDataType) then e = map[keyDataType]valueDataType{...key,values}
    2. var employeeMap = map[keyDataType]valueDataType{...key,values}
    3. employeeMap := map[keyDataType]valueDataType{...key,values}

2. They are copy by reference by default

3. We can use built in functions like len(), delete() to identify length and delete key from map respectively

4. Members can be accessed using [key] (e.g: syntax map["key"] = "value")

5. If key is not found in map, then zero value of valueDataType will be returned

6. We can use second variable to check whether key is found (val = true) or not (val = false) in map (byConvention ok variable is used)
