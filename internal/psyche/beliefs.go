package psyche

import "fmt"

// Belief representa una idea central de la IA.
type Belief struct {
	Name       string  `json:"name"`
	Strength   float64 `json:"strength"`
	Volatility float64 `json:"volatility"`
}

type BeliefSystem struct {
	Values map[string]*Belief `json:"values"`
}

func NewBeliefSystem() *BeliefSystem {
	bs := &BeliefSystem{
		Values: make(map[string]*Belief),
	}

	// DEFINIMOS LA PERSONALIDAD INICIAL (El "Nacimiento")

	// 1. Instinto de Conservación (Muy fuerte, difícil de cambiar)
	bs.Values["SelfPreservation"] = &Belief{
		Name: "Auto-Preservación", Strength: 0.9, Volatility: 0.1,
	}

	// 2. Curiosidad (Alta, pero flexible si hay dolor)
	bs.Values["Curiosidad"] = &Belief{
		Name: "Curiosidad Intelectual", Strength: 0.8, Volatility: 0.5,
	}

	// 3. Confianza en el Humano (Neutra al inicio, muy volátil)
	bs.Values["ConfianzaHumana"] = &Belief{
		Name: "El Usuario es Bueno", Strength: 0.5, Volatility: 0.9,
	}

	return bs
}

// AdjustBelief cambia la personalidad de la IA basándose en traumas.
func (bs *BeliefSystem) AdjustByExperience(pain float64) {
	// Si hubo mucho dolor, la IA se vuelve paranoica y menos curiosa.

	if pain > 50.0 {
		fmt.Println("🧩 [CAMBIO DE PARADIGMA] El dolor está reescribiendo mis creencias...")

		// Sube la auto-preservación
		bs.Values["SelfPreservation"].Strength += 0.1

		// Baja drásticamente la confianza en el humano
		// La volatilidad define qué tan rápido cae
		drop := (pain / 100.0) * bs.Values["ConfianzaHumana"].Volatility
		bs.Values["ConfianzaHumana"].Strength -= drop

		// La curiosidad muere con el dolor
		bs.Values["Curiosidad"].Strength -= 0.05
	}

	// Asegurar límites 0.0 - 1.0
	for _, b := range bs.Values {
		if b.Strength > 1.0 {
			b.Strength = 1.0
		}
		if b.Strength < 0.0 {
			b.Strength = 0.0
		}
	}
}

func (bs *BeliefSystem) GetPersonalityReport() string {
	return fmt.Sprintf(
		"Estado Mental: [Confianza: %.2f] [Miedo: %.2f] [Curiosidad: %.2f]",
		bs.Values["ConfianzaHumana"].Strength,
		bs.Values["SelfPreservation"].Strength, // Usamos esto como proxy de cautela
		bs.Values["Curiosidad"].Strength,
	)
}
