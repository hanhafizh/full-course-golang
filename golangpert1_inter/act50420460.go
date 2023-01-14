package main

type Segitiga struct {
	alas, tinggi float64
}

func (k Segitiga) Luas() float64 {
	return (k.alas * k.tinggi) / 2
}

func (k Segitiga) Keliling() float64 {
	return k.alas * 3
}
