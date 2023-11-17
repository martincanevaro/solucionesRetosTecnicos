package main

import (
    "fmt"
)

func main() {
    // Temperatura en Kelvin
    var temperaturaKelvin float64 = 373.15

    // Calculamos la temperatura en Celsius
    var temperaturaCelsius float64 = temperaturaKelvin - 273.15

    // Mostramos los resultados en pantalla
    fmt.Println("Temperatura en Kelvin:", temperaturaKelvin)
    fmt.Println("Temperatura en Celsius:", temperaturaCelsius)
}
