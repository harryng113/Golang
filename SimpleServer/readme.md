Run server: go run main.go

Create account: curl localhost:8080/signup --include --header "Content-Type: application/json" -d @newuser.json

Sign in: curl -X POST localhost:8080/signin -H "Content-Type: application/json" -d '{"username": "newbie", "password": "what"}'

Edit profile: curl -X POST localhost:8080/editProfile -H "Content-Type: application/json" -d '{"username": "newbie", "profile_pic": "uploads/trees.jpg"}'