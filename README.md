# README
Marketplace Merah Kuning Hijau

## POSTMAN COLLECTION
- Download Postman Collection [Di Sini](https://github.com/ghivarra/backend-ghivarra-senandika-rushdie/blob/main/Backend%20Marketplace.postman_collection.json)

## DEPENDENCY
- Go versi terbaru. Download [Di Sini](https://go.dev/)
- PostgreSQL 15. Lihat cara instalasi [Di Sini](https://www.postgresql.org/download/)

## MIGRASI DATABASE
- Pastikan PostgreSQL sudah terinstall dan berjalan normal.
- Buat User, Password, dan Database yang akan digunakan dengan command di bawah (khusus Linux / MacOS / UNIX-based):

- Masuk ke postgres

```
$ -u postgres psql
```

- Lalu buat user dengan password yang akan digunakan untuk migrasi dan mengakses database

```
CREATE USER nama_user WITH ENCRYPTED PASSWORD 'password';
```

- Buat database

```
CREATE DATABASE nama_db;
```

- Ubah kepemilikan tabel menjadi nama_user

```
ALTER DATABASE nama_db OWNER TO nama_user;
```
- Keluar dari PostgreSQL
```
\q
```

- Lalu selanjutnya jalankan command di bawah ini tepat di root folder aplikasi ini melalui CLI, dan masukkan password:

```
$ psql -U nama_user -d nama_db < marketplace.sql
```

## DEPLOYMENT

Sebelum deployment, pastikan konfigurasi file .env sudah sesuai dengan yang akan dijalankan.

- ### Docker (Any OS)

Install Docker, deploy, dan generate image berdasarkan file Dockerfile yang tersedia.
Manual instalasi docker, pembuatan image, dan lainnya ada [Di Sini](https://www.youtube.com/watch?v=ZyBBv1JmnWQ)

- ### CLI / Bash Command (Linux and MacOS)

Di folder paling utama atau root folder, eksekusi command di bawah ini:

```
bash deploy.sh
```

atau 

```
./deploy.sh
```

Selanjutnya aplikasi akan bisa diakses sesuai dengan konfigurasi pada file .env yang diberikan.