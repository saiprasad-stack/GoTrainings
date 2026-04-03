___
### psql -d gorm_gin_db

INSERT
INSERT INTO items (name, price) VALUES ('Product', 99.99);

-- VIEW ALL
SELECT * FROM items;

-- VIEW ONE
SELECT * FROM items WHERE id = 1;

-- UPDATE
UPDATE items SET price = 149.99 WHERE id = 1;

-- DELETE ONE
DELETE FROM items WHERE id = 1;

-- DELETE ALL
DELETE FROM items;

-- COUNT
SELECT COUNT(*) FROM items;

-- STRUCTURE
\d items

---