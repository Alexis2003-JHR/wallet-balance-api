package models

type Saldo struct {
	ID          int     `json:"id"`
	Item        string  `json:"item"`
	Fecha       string  `json:"fecha"`
	SaldoActual float32 `json:"saldo_actual"`
	Creditos    float32 `json:"creditos"`
	Debitos     float32 `json:"debitos"`
}
