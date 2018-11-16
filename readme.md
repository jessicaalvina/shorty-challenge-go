
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

### Architecture Structure
![architecture diagram](storage/golang%20architecture%20diagram.png)
```
- rl-ms-boilerplate-go
 |- constants
 |- controllers
 |- helpers
 |- middlerware
 |- models
 |- objects
 |- repositories
 |- services
 |- storage
    |- logs
```
#### middleware

Digunakan untuk menyimpan middleware-middleware yang akan digunakan, contoh `cors_middleware` atau `oauth_middleware`.

#### controllers

Controller bertugas untuk menghandle HTTP Request, routing dimasukkan per-controller dan digroup berdasarkan controller, controller terhubung dengan service.

#### service

Service bertugas untuk menghandle business logic, service memungkinkan untuk memanggil banyak repository dan atau service lain.

#### repositories

Repository bertugas untuk menghandle query-query ke database atau storage lainnya, jangan menambahkan logic-logic programming berat pada layer ini.

#### models

Models bertugas untuk menampung model-model representasi database schema yang dapat digunakan untuk kepentingan migration atau enkapsulasi data.

#### objects

Objects bertugas sebagai transporter data antar layer, objects juga bertugas untuk melakukan enkapsulasi data dari HTTP request ataupun sebagai response dari sebuah request.

#### helpers

Bertugas untuk menyimpan helpers atau libraries yang sering digunakan contohnya `error_helper` atau `redis_helper`.

#### constants

Digunakan untuk menyimpan constant-constant seperti `error_constants` atau `configuration_constants`.

#### storage

Storage bertugas untuk menyimpan file-file seperti log error atau temporary file storage.


## TODO
- Endpoint documentation
- Authorization middleware

## How to Setup

### Docker environment

membuat direktori khusus untuk menyimpan workspace untuk project, dengan struktur sepertidibawah ini:
```
- ralali-golang-docker
 |- config
 |- golang
```

pada root directory `ralali-golang-docker` tambahkan konfigurasi docker dengan nama file `docker-compose.yml` dengan isi file sebagai berikut:
```
version: '2.1'
services:
    ralali_golang:
        build:
            context: ./config
            dockerfile: golang-extensions
        container_name: ralali_golang
        command: > 
            sh -c "
            cd /go/src/ralali.com && go get -v -d && go run main.go"
        ports:
            - "3000:3000"
        volumes:
            - './golang:/go:rw'
```

pada konfigurasi diatas, kita membuat sebuah service yang bernama `ralali_golang` dan mengarahkan port 3000 host ke port 3000 docker.

directory mapping kita arahkan ke directory `golang` pada host yang dimapping ke directory `/go` pada docker.

setelah docker configurationnya sudah disiapkan, berikutnya adalah menyiapkan docker file yang diletakkan pada directory config dengan nama file `golang-extension` yang berisi:
```
FROM golang:1.11.1-alpine3.8
   
RUN apk add --no-cache ca-certificates \
       dpkg \
       gcc \
       git \
       musl-dev \
       bash
```
pada dockerfile diatas kita menggunakan golang versio 1.11., dan menambahkan beberapa library kedalam docker image golang yang akan kita jalankan.

jika dockerfile dan dan docker configuration sudah disiapkan, maka tahap selanjutanya adalah menyiapkan projectnya.

### Project Preparation

melakukan preparasi project yang dapat diclone dari github repository kedalam docker volume yang sudah kita mapping tadi, tapi sebelumnya kita menambahkan directory terlebih dahulu kedalam directory `golang` host kita sehingga directorynya menjadi seperti dibawah ini:
```
- ralali-golang-docker
 |- config
    |- golang-extensions
 |- golang
    |- src
```

project yang akan disetup didalam directory `golang -> src` dengan menggunakan command:

```git clone git@github.com:ralali/rl-ms-boilerplate-go.git golang/src/ralali.com```

setelah project sudah berhasil diclone hal berikutnya adalah menyesuaikan environment file yang dapat dicopy dari `.env.example` menjadi `.env`, developer mengisi environment berdasarkan kredensial developer masing-masing.

setelah project dan environment sudah berhasil disiapkan, maka berikutnya adalah melakukan build project golang kita dengan menjalankan docker yang sudah kita setting tadi, untuk menjalankan docker dilakukan dengan menggunakan command berikut ini:

```
docker-compose up --build
``` 

setelah berhasil build, maka akan tampil tampilan seperti dibawah ini:
```
ralali_golang    | username:password@tcp(dev.ralali.xyz:3308)/dbname?parseTime=1&loc=Asia%2FJakarta
ralali_golang    | 0.0.0.0:3000
ralali_golang    | [GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.
ralali_golang    | 
ralali_golang    | [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
ralali_golang    |  - using env:	export GIN_MODE=release
ralali_golang    |  - using code:	gin.SetMode(gin.ReleaseMode)
ralali_golang    | 
ralali_golang    | [GIN-debug] GET    /v1/users/:id             --> ralali.com/controllers.(*V1UserController).GetById-fm (4 handlers)
ralali_golang    | [GIN-debug] POST   /v1/users/:id             --> ralali.com/controllers.(*V1UserController).UpdateById-fm (4 handlers)
ralali_golang    | [GIN-debug] POST   /v2/users/:id             --> ralali.com/controllers.(*V2UserController).UpdateById-fm (4 handlers)
ralali_golang    | [GIN-debug] Listening and serving HTTP on 0.0.0.0:3000
```

dengan tampilan seperti diatas, maka project golang berhasil di serve pada port `3000`, developer dapat mencoba mengakses endpoint-endpoint yang ada.

### Unit Testing
untuk menjalankan unit testing, developer dapat menjalankan command dibawah ini:
```
go test repositories -v -cover
go test services -v -cover
go test controllers -v -cover
``` 

### Code Versioning
versioning level dilakukan pada layer 
- `controllers` 
- `objects` 
- `repositories` 
- `services`

setiap file pada layer-layer tersebut diberi prefix version dengan format snake case, seperti pada contoh yang ada `v1_user_controller.go` yang berarti user_controller versi 1, dan pada level struct diberi prefix versi dalam bentuk upper camel case seperti pada contoh diproject ini `V1UserController` yang berarti controller `UserController` versi 1.

##### Sample Case
terdapat contoh kasus pada saat update data user parameter dan response yang diterima dan diberikan oleh `v1` dan `v2` berbeda, pertama-tama, developer harus melakukan definisi DTO nya terlebih dahulu pada layer `objects`:

- v1_user_object.go
- v2_user_object.go

pada kedua file tersebut terdapat object response dan object request, setelah melakaukan devinisi DTO, developer kemudiam melakukan definisi repository pada layer `repository` yang menggunakan DTO pada masing-masing versi.

setelah melakukan definisi pada `repository`, kemudian dilakukan definisi pada layer `service` dan `controller`, perhatikan routing group pada masing masing controller harus sesuai dengan versi yang didefinisikan.    

### Database Migration
untuk menjalankan database migration, developer dapat menjalankan command dibawah ini:
```
go run database_migration.go
``` 
database migration akan melakukan sinkronisasi skema dan indeks database berdasarkan skema yang dibuat pada directory models dan perintah yang ada di `database_migration.go`
