package main

import "testing"

var (
	segitiga           Segitiga = Segitiga{4, 4}
	luasSeharusnya     float64  = 8
	kelilingSeharusnya float64  = 12
)

func TestHitungLuas(t *testing.T) {
	t.Logf("Luas : %.2f", segitiga.Luas())
	if segitiga.Luas() != luasSeharusnya {
		t.Errorf("SALAH!")
	}
}

func TestHitungkeliling(t *testing.T) {
	t.Logf("Keliling : %.2f", segitiga.Keliling())
	if segitiga.Keliling() != kelilingSeharusnya {
		t.Errorf("SALAH!")
	}
}
