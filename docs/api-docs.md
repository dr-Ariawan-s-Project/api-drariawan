## API DOCS 

<details>
  <summary><kbd>POST /login</kbd></summary>
    > Login feature
    <br>

-  <kbd>Request Body</kbd>
    ```json
    {
    "email": "jhondoe@email.com",
    "password": "supersecret"
    }
    ```
-  <kbd>Response Body</kbd>
    ```json
    {
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2OTU3NTg2MDMsInJvbGUiOiJhZG1pbiIsInVzZXJJZCI6MX0.kzaY2AeYNcT0969dwmPFjEogRY2XLaKN4wDxETOKIJ4"
    },
    "messages": [
        "[success] login"
    ],
    "meta": {
        "code": "200-002-OK",
        "status": "success"
    }
    }   
    ```
</details>
<!-- ==== End Of Section -->

<details>
  <summary><kbd>GET /v1/questioner</kbd></summary>
    > Get all question

-  <kbd>Response Body</kbd>
    ```json
    {
    "data": [
        {
            "id": 1,
            "type": "text",
            "question": "https://linkto.com/video.mp4",
            "description": "berapa tinggi badan anda?",
            "choices": null,
            "goto": 2
        },
        {
            "id": 2,
            "type": "choices",
            "question": "https://linkto.com/video.mp4",
            "description": "seberapa sering anda menggunakan celana ketat?",
            "choices": [
                {
                    "id": 1,
                    "question_id": 2,
                    "option": "1 (tidak pernah sama sekali)",
                    "slugs": "tidak;no;",
                    "score": 0,
                    "goto": 3
                },
                {
                    "id": 2,
                    "question_id": 2,
                    "option": "2 (pernah)",
                    "slugs": "pernah;jarang;",
                    "score": 2,
                    "goto": 3
                },
                {
                    "id": 3,
                    "question_id": 2,
                    "option": "3 (cukup sering)",
                    "slugs": "sering;beberapa kali;",
                    "score": 5,
                    "goto": 3
                },
                {
                    "id": 4,
                    "question_id": 2,
                    "option": "4 (setiap hari)",
                    "slugs": "setiap hari;",
                    "score": 10,
                    "goto": 3
                }
            ],
            "goto": null
        },
        {
            "id": 3,
            "type": "choices",
            "question": "https://linkto.com/video.mp4",
            "description": "apakah anda merokok?",
            "choices": [
                {
                    "id": 5,
                    "question_id": 3,
                    "option": "Ya",
                    "slugs": "ya;yes;iya;",
                    "score": 10,
                    "goto": null
                },
                {
                    "id": 6,
                    "question_id": 3,
                    "option": "Tidak",
                    "slugs": "tidak;no;",
                    "score": 0,
                    "goto": null
                }
            ],
            "goto": null
        }
    ],
    "messages": [
        "[success] read data"
    ],
    "meta": {
        "code": "200-003-OK",
        "status": "success"
    }
    }   
    ```
</details>
<!-- ==== End Of Section -->

<details>
  <summary><kbd>POST /v1/questioner/validate</kbd></summary>
    > Validate the user that want to answer the question. Everyone is just have 2 attemp for answering, as myself and partner.
    <br>

-  <kbd>Request Body as myself</kbd>
    ```json
    {
    "email": "rudi@mail.com",
    "phone":"08123",
    "as":"mysel"
    }
    ```

-  <kbd>Request Body as partner</kbd>
    ```json
   {
    "email": "rudi.partner@mail.com",
    "phone":"08123",
    "as":"partner",
    "partner_email":"rudi@mail.com"
    }
    ```

-  <kbd>Response Body</kbd>
    ```json
    {
    "data": {
        "code_attempt": "3X1dj9HiksJSxYURr2SxLQhDX5vZuRHIsmuBdBqga1tIecz4Hwc8JKHIcIQ7DgX6uvbexSEU4r9xPVUnOZTe1Q==",
        "count_attempt": 0
    },
    "messages": [
        "[success] test attempt added. Start your test."
    ],
    "meta": {
        "code": "200-003-OK",
        "status": "success"
    }
    }
    ```
</details>
<!-- ==== End Of Section -->

<details>
  <summary><kbd>POST /v1/questioner</kbd></summary>
    > Submit the answer of questions
    <br>

-  <kbd>Request Body</kbd>
    ```json
    {
    "code_attempt":"Sz2A7kbp+SoTOF3WhDDF6ybpKQs+bil0d32QH33Dyd34VXPxcTj4LmvI77XBcBmvgjiWZbtaSUGIUGQ+xeApfg==",
    "answer":[
        {
            "question_id": 1,
            "description": "tidak",
            "score":10
        },
         {
            "question_id": 2,
            "description": "ya",
            "score":1
        }
    ]
    }
    ```
-  <kbd>Response Body</kbd>
    ```json
    {
    "data": null,
    "messages": [
        "[success] add answer"
    ],
    "meta": {
        "code": "200-003-OK",
        "status": "success"
    }
    }   
    ```
</details>
<!-- ==== End Of Section -->

<details>
  <summary><kbd>POST /v1/patients</kbd></summary>
  > Add patient
    <br>

-  <kbd>Request Body only email and phone</kbd>
    ```json
    {
    "email": "adi2@mail.com",
    "phone": "0812341"
    }
    ```

-  <kbd>Request Body send all data patient</kbd>
   > if you want to add partner patient, please add partner_email to your json. or if you want to add patient itself, just remove partner_email
    ```json
    {
    "name": "budi partner",
    "email": "budi2.partner@mail.com",
    "password":"qwerty",
    "nik": "12345671",
    "dob": "2023-01-01",
    "phone": "08123456711",
    "gender": "male",
    "marriage_status":"married",
    "nationality": "indonesia",
    "partner_email":"budi@mail.com"
    }
    ```

-  <kbd>Response Body</kbd>
    ```json
    {
    "data": {
        "id": "d7f77642-a6a0-4283-b99a-73e339a16563",
        "name": "budi partner",
        "email": "budi2a.partner@mail.com",
        "nik": "TOBlsuFa3aBs3UfdM3efdL1eFeUNYN5Ptm1wY9+PvVmw35wG",
        "dob": "2023-01-01",
        "phone": "08123456711",
        "gender": "male",
        "marriage_status": "married",
        "nationality": "indonesia"
    },
    "messages": [
        "[success] add patient"
    ],
    "meta": {
        "code": "200-004-OK",
        "status": "success"
    }
    }   
    ```
</details>
<!-- ==== End Of Section -->