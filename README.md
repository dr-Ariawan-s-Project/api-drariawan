<div align="center">
  <h3 align="center">dr. Ariawan Project</h3>

  <p align="center">
    [backend]
  </p>
</div>

## About
dr. Ariawan Project 

## 👨🏽‍💻 Run the Project
1. Setup `.env`
  note: aes-key must have a length of 16, 24, or 32 byte for AES-128, AES-192, or AES-256
    ```
    export JWT_KEY='your-jwt-key'
    export DBUSER='your-db-user'
    export DBPASS='your-db-password'
    export DBHOST='your-db-host'
    export DBPORT='your-db-port'
    export DBNAME='your-db-name'
    export AESGCMSECRET='your-16-byte-aes'
    ```

2. Run app
    ```
    go run .
    ```