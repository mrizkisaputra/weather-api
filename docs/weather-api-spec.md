# Wheater API

## Get Current and Forecasted Weather
- **Endpoint:** ``/api/weather``
- **Method:** ``GET``
- **Description:** Mengambil cuaca saat ini dan prakiraan cuaca di lokasi tertentu.  
Jika tidak ada tanggal1 atau tanggal2 yang ditentukan, permintaan akan mengambil prakiraan di lokasi yang diminta untuk 15 hari ke depan.
- **URL Parameters:**  
  - ``location`` (required) - Lokasi yang akan diambil datanya (alamat, latitude/longitude, atau kode pos).
  - ``date1`` (optional) - Tanggal mulai dalam format yyyy-MM-dd atau waktu spesifik yyyy-MM-ddTHH:mm:ss.
  - ``date2`` (optional) - Tanggal akhir (jika ingin mengambil data rentang waktu).


- **Request Example**:  
  ```
  http://localhost:8080/api/wheater/jakarta, indonesia
  http://localhost:8080/api/wheater/jakarta, indonesia/2024-10-01/2024-10-10
  ```

- **Response Example**:  
  ````json
  {
    "status": 200,
    "latitude": -6.17148,
    "longitude": 106.826,
    "address": "Jakarta, Indonesia",
    "timezone": "Asia/Jakarta",
    "days": [
        {
          "date": "2024-10-04",
          "datetimeEpoch": 1727974800,
          "tempmax": 93.9,
          "tempmin": 76.9,
          "temp": 85.1,
          "hours": [
                {
                  "time": "00:00:00",
                  "datetimeEpoch": 1727974800,
                  "temp": 78.7,
                  "feelslike": 78.7,
                  "humidity": 88.75,
                  "dew": 75.1
                },
                {
                  "time": "01:00:00",
                  "datetimeEpoch": 1727978400,
                  "temp": 78.7,
                  "feelslike": 78.7,
                  "humidity": 88.75,
                  "dew": 75.1
                },
                {...},
                {...}
          ] 
        },
        {...},
        {...},
        {...},
        {...},
        {...},
        {...},
        {...},
        {...},
        {...},
        {...},
        {...},
        {...},
        {...},
        {...},
    ]
  }
  ````