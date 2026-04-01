```markdown
# Go Reflection & Generics

## Reflection
Reflection is a way for your program to **inspect itself while it's running**.  
It lets you ask: "What type is this variable?" and "What is its value?"

- `reflect.Type` – tells you the **type** of a variable (e.g., `int`, `string`, `User`)
- `reflect.Value` – lets you **read or modify** the actual value stored in the variable

### Example
```go
package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := 10

	t := reflect.TypeOf(x)   // t holds the type information: int
	v := reflect.ValueOf(x)  // v holds the value: 10

	fmt.Println(t) // prints "int"
	fmt.Println(v) // prints "10"
}
```

**What's happening?**  
`reflect.TypeOf(x)` returns the type of `x` as a `reflect.Type` object.  
`reflect.ValueOf(x)` returns the value wrapped in a `reflect.Value` object, which you can use to read or even change the value (if it's addressable).

---

### Struct Tags – Adding Metadata
Struct tags are **strings attached to struct fields** that carry extra information.  
They are often used to tell libraries how to handle the field.

```go
type User struct {
	Name string `json:"name"`   // tag tells encoding/json to use "name" in JSON
}
```

Reflection is used to **read these tags** at runtime.

#### Common Uses
- **JSON encoding/decoding** – specify field names, omit empty, etc.
- **Validation** – e.g., `validate:"required,min=3"`
- **ORM tools** – map struct fields to database columns

---

### When to Avoid Reflection
Reflection is powerful but comes with costs:

- **Slower** – because it works at runtime, it’s not as fast as normal code.
- **Runtime errors** – mistakes are caught only when the code runs, not at compile time.
- **Hard to read/debug** – reflection code can be tricky to follow.

#### Prefer These Instead
- **Interfaces** – define behaviour without caring about concrete type.
- **Generics** – write type‑safe code that works with many types (Go 1.18+).

---

## Generics (Go 1.18+)
Generics let you **write one function (or struct) that works with multiple types**, while still being type‑safe.

Without generics, you might need separate functions:
```go
func AddInt(a, b int) int
func AddFloat(a, b float64) float64
```

With generics, you write:
```go
func Add[T any](a, b T) T {
	return a + b  // but note: 'any' doesn't guarantee '+' works
}
```

### Generic Struct
You can also create generic data structures:
```go
type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Usage:
intStack := Stack[int]{}
intStack.Push(10)

stringStack := Stack[string]{}
stringStack.Push("hello")
```
---

## Generics vs Interfaces – Which to Use?

| Feature         | Generics                                | Interfaces                              |
|-----------------|-----------------------------------------|-----------------------------------------|
| **Type safety** | Compile‑time – errors caught early      | Runtime – type assertions can fail      |
| **Speed**       | Faster (no boxing)                      | Slightly slower (indirection)           |
| **Use case**    | Algorithms, data structures, same logic | Polymorphism, different behaviours      |

### Simple Rule
- **Generics** – when you have the **same logic** but for **many types** (e.g., a stack, a sort function).
- **Interfaces** – when you need **different behaviours** for different types (e.g., `io.Reader`, `Stringer`).

```