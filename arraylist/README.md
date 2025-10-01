# Arraylist

This package provides a generic, type-safe and configurable `ArrayList` (dynamic array) built on a standard Go slice.


## Example

```go
package main

import (
	"fmt"
	"github.com/salsgithub/godst/arraylist"
)

func main() {
    // Raw creation
    languages := arraylist.New[string]()
    languages.Add("Go", "C++", "Java")
    fmt.Println(languages.Values()) // [Go C++ Java]

    // Create with options
    fruits := arraylist.New(
        arraylist.WithInitialValues("Apple", "Orange", "Banana"),
        arraylist.WithInitialCapacity[string](20),
        arraylist.WithGrowthFactor[string](1.5),
        arraylist.WithShrinkFactor[string](0.3),
    )
    fruits.Add("Mango")
    fmt.Println(fruits.Values()) // [Apple Orange Banana Mango]

    // Create with fixed size
    shoppingList := arraylist.New(
        arraylist.WithInitialValues("Bread", "Milk", "Yoghurt", "Tea"),
        arraylist.WithFixedSize[string](4),
    )
    shoppingList.Add("Biscuits") // Ignored
    fmt.Println(shoppingList.Values()) // [Bread Milk Yoghurt Tea]
    fmt.Println(shoppingList.Contains("Biscuits")) // False
    fmt.Println(shoppingList.Contains("Bread")) // True
    fmt.Println(shoppingList.ContainsAll("Bread", "Milk")) // True
}
```

## Core API

- `Add(values...T)`: Adds all the values to the array list
- `Remove(index int)`: Remove the value at the given index
- `Replace(index int, value T)`: Replaces the value at the given index
- `Contains(value T)`: Returns true if the value is in the array list
- `ContainsAll(values...T)`: Returns true if all the values are in the array list
- `Value(index int)`: Returns the value at the given index
- `Values()`: Returns the values of the list as a slice
- `Sort(less func(a, b T) int)`: Sorts the array list with the less function
- `IsFixedSize()`: Returns true if the array list is fixed in size
- `Len()`: Returns the number of values in the array list
- `IsEmpty()`: Returns true if the array list is empty
- `Clear()`: Clears the array list
- `Reset()`: Clears the array list but conserves the initial capacity
- `Reverse()`: Reverses the values in the array list