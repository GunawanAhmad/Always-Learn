package main

import "fmt"

type calon struct {namaCalon string; jumPemilih int}
type partai struct {Namap string; jumPemilihp int; namaAnggota[100] calon; noUrutPartai int}
var p[100] partai
type TPS struct{jumGolput int; namaPemilih[100] string }
var arrTps[100]TPS
var jumPartai,anggota int
var jumTps, jumPemilihTps int



func main() {
	var nomor int
	fmt.Println("1. Admin")
	fmt.Println("2. Pemilih")
	fmt.Scan(&nomor)
	fmt.Println()
	fmt.Println()
	if nomor == 1 {
		fmt.Println("1. Masukan Data Partai ")
		fmt.Println("2. Masukan Data Pemilih dan TPS ")
		fmt.Println("3. Hasil Pemilihan")
		fmt.Println("4. Edit Data")
		fmt.Println("5. Cari Data")
		fmt.Println("6. Daftar Partai dan anggota")
		fmt.Println("7. Tambah Data Partai")
		fmt.Println("8. Hasil Perurutan Suara")
		fmt.Println("9. Hasil Golput di TPS")
		fmt.Println("10. Hapus Partai")
		fmt.Println("11. Tambah Data TPS")

		fmt.Scan(&nomor)
		switch {
		case nomor==1 :
			masukan(&jumPartai,&anggota)
			fmt.Println()
			fmt.Println()
			main()
		case nomor==2:
			DataTPS(jumTps,jumPemilihTps,&jumTps,&jumPemilihTps)
		case nomor==3:
			hasilPemilihan(jumPartai,anggota)
		case nomor==4:
			mengedit(jumPartai, anggota)
		case nomor==5:
			searching(jumPartai,anggota)
		case nomor==6:
			daftarpartai(jumPartai,anggota)
		case nomor ==7:
			tambahData(jumPartai,anggota, &jumPartai)
		case nomor==8:
			selSort(jumPartai,anggota)
		case nomor==9:
			suaraTPS(jumTps)
		case nomor ==10:
			delete(jumPartai, &jumPartai)
		case nomor == 11:
			TambahTPS(jumTps,jumPemilihTps,&jumTps)
		default:
			main()
	}
		
	} else if nomor == 2 {
		pemilihan(jumPartai,anggota,jumTps,jumPemilihTps)
	} else {
		main()
	}
		

}



//fungsi masukan untuk menginput daaftar anggota yang ingin dicalonkan
func masukan(jumPartai *int, anggota *int) {
	var n string
	if p[0].noUrutPartai != 0 {
		fmt.Println("Apakah Anda Ingin menghapus Semua Data Partai Yang Ada ? (ya/tidak) ")
		fmt.Scanln(&n)
		if n=="tidak" {
			main()
		}
	}
	var found int
	var urut int
	var stop bool
	stop = false
	fmt.Print("Masukan Jumlah partai : ")
	fmt.Scan(&*jumPartai)
	fmt.Print("Masukan Jumlah Calon : ")
	fmt.Scan(&*anggota)
	for i:= 0; i<*jumPartai; i++ {
		fmt.Printf("Masukan Nama Partai %v : ", i+1)
		fmt.Scan(&p[i].Namap)
		fmt.Print("Masukan Nomor Urut Partai  : ")
		fmt.Scan(&urut)
		SearchUrut(urut,&found)
		if found == 0 {
			p[i].noUrutPartai = urut
		} else if found == 1 {
			for found == 1 && !stop {
				fmt.Print("Nomor Tersebut Sudah Ada, Masukan Lagi  : ")
				fmt.Scan(&urut)
				if urut == string(urut) {
					fmt.Println("Salah masukan")
				SearchUrut(urut,&found)
				if found == 0 {
				  
				   p[i].noUrutPartai = urut
			
				   stop = true
				} 
			}
		}
		for j:=0; j<*anggota; j++ {
			fmt.Printf("Masukan Nama Calon %v : ", j+1)		
			fmt.Scan(&p[i].namaAnggota[j].namaCalon)	
			
		}

	}
	
		
}



//fungsi pemilihan untuk pengguna memilih anggota yang ada dalam array
/*waktu yang digunakan menggunakan input-an "habis", jika string tersebut 
di inputkan maka semua pemilih terhitung golput */
func pemilihan(jumPartai int,anggota int, jumTps int, jumPemilihTps int) {
	var pilih string
	daftarp(jumPartai,anggota)
	var ketemu bool
	var n string

	if jumTps == 0 {
		fmt.Print("Data Pemilih Belum Ada")
		fmt.Println()
		fmt.Println()
		main()
	}
	for a:=0; a<jumTps; a++ {
		for b:=0; b<jumPemilihTps; b++ {
			fmt.Printf("Apakah saudara %v ingin melakukan pemilihan ? (ya/tidak) ", arrTps[a].namaPemilih[b])
			fmt.Scan(&n)
			if n=="ya" {
				fmt.Println("Masukan Pilihan Anda : ")
				fmt.Scan(&pilih)
				if pilih == "habis" {
					for a<jumTps{
						for b<jumPemilihTps {
							arrTps[a].jumGolput = arrTps[a].jumGolput + 1
							b= b+1 
						}
						a= a+1 
					}
					a = jumTps
					b = jumPemilihTps
				} else {
						i:=0
						ketemu = false
						for !ketemu && i<jumPartai {
							j:=0
							for j<anggota && !ketemu {
							if pilih == p[i].namaAnggota[j].namaCalon {
									ketemu = true
									fmt.Println("Pilihan Berhasil")
									p[i].namaAnggota[j].jumPemilih = p[i].namaAnggota[j].jumPemilih + 1
									p[i].jumPemilihp = p[i].jumPemilihp + 1
								} else {
									j++ 
								}
								
							}
							if pilih == p[i].Namap {
								ketemu = true
								p[i].jumPemilihp = p[i].jumPemilihp + 1
								fmt.Println("Pilihan Berhasil")
							} else {
								i++
							}
						} 
						if ketemu == false {
							arrTps[a].jumGolput = arrTps[a].jumGolput + 1
							fmt.Println("Suara Tidak Sah")
						}
				}
			
			} else {
				arrTps[a].jumGolput = arrTps[a].jumGolput + 1
				fmt.Println("Suara anda terhitung Golput")
			}
		} 
	}
	fmt.Println()
	fmt.Println()
			
	main()
	
}

//menampi lkan daftar paprtai yang telah diinput
func daftarp(jumPartai,anggota int) {
	for i:=0; i<jumPartai; i++ {
		fmt.Print("Nama Partai : ")
		fmt.Println(p[i].Namap)
		fmt.Print("No urut : ")
		fmt.Println(p[i].noUrutPartai)
		fmt.Println("Nama Anggota : ")
		for j:=0; j<anggota; j++ {
			fmt.Println(p[i].namaAnggota[j].namaCalon)
		}
	} 
}

func hasilPemilihan(jumPartai int, anggota int) {
	fmt.Println("Hasil Pemilihan Calon Legislatif")
	for i:=0; i<jumPartai; i++ {
		fmt.Println("Nama partai : ", p[i].Namap)
		fmt.Println("Jumlah Pemilih : ", p[i].jumPemilihp)
		for j:=0; j<anggota; j++ {
			fmt.Println("Nama Anggota : ", p[i].namaAnggota[j].namaCalon)
			fmt.Println("Jumlah Pemilih : ", p[i].namaAnggota[j].jumPemilih)
		}
	}
	fmt.Println()
	fmt.Println()
	main()
}

func mengedit(jumPartai int, anggota int) {
 	var cari,edit string
 	var ketemu bool
 	ketemu = false
 	daftarp(jumPartai,anggota)
 	fmt.Print("Masukan nama calon/partai yang ingin di ubah : ")
 	fmt.Scan(&cari)
 	fmt.Print("Masukan Data yang ingin dimasukan : ")
 	fmt.Scan(&edit)
 	i:= 0
 	for !ketemu && i<jumPartai {
				j:=0
				for j<anggota && !ketemu {
				if cari == p[i].namaAnggota[j].namaCalon {
						ketemu = true
						fmt.Println("data berhasil diubah")
						p[i].namaAnggota[j].namaCalon = edit
					} else {
						j++
					}
					
				}
				if cari == p[i].Namap {
					ketemu = true
					fmt.Println("data berhasil diubah")
					p[i].Namap = edit
				} else {
					i++
				}
			} 
			fmt.Println()
			fmt.Println()
 	daftarp(jumPartai,anggota)
 	fmt.Println()
	fmt.Println()
 	main()
}


func delete(Partai int, jumPartai *int) {
	var cari int
	fmt.Print("Masukan No Urut Partai yang ingin dihapus : ")
	fmt.Scan(&cari)
	var arrBaru[100]partai
	j := 0
	for i:=0; i<Partai; i++ {
		if cari != p[i].noUrutPartai {
			arrBaru[j] = p[i]
			j++
		} else if cari == p[i].noUrutPartai {
			i--
		}
		i++
	}
	*jumPartai--
	p = arrBaru
	fmt.Println()
	fmt.Println()
	main()
}

func searching(jumPartai int,anggota int) {
	var ketemu bool
	var cari string
	ketemu = false
	i:=0
	fmt.Print("Masukan Data yang ingin dicari : ")
	fmt.Scan(&cari)
	for i<jumPartai && !ketemu {
		if cari == p[i].Namap {
			ketemu = true
			fmt.Println("Data Ditemukan")
		} else {
			i++
		}
		j:= 0
		for j<anggota && !ketemu {
			if cari == p[i].namaAnggota[j].namaCalon {
				ketemu = true
				fmt.Println("Data Ditemukan")
			} else {
				j++
			}
		} 
	}

	fmt.Println()
	main()
}

func tambahData(Partai int,anggota int, jumPartai *int) {
	var tambah int
	fmt.Print("Masukan Jumlah Partai yang ingin ditambahkan : ")
	fmt.Scan(&tambah)
	var poin int
	var found,urut int
	poin = Partai + tambah
	for Partai < poin {
		fmt.Print("Masukan Nama Partai : ")
		fmt.Scan(&p[Partai].Namap)
		fmt.Println("No urut partai : ")
		fmt.Scan(&urut)
		SearchUrut(urut,&found)
		if found == 0 {
			p[Partai].noUrutPartai = urut
			fmt.Println("No urut partai : ", p[Partai].noUrutPartai)
		} else if found == 1 {
			stop := false
			for found == 1 && !stop {
				SearchUrut(urut,&found)
				fmt.Print("Nomor Tersebut Sudah Ada, Masukan Lagi  : ")
				fmt.Scan(&urut)
				SearchUrut(urut,&found)
				if found == 0 {
				   p[Partai].noUrutPartai = urut
				   fmt.Println("No urut partai : ", p[Partai].noUrutPartai)
				   stop = true
				} 
			}
		}
		for j:=0; j<anggota; j++ {
			fmt.Print("Masukan Nama Anggota : ")
			fmt.Scan(&p[Partai].namaAnggota[j].namaCalon)
		}
		*jumPartai = *jumPartai+1
		Partai = Partai + 1
	}
	fmt.Println()
	main()
}

func selSort(jumPartai int, anggota int) {
	var maxIndex int
	var max int
	var sort[100]partai
	sort = p
	for i:=0; i<jumPartai-1; i++ {
		maxIndex = i 
		for j:=i+1; j<jumPartai; j++ {
			if sort[j].jumPemilihp > sort[maxIndex].jumPemilihp {
				maxIndex = j
			}
			for a:=0; a<anggota-1;a++ {
			max = a
				for b:=a+1; b<anggota; b++ {
					if sort[a].namaAnggota[b].jumPemilih > sort[a].namaAnggota[max].jumPemilih {
					max = b
					}
				}
				tmp := sort[a].namaAnggota[max]
				sort[a].namaAnggota[max] = sort[a].namaAnggota[a]
				sort[a].namaAnggota[a] = tmp
			}
		}

		temp := sort[maxIndex]
		sort[maxIndex] = sort[i]
		sort[i] = temp
	}
	for i:=0; i<jumPartai; i++ {
		fmt.Println("Nama Partai : ", sort[i].Namap)
		fmt.Println("Jumlah Suara : ",sort[i].jumPemilihp)
		for j:=0; j<anggota; j++ {
			fmt.Println("Nama Calon : ", sort[i].namaAnggota[j].namaCalon)
			fmt.Println("Jumlah Suara : ", sort[i].namaAnggota[j].jumPemilih)
		}
	}

	fmt.Println("Partai yang berada dibawah Parlemen Threshold : ")
	Threshold(jumTps, anggota, jumPartai, jumPemilihTps)
	fmt.Println()
	fmt.Println()
	main()
}
func SearchUrut(cari int, ketemu *int) {
	found:= false
	for i:= 0; i<jumPartai && !found 
	{
		if p[i].noUrutPartai == cari {
			*ketemu = 1
			found = true
		} else {
			*ketemu = 0
			i++ 
			}
		}
}

//menampi lkan daftar paprtai yang telah diinput
func daftarpartai(jumPartai,anggota int) {
	for i:=0; i<jumPartai; i++ {
		fmt.Print("Nama Partai : ")
		fmt.Println(p[i].Namap)
		fmt.Print("No urut : ")
		fmt.Println(p[i].noUrutPartai)
		fmt.Println("Nama Anggota : ")
		for j:=0; j<anggota; j++ {
			fmt.Println(p[i].namaAnggota[j].namaCalon)
		}
	} 
	fmt.Println()
	fmt.Println()
	main()
}

func suaraTPS(jumTps int) {
	var jum int
	for i:=0; i<jumTps; i++ {
		fmt.Printf("Jumlah Golput Pada TPS %v adalah %v \n", i+1, arrTps[i].jumGolput )
		jum = jum + arrTps[i].jumGolput
	}

	fmt.Println("Jumlah Total Golput : ", jum)
	fmt.Println()
	fmt.Println()

	main()
}




func DataTPS(jumTps int, jumPemilihTps int, Tps* int, PemilihTps *int) {
	fmt.Print("Masukan Jumlah TPS : ")
	fmt.Scan(&*Tps)
	fmt.Print("Masukan Jmmlah Pemilih : ")
	fmt.Scan(&*PemilihTps)
	for i:=0; i<*Tps; i++ {
		fmt.Println("Pemilih TPS ", i+1)
		for j:=0; j<*PemilihTps; j++ {
			fmt.Print("Masukan Nama Pemilih : ")
			fmt.Scan(&arrTps[i].namaPemilih[j])
		}
		
	}
	fmt.Println()
	fmt.Println()
	main()
}

func Threshold(jumTps int, jumPemilihp int, jumPartai int, jumPemilihTps int) {
	var ParThreshold float64
	ParThreshold = (float64(jumTps) * float64(jumPemilihTps) * 4)/100
	for i:=0; i<jumPartai; i++ {
		if float64(p[i].jumPemilihp) < ParThreshold {
			fmt.Print(p[i].Namap, " ")
		}
	}
}

func TambahTPS(jumTps int, jumPemilihTps int, n *int) {
	var x int
	fmt.Print("Masukan jumlah TPS yang ingin ditambah : ")
	fmt.Scan(&x)
	*n = x + jumTps
	for jumTps < *n {
		fmt.Println("Pemilih TPS ", jumTps)
		for j:= 0; j<jumPemilihTps; j++ {
			fmt.Print("Nama Pemilih : ")
			fmt.Scan(&arrTps[jumTps].namaPemilih[j])
		}
		jumTps = jumTps + 1 
	}
	fmt.Println()
	fmt.Println()
	main()
}







