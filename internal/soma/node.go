package soma

import (
	"fmt"
	"math"
	"sync"
	"time"
)

// Signal es el impulso eléctrico que viaja por la red.
type Signal struct {
	ID         string
	Payload    string
	Complexity float64 // 0.0 a 1.0 (Fuente de estrés)
}

// NodeState define el estado clínico.
type NodeState int

const (
	StateHealthy NodeState = iota
	StateStressed
	StateCritical
	StateDead
)

type Node struct {
	ID        string
	Integrity float64 // 100.0 (Sano) -> 0.0 (Muerto)
	Stress    float64

	Inbox        chan Signal
	Outbox       chan Signal
	PainReceptor chan<- float64 // Conexión al Cortex

	// Estado interno
	state              NodeState
	inRefractoryPeriod bool // Si está en "respiro"
	mu                 sync.Mutex
}

// NewNode nace una nueva célula.
func NewNode(id string, painChannel chan<- float64) *Node {
	return &Node{
		ID:           id,
		Integrity:    100.0,
		Stress:       0.0,
		Inbox:        make(chan Signal, 10),
		Outbox:       make(chan Signal, 10),
		PainReceptor: painChannel,
		state:        StateHealthy,
	}
}

// Start activa el metabolismo del nodo (Bucle de vida).
// Sin esto, el nodo existe pero no hace nada.
func (n *Node) Start() {
	go func() {
		for signal := range n.Inbox {
			// Si está muerto, el canal sigue abierto pero ignoramos la señal
			n.mu.Lock()
			if n.state == StateDead {
				n.mu.Unlock()
				continue
			}
			n.mu.Unlock()

			n.processSignal(signal)
		}
	}()
}

// processSignal es la lógica que tú definiste (Estrés Exponencial).
func (n *Node) processSignal(sig Signal) {
	n.mu.Lock()
	// REGLA: Periodo Refractario (Respiro)
	if n.inRefractoryPeriod {
		n.mu.Unlock()
		// Rechazamos silenciosamente o podríamos loguear "Busy"
		return
	}
	n.mu.Unlock()

	// 1. Lógica de Estrés EXPONENCIAL
	stressImpact := sig.Complexity * 10.0

	n.mu.Lock()
	// Si ya hay ansiedad (>30), el impacto se multiplica (Bola de Nieve)
	if n.Stress > 30.0 {
		stressImpact *= 1.5
	}
	n.Stress += stressImpact
	currentStress := n.Stress
	n.mu.Unlock()

	// 2. Procesamiento simulado (Latencia por carga)
	// Mientras más estrés, más lento piensa
	latency := time.Duration(100+(currentStress*5)) * time.Millisecond
	time.Sleep(latency)

	// 3. Daño Exponencial (Tu regla de oro)
	if currentStress > 50.0 {
		excessStress := currentStress - 50.0

		// FÓRMULA TUYA: Daño = (Exceso ^ 1.5) * 0.05
		damage := math.Pow(excessStress, 1.5) * 0.05

		n.mu.Lock()
		n.Integrity -= damage
		currentIntegrity := n.Integrity
		n.mu.Unlock()

		// Gritar dolor al Cortex
		if currentIntegrity > 0 {
			// Envío no bloqueante para no detener el proceso si el cerebro está saturado
			select {
			case n.PainReceptor <- damage:
			default:
			}
		}
	}

	// 4. Chequeo de Muerte
	n.mu.Lock()
	if n.Integrity <= 0 {
		n.die() // Muerte definitiva
		n.mu.Unlock()
		return
	}
	n.mu.Unlock()

	// 5. El "Respiro" (Recuperación)
	if currentStress > 60.0 {
		n.takeBreather()
	} else {
		// Recuperación pasiva
		n.mu.Lock()
		n.Stress = math.Max(0, n.Stress-5.0)
		n.mu.Unlock()
	}

	// Si sobrevivió, enviamos confirmación (opcional)
	fmt.Printf("   -> [NODO %s] Tarea terminada. (Estrés: %.1f)\n", n.ID, currentStress)
}

func (n *Node) takeBreather() {
	n.mu.Lock()
	n.inRefractoryPeriod = true
	// fmt.Printf("💤 [NODO %s] Sobrecargado. Tomando un respiro...\n", n.ID)
	n.mu.Unlock()

	time.Sleep(500 * time.Millisecond)

	n.mu.Lock()
	// Al volver, el estrés baja drásticamente
	n.Stress = math.Max(0, n.Stress-30.0)
	n.inRefractoryPeriod = false
	n.mu.Unlock()

	// fmt.Printf("🔋 [NODO %s] Listo de nuevo.\n", n.ID)
}

func (n *Node) die() {
	if n.state == StateDead {
		return
	}

	n.state = StateDead
	n.Integrity = 0

	// El grito final de muerte (Dolor máximo)
	select {
	case n.PainReceptor <- 100.0:
	default:
	}

	fmt.Printf("💀 [NODO %s] HA MUERTO. Conexión perdida.\n", n.ID)
}

// Repair repara la integridad del nodo artificialmente.
func (n *Node) Repair(amount float64) {
	n.mu.Lock()
	defer n.mu.Unlock()

	if n.state == StateDead {
		fmt.Printf("🔧 [SISTEMA] No se puede reparar el Nodo %s. Está muerto. Requiere reinicio.\n", n.ID)
		return
	}

	n.Integrity += amount
	if n.Integrity > 100.0 {
		n.Integrity = 100.0
	}

	// Reparar también baja el estrés
	n.Stress = 0

	fmt.Printf("✨ [MANTENIMIENTO] Nodo %s reparado. Integridad: %.1f%%\n", n.ID, n.Integrity)
}
