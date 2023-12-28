# Individual modeling project
Individual modeling project The goal of this work is to create a mathematical model in the form
of a differential equation to estimate the Monkeypox virus dynamics.

> [!INFO]
> The model is based on the **SIERD** because it is useful for predicting the spread of infectious diseases it takes into account the natural history of the disease, including the incubation period, the duration of illness, and the probability of death or recovery. By analyzing data on the spread of the disease, researchers can estimate the parameters of the model and make predictions about future trends.

**Parameters:**
- *Susceptible* — These are individuals who are not yet infected with the disease and are therefore able to contract it.
- *Exposed* — These are individuals who have come into contact with someone who has the disease, but have not yet developed symptoms themselves.
- *Infected* — These are individuals who have the disease and are capable of spreading it to others.
- *Recovered* — These are individuals who have had the disease and have fully recovered, either naturally or after receiving medical treatment.
- *Deceased* — These are individuals who have died from the disease. 


The model can be written in form of Differential Equation:
$$
\frac{dSuspectable}{dt}=-Rate_{transmission}*\frac{Suspectable*Infected}{Population} \\
\frac{dExposed}{dt}=Rate_{transmission}\frac{Suspectable*Infected}{Population}-Rate_{incubation}*Exposed \\
\frac{dInfected}{dt}=Rate_{incubation}*Exposed-Rate_{recovery}*Infected - Rate_{death}*Infected \\
\frac{dRecovered}{dt}= Rate_{recovery}*Infected \\
\frac{dDeceased}{dt} = Rate_{death}*\frac{dInfected}{dt}
$$

Initial params:
- 4003250000 — Initial number of infected individuals
- 0,1— Transmission rate (per day)
- 0,1 — Incubation rate (per day)
- 0,001 — Recovery rate (per day)
- 0,1 — Death rate (per day)
- 7000 — Exposed1 initial number of exposed individuals
- 0 — Recovered1 initial number of recovered individuals

Dataset: [Monkey Pox (Next Covid?)](https://www.kaggle.com/datasets/programmerrdai/monkey-pox-next-covid)