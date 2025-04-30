# go-sqlz

📦 Lightweight SQL-to-struct mapper for Go.  
A simple utility for clean and fast `SELECT` queries using `database/sql`.

---

## ✨ Features

- ✅ Struct mapping via `db` tags  
- ✅ Supports both `[]T` and `[]*T`  
- ✅ Works with standard `database/sql`  
- ✅ No external dependencies (except optional `sqlmock` for testing)  
- ✅ Reflect-based, focused only on `SELECT` use cases  

---

## 📦 Installation

```bash
go get github.com/aquaheyday/go-sqlz
```

---

## 🚀 Usage

### 1. Define your struct

```go
type User struct {
    ID   int    \`db:"id"\`
    Name string \`db:"name"\`
}
```

### 2. Execute query & map results

```go
rows, err := db.Query("SELECT id, name FROM users")
if err != nil {
    log.Fatal(err)
}

var users []User
err = gomapper.Select(rows, &users)
if err != nil {
    log.Fatal(err)
}
```

### 3. For single row

```go
row := db.QueryRow("SELECT id, name FROM users WHERE id = ?", 1)

var user User
err := gomapper.Get(row, &user)
if err != nil {
    log.Fatal(err)
}
```

---

## 🧪 Testing with sqlmock

Install `sqlmock`:

```bash
go get github.com/DATA-DOG/go-sqlmock
```

Example test:

```go
func TestSelect(t *testing.T) {
    db, mock, _ := sqlmock.New()
    defer db.Close()

    rows := sqlmock.NewRows([]string{"id", "name"}).
        AddRow(1, "Alice").
        AddRow(2, "Bob")
    mock.ExpectQuery("SELECT .* FROM users").WillReturnRows(rows)

    result, _ := db.Query("SELECT id, name FROM users")
    var users []User
    _ = gomapper.Select(result, &users)

    if len(users) != 2 || users[0].Name != "Alice" {
        t.Errorf("Unexpected result: %+v", users)
    }
}
```

---

## 📘 Documentation

For detailed documentation and examples, visit the [GoDoc](https://pkg.go.dev/github.com/aquaheyday/go-sqlz) page.

---

## 🛡 License

This project is licensed under the [MIT License](LICENSE).

---

## 🙋‍♂️ Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.
