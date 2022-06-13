SET UP

1. Jika file final_project.db tidak ada pada folder \final-project\backend\database maka jalankan migration pada path \final-project\backend\database\migration\main.go
   cd backend
   go run .\database\migration\main.go

2. Jika ada perubahan pada struktur database maka hapus file final_project.db dan lakukan langkah 1

3. Untuk menjalankan server jalankan main.go pada path \final-project\backend\main.go
   go run .\backend\main.go

TESTING LOGIN

1. Jalankan server pada langkah SET UP nomor 3

2. Untuk testing local dapat dilakukan pada aplikasi postman

3. Untuk Login masukan http://localhost:8080/login dengan methotd POST pada postman

4. Pilih Body pada postman

5. Pilih raw kemudian pilih JSON sebagai inputan

6. Tuliskan :
   {
   "username": "ucup",
   "password": "ucup123"
   }

7. Jika Success maka akan menampilkan data ucup

TESTING REGISTER

1. Jalankan server pada langkah SET UP nomor 3

2. Untuk testing local dapat dilakukan pada aplikasi postman

3. Untuk register masukan http://localhost:8080/register methotd POST pada postman

4. Pilih Body pada postman

5. Pilih raw kemudian pilih JSON sebagai inputan

6. Tuliskan dengan format:
   {
   "username": ....
   "password": ....
   "nama": ....
   "alamat": ....
   "noHP": ....
   "role": ....
   }

7. Jika akun user belum ada di database maka akan success tetapi jika sudah ada maka akan ada pesan error

TESTING LOGOUT

1. Jalankan server pada langkah SET UP nomor 3

2. Untuk testing local dapat dilakukan pada aplikasi postman

3. Untuk logout masukan http://localhost:8080/logout methotd POST pada postman

4. Pastikan sudah login terlebih dahulu, jika belum login maka ada pesan kesalahan

5. Jika berhasil akan ada pesan success
