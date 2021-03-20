#运算符比较库

## 包含 （in）
### 支持的类型
- String
- Int32
- Int64
- Float32
- Float64
### 使用
```
boolValue := xcompare.IN.String("a",[]string{"a","b"})
fmt.Println(boolValue)
//true
```

## 等于（equal）
### 支持的类型
- String
- Int32
- Int64
- Float32
- Float64
### 使用
```
boolValue := xcompare.Equal.String("a","a")
fmt.Println(boolValue)
//true
```

## 不等于（not equal）
### 支持的类型
- String
- Int32
- Int64
- Float32
- Float64
### 使用
```
boolValue := xcompare.NotEqual.String("a","b")
fmt.Println(boolValue)
//true
```

## 小于（little）
### 支持的类型
- Int32
- Int64
- Float32
- Float64
### 使用
```
boolValue := xcompare.Litter.Int64(1,2)
fmt.Println(boolValue)
//true
```

## 小于等于（little equal）
### 支持的类型
- Int32
- Int64
- Float32
- Float64
### 使用
```
boolValue := xcompare.LitterEqual.Int64(1,2)
fmt.Println(boolValue)
//true
```

## 大于（great）
### 支持的类型
- Int32
- Int64
- Float32
- Float64
### 使用
```
boolValue := xcompare.Great.Int64(2,1)
fmt.Println(boolValue)
//true
```


## 大于等于（great equal）
### 支持的类型
- Int32
- Int64
- Float32
- Float64
### 使用
```
boolValue := xcompare.GreatEqual.Int64(2,1)
fmt.Println(boolValue)
//true