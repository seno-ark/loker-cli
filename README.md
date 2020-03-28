# loker-cli

Program loker interaktif sederhana dengan menggunakan command line untuk menyimpan kartu identitas.
Program ini dibuat menggunakan bahasa pemrograman Golang


#### Cara menjalankan program

Masuk ke direktori proyek Go  
```cd $GOPATH/src```

clone Repository  
```git clone https://github.com/seno-ark/loker-cli.git```

install program  
```go install loker-cli```

jalankan program loker-cli  
```$GOPATH/bin/loker-cli```


#### Cara menggunakan program

1. Membuat loker terlebih dahulu, dengan cara menjalankan perintah `init <jumlah loker>` lalu tekan Enter.
2. Memasukkan kartu identitas ke dalam loker, dengan perintah `input <tipe identitas> <nomor identitas>`
3. Untuk melihat semua isi loker, dapat menggunakan perintah `status`
4. Terdapat fitur pencarian nomor loker berdasarkan nomor identitas, yaitu dengan menggunakan perintah `find <nomor identitas>`
5. Selain itu juga terdapat fitur pencarian nomor identitas yang disaring berdasarkan tipe identitas: `search <tipe identitas>`
6. Untuk mengeluarkan kartu identitas dari loker, gunakan perintah `leave <nomor loker>`
7. Untuk keluar dari aplikasi, ketik `exit` lalu tekan Enter.