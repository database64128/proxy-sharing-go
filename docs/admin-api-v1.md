# Admin API v1

The Admin API is a RESTful API that allows you to manage everything on the platform.

## 1. Authentication

To authenticate with the Admin API, configure a separate set of access tokens in the Admin API configuration.

```json
{
    "api": {
        "admin": {
            "accessTokens": [
                "jTGA1dfSmqtyNsLIrM9zkPIdjvw76I1z7LqJAAg13TU="
            ]
        }
    }
}
```

In requests to the Admin API endpoints, include the access token in the `Authorization` header.

```http
GET /api/admin/v1/accounts HTTP/1.1
Host: localhost:18080
Authorization: Bearer jTGA1dfSmqtyNsLIrM9zkPIdjvw76I1z7LqJAAg13TU=
```

The server will respond with `401 Unauthorized` if the access token is missing or invalid.

```json
{
    "error": "invalid access token"
}
```

## 2. Registration Tokens

Create registration tokens to allow users to register on the platform.

### 2.1. List Registration Tokens

#### Request

```
GET /api/admin/v1/registration-tokens
```

#### Response: `200 OK`

```json
{
    "tokens": [
        {
            "id": 1,
            "create_time": "2006-01-02T15:04:05.999999999Z07:00",
            "update_time": "2006-01-02T15:04:05.999999999Z07:00",
            "name": "Test Token",
            "token": "jTGA1dfSmqtyNsLIrM9zkPIdjvw76I1z7LqJAAg13TU="
        }
    ]
}
```

### 2.2. Get a Registration Token

#### Request

```
GET /api/admin/v1/registration-tokens/:id
```

#### Response: `200 OK`

```json
{
    "id": 1,
    "create_time": "2006-01-02T15:04:05.999999999Z07:00",
    "update_time": "2006-01-02T15:04:05.999999999Z07:00",
    "name": "Test Token",
    "token": "jTGA1dfSmqtyNsLIrM9zkPIdjvw76I1z7LqJAAg13TU="
}
```

#### Response: `400 Bad Request`

```json
{
    "error": "invalid token ID"
}
```

#### Response: `404 Not Found`

```json
{
    "error": "token not found"
}
```

### 2.3. Create a Registration Token

#### Request

```
POST /api/admin/v1/registration-tokens
```

```json
{
    "name": "Test Token"
}
```

#### Response: `201 Created`

```json
{
    "id": 1,
    "create_time": "2006-01-02T15:04:05.999999999Z07:00",
    "update_time": "2006-01-02T15:04:05.999999999Z07:00",
    "name": "Test Token",
    "token": "jTGA1dfSmqtyNsLIrM9zkPIdjvw76I1z7LqJAAg13TU="
}
```

#### Response: `400 Bad Request`

```json
{
    "error": "invalid token name"
}
```

#### Response: `409 Conflict`

```json
{
    "error": "a token with the same name already exists"
}
```

### 2.4. Update a Registration Token

#### Request

```
PATCH /api/admin/v1/registration-tokens/:id
```

```json
{
    "name": "Test Token 2"
}
```

#### Response: `200 OK`

```json
{
    "id": 1,
    "create_time": "2006-01-02T15:04:05.999999999Z07:00",
    "update_time": "2006-01-02T15:04:05.999999999Z07:00",
    "name": "Test Token 2",
    "token": "jTGA1dfSmqtyNsLIrM9zkPIdjvw76I1z7LqJAAg13TU="
}
```

#### Response: `400 Bad Request`

```json
{
    "error": "invalid token ID"
}
```

#### Response: `404 Not Found`

```json
{
    "error": "token not found"
}
```

#### Response: `409 Conflict`

```json
{
    "error": "a token with the same name already exists"
}
```

### 2.5. Delete a Registration Token

#### Request

```
DELETE /api/admin/v1/registration-tokens/:id{?purgeRegistrations=true}
```

- `purgeRegistrations`: Optional. If `true`, delete all accounts registered with the token.

#### Response: `204 No Content`

The registration token is successfully deleted.

#### Response: `400 Bad Request`

```json
{
    "error": "invalid token ID"
}
```

#### Response: `404 Not Found`

```json
{
    "error": "token not found"
}
```

## 3. Accounts

Manage accounts on the platform.
