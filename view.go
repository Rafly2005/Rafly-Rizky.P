package view

import (
	"bufio"
	"fmt"
	"os"
	"pegawai/model.go"
	"pegawai/node.go"
)

func Insert() {
	var kota, nama, notelp, email string
	var id, nomer int
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukan ID Pegawai: ")
	fmt.Scanln(&id)

	fmt.Print("Masukan nama Pegawai: ")
	nama, _ = reader.ReadString('\n')
	nama = nama[:len(nama)-1]

	fmt.Print("Masukan Jalan: ")
	jalan, _ = reader.ReadString('\n')
	jalan = jalan[:len(jalan)-1]

	fmt.print("Masukan Nomer rumah: ")
	fmt.Scan(&nomer)

	fmt.Print("Masukan Kota: ")
	fmt.Scan(&kota)

	fmt.Print("Masukan Nomer Telpon: ")
	fmt.Scan(&notelp)

	fmt.Print("Masukan Email: ")
	fmt.Scan(&email)

	// create  new pegawai
	pegawai := node.Pegawai{
		ID:        	id,
		Nama:      	nama,
		Alamat:		node.Address{jalan, kkota, nomer},
		NoTelp:		notelp,
		Email:		email,
	}

	// insert to DaftarPegawai
	cek := model.CreatePegawai(pegawai)
	if cek {
		fmt.Println("== Pegawai berhasil ditambahkan ==")
	} else {
		fmt.Println("Pegawai gagal ditambahkan") 
	}
	fmt.Println()
}

func Views() {
	fmt.Println("Daftar Pegawai")
	for i, emp := range model.DaftarPegawai(){
		fmt.Println("--- Pwgawai ke -", i+1, "---")
		fmt.Println("ID Pegawai\t: ", emp.ID)
		fmt.Println("Nama Pegawai\t: ", emp.Nama)
		fmt.Println("Alamat\t\t: ", emp.Alamat.Jalan, emp.Alamat.Nomer, emp.Alamat.Kota)
		fmt.Println("No Telpon\t:", emp.NoTelp)
		fmt.Println("Email\t\t: ", emp.Email)
		fmt.Println()
	}
}

func Update() {
	var id, nomer int
	var nama, kota, notelp, email string
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukan ID Pegawai yang akan diupdate: ")
	fmt.Scan(&id)

	fmt.Print("Masukan Nama Pegawai: ")
	nama, _ = reader.ReadString('\n')
	nama = nama [:len(nama)-1]

	fmt.Print("Masukan Kota: ")
	fmt.Scan(&kota)

	fmt.Print("Masukan Jalan: ")
	jalan, _ = reader.ReadString('\n')
	jalan = jalan [:len(jalan)-1]

	fmt.Print("Masukan nomer rumah: ")
	fmt.Scan(&nomer)

	fmt.Print("Masukan Nomer Telpon: ")
	fmt.Scan(&notelp)

	fmt.Print("Masukan Email: ")
	fmt.Scan(&email)

	// create new pegawai
	pegawai := node.Pegawai{
		ID: 	id,
		Nama: 	nama,
		Alamat: node.Address{jalan, kota, nomer}
		NoTelp: notelp,
		Email: 	email,
	}

	// update pegawai
	cek := model.UpdatePegawai(pegawai, id)
	if cek {
		fmt.Println("== Pegawai Berhasil Diupdate ==")
	} else {
		fmt.Println("Pegawai gagal diupdate")
		
	}
}

func Delete() {
	var id int
	fmt.Print("Masukan ID Pegawai yang akan dihapus: ")
	fmt.Scan(&id)

	cek := model.DeletePegawai(id)
	if cek {
		fmt.Println("== Pegawai berhasil dihapus ==")
	} else {
		fmt.Println("Pegawai gagal dihapus")
	}
	fmt.Println()
}
