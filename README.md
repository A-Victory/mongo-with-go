# Mongo with Go

Here is a brief web app on using Go to interact with MongoDB server running on local machine

Here we  perform simple CRUD operations on the database
Reading, writing to, as well deleting from the database

For ease of usage, use the following command in your shell(terminal):
    To post to the database:
        curl -X POST -H "Content-Type: application/json" -d '{"name":"Input_Name","gender":"Gender","age":00}' http://localhost:8080/user
    To get a single user, run the following command:
        curl http://localhost:8080/user/<enter-user-id-here>
    To delete a user, run the following command:
        curl -X DELETE http://localhost:8080/user/<enter-user-id-here>
    To get a list of all users, run the following command:
        curl  http://localhost:8080/user
    
You cann also see the list of all users by going to *http://localhost:8080/user* on ur browsers

This example was updated from Todd Mcleod's Golang-Web-Dev course to use more recent mongo drivers.
Go check it out!!
