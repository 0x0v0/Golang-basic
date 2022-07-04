package main

/* Deklarasi Variable Bisa Menggunakan "var" dan ":="
Perbedaannya "var" bisa digunakan diluar atau didalam fungsi main sifatnya (Undepend/Global)
kalau ":=" harus berada dalam fungsi main sifatnya (Depend/Non Global)
*/

// //Mendeklarasikan (Membuat) Variable
// var variablename type = value
// var murid1 string = "Muhammad"
// var numb int = 1
// var a, b = 2, "Faiz"
// //Or
// variablename := value
// c := 3
// d, e := 4, "Aqil"

////Can used outside main function
// var universitas = "Universitas Darma Persada"

////Cannot used outside main function must be inside main function
// x := 20

func main() {
	/* The code below will print Hello World
	   to the screen, and it is amazing */
	// fmt.Println("Hello, World!")


    
	////For Loop Can Be Declared Like This
    // var i int
	// for i = 0; i <= 100; i++ {
	//     fmt.Println("Looping using Golang = ", i)
	// }
    // //Or Like This
    // for i := 0; i < 100000; i++ {
    //     fmt.Println("Looping using Golang 2nd = ", i)
    // }



	// // Variable Declaration With Initial Value
	// var murid1 string = "Muhammad Faiz"
	// var murid2 = "Aqil Fathoni"
	// x := 2

	// fmt.Println("Murid 1 Bernama: ", murid1)
	// fmt.Println("Murid 2 Bernama: ", murid2)
	// fmt.Println("Ini Adalah Angka: ", x)



	// // Variable Declaration Without Initial Value
	// var a string
	// var b int
	// var c bool

	// fmt.Println("String is: ", a)
	// fmt.Println("Int is: ", b)
	// fmt.Println("Bool is: ", c)



	// // Value Assignment After Declaration
	// var murid1 string
	// murid1 = "Muhammad Faiz"

	// fmt.Println("Murid 1 Bernama: ", murid1)

    // Callback
    // x := 20
	// fmt.Println(universitas)
    // fmt.Println(x)

    //Go Multiple Variable Declaration
    // var a, b, c, d int =  1, 2, 3, 4
    //or
    // var a, b, c, d = 1, 2, 3, 4

    // fmt.Println("Ini Adalah a = ", a)
    // fmt.Println("Ini Adalah b = ", b)
    // fmt.Println("Ini Adalah c = ", c)
    // fmt.Println("Ini Adalah d = ", d)



    // //Different Types of Variables in the same line
    // var a, b = 5, "Muhammad"
    // c, d := 6, "Faiz"

    // fmt.Println("Ini adalah int from var a type: ", a)
    // fmt.Println("Ini adalah string from var b type: ", b)
    // fmt.Println("Ini adalah int from := c type: ", c)
    // fmt.Println("Ini adalah string from := d type: ", d)



    /* Go Variable Declaration in a Block
    In Go nameVariable typeData = value
    In C/C++ typeData nameVariable = value
    This is Inverted be Careful!
    */
    // var (
    //     a int
    //     b = 1
    //     c string = "Muhammad Faiz"
    // )

    // fmt.Println("Ini adalah int kosong dari variable a: ", a)
    // fmt.Println("Ini adalah variable b berupa int: ", b)
    // fmt.Println("Ini adalah variable c berupa string: ", c)
}

/* 
Aturan Penamaan Variabel di Golang
Sebuah variabel dapat memiliki nama pendek (seperti x dan y) atau nama yang lebih deskriptif (umur, harga, carname, dll).

Aturan penamaan variabel di Golang sebagai berikut:

- Nama variabel boleh dimulai dengan huruf atau karakter garis bawah (_)
- Nama variabel tidak boleh dimulai dengan angka
- Nama variabel hanya boleh berisi karakter alfanumerik dan garis bawah 
    Contoh: (a-z, A-Z, 0-9, dan _ )
- Nama variabel peka huruf besar/kecil (usia, Usia, dan USIA adalah tiga variabel berbeda)
- Tidak ada batasan panjang nama variabel
- Nama variabel tidak boleh mengandung spasi
- Nama variabel tidak boleh berupa kata kunci Go apa pun
*/
