# Test  Case
## User
### Register User
request:
```shell
curl --request POST \
  --url http://localhost:8080/user/register \
  --header 'content-type: application/json' \
  --data '{
  "username":"tyrone",
  "password":"admin",
  "email":"1014990573@qq.com®"
}'
```
response:
```json
{
  "code": 200,
  "msg": "Success",
  "data": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTAxMjE1MDcsInJvbGVzIjpbInVzZXIiXSwidXNlcklEIjoidHlyb25lIn0.dBqd7AVl9VwB5J6XUbihXUUW2o6dMfjaIvT8682oxiY"
}
```

request:
```shell
curl --request POST \
  --url http://localhost:8080/user/register \
  --header 'content-type: application/json' \
  --data '{
  "username":"tyrone",
  "password":"admin",
  "email":"1014990573@qq.com®"
}'
```
response:
```json
{
  "username":"tyrone",
  "password":"admin",
  "email":"1014990573@qq.com®"
}
```

### Login
request:
```shell
curl --request POST \
  --url http://localhost:8080/user/login \
  --header 'content-type: application/json' \
  --data '{
  "username": "tyrone",
  "password": "admin"
}
'
```

response:
```json
{
  "code": 200,
  "msg": "Success",
  "data": {
    "email": "1014990573@qq.com®",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTAxMjI5NTUsInJvbGVzIjpbInVzZXIiXSwidXNlcklEIjozfQ.x3ZJG4pPqS7-W17QJjCwGjgEhOOD5LX9BVfIvkoFpfg",
    "username": "tyrone"
  }
}
```

request:
```shell
curl --request POST \
  --url http://localhost:8080/user/login \
  --header 'content-type: application/json' \
  --data '{
  "username": "tyrone",
  "password": "adm1in"
}
'
```

response:
```json
{
  "code": 500,
  "msg": "incorrect password",
  "data": null
}
```

## Post
### Create Post
request:
```shell
curl --request POST \
  --url http://localhost:8080/post/add \
  --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTAxMjM4MTMsInJvbGVzIjpbInVzZXIiXSwidXNlcklEIjozfQ.znxZgu5AI_QIG8f79Mw4jFyb_0ybIracq4QAgdHyzDk' \
  --header 'content-type: application/json' \
  --data '{
  "title":"Invitation",
  "content":"Hi Tyrone! Let'\''s start to learn Web3!"
}'
```

response:
```json
{
  "code": 200,
  "msg": "Success",
  "data": {
    "ID": 2,
    "CreatedAt": "2025-06-17T01:40:12.189+08:00",
    "UpdatedAt": "2025-06-17T01:40:12.189+08:00",
    "DeletedAt": null,
    "title": "Invitation",
    "content": "Hi Tyrone! Let's start to learn Web3!",
    "Comment": null,
    "UserId": 3
  }
}
```

### List ALL Posts
request:
```shell
curl --request GET \
  --url http://localhost:8080/post/all
```
response:
```json
{
  "code": 200,
  "msg": "Success",
  "data": [
    {
      "ID": 2,
      "CreatedAt": "2025-06-17T01:43:00.489+08:00",
      "UpdatedAt": "2025-06-17T01:43:00.489+08:00",
      "DeletedAt": null,
      "title": "Invition",
      "content": "Hi Tyrone! Let's start to learn Web3!",
      "Comment": [],
      "UserId": 3
    }
  ]
}
```
### Get Post By ID
request:
```shell
curl --request GET \
  --url http://localhost:8080/post/2
```
response:
```json
{
  "code": 200,
  "msg": "Success",
  "data": {
    "ID": 2,
    "CreatedAt": "2025-06-17T01:43:00.489+08:00",
    "UpdatedAt": "2025-06-17T01:43:00.489+08:00",
    "DeletedAt": null,
    "title": "Invition",
    "content": "Hi Tyrone! Let's start to learn Web3!",
    "Comment": [],
    "UserId": 3
  }
}
```

### Update Post
request:
```shell
curl --request PUT \
  --url http://localhost:8080/post \
  --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTAxMjM4MTMsInJvbGVzIjpbInVzZXIiXSwidXNlcklEIjozfQ.znxZgu5AI_QIG8f79Mw4jFyb_0ybIracq4QAgdHyzDk' \
  --header 'content-type: application/json' \
  --data '{
  "id": 2,
  "title": "try to update title",
  "content": "try to update content"
}
'
```
response:
```json
{
  "code": 200,
  "msg": "Success",
  "data": null
}
```

### Delete Post
request:
```shell
curl --request DELETE \
  --url http://localhost:8080/post/2 \
  --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTAxMjM4MTMsInJvbGVzIjpbInVzZXIiXSwidXNlcklEIjozfQ.znxZgu5AI_QIG8f79Mw4jFyb_0ybIracq4QAgdHyzDk'
```
response:
```json
{
  "code": 200,
  "msg": "Success",
  "data": null
}
```

## Comment
### Create Comment
request:
```shell
curl --request POST \
  --url http://localhost:8080/comment/add \
  --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTAxMjM4MTMsInJvbGVzIjpbInVzZXIiXSwidXNlcklEIjozfQ.znxZgu5AI_QIG8f79Mw4jFyb_0ybIracq4QAgdHyzDk' \
  --header 'content-type: application/json' \
  --data '{
  "content":"this is a comment for testing",
  "postId" : 2
}'
```
response:
```json
{
  "code": 200,
  "msg": "Success",
  "data": {
    "ID": 2,
    "CreatedAt": "2025-06-17T02:55:30.74+08:00",
    "UpdatedAt": "2025-06-17T02:55:30.74+08:00",
    "DeletedAt": null,
    "content": "this is a comment for testing",
    "userId": 3,
    "postId": 2
  }
}
```
### List Comments By Post ID
request:
```shell
curl --request GET \
  --url http://localhost:8080/comment/get/2
```
response:
```json
{
  "code": 200,
  "msg": "Success",
  "data": [
    {
      "ID": 1,
      "CreatedAt": "2025-06-17T02:55:01.412+08:00",
      "UpdatedAt": "2025-06-17T02:55:01.412+08:00",
      "DeletedAt": null,
      "content": "this is a comment for testing",
      "userId": 3,
      "postId": 2
    },
    {
      "ID": 2,
      "CreatedAt": "2025-06-17T02:55:30.74+08:00",
      "UpdatedAt": "2025-06-17T02:55:30.74+08:00",
      "DeletedAt": null,
      "content": "this is a comment for testing",
      "userId": 3,
      "postId": 2
    }
  ]
}
```
