# Devops
A REST API task manager

## API

A sample Task object
```
  {
    "id": 1,
    "name": "Write presentation",
    "completed": true,
    "due": "2017-03-13T14:49:28.740694Z"
  }
```

### commands

`GET 127.0.0.1:8080/tasks` : Lists all Tasks

`POST 127.0.0.1:8080/tasks` : Creates a new Task

`GET 127.0.0.1:8080/tasks/{id}` : Gets a task by it's ID

`PUT 127.0.0.1:8080/tasks/{id}` : Updates a task by it's ID

`DELETE 127.0.0.1:8080/tasks/{id}` : Deletes a task by it's ID
