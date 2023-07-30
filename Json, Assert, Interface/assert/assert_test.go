package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHitungVolume(t *testing.T) {
	var hasil = 2 * 3
	assert.Equal(t, hasil, 10, "Perhitungan Volume Salah")
}

// func TestingHitungLuas(t *testing.T) {
// 	assert.Equal(t, kubus.Luas(), luasSeharusnya, "Perhitungan Luas Salah")
// }

// func TestingHitungKeliling(t *testing.T) {
// 	asser.Equal(t, kubus.Keliling(), kelilingSeharusnya, "Perhitungan Keliling Salah")
// }
