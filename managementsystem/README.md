## Overviews
## Technologies Used

*   **Go:** Programming language.
*   **MongoDB:** Database.
*   **Echo:** Web framework.
*   **JWT:** Authentication and authorization.
*   **Docker:** Containerization.
* **Memcache:** Cache
  
## API Endpoints

### Authorize Module

*default authorize for admin*
```
username: root
password: P@ssw0rd
```

*   **`POST /authorize/login`**
    *   **Description:** Logs in a user and returns access and refresh tokens.
    *   **Request Body:**

        ```json
        {
          "username": "<username>",
          "password": "<password>"
        }
        ```

    *   **Response Body:**

        ```json
        {
          "access_token": "<access-token>",
          "refresh_token": "<refresh-token>"
        }
        ```

*   **`POST /authorize/refresh`**
    *   **Description:** Refreshes an access token using a refresh token.
    *   **Request Body:**

        ```json
        {
          "token": "<refresh-token>"
        }
        ```

    *   **Response Body:**

        ```json
        {
          "access_token": "<access-token>"
        }
        ```

*   **`GET /authorize/logout`**
    *   **Description:** Logs out a user by invalidating their access token.
    *   **Headers**
    ```
     Authorization: Bearer <access-token>
    ```
    *   **Response Body:**

        ```json
        {
          "status": "OK",
          "message": "logout successfully"
        }
        ```

### Management

*   **`GET /management`**
    *   **Description:** Retrieves a list of all managements.
     *   **Headers**
    ```
     Authorization: Bearer <access-token>
    ```
    *   **Response Body:**

        ```json
        [
          {
            "id": "<management-id>",
            "name": "<name>",
            "nickname": "<nickname>",
            "gender": "<gender>",
            "age": <age>,
            "role": "<role>",
            "createDate": "<create-date>",
            "updateDate": "<update-date>"
          }
        ]
        ```

*   **`POST /management`**
    *   **Description:** Creates a new management.
     *   **Headers**
    ```
     Authorization: Bearer <access-token>
    ```
    *   **Request Body:**

        ```json
        {
          "name": "<name>",
          "nickname": "<nickname>",
          "gender": "<gender>",
          "age": <age>,
          "role": "<role>"
        }
        ```

    *   **Response Body:**

        ```json
        {
          "message": "success"
        }
        ```

*   **`PUT /management/:id`**
    *   **Description:** Updates an existing management.
    *   **Path Parameter:** `id` (the ID of the management to update).
     *   **Headers**
    ```
     Authorization: Bearer <access-token>
    ```
    *   **Request Body:**

        ```json
        {
          "name": "<new-name>",
          "nickname": "<new-nickname>",
          "gender": "<new-gender>",
          "age": <new-age>,
          "role": "<new-role>"
        }
        ```

    *   **Response Body:**

        ```json
        {
          "message": "success",
          "data": {
            "id": "<management-id>",
            "name": "<new-name>",
            "nickname": "<new-nickname>",
            "gender": "<new-gender>",
            "age": <new-age>,
            "role": "<new-role>",
            "updateDate": "<update-date>"
          }
        }
        ```

*   **`DELETE /management/:id`**
    *   **Description:** Deletes a management.
    * **Path Parameter:** `id` (the ID of the management to delete).
     *   **Headers**
    ```
     Authorization: Bearer <access-token>
    ```

    *   **Response Body:**

        ```json
        {
          "message": "success"
        }
        ```


