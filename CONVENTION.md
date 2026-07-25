# Backend Coding Convention (Golang + Gin + GORM)

> Phiên bản: 1.0
>
> Mục tiêu của tài liệu này là thống nhất cách viết code trong toàn bộ dự án, giúp code dễ đọc, dễ bảo trì và dễ mở rộng.

---

# 1. Kiến trúc dự án

```
internal/
│
├── constants/
│
├── database/
│
├── dto/
│   ├── requests/
│   └── responses/
│
├── handlers/
│
├── middlewares/
│
├── models/
│
├── repositories/
│
├── services/
│
├── utils/
│
└── routes/
```

Luồng xử lý luôn theo thứ tự:

```
HTTP Request
        │
        ▼
Handler
        │
        ▼
Service
        │
        ▼
Repository
        │
        ▼
Database
```

Không được phép:

- Handler truy cập Database
- Handler xử lý Business Logic
- Repository trả về HTTP Status
- Service trả về JSON

---

# 2. Quy tắc đặt tên

## Struct

```go
type AuthService struct{}
```

Không dùng

```go
type authService struct{}
```

---

## Method

Tên phải là động từ.

Đúng

```
Register()

Login()

CreateUser()

UpdateUser()

FindByID()

Delete()
```

Sai

```
Data()

Info()

Process()
```

---

## File

Tên file viết snake_case.

```
auth_service.go

auth_repository.go

user_handler.go
```

Không dùng

```
AuthService.go

UserHandler.go
```

---

# 3. Handler

Handler chỉ làm đúng 4 việc.

## 1. Bind Request

```go
c.ShouldBindJSON(...)
```

---

## 2. Validate Request

Không viết validate trong Handler.

Chỉ gọi

```go
utils.ParseValidationErrors(...)
```

---

## 3. Gọi Service

```go
service.Register(...)
```

---

## 4. Trả Response

Không viết business logic trong Handler.

Không truy cập database.

Không hash password.

Không generate JWT.

---

# 4. Service

Service là nơi xử lý Business Logic.

Ví dụ

```
Đăng ký

Đăng nhập

Đổi mật khẩu

Đặt phòng

Thanh toán
```

Service được phép:

- gọi Repository
- gọi Utils
- xử lý transaction
- xử lý business rule

Không được:

- trả JSON
- dùng gin.Context
- dùng c.JSON()

---

# 5. Repository

Repository chỉ làm việc với Database.

Được phép

```
SELECT

INSERT

UPDATE

DELETE

Transaction
```

Không được

```
Hash Password

Generate JWT

Validate Request

HTTP Response
```

---

# 6. DTO

## Request

```
dto/requests/
```

Chỉ chứa request.

Ví dụ

```
LoginRequest

RegisterRequest

UpdateProfileRequest
```

---

## Response

```
dto/responses/
```

Chỉ chứa response.

Ví dụ

```
AuthResponse

UserResponse
```

---

# 7. Response Format

Toàn bộ API phải thống nhất.

## Success

```json
{
  "success": true,
  "message": "Success",
  "data": {}
}
```

---

## Validation Error

```json
{
  "success": false,
  "message": "Validation failed",
  "errors": [
    {
      "field": "email",
      "message": "Invalid email"
    }
  ]
}
```

---

## Business Error

```json
{
  "success": false,
  "code": "EMAIL_ALREADY_EXISTS",
  "message": "Email already exists"
}
```

---

## Server Error

```json
{
  "success": false,
  "message": "Internal server error"
}
```

---

# 8. HTTP Status

Không hardcode message.

Không viết

```go
"Login failed"
```

Thay vào đó sử dụng helper chung.

Ví dụ

```
200 OK

201 Created

400 Bad Request

401 Unauthorized

403 Forbidden

404 Not Found

409 Conflict

500 Internal Server Error
```

---

# 9. Business Error

Không dùng

```go
errors.New(...)
```

để biểu diễn lỗi nghiệp vụ.

Thay vào đó tạo AppError.

Ví dụ

```
EMAIL_ALREADY_EXISTS

PHONE_ALREADY_EXISTS

USER_NOT_FOUND

INVALID_PASSWORD

INVALID_TOKEN
```

Service chỉ return AppError.

Handler sẽ tự map sang HTTP Status.

---

# 10. Validation

Validation chỉ khai báo bằng struct tag.

Ví dụ

```go
Email string `validate:"required,email"`
```

Không validate trong Handler.

---

Validation message được convert tại

```
utils/validation.go
```

---

# 11. Password

Toàn bộ hash password đi qua

```
utils/password.go
```

Không được gọi bcrypt trực tiếp ở Service.

---

# 12. JWT

Toàn bộ JWT xử lý tại

```
utils/jwt.go
```

Không generate token ở Handler.

Không parse token ở Service.

---

# 13. Constants

Constants chỉ chứa giá trị không thay đổi.

Ví dụ

```
UserStatusActive

UserStatusInactive

AccountRoleAdmin

AccountRoleTenant
```

Không lưu Business Logic.

---

# 14. Utils

Utils chỉ chứa các hàm dùng chung.

Ví dụ

```
Hash Password

Snake Case

JWT

Validation Helper
```

Không chứa:

- Query Database
- Business Logic

---

# 15. Transaction

Transaction phải được mở ở Service.

Ví dụ

```
Begin()

Repository()

Repository()

Commit()
```

Rollback ngay khi có lỗi.

---

# 16. Dependency Rule

Luồng phụ thuộc phải luôn là

```
Handler
    ↓

Service
    ↓

Repository
    ↓

Database
```

Không được đi ngược.

Ví dụ:

Repository không được gọi Service.

Service không được gọi Handler.

---

# 17. Comment

Comment giải thích "tại sao", không giải thích "làm gì".

Đúng

```go
// Update last login time after successful authentication.
```

Sai

```go
// Set variable.
```

---

# 18. Không Hardcode

Không hardcode:

```
Message

Role

Status

JWT Secret

Magic Number
```

Đưa vào:

```
constants/

config/

env
```

---

# 19. Clean Code

Một function chỉ nên làm một việc.

Nếu function dài trên khoảng 50 dòng:

→ cân nhắc tách nhỏ.

---

Không viết

```go
if ...

if ...

if ...

if ...

if ...

if ...

if ...
```

liên tục trong một function.

Tách thành helper.

---

# 20. Nguyên tắc DRY

Không copy cùng một đoạn code nhiều nơi.

Ví dụ:

```
Validation Response

Success Response

Error Response
```

nên gom thành helper.

---

# 21. Nguyên tắc KISS

Ưu tiên code đơn giản.

Không tạo abstraction nếu chỉ dùng đúng một nơi.

Không tạo interface khi chưa cần.

---

# 22. SOLID

Áp dụng vừa đủ.

Không lạm dụng interface.

Ví dụ:

Repository có thể là struct nếu chỉ có một implementation.

Không cần interface ngay từ đầu.

---

# 23. Code Review Checklist

Trước khi merge, kiểm tra:

- Handler có business logic không?
- Có hardcode message không?
- Có dùng response helper không?
- Có dùng AppError không?
- Có rollback transaction không?
- Có kiểm tra nil không?
- Có comment cần thiết không?
- Có duplicate code không?
- Có function quá dài không?
- Có đúng naming convention không?

---

# 24. Những điều không nên làm

❌ errors.New("Email exists")

❌ c.JSON(...) ở Service

❌ bcrypt.GenerateFromPassword() trong Handler

❌ DB.Create() trong Handler

❌ JWT.Generate() trong Handler

❌ Hardcode message nhiều nơi

❌ Hardcode HTTP Response

❌ Repository gọi Service

❌ Handler gọi Database

---

# 25. Quy trình tạo một API mới

Ví dụ thêm API:

```
POST /users
```

Thứ tự thực hiện:

```
1.
Request DTO

↓

2.
Response DTO

↓

3.
Repository

↓

4.
Service

↓

5.
Handler

↓

6.
Route

↓

7.
Test API
```

Không được làm ngược.

---

# 26. Mục tiêu cuối cùng

Toàn bộ project phải đạt được các tiêu chí sau:

- Mỗi tầng chỉ có một trách nhiệm (Single Responsibility).
- Business Logic chỉ tồn tại trong Service.
- Database chỉ được truy cập qua Repository.
- Response trả về thống nhất trên toàn bộ API.
- Error được chuẩn hóa bằng AppError, không dùng `errors.New()` cho lỗi nghiệp vụ.
- Không hardcode message trong Handler hoặc Service.
- Tái sử dụng tối đa các helper và utility để tránh lặp code.
- Dễ mở rộng, dễ test và dễ bảo trì khi dự án phát triển.
