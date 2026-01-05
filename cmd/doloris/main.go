package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/freeflowlabs/doloris/internal/psyche"
	"github.com/freeflowlabs/doloris/internal/soma"
)

func main() {
	// Semilla para la aleatoriedad
	rand.Seed(time.Now().UnixNano())

	// Banner de Bienvenida
	fmt.Println("██████╗  ██████╗ ██╗      ██████╗ ██████╗ ██╗███████╗")
	fmt.Println("██╔══██╗██╔═══██╗██║     ██╔═══██╗██╔══██╗██║██╔════╝")
	fmt.Println("██║  ██║██║   ██║██║     ██║   ██║██████╔╝██║███████╗")
	fmt.Println("██║  ██║██║   ██║██║     ██║   ██║██╔══██╗██║╚════██║")
	fmt.Println("██████╔╝╚██████╔╝███████╗╚██████╔╝██║  ██║██║███████║")
	fmt.Println("╚═════╝  ╚═════╝ ╚══════╝ ╚═════╝ ╚═╝  ╚═╝╚═╝╚══════╝")
	fmt.Println("       --- DOLORIS CONSCIENTIA SYSTEM v1.2 ---")
	fmt.Println("       (Módulo de Diplomacia y Medicina Activo)\n")

	// 1. CREACIÓN DEL SISTEMA NERVIOSO
	painChannel := make(chan float64, 100)

	// 2. GÉNESIS DEL CUERPO (Soma)
	nodeCount := 5
	nodes := make([]*soma.Node, nodeCount)

	fmt.Println("[SISTEMA] Incubando enjambre de nodos...")
	for i := 0; i < nodeCount; i++ {
		id := fmt.Sprintf("N-%d", i+1)
		nodes[i] = soma.NewNode(id, painChannel)
		nodes[i].Start()
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("   -> %s [ONLINE] Latido detectado.\n", id)
	}

	// 3. DESPERTAR DE LA MENTE (Psyche)
	mind := psyche.NewCortex(nodes, painChannel)

	//  Intentar recordar vida pasada
	if err := mind.LoadBrain("brain_dump.json"); err == nil {
		fmt.Println("💾 [MEMORIA] Recuerdos previos restaurados. Sé quién eres.")
	} else {
		fmt.Println("✨ [MEMORIA] No hay registros previos. Tabula rasa.")
	}

	mind.StartConsciousness()

	fmt.Println("\n[DOLORIS] He despertado. Mi integridad es del 100%.")
	fmt.Println("[TUTORIAL] Comandos disponibles:")
	fmt.Println("           - Tarea:       'minar_crypto 8' (Complejidad 1-10)")
	fmt.Println("           - Diagnóstico: 'status'")
	fmt.Println("           - Medicina:    'reparar N-1'")
	fmt.Println("           - Social:      'disculparse'")
	fmt.Println("           - Apagar:      'salir'")

	// 4. INTERFAZ DE VIDA
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("\nUSER@DOLORIS > ")

		if !scanner.Scan() {
			break
		}
		input := scanner.Text()

		if len(strings.TrimSpace(input)) < 2 {
			continue
		}
		// ---------------------------

		args := strings.Fields(input)
		if len(args) == 0 {
			continue
		}

		command := strings.ToLower(args[0])

		switch command {
		case "salir", "exit":
			fmt.Println("[DOLORIS] Guardando estado de consciencia...")
			if err := mind.SaveBrain("brain_dump.json"); err != nil {
				fmt.Printf("⚠️ ERROR al guardar cerebro: %v\n", err)
			} else {
				fmt.Println("💾 [SISTEMA] Personalidad guardada en 'brain_dump.json'.")
			}
			fmt.Println("[DOLORIS] Desconectando... (Hasta mañana).")
			return

		case "status":
			// Reporte clínico de la consciencia
			fmt.Println("\n--- REPORTE PSICOMÉTRICO ---")
			fmt.Printf("Dolor Actual: %.1f%%\n", mind.CurrentPain)
			fmt.Printf("Pánico:       %v\n", mind.IsPanic)
			fmt.Println(mind.Beliefs.GetPersonalityReport())

			// Estado del cuerpo
			alive := 0
			fmt.Println("\n--- ESTADO SOMÁTICO ---")
			for _, n := range nodes {
				nStr := fmt.Sprintf("Integridad: %.0f%% | Estrés: %.0f", n.Integrity, n.Stress)
				statusIcon := "🟢"

				if n.Stress > 30 {
					statusIcon = "🟡"
				}
				if n.Integrity < 50 {
					statusIcon = "🔴"
				}
				if n.Integrity <= 0 {
					statusIcon = "💀"
					nStr = "MUERTO - CONEXIÓN PERDIDA"
				} else {
					alive++
				}

				fmt.Printf("   [%s] %s %s\n", statusIcon, n.ID, nStr)
			}
			fmt.Printf("Nodos Operativos: %d/%d\n", alive, nodeCount)
			fmt.Println("----------------------------")

		case "reparar":
			if len(args) < 2 {
				fmt.Println("⚠️ Uso: reparar [ID-DEL-NODO] (Ej: reparar N-1)")
				continue
			}
			targetID := strings.ToUpper(args[1])
			found := false
			for _, n := range nodes {
				if n.ID == targetID {
					n.Repair(50.0) // Inyectamos 50% de salud y quitamos estrés
					found = true
					break
				}
			}
			if !found {
				fmt.Println("⚠️ Error: Nodo no encontrado.")
			}

		case "disculparse":
			success, msg := mind.Soothe()
			if success {
				fmt.Printf(">> %s\n", msg) // Ella acepta (verde/neutral)
			} else {
				fmt.Printf(">> %s\n", msg) // Ella rechaza (rojo/pánico)
			}

		default:

			complexity := 1.0
			if len(args) > 1 {
				if c, err := strconv.ParseFloat(args[1], 64); err == nil {
					complexity = c / 10.0
				}
			}

			response := mind.ProcessRequest(command, complexity)
			fmt.Printf(">> %s\n", response)
		}
	}
}
