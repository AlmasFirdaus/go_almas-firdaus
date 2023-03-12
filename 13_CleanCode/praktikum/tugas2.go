package main

type Kendaraan struct {
	roda      int
	kecepatan int
}

type Mobil struct {
	Kendaraan
}

func (m *Mobil) berjalan() {
	m.tambahKecepatan(10)
}

func (m *Mobil) tambahKecepatan(kecepatanBaru int) {
	m.kecepatan = m.kecepatan + kecepatanBaru
}

func main() {
	mobilcepat := Mobil{}
	mobilcepat.berjalan()
	mobilcepat.berjalan()
	mobilcepat.berjalan()

	mobillamban := Mobil{}
	mobillamban.berjalan()
}
