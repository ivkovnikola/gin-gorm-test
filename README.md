Sample application written in golang with gin and gorm.

To run the application, run go run main.go from home folder.

.evn file contains all the environment variables - it can connect either to a test sql lite database, or to a real postgres instance if provided.

To run tests run:
from the /tests directory you can run go test -v ./...
or from directory where test file is run: go test -v --run name of the test file.
If running in visual studio code for example, with go plugin, it will allow you to run any test from the editor itself.


Some of the endpoints are private and some are public. You can test endpoints using postman.
By default DB is seeded with values(in seeder.go file)
You can use test user with credentials 
            email:     testuser@gmail.com
			Password: password,

Query localhost:8085/users or localhost:8085/posts, and it will return values in database. 
To run endpoints that are behind auth server, you need to first get an auth token by calling /login endpoint with body:
            {
                email: "testuser@gmail.com",
                password: "password"
            }

It will return a token, that you can then place in Authorization: Bearer Token, and test any other endpoint.
All the endpoints can be seen in routes.go