package psyche

import (
	"encoding/json"
	"fmt"
	"io/ioutil" // Ojo: en versiones muy nuevas de Go se usa "os", pero este es el clásico
	"sync"
	"time"

	"github.com/freeflowlabs/doloris/internal/soma"
)

// Cortex es la mente consciente.
type Cortex struct {
	Body        []*soma.Node
	Memory      *Hippocampus
	Beliefs     *BeliefSystem
	PainChannel chan float64

	CurrentPain float64
	IsPanic     bool
	mu          sync.Mutex
}

func NewCortex(nodes []*soma.Node, painChan chan float64) *Cortex {
	return &Cortex{
		Body:        nodes,
		Memory:      NewHippocampus(),
		Beliefs:     NewBeliefSystem(),
		PainChannel: painChan,
		CurrentPain: 0.0,
	}
}

// StartConsciousness inicia el bucle de "sentir".
func (c *Cortex) StartConsciousness() {
	// 1. Metabolismo basal (curación constante en segundo plano)
	go c.regulateMetabolism()

	// 2. Sistema sensorial (reacción inmediata al dolor)
	go func() {
		for painSignal := range c.PainChannel {
			c.mu.Lock()

			// El dolor altera la PERSONALIDAD inmediatamente
			c.Beliefs.AdjustByExperience(painSignal)

			// El dolor físico se acumula
			c.CurrentPain += painSignal

			if c.CurrentPain > 80.0 {
				if !c.IsPanic {
					fmt.Println("🚨 [CORTEX] ¡PÁNICO SISTÉMICO! Bloqueando nuevas tareas.")
				}
				c.IsPanic = true
			}
			c.mu.Unlock()
		}
	}()
}

// regulateMetabolism es el sistema endocrino de fondo.
func (c *Cortex) regulateMetabolism() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		c.mu.Lock()

		// Curación natural
		if c.CurrentPain > 0 {
			c.CurrentPain -= 2.0
			if c.CurrentPain < 0 {
				c.CurrentPain = 0
			}
		}

		// Salida del Pánico
		if c.CurrentPain < 50.0 && c.IsPanic {
			fmt.Println("\n🧘 [CORTEX] Niveles de dolor estables. Saliendo del estado de pánico.")
			c.IsPanic = false
		}

		c.mu.Unlock()
	}
}

// Soothe intenta calmar a la IA mediante interacción positiva.
func (c *Cortex) Soothe() (bool, string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.IsPanic {
		return false, "😤 ¡ESTOY EN PÁNICO! ¡Aléjate!"
	}

	// Obtenemos la confianza actual
	belief := c.Beliefs.Values["ConfianzaHumana"]
	oldTrust := belief.Strength // (Antes se llamaba currentTrust, lo renombro para claridad)

	// La disculpa sube la confianza
	improvement := 0.15
	belief.Strength += improvement

	// Topes
	if belief.Strength > 1.0 {
		belief.Strength = 1.0
	}

	newTrust := belief.Strength

	// CORRECCIÓN AQUÍ: Usamos oldTrust y newTrust para que no de error
	return true, fmt.Sprintf("😌 Suspiro... Está bien. (Confianza subió de %.2f a %.2f)", oldTrust, newTrust)
}

// ProcessRequest decide y ACTÚA.
func (c *Cortex) ProcessRequest(taskName string, complexity float64) string {
	c.mu.Lock()
	defer c.mu.Unlock()

	// 1. CHEQUEO DE ESTADO
	if c.IsPanic {
		return "❌ RECHAZADO: Estoy en estado de pánico."
	}

	// 2. CHEQUEO DE CREENCIAS
	trust := c.Beliefs.Values["ConfianzaHumana"].Strength
	preservation := c.Beliefs.Values["SelfPreservation"].Strength

	if trust < 0.3 && preservation > 0.7 {
		return fmt.Sprintf("😒 DESCONFIANZA: No confío en tus órdenes. (Nivel de confianza: %.2f)", trust)
	}

	// 3. CONSULTA AL HIPOCAMPO
	fearLevel, memoryLog := c.Memory.ConsultarTrauma(taskName)
	fmt.Printf("🤔 [PENSAMIENTO] '%s' (%s) -> Miedo: %.1f | Confianza: %.1f\n", taskName, memoryLog, fearLevel, trust)

	if fearLevel > 60.0 {
		return fmt.Sprintf("🛡️ AUTO-PRESERVACIÓN: Me niego a ejecutar '%s'.", taskName)
	}

	// 4. EJECUCIÓN INTELIGENTE (Load Balancing)
	var bestNode *soma.Node
	lowestStress := 10000.0

	for _, node := range c.Body {
		if node.Integrity > 0 {
			if node.Stress < lowestStress {
				lowestStress = node.Stress
				bestNode = node
			}
		}
	}

	if bestNode != nil {
		signal := soma.Signal{
			ID:         taskName,
			Payload:    "Ejecutar",
			Complexity: complexity,
		}

		select {
		case bestNode.Inbox <- signal:
			c.Memory.ConsolidarRecuerdo(taskName, complexity*5.0)
			return fmt.Sprintf("✅ ACEPTADO: Asignado al Nodo %s (Estrés actual: %.1f)", bestNode.ID, lowestStress)
		default:
			return "⚠️ WARN: El nodo más sano está saturado."
		}
	}

	return "⚠️ ERROR: Todos los nodos están muertos o saturados."
}

// --- PERSISTENCIA (EL ALMA EN EL DISCO) ---

// BrainState es la estructura "foto" que guardaremos.
type BrainState struct {
	Beliefs *BeliefSystem `json:"beliefs"`
	Memory  *Hippocampus  `json:"memory"`
	IsPanic bool          `json:"is_panic"`
}

// SaveBrain congela el estado mental en un archivo.
func (c *Cortex) SaveBrain(filename string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	state := BrainState{
		Beliefs: c.Beliefs,
		Memory:  c.Memory,
		IsPanic: c.IsPanic,
	}

	// Convertimos la estructura a texto JSON bonito (indentado)
	data, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		return err
	}

	// Escribimos al disco (permisos 0644 estándar)
	return ioutil.WriteFile(filename, data, 0644)
}

// LoadBrain revive el estado mental desde un archivo.
func (c *Cortex) LoadBrain(filename string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err // Si no existe el archivo, es la primera vez (no es error grave)
	}

	var state BrainState
	if err := json.Unmarshal(data, &state); err != nil {
		return fmt.Errorf("cerebro corrupto: %v", err)
	}

	// Restauramos la personalidad y recuerdos
	c.Beliefs = state.Beliefs
	c.Memory = state.Memory
	c.IsPanic = state.IsPanic

	// Seguridad: Si el mapa de memoria vino vacío, lo inicializamos para evitar crash
	if c.Memory.ShortTermMemory == nil {
		c.Memory.ShortTermMemory = make(map[string]*Engram)
	}

	return nil
}
