## Maps
A map in Golang is the language implementation of an 
[associative array](https://brilliant.org/wiki/associative-arrays/) consisting
of key: value pairs, and give the user the ability to retrieve values by specifying
a key, without having to iterate through an entire collection.

#### Key Takeaways About Maps
 * Maps are a complex data type that allow fast lookups of values using a key and are implemented using a hash table.
 * If you preform a lookup on a key that does not exist, the map will return the **zero value** for the values type.
 * Map iterations are not consistent and will not be in the order the map is defined.
 * Maps are a reference type. Meaning a map shares the same properties as a pointer.

#### Using Maps in Practice
##### Creating a Map
To create a map, you can either initialize it by assigning an empty declaration:
```go
var myMap = map[string]string{}
```
Or by using the built-in [make](https://github.com/josh5276/go-course/tree/master/reference#the-make-built-in)
function, which gives you the ability to initialize a variable with a preallocated amount of memory:
Note: Memory is allocated into up-to 8 key:value pairs called hash buckets. If more 
values are needed, Go automatically allocates new buckets.
```go
var myMap = make(map[string]string, 8)
```

Alternatively, like other Go types, you can assign data to the map at initialization:
```go
var platformMap = map[string]string{
    "Cisco ASA 5505": "asa",
    "Cisco ASA 5510": "asa",
    "Cisco ASA 5508": "asa",
    "F5- Big-IP 2000s": "f5",
    "F5- Big-IP 1600": "f5",
    "F5- Big-IP 2200": "f5",
    "F5- Big-IP 5000v": "f5",
    "Brocade ADX 1000": "adx",
}
```

##### Setting Values
Assigning a new key:value to a map is very similar to methods of other languages 
with similar associative arrays. To set a new value:
```go
platformMap["F5- Big-IP i2600"] = "f5"
```
##### Getting Values
Getting an existing value from a map is also very similar to methods of other languages 
with similar associative arrays. To get a value:
```go
value := platformMap["F5- Big-IP i2600"]
```
It's important to note that you can expect maps to return the empty value of the value type if a key
does not exist. For example, if a key from `map[string]string` does not exist, the value returned will be `""`
_the empty value for a string_. As an example

```go 
func main() {
	ms := map[string]string{"key": "value"}
	mi := map[string]int{"key": 5}
	mb := map[string]bool{"key": true}

	fmt.Println(ms["different_key"])
	fmt.Println(mi["different_key"])
	fmt.Println(mb["different_key"])
}
Result:
---------------
""   
0
false
```
There are scenarios where the actual value of a key is the zero value of a type. Because of this, it is usually better 
to use Go's two-value assignment of a key retrieval.
```go
value, exist := ms["different_key"]
```
Using a two-value retrieval, the first value is the actual value returned from the key lookup (or it's zero value). 
The second value will be a boolean that returns `true` if the key exist or `false` if it does not exist. 

We can now check to see if a key exist in a map prior to using it. Or just check if a key exist.
```go
func main() {
	m := map[string]string{"key1": "value1"}

	// Check if key1 exist, if so, print the value of key1.
	if val, ok := m["key1"]; ok {
		fmt.Printf("found key1 with value of %s\n", val)
	}

	// Check if key2 exist, if so, print the value of key2.
	// In this scenario, this value would not exist, so the print statement
	// would never get executed.
	if val, ok := m["key2"]; ok {
		fmt.Printf("found key2 with value of %s\n", val)
		return
	}

	// Add key2 to the map and perform another key2 check.
	m["key2"] = "value2"
	if val, ok := m["key2"]; ok {
		fmt.Printf("added and found key2 with value of %s\n", val)
	}
}

Result:
----------------
found key1 with value of value1
added and found key2 with value of value2
```

##### Deleting a Key
Go has a `delete()` function built into the language to remove a value from a key.
From the [builtin source](https://github.com/golang/go/blob/default/src/builtin/builtin.go#L146):

> The delete built-in function deletes the element with the specified key from 
the map. If the key value is nil or there is no such element, delete is a no-op.
```go
delete(platformMap, "F5- Big-IP i2600")
```

##### Iterating through a Map
Iterations are pretty simple here and very similar to the way you would iterate 
a slice. The for/range syntax assigns a key and value variable for each iteration that is
only accessible inside of the loop.
```go
var platformMap = map[string]string{
	"Cisco ASA 5505":   "asa",
	"Cisco ASA 5510":   "asa",
	"Cisco ASA 5508":   "asa",
	"F5- Big-IP 2000s": "f5",
	"F5- Big-IP 1600":  "f5",
	"F5- Big-IP 2200":  "f5",
	"F5- Big-IP 5000v": "f5",
	"Brocade ADX 1000": "adx",
}

for key, value := range platformMap {
	fmt.Printf("%-25s %s\n", key, strings.ToUpper(value))
}
```
+ [Go Playground](https://play.golang.org/p/HEX1rnnbprq)

#### More About Map Iteration
One important note about map iteration is that the iteration order is not consistent. That is, each time
you iterate through a map, the order in which range retrieves a key:value pair could be different. 
**It is important that you do not rely on map ordered iteration**. 

You can see this in a simple example of map iteration:
```go
func main() {
	intMap := make(map[int]string)
	intMap[0] = "value0"
	intMap[1] = "value1"
	intMap[2] = "value2"
	intMap[3] = "value3"

	for key, value := range intMap {
		fmt.Printf("%d:%s\n", key, value)
	}
}
```
Where the output of this snippet of code would differ on each run:
```go
Run 1:           Run 2:           Run 3:
----------       ----------       ----------
1:value1         2:value2         0:value0
2:value2         3:value3         1:value1
3:value3         0:value0         2:value2
0:value0         1:value1         3:value3
```
So why is map iteration not consistent? Because the language team intentionally designed maps, so that it is
random in order. Previous to the Go 1 release, map iteration was not forcibly random, but because of the underlying 
system, could vary iteration order depending on the actual implementation. So instead of trying to be overly flexible 
and allow programming that would be considered sloppy and inconsistent, Go forces you to handle the sorting of your 
map iteration (if needed) in your code, thus making it rock solid. 

Specifically stated in the [Go1 Release Notes](https://golang.org/doc/go1#iteration):
> This caused tests that iterated over maps to be fragile and non-portable, with the unpleasant property that a 
test might always pass on one machine but break on another.

Below is an example of how hashing works with maps in Go. When a lookup on a key is performed, the language will:
1. Generate the hash-key from the key that you specify to lookup (`platformMap["Cisco ASA 5505"]` as an example)
2. Use the low-order of bits from the hash to determine the bucket the value lives in.
3. Use the high-order of bits from the same hash to select the element in the HOB array that references the key:value you are looking for.
4. Finally, the last array is just a slice of bytes that actually contain the key:value data.

![Hash Table Example](https://github.com/josh5276/go-course/blob/default/common/img/Map_Hash_Table.png)
 - _[From: Macro View of Map Internals](https://www.ardanlabs.com/blog/2013/12/macro-view-of-map-internals-in-go.html)_

#### How to Iterate in Order
If you need to iterate a map in a specific order, you can do so by creating a slice 
of keys, then ordering the keys in the manner you want, and finally iterate through the
keys to gather the map value for each key.

```go
var keys []string
for k := range platformMap {
	keys = append(keys, k)
}
sort.Strings(keys)

for _, key := range keys {
	fmt.Printf("%-25s %s\n", key, strings.ToUpper(platformMap[key]))
}
```

+ [Go Playground](https://play.golang.org/p/ojiEG22k1q-)

#### References
+ [Maps In Action](https://blog.golang.org/go-maps-in-action)
+ [Go Runtime#Map](https://github.com/golang/go/blob/default/src/runtime/map.go)
+ [How the Go Runtime Implements Maps](https://dave.cheney.net/2018/05/29/how-the-go-runtime-implements-maps-efficiently-without-generics)
