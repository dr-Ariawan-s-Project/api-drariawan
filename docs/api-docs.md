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
            "question": "Silakan memberi masukkan dan saran. (Bila tidak ada harap isi dg \"-\".)",
            "description": "",
            "url_video": "",
            "section": "saran",
            "choices": null,
            "goto": 2
        },
        {
            "id": 2,
            "type": "choices",
            "question": "Saya telah mengerti dan memahami maksud dan tujuan pengisian kuesioner ini.\nDengan ini saya sukarela bersedia untuk menjadi responden dalam penelitian ini\ntanpa adanya paksaan atau tekanan dari siapapun.",
            "description": "",
            "url_video": "",
            "section": "konfirmasi",
            "choices": [
                {
                    "id": 1,
                    "question_id": 1,
                    "option": "Setuju",
                    "slugs": "setuju;ya;",
                    "score": 0,
                    "goto": 2
                },
                {
                    "id": 2,
                    "question_id": 1,
                    "option": "Tidak setuju",
                    "slugs": "tidak;",
                    "score": 1,
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

<details>
  <summary><kbd>PUT /v1/patients/{patient_id}</kbd></summary>
  > Edit patient
    <br>

-  <kbd>Parameter</kbd>
    ```
    patient_id
    ```

-  <kbd>Request Body update data patient</kbd>

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
        "[success] update patient"
    ],
    "meta": {
        "code": "200-004-OK",
        "status": "success"
    }
    }   
    ```
</details>
<!-- ==== End Of Section -->

<details>
  <summary><kbd>GET /v1/patients</kbd></summary>
  > GET all patient
    <br>

-  <kbd>Query params</kbd>
    ```
    - page
    - limit
    - search
    ```

-  <kbd>Response Body</kbd>
    ```json
    {
    "data": [
        {
            "id": "022b7a27-6890-403c-978c-aa33448d78bf",
            "name": "rudi partner 2",
            "email": "rudi.partner@mail.com",
            "phone": "08123",
            "partner_id": "7c0706b3-8cdd-43b2-8262-c542e2cae870",
            "partner": {
                "id": "7c0706b3-8cdd-43b2-8262-c542e2cae870"
            }
        },
        {
            "id": "38e82d68-cceb-4063-8c99-ce4e2676f26d",
            "name": "budi 2",
            "email": "budi@mail.com",
            "nik": "OKdIIuAJgvgRJDGb97E5cA/hPNROZsbpVqz9KcVv/G+EPxUi",
            "dob": "2023-01-01",
            "phone": "0812345671",
            "gender": "male",
            "marriage_status": "married",
            "nationality": "indonesia"
        }
    ],
    "messages": [
        "[success] read data"
    ],
    "meta": {
        "code": "200-004-OK",
        "status": "success"
    }
    }  
    ```
</details>
<!-- ==== End Of Section -->

<details>
  <summary><kbd>GET /v1/patients/{patient_id}</kbd></summary>
  > GET patient by ID
    <br>

-  <kbd>Parameter</kbd>
    ```
    patient_id
    ```

-  <kbd>Response Body</kbd>
    ```json
    {
    "data": {
        "id": "022b7a27-6890-403c-978c-aa33448d78bf",
        "email": "rudi.partner@mail.com",
        "phone": "08123",
        "partner_id": "7c0706b3-8cdd-43b2-8262-c542e2cae870"
    },
    "messages": [
        "[success] read data"
    ],
    "meta": {
        "code": "200-004-OK",
        "status": "success"
    }
    }  
    ```
</details>
<!-- ==== End Of Section -->

<details>
  <summary><kbd>GET /v1/patients/profile</kbd></summary>
  > GET patient profile
    <br>

-  <kbd>Response Body</kbd>
    ```json
    {
    "data": {
        "id": "022b7a27-6890-403c-978c-aa33448d78bf",
        "email": "rudi.partner@mail.com",
        "phone": "08123",
        "partner_id": "7c0706b3-8cdd-43b2-8262-c542e2cae870"
    },
    "messages": [
        "[success] read data"
    ],
    "meta": {
        "code": "200-004-OK",
        "status": "success"
    }
    }  
    ```
</details>
<!-- ==== End Of Section -->

<details>
  <summary><kbd>DELETE /v1/patients/{patient_id}</kbd></summary>
  > DELETE patient by ID
    <br>

-  <kbd>Parameter</kbd>
    ```
    patient_id
    ```

-  <kbd>Response Body</kbd>
    ```json
    {
    "data": null,
    "messages": [
        "[success] delete data"
    ],
    "meta": {
        "code": "200-004-OK",
        "status": "success"
    }
    }
    ```
</details>
<!-- ==== End Of Section -->

<details>
  <summary><kbd>GET /v1/dashboard</kbd></summary>
    > Dashboard feature
    <br>

-  <kbd>Response Body</kbd>
    ```json
   {
    "data": {
        "questioner_all": 7,
        "questioner_need_assess": 6,
        "questioner_this_month": 1
    },
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
  <summary><kbd>GET /v1/dashboard/questioner</kbd></summary>
    > Dashboard feature - grafik
    <br>

-  <kbd>Response Body</kbd>
    ```json
   {
    "data": [
        {
            "month": "januari",
            "count": 0
        },
        {
            "month": "februari",
            "count": 0
        },
        {
            "month": "maret",
            "count": 0
        },
        {
            "month": "april",
            "count": 0
        },
        {
            "month": "mei",
            "count": 0
        },
        {
            "month": "juni",
            "count": 0
        },
        {
            "month": "juli",
            "count": 0
        },
        {
            "month": "agustus",
            "count": 0
        },
        {
            "month": "september",
            "count": 0
        },
        {
            "month": "oktober",
            "count": 40
        },
        {
            "month": "november",
            "count": 0
        },
        {
            "month": "desember",
            "count": 0
        }
    ],
    "messages": [
        "[success] read data"
    ],
    "meta": {
        "code": "200-007-OK",
        "status": "success"
    }
    }      
    ```
</details>
<!-- ==== End Of Section -->
