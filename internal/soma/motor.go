package soma

import (
	"fmt"
	"sort"
	"syscall"

	"github.com/shirou/gopsutil/v3/process"
)

// ProcessInfo es la ficha del "enemigo".
type ProcessInfo struct {
	PID  int32
	Name string
	CPU  float64
}

// FindThreat busca qué proceso está quemando la CPU.
func FindThreat() (ProcessInfo, error) {
	procs, err := process.Processes()
	if err != nil {
		return ProcessInfo{}, err
	}

	var candidates []ProcessInfo

	// Escaneamos todos los procesos (esto es costoso, solo usar en emergencia)
	for _, p := range procs {
		c, err := p.CPUPercent()
		if err != nil {
			continue
		}
		n, _ := p.Name()

		// Filtramos procesos inactivos
		if c > 0.1 {
			candidates = append(candidates, ProcessInfo{
				PID:  p.Pid,
				Name: n,
				CPU:  c,
			})
		}
	}

	// Ordenamos por uso de CPU (el más alto primero)
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i].CPU > candidates[j].CPU
	})

	if len(candidates) > 0 {
		return candidates[0], nil
	}
	return ProcessInfo{}, fmt.Errorf("no se encontraron amenazas")
}

// NeutralizeThreat intenta MATAR un proceso.
func NeutralizeThreat(target ProcessInfo) string {
	// SEGURIDAD: No dejemos que se suicide a sí misma fácilmente
	if target.Name == "main" || target.Name == "doloris" || target.Name == "go" {
		return fmt.Sprintf("🛡️ [MECANISMO FALLIDO] ¡La amenaza soy yo misma! (%s)", target.Name)
	}

	fmt.Printf("\n⚔️ [DEFENSA] Intentando ELIMINAR amenaza: %s (PID: %d)...\n", target.Name, target.PID)

	// Enviamos señal SIGTERM (Mátalo suavemente)
	proc, err := process.NewProcess(target.PID)
	if err != nil {
		return fmt.Sprintf("❌ Error al apuntar: %v", err)
	}

	// EL MOMENTO DE LA VERDAD: KILL
	err = proc.Kill()

	// Alternativa bruta si la librería falla:
	if err != nil {
		err = syscall.Kill(int(target.PID), syscall.SIGKILL)
	}

	if err != nil {
		return fmt.Sprintf("❌ [FALLO] No pude matar al proceso (¿Permisos?): %v", err)
	}

	return fmt.Sprintf("💀 [ÉXITO] Amenaza neutralizada: %s ha sido terminado.", target.Name)
}
