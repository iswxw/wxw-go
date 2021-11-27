## ` package/encoding/json` 

### 初识`json` 

```json
Package json implements encoding and decoding of JSON as defined in RFC 7159. The mapping between JSON and Go values is described in the documentation for the Marshal and Unmarshal functions.
```

- 出处：https://golang.google.cn/pkg/encoding/json/

### 核心函数

#### 1. ` func Marshal    ` 

##### 1.1 使用 ` Marshal` 

```go
// Marshal returns the JSON encoding of v.
func Marshal(v interface{}) ([]byte, error)
```

**使用示例** 

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

// User 定义一个结构体
// 首字母大写：表示 public
// 首字母小写：表示 private
type User struct {
	UserName string    `json:"user_name,omitempty"` // omitempty 表示忽略空值,
	Age      int       `json:"age"`
	Age1     int       `json:"-"`                   //`json:"-"` 表示不进行序列化,忽略这个字段
	Gender   string    `json:"gender,string"`
	Birthday time.Time `json:"birthday" `
}

func NewUser() *User {
	return &User{
		UserName: "Java半颗糖",
		Age:      18,
		Gender:   "男",
	}
}

func NewEmptyUser() *User {
	return &User{}
}

func main() {
	user := NewUser()
	fmt.Printf("user: %#v\n",user)

	// marshal
	byteUser, err := json.Marshal(user)
	if err != nil {
		 log.Fatal("json marshal error:",err)
	}
	fmt.Printf(" user: %#v\n",byteUser)

	// unmarshal
	user1 := NewEmptyUser()
	if err = json.Unmarshal(byteUser, &user1); err != nil {
		log.Fatal("json unmarshal error:",err)
	}
	fmt.Printf(" user: %#v\n",user1)
}

```

##### 1.2 `marshal struct` 

` Examples of struct field tags and their meanings:` 

```go
// Field appears in JSON as key "myName".
Field int `json:"myName"`

// Field appears in JSON as key "myName" and
// the field is omitted from the object if its value is empty,
// as defined above.
Field int `json:"myName,omitempty"`

// Field appears in JSON as key "Field" (the default), but
// the field is skipped if empty.
// Note the leading comma. (注意前面的逗号)
Field int `json:",omitempty"`

// Field is ignored by this package.（这个包忽略这个字段）
Field int `json:"-"`

// Field appears in JSON as key "-".（这个字段在json中显示key是"-"） 
Field int `json:"-,"`
```

**序列化时字段限定 tag 标签：** 

- `json:"myName"` 指定 序列化后在 ` json` 字符串中显示的 key 名称
- `json:"-"` 表示不进行序列化，这样在 marshal 后 可以看见值
- `json:"myName,omitempty"` 表示 序列化的时候，忽略控制，也不展示在序列化后的 ` json`字符串中
- ` json:"product_id,string"` 表示序列化时指定 key 的类型，帮助转换为指定类型

**案例分析** 

- 一个结构体正常序列化过后是什么样的呢？ 

  ```go
  package main
  import (
      "encoding/json"
      "fmt"
  )
  
  // Product 商品信息
  type Product struct {
      Name      string
      ProductID int64
      Number    int
      Price     float64
      IsOnSale  bool
  }
  
  func main() {
      p := &Product{}
      p.Name = "Xiao mi 6"
      p.IsOnSale = true
      p.Number = 10000
      p.Price = 2499.00
      p.ProductID = 1
      data, _ := json.Marshal(p)
      fmt.Println(string(data))
  }
  
  
  //结果
  {"Name":"Xiao mi 6","ProductID":1,"Number":10000,"Price":2499,"IsOnSale":true}
  ```

- 何为Tag，tag就是标签，给结构体的每个字段打上一个标签，标签冒号前是类型，后面是标签名。

  ```go
  // Product _
  type Product struct {
      Name      string  `json:"name"`
      ProductID int64   `json:"-"` // 表示不进行序列化
      Number    int     `json:"number"`
      Price     float64 `json:"price"`
      IsOnSale  bool    `json:"is_on_sale,string"`
  }
  
  // 序列化过后，可以看见
  {"name":"Xiao mi 6","number":10000,"price":2499,"is_on_sale":"false"}
  ```

- omitempty，tag里面加上omitempy，可以在序列化的时候忽略0值或者空值

  ```go
  package main
  
  import (
      "encoding/json"
      "fmt"
  )
  
  // Product _
  type Product struct {
      Name      string  `json:"name"`
      ProductID int64   `json:"product_id,omitempty"` 
      Number    int     `json:"number"`
      Price     float64 `json:"price"`
      IsOnSale  bool    `json:"is_on_sale,omitempty"`
  }
  
  func main() {
      p := &Product{}
      p.Name = "Xiao mi 6"
      p.IsOnSale = false
      p.Number = 10000
      p.Price = 2499.00
      p.ProductID = 0
  
      data, _ := json.Marshal(p)
      fmt.Println(string(data))
  }
  // 结果
  {"name":"Xiao mi 6","number":10000,"price":2499}
  ```

- type，有些时候，我们在序列化或者反序列化的时候，可能结构体类型和需要的类型不一致，这个时候可以指定,支持string,number和boolean

  ```go
  package main
  
  import (
      "encoding/json"
      "fmt"
  )
  
  // Product _
  type Product struct {
      Name      string  `json:"name"`
      ProductID int64   `json:"product_id,string"`
      Number    int     `json:"number,string"`
      Price     float64 `json:"price,string"`
      IsOnSale  bool    `json:"is_on_sale,string"`
  }
  
  func main() {
      var data = `{"name":"Xiao mi 6","product_id":"10","number":"10000","price":"2499","is_on_sale":"true"}`
      p := &Product{}
      err := json.Unmarshal([]byte(data), p)
      fmt.Println(err)
      fmt.Println(*p)
  }
  // 结果
  <nil>
  {Xiao mi 6 10 10000 2499 true}
  ```

##### 1.3 ` marshal` 源码

```go
// Marshal returns the JSON encoding of v.
//
// Marshal traverses the value v recursively.
// If an encountered value implements the Marshaler interface
// and is not a nil pointer, Marshal calls its MarshalJSON method
// to produce JSON. If no MarshalJSON method is present but the
// value implements encoding.TextMarshaler instead, Marshal calls
// its MarshalText method and encodes the result as a JSON string.
// The nil pointer exception is not strictly necessary
// but mimics a similar, necessary exception in the behavior of
// UnmarshalJSON.
//
// Otherwise, Marshal uses the following type-dependent default encodings:
//
// Boolean values encode as JSON booleans.
//
// Floating point, integer, and Number values encode as JSON numbers.
//
// String values encode as JSON strings coerced to valid UTF-8,
// replacing invalid bytes with the Unicode replacement rune.
// So that the JSON will be safe to embed inside HTML <script> tags,
// the string is encoded using HTMLEscape,
// which replaces "<", ">", "&", U+2028, and U+2029 are escaped
// to "\u003c","\u003e", "\u0026", "\u2028", and "\u2029".
// This replacement can be disabled when using an Encoder,
// by calling SetEscapeHTML(false).
//
// Array and slice values encode as JSON arrays, except that
// []byte encodes as a base64-encoded string, and a nil slice
// encodes as the null JSON value.
//
// Struct values encode as JSON objects.
// Each exported struct field becomes a member of the object, using the
// field name as the object key, unless the field is omitted for one of the
// reasons given below.
//
// The encoding of each struct field can be customized by the format string
// stored under the "json" key in the struct field's tag.
// The format string gives the name of the field, possibly followed by a
// comma-separated list of options. The name may be empty in order to
// specify options without overriding the default field name.
//
// The "omitempty" option specifies that the field should be omitted
// from the encoding if the field has an empty value, defined as
// false, 0, a nil pointer, a nil interface value, and any empty array,
// slice, map, or string.
//
// As a special case, if the field tag is "-", the field is always omitted.
// Note that a field with name "-" can still be generated using the tag "-,".
//
// Examples of struct field tags and their meanings:
//
//   // Field appears in JSON as key "myName".
//   Field int `json:"myName"`
//
//   // Field appears in JSON as key "myName" and
//   // the field is omitted from the object if its value is empty,
//   // as defined above.
//   Field int `json:"myName,omitempty"`
//
//   // Field appears in JSON as key "Field" (the default), but
//   // the field is skipped if empty.
//   // Note the leading comma.
//   Field int `json:",omitempty"`
//
//   // Field is ignored by this package.
//   Field int `json:"-"`
//
//   // Field appears in JSON as key "-".
//   Field int `json:"-,"`
//
// The "string" option signals that a field is stored as JSON inside a
// JSON-encoded string. It applies only to fields of string, floating point,
// integer, or boolean types. This extra level of encoding is sometimes used
// when communicating with JavaScript programs:
//
//    Int64String int64 `json:",string"`
//
// The key name will be used if it's a non-empty string consisting of
// only Unicode letters, digits, and ASCII punctuation except quotation
// marks, backslash, and comma.
//
// Anonymous struct fields are usually marshaled as if their inner exported fields
// were fields in the outer struct, subject to the usual Go visibility rules amended
// as described in the next paragraph.
// An anonymous struct field with a name given in its JSON tag is treated as
// having that name, rather than being anonymous.
// An anonymous struct field of interface type is treated the same as having
// that type as its name, rather than being anonymous.
//
// The Go visibility rules for struct fields are amended for JSON when
// deciding which field to marshal or unmarshal. If there are
// multiple fields at the same level, and that level is the least
// nested (and would therefore be the nesting level selected by the
// usual Go rules), the following extra rules apply:
//
// 1) Of those fields, if any are JSON-tagged, only tagged fields are considered,
// even if there are multiple untagged fields that would otherwise conflict.
//
// 2) If there is exactly one field (tagged or not according to the first rule), that is selected.
//
// 3) Otherwise there are multiple fields, and all are ignored; no error occurs.
//
// Handling of anonymous struct fields is new in Go 1.1.
// Prior to Go 1.1, anonymous struct fields were ignored. To force ignoring of
// an anonymous struct field in both current and earlier versions, give the field
// a JSON tag of "-".
//
// Map values encode as JSON objects. The map's key type must either be a
// string, an integer type, or implement encoding.TextMarshaler. The map keys
// are sorted and used as JSON object keys by applying the following rules,
// subject to the UTF-8 coercion described for string values above:
//   - keys of any string type are used directly
//   - encoding.TextMarshalers are marshaled
//   - integer keys are converted to strings
//
// Pointer values encode as the value pointed to.
// A nil pointer encodes as the null JSON value.
//
// Interface values encode as the value contained in the interface.
// A nil interface value encodes as the null JSON value.
//
// Channel, complex, and function values cannot be encoded in JSON.
// Attempting to encode such a value causes Marshal to return
// an UnsupportedTypeError.
//
// JSON cannot represent cyclic data structures and Marshal does not
// handle them. Passing cyclic structures to Marshal will result in
// an error.
//
func Marshal(v interface{}) ([]byte, error) {
	e := newEncodeState()

	err := e.marshal(v, encOpts{escapeHTML: true})
	if err != nil {
		return nil, err
	}
	buf := append([]byte(nil), e.Bytes()...)

	encodeStatePool.Put(e)

	return buf, nil
}
```

相关资料

1. 源码：https://golang.google.cn/src/encoding/json/encode.go?s=6458:6501#L148 

#### 2. `func unMarshal` 

##### 2.1 使用 ` unmarshal` 

```go
func Unmarshal(data []byte, v interface{}) error
```





##### 2.2 `unmarshal struct` 

##### 2.3 `unmarshal 源码` 

```go
// Unmarshal parses the JSON-encoded data and stores the result
// in the value pointed to by v. If v is nil or not a pointer,
// Unmarshal returns an InvalidUnmarshalError.
//
// Unmarshal uses the inverse of the encodings that
// Marshal uses, allocating maps, slices, and pointers as necessary,
// with the following additional rules:
//
// To unmarshal JSON into a pointer, Unmarshal first handles the case of
// the JSON being the JSON literal null. In that case, Unmarshal sets
// the pointer to nil. Otherwise, Unmarshal unmarshals the JSON into
// the value pointed at by the pointer. If the pointer is nil, Unmarshal
// allocates a new value for it to point to.
//
// To unmarshal JSON into a value implementing the Unmarshaler interface,
// Unmarshal calls that value's UnmarshalJSON method, including
// when the input is a JSON null.
// Otherwise, if the value implements encoding.TextUnmarshaler
// and the input is a JSON quoted string, Unmarshal calls that value's
// UnmarshalText method with the unquoted form of the string.
//
// To unmarshal JSON into a struct, Unmarshal matches incoming object
// keys to the keys used by Marshal (either the struct field name or its tag),
// preferring an exact match but also accepting a case-insensitive match. By
// default, object keys which don't have a corresponding struct field are
// ignored (see Decoder.DisallowUnknownFields for an alternative).
//
// To unmarshal JSON into an interface value,
// Unmarshal stores one of these in the interface value:
//
//	bool, for JSON booleans
//	float64, for JSON numbers
//	string, for JSON strings
//	[]interface{}, for JSON arrays
//	map[string]interface{}, for JSON objects
//	nil for JSON null
//
// To unmarshal a JSON array into a slice, Unmarshal resets the slice length
// to zero and then appends each element to the slice.
// As a special case, to unmarshal an empty JSON array into a slice,
// Unmarshal replaces the slice with a new empty slice.
//
// To unmarshal a JSON array into a Go array, Unmarshal decodes
// JSON array elements into corresponding Go array elements.
// If the Go array is smaller than the JSON array,
// the additional JSON array elements are discarded.
// If the JSON array is smaller than the Go array,
// the additional Go array elements are set to zero values.
//
// To unmarshal a JSON object into a map, Unmarshal first establishes a map to
// use. If the map is nil, Unmarshal allocates a new map. Otherwise Unmarshal
// reuses the existing map, keeping existing entries. Unmarshal then stores
// key-value pairs from the JSON object into the map. The map's key type must
// either be any string type, an integer, implement json.Unmarshaler, or
// implement encoding.TextUnmarshaler.
//
// If a JSON value is not appropriate for a given target type,
// or if a JSON number overflows the target type, Unmarshal
// skips that field and completes the unmarshaling as best it can.
// If no more serious errors are encountered, Unmarshal returns
// an UnmarshalTypeError describing the earliest such error. In any
// case, it's not guaranteed that all the remaining fields following
// the problematic one will be unmarshaled into the target object.
//
// The JSON null value unmarshals into an interface, map, pointer, or slice
// by setting that Go value to nil. Because null is often used in JSON to mean
// ``not present,'' unmarshaling a JSON null into any other Go type has no effect
// on the value and produces no error.
//
// When unmarshaling quoted strings, invalid UTF-8 or
// invalid UTF-16 surrogate pairs are not treated as an error.
// Instead, they are replaced by the Unicode replacement
// character U+FFFD.
//
func Unmarshal(data []byte, v interface{}) error {
	// Check for well-formedness.
	// Avoids filling out half a data structure
	// before discovering a JSON syntax error.
	var d decodeState
	err := checkValid(data, &d.scan)
	if err != nil {
		return err
	}

	d.init(data)
	return d.unmarshal(v)
}
```

相关资料

1. 源码：https://golang.google.cn/src/encoding/json/decode.go?s=4081:4129#L86

