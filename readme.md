
# Ralali Golang Microservice Boilerplate

### Pendahuluan
Dengan adanya kebutuhan untuk memecah Arsitektur Ralali yang Monolitik menjadi microservice, maka hadirlan boilerplate ini yang dapat digunakan oleh internal tim ralali untuk menunjang pembangunan microservice menggunakan bahasa pemrograman Go, arsitektur pada mikroservice ini diadoptasi berdasarkan teori yang ada pada link-link berikut ini:

- https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html
- http://www0.cs.ucl.ac.uk/staff/A.Finkelstein/crsenotes/1B1499REQTS.pdf
- https://blog.alexellis.io/golang-writing-unit-tests/
- http://doc.gorm.io

### Main Open Source Library
- https://github.com/gin-gonic/gin
- https://github.com/jinzhu/gorm -> (https://github.com/go-sql-driver/mysql)
- https://github.com/joho/godotenv

### Directory Structure
```
- rl-ms-boilerplate-go
 |- constants
 |- controllers
 |- helpers
 |- middlerware
 |- models
 |- repositories
 |- requests
 |- services
 |- storage
   |- logs
```
#### constants

Digunakan untuk menyimpan constant-constant seperti `error_constants` atau `configuration_constants`.

#### controllers

Controller bertugas untuk menghandle HTTP Request, routing dimasukkan per-controller dan digroup berdasarkan controller, controller terhubung dengan repository dan service.

#### helpers

Helper bertugas untuk menampung fungsi-fungsi yang sering digunakan, contohnya `error_handling`, `http_response_handling`, atau `language_handling`.

#### middleware

Digunakan untuk menyimpan middleware-middleware yang akan digunakan, contoh `cors_middleware` atau `oauth_middleware`.

#### models

Models bertugas untuk menampung model-model representasi database schema yang dapat digunakan untuk kepentingan migration atau enkapsulasi data.

#### repositories

Repository bertugas untuk menghandle use-case use-case dari proses bisnis, misalnya untuk mendapatkan order detail, payment detail, atau menjalankan proses-proses logic kebutuhan bisnis.

#### request

Requests bertugas menampung struct-struct untuk mengenkapsulasi request dari client kepada HTTP handler (controller).

#### service

Service digunakan untuk menghandle service-service seperti 3rd party service atau contohnya cache service.

#### storage

Storage bertugas untuk menyimpan file-file seperti log error atau temporary file storage.


## TODO
- Endpoint documentation
- Authorization middleware
- Ralali Old OAuth middleware
- Ralali New OAuth middleware
