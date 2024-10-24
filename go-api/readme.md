### Inicia la base de datos y corre la aplicación
```bash
go run main.go
```

### Crear un usuario:
```bash
curl -X POST http://localhost:8888/users -d '{"name": "John", "email": "john@example.com"}' -H "Content-Type: application/json"
```

### Obtener todos los usuarios
```bash
curl -X GET http://localhost:8888/users
```

### Actualizar un usuario
```bash
curl -X PUT http://localhost:8888/users/1 -d '{"name": "John Updated", "email": "john.updated@example.com"}' -H "Content-Type: application/json"
```

### Eliminar un usuario:
```bash
curl -X DELETE http://localhost:8888/users/1
```