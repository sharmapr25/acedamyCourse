# acedamyCourse
It is the repo for instructor and student to see and view all course, subjects those are available

# Requirement
* Go
* mysql

# Apis
* Create tag
`
curl --location --request POST 'localhost:8081/tags' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "tag-2",
    "created_at": "1639225436",
    "is_active": true
}'
`

# Run Test
Specific test
```
go test -timeout 30s -run ^TestShouldCreateTag$ sample/acedamyCourse/models
```
