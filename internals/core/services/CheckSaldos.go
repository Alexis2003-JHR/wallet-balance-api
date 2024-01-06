package services

import (
	"net/http"

	"wallet/internals/core/models"

	"github.com/gin-gonic/gin"
)

func CheckSaldos(c *gin.Context) {
	c.JSON(http.StatusOK, saldos)
}

var saldos = []models.Saldo{
	{ID: 1, Item: "001", Fecha: "2024-01-06", SaldoActual: 300000.32, Creditos: 45500.34, Debitos: 390000.23},
	{ID: 2, Item: "023", Fecha: "2024-01-06", SaldoActual: 300000.32, Creditos: 45500.34, Debitos: 390000.23},
	{ID: 3, Item: "345", Fecha: "2024-01-06", SaldoActual: 300000.32, Creditos: 45500.34, Debitos: 390000.23},
	{ID: 4, Item: "034", Fecha: "2024-01-06", SaldoActual: 300000.32, Creditos: 45500.34, Debitos: 390000.23},
	{ID: 5, Item: "430", Fecha: "2024-01-06", SaldoActual: 300000.32, Creditos: 45500.34, Debitos: 390000.23},
	{ID: 6, Item: "2323", Fecha: "2024-01-06", SaldoActual: 300000.32, Creditos: 45500.34, Debitos: 390000.23},
	{ID: 7, Item: "107", Fecha: "2024-01-06", SaldoActual: 300000.32, Creditos: 45500.34, Debitos: 390000.23},
	{ID: 8, Item: "122", Fecha: "2024-01-06", SaldoActual: 300000.32, Creditos: 45500.34, Debitos: 390000.23},
	{ID: 9, Item: "233", Fecha: "2024-01-06", SaldoActual: 300000.32, Creditos: 45500.34, Debitos: 390000.23},
	{ID: 10, Item: "102", Fecha: "2024-01-06", SaldoActual: 300000.32, Creditos: 45500.34, Debitos: 390000.23},
}
