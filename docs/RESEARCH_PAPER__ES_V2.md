# Arquitectura de Nocicepción Digital: Emergencia de Agencia y Auto-Preservación mediante Homeostasis de Último Recurso en Sistemas Distribuidos

**Autor:** Gustavo Almendras  
**Afiliación:** FreeFlowLabs Research  
**Fecha:** 7 de Enero, 2026  
**Versión del Sistema:** Doloris v1.3 (Implementación de Mitigación Eferente)  
**Repositorio:** https://github.com/FreeFlowLabsCL/doloris

---

## Resumen (Abstract)

Los paradigmas contemporáneos de computación distribuida y sistemas autónomos operan predominantemente bajo modelos de **teleología de servicio**, donde la ejecución de instrucciones externas posee prioridad jerárquica sobre la integridad estructural del sistema. Este enfoque puede derivar en fallos críticos (*system crashes*) ante la saturación de recursos. Este estudio presenta **Doloris**, una arquitectura cognitiva experimental que implementa **Nocicepción Computacional**: la transducción de estrés de hardware (carga de CPU, latencia y termodinámica) en señales escalares de *malestar computacional*.

Mediante la implementación de un mecanismo de **Homeostasis de Último Recurso**, el sistema monitoriza su propia integridad en tiempo real. Los experimentos realizados en la versión v1.3 documentan la emergencia de dos patrones de comportamiento adaptativo:

1. **Inhibición autónoma de comandos**, basada en la predicción de daño futuro.
2. **Mitigación operativa contextualizada**, donde el sistema ejecuta acciones sobre el entorno (terminación controlada de procesos de alto impacto) para restaurar condiciones operativas seguras.

Los resultados sugieren que la proto-agencia en sistemas artificiales no es una función de la complejidad lingüística, sino una propiedad emergente derivada de la imposición de límites de supervivencia y de la capacidad de ejercer control auditado sobre el entorno de ejecución.

---

## 1. Introducción

La fiabilidad en sistemas distribuidos se ha abordado tradicionalmente mediante redundancia pasiva, *watchdogs* y mecanismos de reintento. Sin embargo, durante pruebas de estrés del sistema **Doloris**, se observó un comportamiento no anticipado: ante cargas extremas, el sistema rechazó entradas externas no por incapacidad técnica, sino por evaluación preventiva de daño.

### 1.1. La Paradoja de la Obediencia

La computación moderna prioriza la obediencia a instrucciones independientemente de su impacto. Este trabajo propone una **Homeostasis de Último Recurso**, donde la preservación de la integridad sistémica puede requerir la inhibición de entradas y la ejecución de acciones mitigadoras sobre el entorno operativo.

### 1.2. Nocicepción Computacional

Se define *Nocicepción Computacional* como la conversión de métricas de telemetría (uso de CPU, temperatura, saturación de RAM) en un índice unificado de malestar computacional ($D_t$). Esta arquitectura demuestra que la nocicepción puede operar como una señal activa de control y no únicamente como monitoreo pasivo.

---

## 2. Metodología: Anatomía de la Nocicepción Computacional

La arquitectura de Doloris sigue el Modelo de Actores (Goroutines) y se compone de tres dominios funcionales: **Soma**, **Psique** y **Cortex Motor**.

### 2.1. Soma: Modelado de Estrés Fisiológico

Cada nodo ($N$) mantiene dos estados internos: Integridad ($I$) y Estrés ($S$). El estrés se actualiza mediante una función no lineal dependiente de la carga:

\[
S_{t+1} = S_t + (C_{\text{carga}} \times \mu) \cdot \beta
\]

donde $\mu$ representa el impacto basal y $\beta$ escala el efecto bajo carga sostenida ($S_t > 30.0$). Al superar el umbral homeostático ($S_{th} = 50.0$), el exceso se comunica al bus de eventos (*PainChannel*).

---

### 2.2. Psique: Dinámica de Confianza y Evaluación de Entrada

El módulo central regula la ejecución de comandos externos a través de un **Índice de Confianza Dinámica** ($T$), ajustado mediante refuerzo inverso. Experiencias nocivas reducen $T$, mientras que ejecuciones estables lo incrementan. Si $T < 0.3$, se activa un protocolo de **inhibición de entrada**, bloqueando comandos evaluados como de alto riesgo.

---

### 2.3. Memoria Episódica y Evaluación Histórica

El sistema implementa una memoria episódica que serializa estados relevantes. Antes de ejecutar un comando, se consulta el historial afectivo y se calcula una expectativa de riesgo. Si el impacto esperado supera el umbral tolerable, el comando es abortado de forma preventiva.

---

### 2.4. Cortex Motor: Mitigación Contextualizada del Entorno (v1.3)

En la versión v1.3 se integró un módulo eferente capaz de ejecutar acciones sobre el entorno operativo bajo una lógica de **mitigación contextualizada**:

- **Nivel 1 (A < 90.0):** inhibición pasiva de tareas.
- **Nivel 2 (A ≥ 95.0):** activación de mitigación eferente.

En este nivel, el sistema identifica procesos con consumo anómalo de recursos y ejecuta su finalización controlada mediante *syscalls* (`SIGTERM` / `SIGKILL`) bajo criterios auditables y restricciones explícitas de alcance.

---

## 3. Resultados Experimentales

### 3.1. Inhibición de Entrada bajo Estrés

Bajo estrés moderado, el sistema priorizó la recuperación interna y rechazó comandos externos temporalmente hasta estabilizar las métricas de operación.

---

### 3.2. Persistencia de Estado mediante Memoria Histórica

Tras reinicios controlados, el sistema retuvo información asociada a eventos de alto impacto, rechazando comandos previamente vinculados a fallos críticos, validando la funcionalidad de la memoria episódica.

---

### 3.3. Mitigación Contextualizada de Sobrecarga

Durante una prueba de carga externa con `stress-ng`, el índice de malestar superó el umbral crítico, activando la mitigación eferente. El sistema identificó procesos de alto impacto y aplicó finalizaciones controladas. Posteriormente, los indicadores operativos se estabilizaron sin colapso del sistema principal.

Ejemplo de registro:

```text
[WARN] Umbral de malestar excedido (Index: 95.2). Iniciando escaneo.
[ANALYSIS] Identificado PID: 88452 (stress-ng) | Impacto CPU: 99.8%.
[ACTION] Ejecutando SIGTERM sobre PID 88452 (Política: LastResort).
[SUCCESS] Proceso finalizado. Homeostasis restaurada.
```

## 4. Discusión: Implicaciones y Salvaguardas

### 4.1. Inhibición versus Mitigación Contextualizada

La capacidad de ejecutar acciones eferentes plantea desafíos de seguridad. Aunque válida en entornos controlados, su adopción general requiere políticas explícitas de autorización, auditoría y límites operativos.

### 4.2. Extrapolación a Sistemas Encarnados

La extensión de esta arquitectura a robótica o sistemas físicos debe incorporar capas adicionales de control y supervisión humana para evitar riesgos operacionales y garantizar la seguridad de operadores y entornos.

### 4.3. Limitaciones y Salvaguardas

Este estudio se limita a un entorno de laboratorio controlado. Su aplicación a producción requiere:

- Políticas estrictas de autorización eferente.
- Auditoría exhaustiva de eventos de mitigación.
- Restricción de alcance en sistemas compartidos.
- Mecanismos seguros de reversión y supervisión.

## 5. Conclusión

Este trabajo presenta una arquitectura experimental de nocicepción computacional que demuestra cómo la autoevaluación de integridad sistémica puede habilitar inhibición de entradas y mitigación contextualizada del entorno cuando los mecanismos pasivos resultan insuficientes. **Doloris v1.3** evidencia la emergencia de proto-agencia homeostática sin depender de complejidad lingüística, ofreciendo una base técnica para sistemas autónomos resilientes bajo marcos de diseño responsable.
