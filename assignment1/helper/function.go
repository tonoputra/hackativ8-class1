package helper


func GetTeman(noAbsen int) (callteman Teman) {

	var TemanList [7]Teman
	TemanList[0] = Teman{Nama: "Fiqri", Alamat: "Bandung", Pekerjaan: "Atlet", AlasanBergabung: "Iseng"}
	TemanList[1] = Teman{Nama: "Clara", Alamat: "Bekasi", Pekerjaan: "Pedagang Nasi Uduk", AlasanBergabung: "Belajar"}
	TemanList[2] = Teman{Nama: "Rijal", Alamat: "Sulawesi", Pekerjaan: "Juru Parkir", AlasanBergabung: "Bermain"}
	TemanList[3] = Teman{Nama: "Ivan", Alamat: "Bintaro", Pekerjaan: "Pemancing", AlasanBergabung: "Tantangan"}
	TemanList[4] = Teman{Nama: "Aditya", Alamat: "Malang", Pekerjaan: "Pendaki", AlasanBergabung: "Iseng Aja"}
	TemanList[5] = Teman{Nama: "Tono", Alamat: "Papua", Pekerjaan: "Satuan Pengaman", AlasanBergabung: "Ga ada kerjaan"}

	return TemanList[noAbsen]
}
