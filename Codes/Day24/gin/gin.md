--

## 1. Basic Server

## `gin.Default()`

**Theory:** Preconfigured router (logger + recovery)

```go
r := gin.Default()

r.GET("/", func(c *gin.Context) {
    c.JSON(200, gin.H{"message": "Hello"})
})

r.Run(":8080")
```

**Request:**

```
GET / 
```

**Output:**

```json
{
  "message": "Hello"
}
```

---

# 2. Routing

## `GET / POST / PUT / DELETE`

**Theory:** Map HTTP method → handler

```go
r.GET("/ping", func(c *gin.Context) {
    c.JSON(200, gin.H{"msg": "pong"})
})
```

**Request:**

```
GET /ping
```

**Output:**

```json
{"msg": "pong"}
```

---

## `Group()`

**Theory:** Logical grouping of routes

```go
api := r.Group("/api")

api.GET("/users", func(c *gin.Context) {
    c.JSON(200, gin.H{"users": []string{"A", "B"}})
})
```

**Request:**

```
GET /api/users
```

**Output:**

```json
{"users": ["A", "B"]}
```

---

# 3. Path & Query Parameters

## `c.Param()`

```go
r.GET("/user/:id", func(c *gin.Context) {
    id := c.Param("id")
    c.JSON(200, gin.H{"id": id})
})
```

**Request:**

```
GET /user/101
```

**Output:**

```json
{"id": "101"}
```

---

## `c.Query()`

```go
r.GET("/search", func(c *gin.Context) {
    name := c.Query("name")
    c.JSON(200, gin.H{"name": name})
})
```

**Request:**

```
GET /search?name=John
```

**Output:**

```json
{"name": "John"}
```

---

## `c.DefaultQuery()`

```go
name := c.DefaultQuery("name", "guest")
```

**Request:**

```
GET /search
```

**Output:**

```json
{"name": "guest"}
```

---

# 4. Binding (JSON → Struct)

## `c.ShouldBindJSON()`

```go
type User struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

r.POST("/user", func(c *gin.Context) {
    var u User

    if err := c.ShouldBindJSON(&u); err != nil {
        c.JSON(400, gin.H{"error": "invalid input"})
        return
    }

    c.JSON(200, u)
})
```

**Request:**

```json
POST /user
{
  "name": "John",
  "age": 25
}
```

**Output:**

```json
{
  "name": "John",
  "age": 25
}
```

---

# 5. Response Functions

## `c.JSON()`

```go
c.JSON(200, gin.H{"status": "ok"})
```

**Output:**

```json
{"status": "ok"}
```

---

## `c.String()`

```go
c.String(200, "Hello World")
```

**Output:**

```
Hello World
```

---

## `c.XML()`

```go
type User struct {
    Name string `xml:"name"`
}

c.XML(200, User{Name: "John"})
```

**Output:**

```xml
<User>
  <name>John</name>
</User>
```

---

# 6. Middleware

## `r.Use()`

```go
func Logger() gin.HandlerFunc {
    return func(c *gin.Context) {
        println("Before request")
        c.Next()
        println("After request")
    }
}

r.Use(Logger())
```

**Request Flow Output (console):**

```
Before request
After request
```

---

## `c.Abort()`

```go
r.GET("/secure", func(c *gin.Context) {
    c.AbortWithStatusJSON(401, gin.H{"error": "unauthorized"})
})
```

**Output:**

```json
{"error": "unauthorized"}
```

---

# 7. Logging & Recovery

## `gin.Logger()` and `gin.Recovery()`

Already included in `gin.Default()`

**Console Output Example:**

```
[GIN] 200 | GET "/ping"
```

---

# 8. Server Configuration

## `r.Run()`

```go
r.Run(":8080")
```

Server starts at:

```
http://localhost:8080
```

---

## Custom Server (using net/http)

```go
srv := &http.Server{
    Addr:    ":8080",
    Handler: r,
}
srv.ListenAndServe()
```

---

# 9. Testing

```go
w := httptest.NewRecorder()
req, _ := http.NewRequest("GET", "/ping", nil)

r.ServeHTTP(w, req)
```

**Output Check:**

```go
if w.Code == 200 {
    fmt.Println("PASS")
}
```
---