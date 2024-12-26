package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type dataKuis struct {
	length int
	lastID int
	data   [99]kuis
}

type soals struct {
	idSoal  int
	soal    string
	jawaban string
}

type kuis struct {
	idKuis     int
	namaKuis   string
	kodeMatkul string
	tanggal    string
	banyakSoal int
	soal       [99]soals
	skor       int
}

type tugas struct {
	idTugas    int
	kodeMatkul string
	tanggal    string
	deadline   string
	isiTugas   string
}

type dataTugas struct {
	length int
	lastID int
	data   [99]tugas
}

type pengumpulan struct {
	idPengumpulan  int
	idTugas        int
	idKuis         int
	namaTugas      string
	idUser         string
	tanggal        string
	isiPengumpulan string
	nilai          int
}

type dataPengumpulan struct {
	length int
	lastID int
	data   [99]pengumpulan
}

type chat struct {
	idChat   int
	fromUser string
	chatData string
}

type forum struct {
	idForum    int
	kodeMatkul string
	judulTopik string
	lengthChat int
	chatData   [99]chat
}

type dataForum struct {
	length int
	lastID int
	data   [99]forum
}

type matkul struct {
	kodeMatkul string
	nama       string
}

type dataMatkul struct {
	length int
	data   [99]matkul
}

type User struct {
	idUser       string
	nama         string
	password     string
	lengthMatkul int
	role         string
	kodeMatkul   [99]string
}

type dataUser struct {
	length int
	data   [99]User
}

var dataKuiss = dataKuis{length: 0, lastID: 0, data: [99]kuis{}}
var dataTugass = dataTugas{length: 0, lastID: 0, data: [99]tugas{}}
var dataPengumpulans = dataPengumpulan{length: 0, lastID: 0, data: [99]pengumpulan{}}
var dataMatkuls = dataMatkul{length: 1, data: [99]matkul{{kodeMatkul: "IF1234", nama: "Pemrograman Dasar"}}}
var dataUsr dataUser = dataUser{length: 3, data: [99]User{
	{idUser: "1", nama: "Aldi", password: "123", lengthMatkul: 1, role: "Dosen", kodeMatkul: [99]string{"IF1234"}},
	{idUser: "2", nama: "Budi", password: "123", lengthMatkul: 0, role: "Mahasiswa"},
	{idUser: "3", nama: "Caca", password: "123", lengthMatkul: 0, role: "Mahasiswa"},
}}
var dataForums = dataForum{length: 0, data: [99]forum{}}
var datalogin User = User{}
var acessedMatkul matkul = matkul{}

func HandleLongInput(text *string) {
	reader := bufio.NewReader(os.Stdin)
	dataInput, _ := reader.ReadString('\n')
	*text = strings.TrimSpace(dataInput)
}

func Clrscr() {
	fmt.Print("\033[H\033[2J")
}

func border(typeBound string, text string, length int) string {
	var merger string
	if len(text) != 0 {
		text = " " + text + " "
	}
	length -= len(text)

	var leftBorder = length / 2
	var rightBorder = length - leftBorder

	for i := 0; i < leftBorder; i++ {
		merger += typeBound
	}
	merger += text
	for i := 0; i < rightBorder; i++ {
		merger += typeBound
	}
	return merger
}

func delay(sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
}

func login() User {
	var idUser, password string
	Clrscr()
	fmt.Println(border("=", "Login", 50))
	fmt.Print("ID User: ")
	HandleLongInput(&idUser)
	fmt.Print("Password: ")
	HandleLongInput(&password)
	for i := 0; i < dataUsr.length; i++ {
		if dataUsr.data[i].idUser == idUser && dataUsr.data[i].password == password {
			return dataUsr.data[i]
		}
	}
	return User{}

}

func getUserbyID(id string) User {
	for i := 0; i < dataUsr.length; i++ {
		if dataUsr.data[i].idUser == id {
			return dataUsr.data[i]
		}
	}
	return User{}
}

func registMahasiswa() {
	var idUser, nama, password string
	Clrscr()
	fmt.Println(border("=", "Registrasi Mahasiswa", 50))
	for {
		fmt.Print("ID User: ")
		fmt.Scanln(&idUser)
		duplicate := false
		for i := 0; i < dataUsr.length; i++ {
			if dataUsr.data[i].idUser == idUser {
				duplicate = true
				break
			}
		}
		if duplicate {
			fmt.Println("ID User sudah terdaftar, silakan coba lagi.")
		} else {
			break
		}
	}
	fmt.Print("Nama: ")
	HandleLongInput(&nama)
	fmt.Print("Password: ")
	HandleLongInput(&password)
	dataUsr.data[dataUsr.length].idUser = idUser
	dataUsr.data[dataUsr.length].nama = nama
	dataUsr.data[dataUsr.length].password = password
	dataUsr.data[dataUsr.length].role = "Mahasiswa"
	dataUsr.data[dataUsr.length].lengthMatkul = 0
	dataUsr.length++
}

func registDosen() {
	//regist dosen + regist matkul with check is matkul and dosen duplicate before create matkul and user dosen
	var idUser, nama, password string
	Clrscr()
	fmt.Println(border("=", "Registrasi Dosen", 50))
	for {
		fmt.Print("ID User: ")
		fmt.Scanln(&idUser)
		duplicate := false
		for i := 0; i < dataUsr.length; i++ {
			if dataUsr.data[i].idUser == idUser {
				duplicate = true
				break
			}
		}
		if duplicate {
			fmt.Println("ID User sudah terdaftar, silakan coba lagi.")
		} else {
			break
		}
	}
	fmt.Print("Nama: ")
	HandleLongInput(&nama)
	fmt.Print("Password: ")
	HandleLongInput(&password)
	dataUsr.data[dataUsr.length].idUser = idUser
	dataUsr.data[dataUsr.length].nama = nama
	dataUsr.data[dataUsr.length].password = password
	dataUsr.data[dataUsr.length].role = "Dosen"
	dataUsr.data[dataUsr.length].lengthMatkul = 0
	dataUsr.length++
}

func registMenu() {
	var choice int
	Clrscr()
	fmt.Println(border("=", "Register", 50))
	fmt.Println("1. Registrasi sebagai Mahasiswa")
	fmt.Println("2. Registrasi sebagai Dosen")
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		registMahasiswa()
	case 2:
		registDosen()
	default:
		fmt.Println("Pilihan tidak tersedia")
		delay(2)
		registMenu()
	}
}

func kuisTerdekat() {
	Clrscr()
	fmt.Println(border("=", "Kuis Terdekat", 50))
	for i := 0; i < dataKuiss.length; i++ {
		for j := 0; j < datalogin.lengthMatkul; j++ {
			if dataKuiss.data[i].kodeMatkul == datalogin.kodeMatkul[j] {
				alreadyTaken := false
				for k := 0; k < dataPengumpulans.length; k++ {
					if dataPengumpulans.data[k].idUser == datalogin.idUser && dataPengumpulans.data[k].idKuis == dataKuiss.data[i].idKuis {
						alreadyTaken = true
						break
					}
				}
				if !alreadyTaken {
					fmt.Println("ID Kuis: ", dataKuiss.data[i].idKuis)
					fmt.Println("Nama Kuis: ", dataKuiss.data[i].namaKuis)
					fmt.Println("Kode Matkul: ", dataKuiss.data[i].kodeMatkul)
					fmt.Println("Tanggal: ", dataKuiss.data[i].tanggal)
					fmt.Println("Banyak Soal: ", dataKuiss.data[i].banyakSoal)
					fmt.Println(border("=", "", 50))
				}
			}
		}
	}
	delay(3)
}

func tugasTerdekat() {
	Clrscr()
	fmt.Println(border("=", "Tugas Terdekat", 50))
	for i := 0; i < dataTugass.length; i++ {
		for j := 0; j < datalogin.lengthMatkul; j++ {
			if dataTugass.data[i].kodeMatkul == datalogin.kodeMatkul[j] {
				alreadyTaken := false
				for k := 0; k < dataPengumpulans.length; k++ {
					if dataPengumpulans.data[k].idUser == datalogin.idUser && dataPengumpulans.data[k].idTugas == dataTugass.data[i].idTugas {
						alreadyTaken = true
						break
					}
				}
				if !alreadyTaken {
					fmt.Println("ID Tugas: ", dataTugass.data[i].idTugas)
					fmt.Println("Kode Matkul: ", dataTugass.data[i].kodeMatkul)
					fmt.Println("Tanggal: ", dataTugass.data[i].tanggal)
					fmt.Println("Deadline: ", dataTugass.data[i].deadline)
					fmt.Println("Isi Tugas: ", dataTugass.data[i].isiTugas)
					fmt.Println(border("=", "", 50))
				}
			}
		}
	}
	delay(3)
}

func menuMahasiswa() {
	var choice int
	for {
		Clrscr()
		fmt.Println(border("=", "Halo "+datalogin.nama+", Semoga harimu menyenangkan!", 50))
		fmt.Println("1. Enroll Matkul")
		fmt.Println("2. Pilih Matkul")
		fmt.Println("3. Lihat Kuis terdekat")
		fmt.Println("4. Lihat Tugas terdekat")
		fmt.Println("5. Logout")
		fmt.Println(border("=", "", 50))
		fmt.Print("Pilihan: ")
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			enrollMatkul()
		case 2:
			selectManageMatkul()
		case 3:
			kuisTerdekat()
		case 4:
			tugasTerdekat()
		case 5:
			datalogin = User{}
			return
		default:
			fmt.Println("Pilihan tidak tersedia")
			delay(2)
		}
	}
}

func getMatkulbyKode(kode string) matkul {
	for i := 0; i < dataMatkuls.length; i++ {
		if dataMatkuls.data[i].kodeMatkul == kode {
			return dataMatkuls.data[i]
		}
	}
	return matkul{}
}

func manageMatkul() {
	var choice int
	for {
		Clrscr()
		fmt.Println(border("=", "Manage Matkul "+acessedMatkul.nama, 50))
		fmt.Println("1. Kerjakan Kuis")
		fmt.Println("2. Kerjakan Tugas")
		fmt.Println("3. Masuk Forum")
		fmt.Println("4. Lihat tugas yang sudah dan dinilai")
		fmt.Println("5. Kembali")
		fmt.Println(border("=", "", 50))
		fmt.Print("Pilihan: ")
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			kerjakanKuis()
		case 2:
			kerjakanTugas()
		case 3:
			masukForum()
		case 4:
			lihatTugasYangSudah()
		case 5:
			acessedMatkul = matkul{}
			return
		default:
			fmt.Println("Pilihan tidak tersedia")
			delay(2)
		}

	}
}

func lihatTugasYangSudah() {
	Clrscr()
	fmt.Println(border("=", "Tugas yang sudah dan dinilai", 50))
	for i := 0; i < dataPengumpulans.length; i++ {
		if dataPengumpulans.data[i].idUser == datalogin.idUser {
			//check if score not -1
			if dataPengumpulans.data[i].nilai != -1 {
				fmt.Println("ID Pengumpulan: ", dataPengumpulans.data[i].idPengumpulan)
				fmt.Println("Nama Tugas: ", dataPengumpulans.data[i].namaTugas)
				fmt.Println("Tanggal: ", dataPengumpulans.data[i].tanggal)
				fmt.Println("Isi Pengumpulan: ", dataPengumpulans.data[i].isiPengumpulan)
				fmt.Println("Nilai: ", dataPengumpulans.data[i].nilai)
				fmt.Println(border("=", "", 50))
			} else {
				fmt.Println("ID Pengumpulan: ", dataPengumpulans.data[i].idPengumpulan)
				fmt.Println("Nama Tugas: ", dataPengumpulans.data[i].namaTugas)
				fmt.Println("Tanggal: ", dataPengumpulans.data[i].tanggal)
				fmt.Println("Isi Pengumpulan: ", dataPengumpulans.data[i].isiPengumpulan)
				fmt.Println("Nilai: Belum dinilai")
				fmt.Println(border("=", "", 50))
			}
		}
	}
	delay(3)
}

func masukForum() {
	//make mahasiswa can select topic in forum to chat in list forum
	Clrscr()
	var arrChoice [99]int
	var count int
	fmt.Println(border("=", "Masuk Forum", 50))
	for i := 0; i < dataForums.length; i++ {
		if dataForums.data[i].kodeMatkul == acessedMatkul.kodeMatkul {
			arrChoice[count] = i
			count++
			fmt.Println(count, " Topik: ", dataForums.data[i].judulTopik)
		}
	}
	fmt.Println("0. Kembali")
	fmt.Println(border("=", "", 50))
	fmt.Print("Pilihan: ")
	var choice int
	fmt.Scanln(&choice)
	if choice == 0 {
		return
	}
	chatInForum(arrChoice[choice-1])
}

func chatInForum(forumIndex int) {
	for {
		Clrscr()
		fmt.Println(border("=", "Masuk Forum", 50))
		fmt.Println("Topik: ", dataForums.data[forumIndex].judulTopik)
		for i := 0; i < dataForums.data[forumIndex].lengthChat; i++ {
			fmt.Println(dataForums.data[forumIndex].chatData[i].fromUser, ": ", dataForums.data[forumIndex].chatData[i].chatData)
		}
		fmt.Println(border("=", "", 50))
		fmt.Print("Chat (-1 untuk kembali): ")
		var chat string
		HandleLongInput(&chat)
		if chat == "-1" {
			return
		}
		dataForums.data[forumIndex].chatData[dataForums.data[forumIndex].lengthChat].fromUser = datalogin.nama
		dataForums.data[forumIndex].chatData[dataForums.data[forumIndex].lengthChat].chatData = chat
		dataForums.data[forumIndex].chatData[dataForums.data[forumIndex].lengthChat].idChat = dataForums.data[forumIndex].lengthChat
		dataForums.data[forumIndex].lengthChat++
	}
}

func kerjakanTugas() {
	Clrscr()
	var arrChoice [99]int
	var count int
	fmt.Println(border("=", "Kerjakan Tugas", 50))
	for i := 0; i < dataTugass.length; i++ {
		if dataTugass.data[i].kodeMatkul == acessedMatkul.kodeMatkul {
			arrChoice[count] = i
			count++
			fmt.Println(count, " Tanggal: ", dataTugass.data[i].tanggal)
		}
	}
	fmt.Println("0. Kembali")
	fmt.Println(border("=", "", 50))
	fmt.Print("Pilihan: ")
	var choice int
	fmt.Scanln(&choice)
	if choice == 0 {
		return
	}
	Clrscr()
	fmt.Println(border("=", "Kerjakan Tugas", 50))
	fmt.Println("Tanggal: ", dataTugass.data[arrChoice[choice-1]].tanggal)
	fmt.Println("Deadline: ", dataTugass.data[arrChoice[choice-1]].deadline)
	fmt.Println("Detail Tugas: ", dataTugass.data[arrChoice[choice-1]].isiTugas)
	fmt.Println(border("=", "", 50))
	fmt.Println("Pengumpulan: ")
	var jawaban string
	fmt.Print("Jawaban: ")
	HandleLongInput(&jawaban)
	dataPengumpulans.lastID++
	dataPengumpulans.data[dataPengumpulans.length] = pengumpulan{
		idPengumpulan:  dataPengumpulans.lastID,
		idTugas:        dataTugass.data[arrChoice[choice-1]].idTugas,
		namaTugas:      dataTugass.data[arrChoice[choice-1]].isiTugas,
		idUser:         datalogin.idUser,
		tanggal:        time.Now().Format("2006-01-02"),
		isiPengumpulan: jawaban,
		nilai:          -1,
	}
	dataPengumpulans.length++
	delay(3)
}

func kerjakanKuis() {
	//make mahasiswa can select kuis to do
	Clrscr()
	var arrChoice [99]int
	var count int
	fmt.Println(border("=", "Kerjakan Kuis", 50))
	for i := 0; i < dataKuiss.length; i++ {
		if dataKuiss.data[i].kodeMatkul == acessedMatkul.kodeMatkul {
			arrChoice[count] = i
			count++
			fmt.Println(count, "Nama Kuis : ", dataKuiss.data[i].namaKuis)
		}
	}
	fmt.Println("0. Kembali")
	fmt.Println(border("=", "", 50))
	fmt.Print("Pilihan: ")
	var choice int
	fmt.Scanln(&choice)
	if choice == 0 {
		return
	}
	Clrscr()
	fmt.Println(border("=", "Kerjakan Kuis", 50))
	fmt.Println("Tanggal: ", dataKuiss.data[arrChoice[choice-1]].tanggal)
	var correctAnswers int
	for i := 0; i < dataKuiss.data[arrChoice[choice-1]].banyakSoal; i++ {
		var jawaban string
		fmt.Println("Soal ", i+1, ": ", dataKuiss.data[arrChoice[choice-1]].soal[i].soal)
		fmt.Print("Jawaban: ")
		HandleLongInput(&jawaban)
		if jawaban == dataKuiss.data[arrChoice[choice-1]].soal[i].jawaban {
			correctAnswers++
		}
	}
	dataKuiss.data[arrChoice[choice-1]].skor = correctAnswers
	fmt.Println("Skor: ", correctAnswers, "/", dataKuiss.data[arrChoice[choice-1]].banyakSoal)
	fmt.Println(border("=", "", 50))
	dataPengumpulans.lastID++
	dataPengumpulans.data[dataPengumpulans.length] = pengumpulan{
		idPengumpulan:  dataPengumpulans.lastID,
		idTugas:        -1,
		idKuis:         dataKuiss.data[arrChoice[choice-1]].idKuis,
		namaTugas:      dataKuiss.data[arrChoice[choice-1]].namaKuis,
		idUser:         datalogin.idUser,
		tanggal:        time.Now().Format("2006-01-02"),
		isiPengumpulan: fmt.Sprintf("%d/%d", correctAnswers, dataKuiss.data[arrChoice[choice-1]].banyakSoal),
		nilai:          correctAnswers,
	}
	dataPengumpulans.length++
	delay(3)
}

func selectManageMatkul() {
	var choice int
	Clrscr()
	fmt.Println(border("=", "Pilih Matkul", 50))
	for i := 0; i < datalogin.lengthMatkul; i++ {
		fmt.Println(i+1, ". ", getMatkulbyKode(datalogin.kodeMatkul[i]).nama)
	}
	fmt.Println("0. Kembali")
	fmt.Println(border("=", "", 50))
	fmt.Print("Pilihan: ")
	fmt.Scanln(&choice)
	if choice == 0 {
		return
	} else if choice > datalogin.lengthMatkul {
		fmt.Println("Pilihan tidak tersedia")
		delay(2)
		selectManageMatkul()
	} else {
		acessedMatkul = getMatkulbyKode(datalogin.kodeMatkul[choice-1])
		manageMatkul()
	}
}

func enrollMatkul() {
	//enroll matkul berdasar kode matkul. mahasiswa menginput kode matkuls
	var kodeMatkul string
	Clrscr()
	fmt.Println(border("=", "Enroll Matkul", 50))
	fmt.Print("Masukkan kode matkul: ")
	HandleLongInput(&kodeMatkul)
	for i := 0; i < dataMatkuls.length; i++ {
		if dataMatkuls.data[i].kodeMatkul == kodeMatkul {
			for j := 0; j < datalogin.lengthMatkul; j++ {
				if datalogin.kodeMatkul[j] == kodeMatkul {
					fmt.Println("Anda sudah terdaftar di matkul ini")
					delay(2)
					return
				}
			}
			for j := 0; j < dataUsr.length; j++ {
				if dataUsr.data[j].idUser == datalogin.idUser {
					dataUsr.data[j].kodeMatkul[datalogin.lengthMatkul] = kodeMatkul
					dataUsr.data[j].lengthMatkul++
					datalogin.kodeMatkul[datalogin.lengthMatkul] = kodeMatkul
					datalogin.lengthMatkul++
					fmt.Println("Berhasil terdaftar di matkul ", dataMatkuls.data[i].nama)
					delay(2)
					return
				}
			}
		}
	}
	fmt.Println("Kode matkul tidak ditemukan")
	delay(2)
}

func tambahKuis() {
	var tanggal, nama string
	var banyakSoal int
	Clrscr()
	fmt.Println(border("=", "Tambah Kuis", 50))
	fmt.Print("Nama Kuis: ")
	HandleLongInput(&nama)
	fmt.Print("Tanggal: ")
	HandleLongInput(&tanggal)
	fmt.Print("Banyak Soal: ")
	fmt.Scanln(&banyakSoal)
	dataKuiss.lastID++
	dataKuiss.data[dataKuiss.length].idKuis = dataKuiss.lastID
	dataKuiss.data[dataKuiss.length].namaKuis = nama
	dataKuiss.data[dataKuiss.length].kodeMatkul = datalogin.kodeMatkul[0] //dosen cuma bisa 1 matkul
	dataKuiss.data[dataKuiss.length].tanggal = tanggal
	dataKuiss.data[dataKuiss.length].banyakSoal = banyakSoal
	for i := 0; i < banyakSoal; i++ {
		Clrscr()
		var soaln, jawaban string
		fmt.Println(border("=", "Soal "+fmt.Sprint(i+1), 50))
		fmt.Print("Soal: ")
		HandleLongInput(&soaln)
		fmt.Print("Jawaban : ")
		HandleLongInput(&jawaban)
		dataKuiss.data[dataKuiss.length].soal[i].idSoal = i
		dataKuiss.data[dataKuiss.length].soal[i].soal = soaln
		dataKuiss.data[dataKuiss.length].soal[i].jawaban = jawaban
		fmt.Println(border("=", "", 50))
	}
	dataKuiss.length++
}

func tambahTugas() {
	var tanggal, deadline, isiTugas string
	Clrscr()
	fmt.Println(border("=", "Tambah Tugas", 50))
	fmt.Print("Tanggal: ")
	HandleLongInput(&tanggal)
	fmt.Print("Deadline: ")
	HandleLongInput(&deadline)
	fmt.Print("Isi Tugas: ")
	HandleLongInput(&isiTugas)
	dataTugass.lastID++
	dataTugass.data[dataTugass.length].idTugas = dataTugass.lastID
	dataTugass.data[dataTugass.length].kodeMatkul = datalogin.kodeMatkul[0] //dosen cuma bisa 1 matkul
	dataTugass.data[dataTugass.length].tanggal = tanggal
	dataTugass.data[dataTugass.length].deadline = deadline
	dataTugass.data[dataTugass.length].isiTugas = isiTugas
	dataTugass.length++
}

func tambahForum() {
	var judulTopik string
	Clrscr()
	fmt.Println(border("=", "Tambah Forum", 50))
	fmt.Print("Judul Topik: ")
	HandleLongInput(&judulTopik)
	dataForums.lastID++
	dataForums.data[dataForums.length].idForum = dataForums.lastID
	dataForums.data[dataForums.length].kodeMatkul = datalogin.kodeMatkul[0] //dosen cuma bisa 1 matkul
	dataForums.data[dataForums.length].judulTopik = judulTopik
	dataForums.length++
}

func tambahKonten() {
	var choice int
	Clrscr()
	fmt.Println(border("=", "Tambah Konten", 50))
	fmt.Println("1. Kuis")
	fmt.Println("2. Tugas")
	fmt.Println("3. Forum")
	fmt.Println("4. Kembali")
	fmt.Println(border("=", "", 50))
	fmt.Print("Pilihan: ")
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		tambahKuis()
	case 2:
		tambahTugas()
	case 3:
		tambahForum()
	case 4:
		return
	default:
		fmt.Println("Pilihan tidak tersedia")
		delay(2)
		tambahKonten()
	}
}

func lihatKuis() {
	Clrscr()
	fmt.Println(border("=", "Lihat Kuis", 50))
	for i := 0; i < dataKuiss.length; i++ {
		if dataKuiss.data[i].kodeMatkul == datalogin.kodeMatkul[0] {
			fmt.Println("ID Kuis: ", dataKuiss.data[i].idKuis)
			fmt.Println("Nama Kuis: ", dataKuiss.data[i].namaKuis)
			fmt.Println("Kode Matkul: ", dataKuiss.data[i].kodeMatkul)
			fmt.Println("Tanggal: ", dataKuiss.data[i].tanggal)
			fmt.Println("Banyak Soal: ", dataKuiss.data[i].banyakSoal)
			for j := 0; j < dataKuiss.data[i].banyakSoal; j++ {
				fmt.Println("Soal ", j+1, ": ", dataKuiss.data[i].soal[j].soal, "(", dataKuiss.data[i].soal[j].jawaban, ")")
			}
			fmt.Println(border("=", "", 50))
		}
	}
	delay(3)
}

func lihatTugas() {
	Clrscr()
	fmt.Println(border("=", "Lihat Tugas", 50))
	for i := 0; i < dataTugass.length; i++ {
		if dataTugass.data[i].kodeMatkul == datalogin.kodeMatkul[0] {
			fmt.Println("ID Tugas: ", dataTugass.data[i].idTugas)
			fmt.Println("Kode Matkul: ", dataTugass.data[i].kodeMatkul)
			fmt.Println("Tanggal: ", dataTugass.data[i].tanggal)
			fmt.Println("Deadline: ", dataTugass.data[i].deadline)
			fmt.Println("Isi Tugas: ", dataTugass.data[i].isiTugas)
			fmt.Println(border("=", "", 50))
		}
	}
	delay(3)
}

func lihatForum() {
	//make dosen can select topic in forum to look chat in list forum
	Clrscr()
	var arrChoice [99]int
	var count int
	fmt.Println(border("=", "Lihat Forum", 50))
	for i := 0; i < dataForums.length; i++ {
		if dataForums.data[i].kodeMatkul == datalogin.kodeMatkul[0] {
			arrChoice[count] = i
			count++
			fmt.Println(count, " Topik: ", dataForums.data[i].judulTopik)
		}
	}
	fmt.Println("0. Kembali")
	fmt.Println(border("=", "", 50))
	fmt.Print("Pilihan: ")
	var choice int
	fmt.Scanln(&choice)
	if choice == 0 {
		return
	}
	Clrscr()
	fmt.Println(border("=", "Lihat Forum", 50))
	fmt.Println("Topik: ", dataForums.data[arrChoice[choice-1]].judulTopik)
	for i := 0; i < dataForums.data[arrChoice[choice-1]].lengthChat; i++ {
		fmt.Println(dataForums.data[arrChoice[choice-1]].chatData[i].fromUser, ": ", dataForums.data[arrChoice[choice-1]].chatData[i].chatData)
	}
	fmt.Println(border("=", "", 50))
	delay(3)
}

func lihatKonten() {
	var choice int
	Clrscr()
	fmt.Println(border("=", "Lihat Konten", 50))
	fmt.Println("1. Kuis")
	fmt.Println("2. Tugas")
	fmt.Println("3. Forum")
	fmt.Println("4. Kembali")
	fmt.Println(border("=", "", 50))
	fmt.Print("Pilihan: ")
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		lihatKuis()
	case 2:
		lihatTugas()
	case 3:
		masukForum()
	case 4:
		return
	default:
		fmt.Println("Pilihan tidak tersedia")
		delay(2)
		lihatKonten()
	}
}

func hapusKonten() {
	var choice int
	Clrscr()
	fmt.Println(border("=", "Hapus Konten", 50))
	fmt.Println("1. Kuis")
	fmt.Println("2. Tugas")
	fmt.Println("3. Forum")
	fmt.Println("4. Kembali")
	fmt.Println(border("=", "", 50))
	fmt.Print("Pilihan: ")
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		lihatKuis()
		fmt.Print("Masukkan ID Kuis: ")
		var idKuis int
		fmt.Scanln(&idKuis)
		for i := 0; i < dataKuiss.length; i++ {
			if dataKuiss.data[i].idKuis == idKuis {
				dataKuiss.data[i] = dataKuiss.data[dataKuiss.length-1]
				dataKuiss.length--
				break
			}
		}
	case 2:
		lihatTugas()
		fmt.Print("Masukkan ID Tugas: ")
		var idTugas int
		fmt.Scanln(&idTugas)
		for i := 0; i < dataTugass.length; i++ {
			if dataTugass.data[i].idTugas == idTugas {
				dataTugass.data[i] = dataTugass.data[dataTugass.length-1]
				dataTugass.length--
				break
			}
		}
	case 3:
		lihatForum()
		fmt.Print("Masukkan ID Forum: ")
		var idForum int
		fmt.Scanln(&idForum)
		for i := 0; i < dataForums.length; i++ {
			if dataForums.data[i].idForum == idForum {
				dataForums.data[i] = dataForums.data[dataForums.length-1]
				dataForums.length--
				break
			}
		}
	case 4:
		return
	default:
		fmt.Println("Pilihan tidak tersedia")
		delay(2)
		hapusKonten()
	}
}

func GiveNilaiMahasiswa() {
	//make dosen can select tugas to give nilai
	Clrscr()
	var arrChoice [99]int
	var count int
	fmt.Println(border("=", "Nilai Mahasiswa", 50))
	for i := 0; i < dataTugass.length; i++ {
		if dataTugass.data[i].kodeMatkul == datalogin.kodeMatkul[0] {
			arrChoice[count] = i
			count++
			fmt.Println(count, " Tugas: ", dataTugass.data[i].isiTugas)
		}
	}
	fmt.Println("0. Kembali")
	fmt.Println(border("=", "", 50))
	fmt.Print("Pilihan: ")
	var choice int
	fmt.Scanln(&choice)
	if choice == 0 {
		return
	}
	Clrscr()
	fmt.Println(border("=", "Nilai Mahasiswa", 50))
	fmt.Println("Tugas: ", dataTugass.data[arrChoice[choice-1]].isiTugas)
	for i := 0; i < dataPengumpulans.length; i++ {
		if dataPengumpulans.data[i].idTugas == dataTugass.data[arrChoice[choice-1]].idTugas && dataPengumpulans.data[i].nilai == -1 && dataPengumpulans.data[i].idTugas != -1 {
			fmt.Println("Nama Mahasiswa: ", getUserbyID(dataPengumpulans.data[i].idUser).nama)
			fmt.Println("Isi Pengumpulan: ", dataPengumpulans.data[i].isiPengumpulan)
			fmt.Print("Nilai: ")
			fmt.Scanln(&dataPengumpulans.data[i].nilai)
			fmt.Println("Nilai berhasil disimpan")
		}
	}
	fmt.Println(border("=", "", 50))
	delay(2)
}

func nilaiKuis() {
	Clrscr()
	///make dosen can select which kuis and show nilai kuis dari datapengumpulan
	var arrChoice [99]int
	var count int
	fmt.Println(border("=", "Nilai Kuis", 50))
	for i := 0; i < dataKuiss.length; i++ {
		if dataKuiss.data[i].kodeMatkul == datalogin.kodeMatkul[0] {
			arrChoice[count] = i
			count++
			fmt.Println(count, " Kuis: ", dataKuiss.data[i].namaKuis)
		}
	}
	fmt.Println("0. Kembali")
	fmt.Println(border("=", "", 50))
	fmt.Print("Pilihan: ")
	var choice int
	fmt.Scanln(&choice)
	if choice == 0 {
		return
	}
	Clrscr()
	fmt.Println(border("=", "Nilai Kuis", 50))
	fmt.Println("Kuis: ", dataKuiss.data[arrChoice[choice-1]].namaKuis)
	for i := 0; i < dataPengumpulans.length; i++ {
		if dataPengumpulans.data[i].idKuis == dataKuiss.data[arrChoice[choice-1]].idKuis {
			fmt.Println("Nama Mahasiswa: ", getUserbyID(dataPengumpulans.data[i].idUser).nama)
			fmt.Println("Benar: ", dataPengumpulans.data[i].isiPengumpulan)
			fmt.Println("Nilai: ", dataPengumpulans.data[i].nilai)
		}
	}
	fmt.Println(border("=", "", 50))
	delay(3)
}

func nilaiTugas() {
	//make dosen can select which tugas and show nilai tugas dari datapengumpulan
	Clrscr()
	var arrChoice [99]int
	var count int
	fmt.Println(border("=", "Nilai Tugas", 50))
	for i := 0; i < dataTugass.length; i++ {
		if dataTugass.data[i].kodeMatkul == datalogin.kodeMatkul[0] {
			arrChoice[count] = i
			count++
			fmt.Println(count, " Tugas: ", dataTugass.data[i].isiTugas)
		}
	}
	fmt.Println("0. Kembali")
	fmt.Println(border("=", "", 50))
	fmt.Print("Pilihan: ")
	var choice int
	fmt.Scanln(&choice)
	if choice == 0 {
		return
	}
	Clrscr()
	fmt.Println(border("=", "Nilai Tugas", 50))
	fmt.Println("Tugas: ", dataTugass.data[arrChoice[choice-1]].isiTugas)
	for i := 0; i < dataPengumpulans.length; i++ {
		if dataPengumpulans.data[i].idTugas == dataTugass.data[arrChoice[choice-1]].idTugas {
			fmt.Println("Nama Mahasiswa: ", getUserbyID(dataPengumpulans.data[i].idUser).nama)
			fmt.Println("Isi Pengumpulan: ", dataPengumpulans.data[i].isiPengumpulan)
			fmt.Println("Nilai: ", dataPengumpulans.data[i].nilai)
		}
	}
	fmt.Println(border("=", "", 50))
	delay(3)
}

func lihatNilai() {
	//make lecture can see nilai kuis and tugas
	Clrscr()
	fmt.Println(border("=", "Lihat Nilai", 50))
	fmt.Println("1. Kuis")
	fmt.Println("2. Tugas")
	fmt.Println("3. Kembali")
	fmt.Println(border("=", "", 50))
	fmt.Print("Pilihan: ")
	var choice int
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		nilaiKuis()
	case 2:
		nilaiTugas()
	case 3:
		return
	default:
		fmt.Println("Pilihan tidak tersedia")
		delay(2)
		lihatNilai()
	}

}

func addNewMatkul() {
	//add new matkul and make sure matkul is unique retry if not uniq and add to datauser dosen
	var kodeMatkul, namaMatkul string
	for {
		Clrscr()
		fmt.Println(border("=", "Tambah Matkul", 50))
		fmt.Print("Kode Matkul: ")
		HandleLongInput(&kodeMatkul)
		duplicate := false
		for i := 0; i < dataMatkuls.length; i++ {
			if dataMatkuls.data[i].kodeMatkul == kodeMatkul {
				duplicate = true
				break
			}
		}
		if duplicate {
			fmt.Println("Kode Matkul sudah terdaftar, silakan coba lagi.")
		} else {
			break
		}
	}
	fmt.Print("Nama Matkul: ")
	HandleLongInput(&namaMatkul)
	dataMatkuls.data[dataMatkuls.length].kodeMatkul = kodeMatkul
	dataMatkuls.data[dataMatkuls.length].nama = namaMatkul
	dataMatkuls.length++
	for i := 0; i < dataUsr.length; i++ {
		if dataUsr.data[i].idUser == datalogin.idUser {
			dataUsr.data[i].kodeMatkul[datalogin.lengthMatkul] = kodeMatkul
			dataUsr.data[i].lengthMatkul++
			datalogin.kodeMatkul[datalogin.lengthMatkul] = kodeMatkul
			datalogin.lengthMatkul++
			break
		}
	}
	fmt.Println("Berhasil menambahkan matkul")
	delay(2)
}

func menuDosen() {
	if datalogin.lengthMatkul == 0 {
		addNewMatkul()
	}
	var choice int
	for {
		Clrscr()
		fmt.Println(border("=", "Halo "+datalogin.nama+", Semoga harimu menyenangkan!", 50))
		fmt.Println("1. Tambah Konten")
		fmt.Println("2. Lihat Konten")
		fmt.Println("3. Ubah Konten")
		fmt.Println("4. Hapus Konten")
		fmt.Println("5. Nilai mahasiswa")
		fmt.Println("6. Lihat Nilai")
		fmt.Println("7. Logout")
		fmt.Println(border("=", "", 50))
		fmt.Print("Pilihan: ")
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			tambahKonten()
		case 2:
			lihatKonten()
		case 3:
			//ubahKonten()
		case 4:
			hapusKonten()
		case 5:
			GiveNilaiMahasiswa()
		case 6:
			lihatNilai()
		case 7:
			datalogin = User{}
			acessedMatkul = matkul{}
			return
		default:
			fmt.Println("Pilihan tidak tersedia")
			delay(2)
		}
	}
}

func main() {
	var choice int
	for {
		Clrscr()
		fmt.Println(border("=", "Menu LMS", 50))
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. Exit")
		fmt.Println(border("=", "", 50))
		fmt.Print("Pilihan: ")
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			var checkd = login()
			if checkd.idUser != "" {
				fmt.Println("Login Berhasil")
				datalogin = checkd
				if datalogin.role == "Dosen" {
					acessedMatkul = getMatkulbyKode(datalogin.kodeMatkul[0])
					menuDosen()
				} else {
					menuMahasiswa()
				}
				delay(2)
			} else {
				fmt.Println("Login Gagal")
				delay(2)
			}
		case 2:
			registMenu()
		case 3:
			fmt.Println("Terima kasih telah menggunakan aplikasi ini")
			os.Exit(0)
		default:
			fmt.Println("Pilihan tidak tersedia")
			delay(2)
		}
	}
}
