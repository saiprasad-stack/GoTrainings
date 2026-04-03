```markdown
# Backend Project 

## Components 

| Piece | Role (what it’s supposed to do) |
|-------|--------------------------------|
| **Client** | You (or a browser, or `curl`) – sends a request (e.g., “create an item”). |
| **Router** | Looks at the URL and method (`POST /items`) and decides which function should handle it. |
| **Handler** | The function that does the actual work for that request: reads the data, calls the database, sends back a response. |
| **Model** | A blueprint of what an “item” looks like (has a name, price, ID, etc.). |
| **Database (DB)** | Where data is stored permanently (PostgreSQL in our case). |
| **ORM (GORM)** | A helper that translates Go code into SQL so you don’t have to write raw database queries. |

---

## How the flow works – step by step (with arrow diagram)

Let’s take **“Create a new item”** as an example.

```
1. You send a POST request with JSON:
   → {"name":"Laptop","price":999}

2. The Router receives the request.
   It says: "POST /api/items → that belongs to CreateItem handler."

3. The CreateItem handler runs:
   - It takes the JSON from the request.
   - It uses the Item model as a template to create a Go object.

4. The handler asks the ORM (GORM) to save that object into the database.

5. The ORM translates the request into SQL:
   → INSERT INTO items (name, price) VALUES ('Laptop', 999);

6. PostgreSQL runs the SQL and stores the data.

7. PostgreSQL tells the ORM: "Saved successfully, here’s the new ID and timestamps."

8. The ORM gives the updated object back to the handler.

9. The handler sends a response (JSON) back to you:
   → {"id":1, "name":"Laptop", "price":999, ...}

10. You see the response.
```

### Visual flow

```
YOU (Client)
   │
   │ POST /api/items  {"name":"Laptop","price":999}
   ▼
ROUTER (routes.go)
   │
   │ "This URL goes to CreateItem handler"
   ▼
HANDLER (item_handler.go)
   │
   │ 1. Parse JSON
   │ 2. Call ORM to save
   ▼
ORM (GORM)
   │
   │ Translate to SQL: INSERT ...
   ▼
DATABASE (PostgreSQL)
   │
   │ Store data, return new ID
   ▼
ORM
   │
   │ Give back saved object
   ▼
HANDLER
   │
   │ Wrap object into JSON response
   ▼
ROUTER
   │
   │ Send HTTP response
   ▼
YOU receive: {"id":1,"name":"Laptop","price":999}
```

---

## What about GET, UPDATE, DELETE?

The idea is the same – only the SQL changes:

| Action | What the handler does |
|--------|----------------------|
| **GET** | Asks ORM: “find item with ID = 1” → ORM runs `SELECT` → returns data to you. |
| **UPDATE** | Asks ORM: “find item 1, change its name and price” → ORM runs `UPDATE` → returns updated data. |
| **DELETE** | Asks ORM: “remove item 1” → ORM runs `DELETE` → returns “success” message. |

---

## Why this order makes sense

- **Router** is the receptionist – directs traffic.  
- **Handler** is the worker – does the task.  
- **ORM** is the translator – turns Go into SQL.  
- **Database** is the storage room – keeps data safe.  

A beginner should think:  
1. Define your **model** (what data looks like).  
2. Set up the **database connection**.  
3. Write **handlers** for each action.  
4. Connect handlers to URLs in the **router**.  

Then the flow from client → router → handler → ORM → DB → back becomes natural.
```
