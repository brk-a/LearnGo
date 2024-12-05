# ERD

### in plain text...

```plaintext
    [User ] 1 ----< [Token]
   |
   | 1
   |
   | 
    [Order] 1 ----< [Invoice]
    |
    | 1
    |
    | 
    [Table] 1 ----< [Order]
    |
    | 1
    |
    | 
    [OrderItem] >---- [Food] >---- [Menu]
    |
    | 1
    |
    | 
    [Order] >---- [Note]
```

---

### better...

```mermaid
    erDiagram
        USER {
            string ID PK
            string First_name
            string Last_name
            string Password
            string Email
            string Avatar
            string Phone
            string Role
            datetime Created_at
            datetime Updated_at
        }

        TOKEN {
            string ID PK
            string Token
            string Refresh_token
            string Reset_token
            datetime Created_at
            datetime Updated_at
            datetime Expiry
            datetime TTL
            string User_id FK
            string Scope
        }

        MENU {
            string ID PK
            string Name
            string Category
            datetime Start_date
            datetime End_date
            datetime Created_at
            datetime Updated_at
        }

        FOOD {
            string ID PK
            string Name
            float Price
            string Food_image
            string Description
            datetime Created_at
            datetime Updated_at
            string Menu_id FK
        }

        TABLE {
            string ID PK
            int Number_of_guests
            int Table_number
            string Status
            datetime Created_at
            datetime Updated_at
        }

        ORDER {
            string ID PK
            datetime Order_date
            string Status
            datetime Created_at
            datetime Updated_at
            string Table_id FK
            string Customer_id FK
        }

        INVOICE {
            string ID PK
            string InvoiceID
            string OrderID FK
            string Payment_method
            string Payment_status
            datetime Payment_due_date
            float Total_amount
            datetime Created_at
            datetime Updated_at
        }

        ORDERITEM {
            string ID PK
            string Quantity
            float Unit_price
            float Total_price
            datetime Created_at
            datetime Updated_at
            string Food_id FK
            string Order_id FK
        }

        NOTE {
            string ID PK
            string Text
            string Title
            datetime Created_at
            datetime Updated_at
            string Order_id FK
        }

        %% Relationships
        USER ||--o{ TOKEN : has
        USER ||--o{ ORDER : places
        ORDER ||--o{ INVOICE : generates
        TABLE ||--o{ ORDER : serves
        ORDER ||--o{ ORDERITEM : contains
        FOOD ||--o{ ORDERITEM : includes
        MENU ||--o{ FOOD : contains
        ORDER ||--o{ NOTE : has
```

---