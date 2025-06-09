// main.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort" // Dipertahankan untuk digunakan oleh Binary Search
	"strconv"
	"strings"
)

// Struct adalah blueprint untuk data kita.
type Sampah struct {
	ID              int
	Jenis           string
	JumlahKg        float64
	StatusDaurUlang string
	MetodeDaurUlang string
}

var dataSampah []Sampah
var idTerakhir int

// Mengisi data awal saat program dimulai.
func init() {
	dataSampah = []Sampah{
		{ID: 1, Jenis: "Botol Plastik", JumlahKg: 25.5, StatusDaurUlang: "Sudah", MetodeDaurUlang: "Dilebur menjadi biji plastik"},
		{ID: 2, Jenis: "Kertas Karton", JumlahKg: 50.0, StatusDaurUlang: "Sudah", MetodeDaurUlang: "Diolah menjadi bubur kertas"},
		{ID: 3, Jenis: "Kaleng Aluminium", JumlahKg: 15.2, StatusDaurUlang: "Belum", MetodeDaurUlang: "-"},
		{ID: 4, Jenis: "Sampah Organik", JumlahKg: 75.8, StatusDaurUlang: "Sudah", MetodeDaurUlang: "Dijadikan kompos"},
		{ID: 5, Jenis: "Kaca", JumlahKg: 10.0, StatusDaurUlang: "Belum", MetodeDaurUlang: "-"},
	}
	idTerakhir = 5
}

// --- FUNGSI UTAMA ---
func main() {
	for {
		fmt.Println("\n=====================================================")
		fmt.Println("   Aplikasi Pengelolaan Data Sampah & Daur Ulang")
		fmt.Println("       (Sesuai Spesifikasi Algoritma)")
		fmt.Println("=====================================================")
		fmt.Println("1. Tampilkan Semua Data")
		fmt.Println("2. Tambah Data")
		fmt.Println("3. Ubah Data")
		fmt.Println("4. Hapus Data")
		fmt.Println("5. Urutkan Data (Selection / Insertion Sort)")
		fmt.Println("6. Cari Data (Sequential / Binary Search)")
		fmt.Println("7. Lihat Statistik")
		fmt.Println("8. Keluar")
		fmt.Println("-----------------------------------------------------")

		pilihan := bacaInput("Masukkan pilihan Anda (1-8): ")

		switch pilihan {
		case "1":
			tampilkanData(dataSampah)
		case "2":
			tambahData()
		case "3":
			ubahData()
		case "4":
			hapusData()
		case "5":
			menuUrutkan() // Diubah ke menu baru
		case "6":
			menuCari() // Diubah ke menu baru
		case "7":
			tampilkanStatistik()
		case "8":
			fmt.Println("\nTerima kasih! Sampai jumpa.")
			return
		default:
			fmt.Println("\nPilihan tidak valid.")
		}
		bacaInput("\nTekan Enter untuk melanjutkan...")
	}
}

// --- FUNGSI PENGURUTAN (SESUAI SPESIFIKASI) ---

// Menampilkan menu untuk memilih algoritma pengurutan.
func menuUrutkan() {
	fmt.Println("\n--- Menu Urutkan Data ---")
	fmt.Println("Pilih Kriteria Pengurutan:")
	fmt.Println("1. Berdasarkan Jenis (A-Z)")
	fmt.Println("2. Berdasarkan Jumlah (Terkecil ke Terbesar)")
	kriteria := bacaInput("Pilihan Anda (1-2): ")

	if kriteria != "1" && kriteria != "2" {
		fmt.Println("Pilihan kriteria tidak valid.")
		return
	}

	fmt.Println("\nPilih Algoritma Pengurutan:")
	fmt.Println("1. Selection Sort")
	fmt.Println("2. Insertion Sort")
	algo := bacaInput("Pilihan Anda (1-2): ")

	berdasarkan := "jumlah"
	if kriteria == "1" {
		berdasarkan = "jenis"
	}

	switch algo {
	case "1":
		fmt.Println("\nMengurutkan menggunakan Selection Sort...")
		selectionSort(berdasarkan)
		fmt.Println(">> Data berhasil diurutkan.")
	case "2":
		fmt.Println("\nMengurutkan menggunakan Insertion Sort...")
		insertionSort(berdasarkan)
		fmt.Println(">> Data berhasil diurutkan.")
	default:
		fmt.Println("Pilihan algoritma tidak valid.")
		return
	}
	tampilkanData(dataSampah)
}

// Algoritma Selection Sort.
// Konsep: Cari elemen terkecil di sisa array, lalu tukar dengan elemen saat ini.
func selectionSort(berdasarkan string) {
	n := len(dataSampah)
	for i := 0; i < n-1; i++ {
		// Asumsikan elemen terkecil ada di posisi saat ini (i)
		indeksTerkecil := i
		for j := i + 1; j < n; j++ {
			// Bandingkan elemen j dengan elemen terkecil yang sudah ditemukan
			harusTukar := false
			if berdasarkan == "jenis" {
				if dataSampah[j].Jenis < dataSampah[indeksTerkecil].Jenis {
					harusTukar = true
				}
			} else { // berdasarkan == "jumlah"
				if dataSampah[j].JumlahKg < dataSampah[indeksTerkecil].JumlahKg {
					harusTukar = true
				}
			}

			// Jika ditemukan elemen yang lebih kecil, update indeksTerkecil
			if harusTukar {
				indeksTerkecil = j
			}
		}
		// Tukar elemen di posisi i dengan elemen terkecil yang ditemukan
		dataSampah[i], dataSampah[indeksTerkecil] = dataSampah[indeksTerkecil], dataSampah[i]
	}
}

// Algoritma Insertion Sort.
// Konsep: Ambil satu elemen, lalu sisipkan ke posisi yang benar di bagian array yang sudah terurut.
func insertionSort(berdasarkan string) {
	n := len(dataSampah)
	for i := 1; i < n; i++ {
		// Ambil elemen yang akan disisipkan
		kunci := dataSampah[i]
		j := i - 1

		// Pindahkan elemen-elemen yang lebih besar dari 'kunci' ke kanan
		for j >= 0 {
			harusGeser := false
			if berdasarkan == "jenis" {
				if dataSampah[j].Jenis > kunci.Jenis {
					harusGeser = true
				}
			} else { // berdasarkan == "jumlah"
				if dataSampah[j].JumlahKg > kunci.JumlahKg {
					harusGeser = true
				}
			}

			if !harusGeser {
				break // Berhenti jika posisi yang benar sudah ditemukan
			}

			dataSampah[j+1] = dataSampah[j]
			j = j - 1
		}
		// Tempatkan 'kunci' di posisi yang benar
		dataSampah[j+1] = kunci
	}
}

// --- FUNGSI PENCARIAN (SESUAI SPESIFIKASI) ---

// Menampilkan menu untuk memilih algoritma pencarian.
func menuCari() {
	fmt.Println("\n--- Menu Cari Data ---")
	fmt.Println("Pilih Algoritma Pencarian:")
	fmt.Println("1. Sequential Search (Pencarian Berurutan)")
	fmt.Println("2. Binary Search (Hanya untuk jenis, data harus terurut)")
	pilihan := bacaInput("Pilihan Anda (1-2): ")

	switch pilihan {
	case "1":
		sequentialSearch()
	case "2":
		binarySearch()
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}

// Algoritma Sequential Search.
// Konsep: Periksa setiap elemen satu per satu dari awal sampai akhir.
func sequentialSearch() {
	fmt.Println("\n--- Pencarian Berurutan (Sequential Search) ---")
	target := bacaInput("Masukkan jenis sampah yang dicari: ")
	var hasil []Sampah

	for _, s := range dataSampah {
		// strings.Contains mencari apakah 'target' ada di dalam 's.Jenis'
		if strings.Contains(strings.ToLower(s.Jenis), strings.ToLower(target)) {
			hasil = append(hasil, s)
		}
	}

	if len(hasil) > 0 {
		fmt.Printf("\n>> Ditemukan %d data yang cocok untuk '%s':\n", len(hasil), target)
		tampilkanData(hasil)
	} else {
		fmt.Printf("\n>> Data dengan jenis '%s' tidak ditemukan.\n", target)
	}
}

// Algoritma Binary Search.
// Konsep: Hanya bekerja pada data terurut. Membagi data menjadi dua bagian berulang kali.
func binarySearch() {
	fmt.Println("\n--- Pencarian Biner (Binary Search) ---")
	target := bacaInput("Masukkan jenis sampah yang dicari (harus sama persis): ")

	// PENTING: Binary Search memerlukan data yang sudah terurut.
	// Kita buat salinan data agar tidak mengubah urutan data asli.
	dataUrut := make([]Sampah, len(dataSampah))
	copy(dataUrut, dataSampah)

	// Urutkan salinan data berdasarkan jenis. Boleh pakai sort.Slice karena ini hanya persiapan.
	sort.Slice(dataUrut, func(i, j int) bool {
		return dataUrut[i].Jenis < dataUrut[j].Jenis
	})

	fmt.Println("(Info: Data diurutkan sementara untuk keperluan pencarian...)")

	var hasil []Sampah
	low, high := 0, len(dataUrut)-1
	indeksDitemukan := -1

	for low <= high {
		mid := low + (high-low)/2
		// Bandingkan dengan tidak peka huruf besar/kecil
		if strings.EqualFold(dataUrut[mid].Jenis, target) {
			indeksDitemukan = mid
			break // Ditemukan!
		} else if strings.ToLower(dataUrut[mid].Jenis) < strings.ToLower(target) {
			low = mid + 1 // Cari di setengah kanan
		} else {
			high = mid - 1 // Cari di setengah kiri
		}
	}

	if indeksDitemukan != -1 {
		// Tambahkan hasil yang ditemukan
		hasil = append(hasil, dataUrut[indeksDitemukan])
		fmt.Printf("\n>> Data dengan jenis '%s' ditemukan:\n", target)
		tampilkanData(hasil)
	} else {
		fmt.Printf("\n>> Data dengan jenis '%s' tidak ditemukan.\n", target)
	}
}

// --- FUNGSI-FUNGSI LAIN (TIDAK BERUBAH) ---

// Menampilkan data dalam format tabel.
func tampilkanData(data []Sampah) {
	fmt.Println("\n--------------------------------------------------------------------------------------")
	fmt.Printf("%-5s | %-20s | %-12s | %-18s | %s\n", "ID", "Jenis", "Jumlah (Kg)", "Status Daur Ulang", "Metode")
	fmt.Println("--------------------------------------------------------------------------------------")
	if len(data) == 0 {
		fmt.Println(">> Belum ada data.")
	} else {
		for _, s := range data {
			fmt.Printf("%-5d | %-20s | %-12.2f | %-18s | %s\n", s.ID, s.Jenis, s.JumlahKg, s.StatusDaurUlang, s.MetodeDaurUlang)
		}
	}
	fmt.Println("--------------------------------------------------------------------------------------")
}

// Menambah data.
func tambahData() {
	fmt.Println("\n--- Tambah Data Baru ---")
	jenis := bacaInput("Masukkan Jenis Sampah: ")
	jumlahKgStr := bacaInput("Masukkan Jumlah Sampah (Kg): ")
	jumlahKg, err := strconv.ParseFloat(jumlahKgStr, 64)
	if err != nil {
		fmt.Println(">> Gagal: Jumlah harus berupa angka.")
		return
	}
	status := "Belum"
	metode := "-"
	statusInput := bacaInput("Sudah didaur ulang? (y/n): ")
	if strings.ToLower(statusInput) == "y" {
		status = "Sudah"
		metode = bacaInput("Metode Daur Ulang: ")
	}
	idTerakhir++
	dataBaru := Sampah{idTerakhir, jenis, jumlahKg, status, metode}
	dataSampah = append(dataSampah, dataBaru)
	fmt.Println("\n>> Data berhasil ditambahkan!")
}

// Mengubah data.
func ubahData() {
	tampilkanData(dataSampah)
	if len(dataSampah) == 0 {
		return
	}
	idStr := bacaInput("\nMasukkan ID data yang ingin diubah: ")
	idUbah, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println(">> Gagal: ID harus berupa angka.")
		return
	}
	for i, s := range dataSampah {
		if s.ID == idUbah {
			fmt.Printf("\nMengubah data untuk '%s'...\n", s.Jenis)
			jenisBaru := bacaInput(fmt.Sprintf("Jenis Baru (sebelumnya: %s): ", s.Jenis))
			if jenisBaru != "" {
				dataSampah[i].Jenis = jenisBaru
			}
			fmt.Println("\n>> Data berhasil diubah!")
			return
		}
	}
	fmt.Printf("\n>> Gagal: Data dengan ID %d tidak ditemukan.\n", idUbah)
}

// Menghapus data.
func hapusData() {
	tampilkanData(dataSampah)
	if len(dataSampah) == 0 {
		return
	}
	idStr := bacaInput("\nMasukkan ID data yang akan dihapus: ")
	idHapus, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println(">> Gagal: ID harus berupa angka.")
		return
	}
	indexDitemukan := -1
	for i, s := range dataSampah {
		if s.ID == idHapus {
			indexDitemukan = i
			break
		}
	}
	if indexDitemukan != -1 {
		konfirmasi := bacaInput(fmt.Sprintf("Yakin ingin menghapus '%s'? (y/n): ", dataSampah[indexDitemukan].Jenis))
		if strings.ToLower(konfirmasi) == "y" {
			dataSampah = append(dataSampah[:indexDitemukan], dataSampah[indexDitemukan+1:]...)
			fmt.Println("\n>> Data berhasil dihapus.")
		} else {
			fmt.Println("\n>> Penghapusan dibatalkan.")
		}
	} else {
		fmt.Printf("\n>> Gagal: Data dengan ID %d tidak ditemukan.\n", idHapus)
	}
}

// Menampilkan statistik.
func tampilkanStatistik() {
	fmt.Println("\n--- Statistik Sampah ---")
	if len(dataSampah) == 0 {
		fmt.Println(">> Belum ada data untuk dihitung.")
		return
	}
	var totalTerkumpul, totalDidaurUlang float64
	for _, s := range dataSampah {
		totalTerkumpul += s.JumlahKg
		if s.StatusDaurUlang == "Sudah" {
			totalDidaurUlang += s.JumlahKg
		}
	}
	fmt.Println("-------------------------------------------------")
	fmt.Printf("Total Sampah Terkumpul      : %.2f Kg\n", totalTerkumpul)
	fmt.Printf("Total Sampah Didaur Ulang   : %.2f Kg\n", totalDidaurUlang)
	fmt.Println("-------------------------------------------------")
}

// Fungsi bantuan untuk membaca input.
func bacaInput(prompt string) string {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}
