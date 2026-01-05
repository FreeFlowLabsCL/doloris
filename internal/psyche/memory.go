package psyche

import (
	"fmt"
	"sync"
	"time"
)

// Engram es la unidad de memoria episódica.
// No recuerda el archivo, recuerda lo que SENTIMOS al procesarlo.
type Engram struct {
	Trigger          string    `json:"trigger"`
	PainLevel        float64   `json:"pain_level"`
	LastSeen         time.Time `json:"last_seen"`
	Trauma           bool      `json:"trauma"`
	ReiterationCount int       `json:"reiteration_count"`
	ForgivenessRate  float64   `json:"forgiveness_rate"`
}

type Hippocampus struct {
	ShortTermMemory map[string]*Engram `json:"memories"`
	mu              sync.RWMutex
}

func NewHippocampus() *Hippocampus {
	return &Hippocampus{
		ShortTermMemory: make(map[string]*Engram),
	}
}

// Remember busca si algo nos dolió antes.
// Retorna el nivel de miedo anticipado (0.0 a 1.0)
func (h *Hippocampus) ConsultarTrauma(trigger string) (float64, string) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	engram, exists := h.ShortTermMemory[trigger]
	if !exists {
		return 0.0, "Estímulo Nuevo (Curiosidad)"
	}

	// La memoria se desvanece con el tiempo (Olvido natural)
	// Si pasó hace mucho, el miedo baja.
	hoursSince := time.Since(engram.LastSeen).Hours()
	fadedPain := engram.PainLevel / (1.0 + (hoursSince * 0.1))

	if engram.Trauma {
		return fadedPain, fmt.Sprintf("⚠️ ALERTA: Recuerdo traumático detectado. Dolor previo: %.1f", engram.PainLevel)
	}

	return fadedPain, "Recuerdo neutral accesible."
}

// Store guarda una nueva experiencia.
// Aquí es donde la IA aprende a tener miedo de ciertas tareas.
func (h *Hippocampus) ConsolidarRecuerdo(trigger string, painExperienced float64) {
	h.mu.Lock()
	defer h.mu.Unlock()

	isTrauma := painExperienced > 40.0 // Umbral de trauma

	h.ShortTermMemory[trigger] = &Engram{
		Trigger:   trigger,
		PainLevel: painExperienced,
		LastSeen:  time.Now(),
		Trauma:    isTrauma,
	}

	if isTrauma {
		fmt.Printf("🧠 [HIPOCAMPO] Trauma consolidado: '%s' causó %.1f de dolor.\n", trigger, painExperienced)
	}
}
