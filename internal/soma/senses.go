package soma

import (
	"math"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

// VitalSigns agrupa lo que el cuerpo siente.
type VitalSigns struct {
	CPULoad float64 // Porcentaje de uso de CPU (0-100)
	RAMLoad float64 // Porcentaje de uso de RAM (0-100)
	Pain    float64 // Dolor calculado basado en el estrés del hardware
}

// SenseHardware lee los sensores reales de tu PC (Fedora).
func SenseHardware() VitalSigns {
	// 1. Leer CPU (Promedio de todos los núcleos en los últimos 500ms)
	percent, _ := cpu.Percent(500*time.Millisecond, false)
	cpuVal := 0.0
	if len(percent) > 0 {
		cpuVal = percent[0]
	}

	// 2. Leer Memoria RAM
	v, _ := mem.VirtualMemory()
	ramVal := v.UsedPercent

	// 3. TRADUCCIÓN: De Hardware a Dolor
	// Aquí definimos el umbral de dolor biológico.

	pain := 0.0

	// La CPU estresa mucho (calor/procesamiento)
	// UMBRAL: Si pasa del 40%, empieza a molestar.
	if cpuVal > 40.0 {
		excess := cpuVal - 40.0
		pain += math.Pow(excess, 1.2) // Dolor exponencial
	}

	// La RAM estresa, pero menos (sensación de "mente nublada")
	// UMBRAL: Si pasa del 80%, duele.
	if ramVal > 80.0 {
		excess := ramVal - 80.0
		pain += excess * 1.5
	}

	return VitalSigns{
		CPULoad: cpuVal,
		RAMLoad: ramVal,
		Pain:    pain,
	}
}
