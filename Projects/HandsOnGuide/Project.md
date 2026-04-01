# Go Hands on Guide
---

# Step 1: Project Initialization

## Create Project

```bash
mkdir go-learning-project
cd go-learning-project
```

## Initialize Module

```bash
go mod init go-learning-project
```

### Output

```go
module go-learning-project

go 1.22
```

## go.mod

`go.mod` is the **source of truth** for your project.

It contains:

* Module path
* Go version
* Direct dependencies

Example after adding dependencies:

```go
module go-learning-project

go 1.22

require (
	github.com/fatih/color v1.16.0
)
```

### Why it exists

Before modules, Go used GOPATH which caused:

* Version conflicts
* No reproducibility

`go.mod` solves this by locking dependencies.

---

# Step 2: go.sum (VERY IMPORTANT)

Example:

```
github.com/fatih/color v1.16.0 h1:abc123...
```

## What it stores

* Cryptographic hashes of dependencies

## Why it exists

Ensures:

* Security (no tampered packages)
* Reproducible builds

### Analogy

* go.mod = shopping list
* go.sum = bill receipt + verification

---

# Step 3: Project Structure

```
go-learning-project/
  go.mod
  go.sum
  main.go
  pkg/
    mathutils/
      math.go
    stringutils/
      string.go
  internal/
    logger/
      logger.go
```

## Explanation

### pkg/

Reusable code

### internal/

Private code

* Cannot be imported outside module

### main.go

Entry point

---

# Step 4: Install Dependencies

## Install external library

```bash
go get github.com/fatih/color
```

## What happens

* Downloads package
* Updates go.mod
* Updates go.sum

### Updated go.mod

```go
require github.com/fatih/color v1.16.0
```

---

## go install

```bash
go install golang.org/x/tools/cmd/stringer@latest
```

### Purpose

* Installs CLI tool globally
* Does NOT affect go.mod

---

## go mod tidy

```bash
go mod tidy
```

### What it does

* Removes unused dependencies
* Adds missing dependencies

---

# Step 5: Full Code Implementation

## pkg/mathutils/math.go

```go
package mathutils

func Add(a, b int) int {
	return a + b
}

func Multiply(a, b int) int {
	return a * b
}

func Subtract(a, b int) int {
	return a - b
}
```

---

## pkg/stringutils/string.go

```go
package stringutils

import "strings"

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func CountVowels(s string) int {
	count := 0
	for _, ch := range strings.ToLower(s) {
		if strings.ContainsRune("aeiou", ch) {
			count++
		}
	}
	return count
}
```

---

## internal/logger/logger.go

```go
package logger

import (
	"fmt"
	"time"
)

func Info(msg string) {
	fmt.Println(time.Now().Format("15:04:05"), "[INFO]", msg)
}

func Error(msg string) {
	fmt.Println(time.Now().Format("15:04:05"), "[ERROR]", msg)
}
```

---

## main.go

```go
package main

import (
	"fmt"

	"github.com/fatih/color"

	"go-learning-project/internal/logger"
	"go-learning-project/pkg/mathutils"
	str "go-learning-project/pkg/stringutils"
)

func main() {

	logger.Info("Application Started")

	fmt.Println("Add:", mathutils.Add(2, 3))
	fmt.Println("Multiply:", mathutils.Multiply(4, 5))
	fmt.Println("Subtract:", mathutils.Subtract(10, 4))

	fmt.Println("Upper:", str.ToUpper("golang"))
	fmt.Println("Reverse:", str.Reverse("golang"))
	fmt.Println("Vowels:", str.CountVowels("golang"))

	color.Red("Let's color it. Red Output")

	logger.Info("Application Finished")
}
```

---

# Step 6: Run Project

## Run

```bash
go run main.go
```

## Output

```
[INFO] Application Started
Add: 5
Multiply: 20
Subtract: 6
Upper: GOLANG
Reverse: gnalog
Vowels: 2
Let's color it. Red Output
[INFO] Application Finished
```

---

## Build

```bash
go build
```

Run binary:

```bash
./go-learning-project
```

---

# Step 7: Key Observations

* go.mod grows when adding dependencies
* go.sum ensures integrity
* internal package is restricted
* pkg is reusable

---

# Step 8: Advanced Understanding

## Why modules matter

Without modules:

* No version control
* Conflicts

With modules:

* Versioned dependencies
* Reproducible builds

---


