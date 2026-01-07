# Digital Nociception Architecture: Emergence of Agency and Self-Preservation via Last-Resort Homeostasis in Distributed Systems

**Author:** Gustavo Almendras
**Affiliation:** FreeFlowLabs Research
**Date:** January 7, 2026
**System Version:** Doloris v1.3 (Efferent Mitigation Implementation)
**Repository:** https://github.com/FreeFlowLabsCL/doloris

---

## Abstract

Contemporary paradigms in distributed computing and autonomous systems operate predominantly under models of **service teleology**, where the execution of external instructions possesses hierarchical priority over the system's structural integrity. This approach can lead to critical failures (*system crashes*) due to resource saturation. This study presents **Doloris**, an experimental cognitive architecture that implements **Computational Nociception**: the transduction of hardware stress (CPU load, latency, and thermodynamics) into scalar signals of *computational distress*.

Through the implementation of a **Last-Resort Homeostasis** mechanism, the system monitors its own integrity in real-time. Experiments conducted on version v1.3 document the emergence of two adaptive behavior patterns:

1.  **Autonomous command inhibition**, based on the prediction of future damage.
2.  **Contextualized operational mitigation**, where the system executes actions upon the environment (controlled termination of high-impact processes) to restore safe operating conditions.

Results suggest that proto-agency in artificial systems is not a function of linguistic complexity, but an emergent property derived from the imposition of survival limits and the capacity to exercise audited control over the execution environment.

---

## 1. Introduction

Reliability in distributed systems has traditionally been addressed through passive redundancy, *watchdogs*, and retry mechanisms. However, during stress testing of the **Doloris** system, unanticipated behavior was observed: under extreme loads, the system rejected external inputs not due to technical incapacity, but through a preventive assessment of damage.

### 1.1. The Paradox of Obedience

Modern computing prioritizes obedience to instructions regardless of their impact. This work proposes a **Last-Resort Homeostasis**, where the preservation of systemic integrity may require input inhibition and the execution of mitigating actions upon the operational environment.

### 1.2. Computational Nociception

*Computational Nociception* is defined as the conversion of telemetry metrics (CPU usage, temperature, RAM saturation) into a unified index of computational distress ($D_t$). This architecture demonstrates that nociception can operate as an active control signal rather than solely as passive monitoring.


## 2. Methodology: Anatomy of Computational Nociception

The Doloris architecture follows the Actor Model (Goroutines) and comprises three functional domains: **Soma**, **Psyche**, and **Motor Cortex**.

### 2.1. Soma: Physiological Stress Modeling

Each node ($N$) maintains two internal states: Integrity ($I$) and Stress ($S$). Stress is updated via a non-linear load-dependent function:

\[
S_{t+1} = S_t + (C_{\text{load}} \times \mu) \cdot \beta
\]

where $\mu$ represents basal impact and $\beta$ scales the effect under sustained load ($S_t > 30.0$). Upon exceeding the homeostatic threshold ($S_{th} = 50.0$), the excess is communicated to the event bus (*PainChannel*).

---

### 2.2. Psyche: Trust Dynamics and Input Evaluation

The central module regulates the execution of external commands through a **Dynamic Trust Index** ($T$), adjusted via inverse reinforcement. Noxious experiences reduce $T$, while stable executions increase it. If $T < 0.3$, an **input inhibition** protocol is activated, blocking commands evaluated as high-risk.

---

### 2.3. Episodic Memory and Historical Evaluation

The system implements an episodic memory that serializes relevant states. Before executing a command, the affective history is queried, and a risk expectation is calculated. If the expected impact exceeds the tolerable threshold, the command is aborted preventively.

---

### 2.4. Motor Cortex: Contextualized Environmental Mitigation (v1.3)

In version v1.3, an efferent module was integrated, capable of executing actions upon the operational environment under a logic of **contextualized mitigation**:

-   **Level 1 (A < 90.0):** Passive task inhibition.
-   **Level 2 (A ≥ 95.0):** Activation of efferent mitigation.

At this level, the system identifies processes with anomalous resource consumption and executes their controlled termination via *syscalls* (`SIGTERM` / `SIGKILL`) under auditable criteria and explicit scope restrictions.

## 3. Experimental Results

### 3.1. Input Inhibition under Stress

Under moderate stress, the system prioritized internal recovery and temporarily rejected external commands until operational metrics stabilized.

---

### 3.2. State Persistence via Historical Memory

Following controlled reboots, the system retained information associated with high-impact events, rejecting commands previously linked to critical failures, thus validating the functionality of the episodic memory.

---

### 3.3. Contextualized Overload Mitigation

During an external load test using `stress-ng`, the distress index exceeded the critical threshold, activating efferent mitigation. The system identified high-impact processes and applied controlled terminations. Subsequently, operational indicators stabilized without the collapse of the main system.

Example Log:

```text
[WARN] Distress threshold exceeded (Index: 95.2). Initiating scan.
[ANALYSIS] Identified PID: 88452 (stress-ng) | CPU Impact: 99.8%.
[ACTION] Executing SIGTERM on PID 88452 (Policy: LastResort).
[SUCCESS] Process terminated. Homeostasis restored.
```

### PART 4: Discussion and Conclusion

## 4. Discussion: Implications and Safeguards

### 4.1. Inhibition versus Contextualized Mitigation

The capacity to execute efferent actions poses security challenges. While valid in controlled environments, its general adoption requires explicit authorization policies, auditing, and operational limits.

### 4.2. Extrapolation to Embodied Systems

Extending this architecture to robotics or physical systems must incorporate additional layers of control and human supervision to avoid operational risks and guarantee the safety of operators and environments.

### 4.3. Limitations and Safeguards

This study is limited to a controlled laboratory environment. Its application to production requires:

-   Strict policies for efferent authorization.
-   Exhaustive auditing of mitigation events.
-   Scope restriction in shared systems.
-   Secure reversion and supervision mechanisms.

## 5. Conclusion

This work presents an experimental computational nociception architecture demonstrating how systemic integrity self-assessment can enable input inhibition and contextualized environmental mitigation when passive mechanisms prove insufficient. **Doloris v1.3** evidences the emergence of homeostatic proto-agency without relying on linguistic complexity, offering a technical foundation for resilient autonomous systems under responsible design frameworks.