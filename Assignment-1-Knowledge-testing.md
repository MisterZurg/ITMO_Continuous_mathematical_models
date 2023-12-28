# Assignment 1 — Knowledge testing

## Part A

### 1) Please continue the definition:

$$
f'(x_0) = \lim_{\Delta x \to 0} \frac{f(x_0 + \Delta x) - f(x_0)}{\Delta x}
$$



### 2) Evaluate:

$$
\Big( 2x + (5-x) x^2 - \mathrm{e}^{2x} \Big)' = \Big( 2x + 5x^2 - x^3 - \mathrm{e}^{2x} \Big)' = 2 + 10x - 3x^2 - 2\mathrm{e}^{2x}
$$



### 3) Evaluate:

**Note:** evaluating $\Big( 1 + \frac{1}{x} \Big)^x$ makes sense only if evaluation of the function at a certain point was needed. I assume, differentiation is the task, so:
$$
y = \Big( 1 + \frac{1}{x}\Big)^x \\
\ln y = x \ln \Big( 1 + \frac{1}{x}\Big) \\
\frac{\mathrm{d}}{\mathrm{d}x} \ln y = \frac{\mathrm{d}}{\mathrm{d}x} \Big[ x \ln \Big( 1 + \frac{1}{x}\Big) \Big] \\
\frac{1}{y} \frac{\mathrm{d}y}{\mathrm{d}x} = \Big[ x \underbrace{\frac{\mathrm{d}}{\mathrm{d}x} \ln \Big( 1 + \frac{1}{x} \Big)}_{*} + \ln \Big( 1 + \frac{1}{x} \Big) \Big]
$$

Let's differentiate $*$:
$$
\frac{\mathrm{d}}{\mathrm{d}x} \ln \Big( 1 + \frac{1}{x} \Big) = - \frac{1}{x^2} \frac{1}{\Big( 1 + \frac{1}{x} \Big)} = - \frac{1}{x^2\Big( 1 + \frac{1}{x} \Big)}
$$
Let's continue:
$$
\frac{1}{y} \frac{\mathrm{d}y}{\mathrm{d}x} = \Big[ -x \frac{1}{x^2\Big( 1 + \frac{1}{x} \Big)} + \ln \Big( 1 + \frac{1}{x} \Big) \Big] \\
\frac{1}{y} \frac{\mathrm{d}y}{\mathrm{d}x} = \Big[ - \frac{1}{(x + 1)} + \ln \Big( 1 + \frac{1}{x} \Big) \Big] \qquad \mathrm{Let's~multiply~by~} y \\
\frac{\mathrm{d}y}{\mathrm{d}x} = \Big( 1 + \frac{1}{x}\Big)^x \Big[\frac{-1}{(x + 1)} + \ln \Big( 1 + \frac{1}{x} \Big) \Big]
$$


### 4) Evaluate:

$$
\int\limits_{-5}^5 x^2 \mathrm{d}x = \Big( \frac{x^3}{3} + C \Big) \bigg|_{-5}^5 = \frac{125 - (-125)}{3} = \frac{250}{3}
$$



### 5) Evaluate:

$$
\mathrm{d}(2 - 5^x) = \frac{\mathrm{d}}{\mathrm{d}x} \Big[ 2 - 5^x \Big] \mathrm{d}x = \Big( 0 - 5^x \ln 5 \Big) \mathrm{d}x = - 5^x \ln 5 \cdot \mathrm{d}x
$$



### 6) Please write the Taylor formula for the general case (for the function $f(x)$ in the point $x_0$)

Taylor series for the general case:
$$
f(x_0) + \frac{f'(x_0)}{1!} (x - x_0) + \frac{f''(x_0)}{2!} (x - x_0)^2 + \frac{f'''(x_0)}{3!} (x - x_0)^3 + \cdots
$$


### 7) Find roots:

$$
x^2 + 2(x + 1) = 5 \qquad \mathrm{find~roots} \\
x^2 + 2x - 3 = 0 \\
x_0 = \frac{-2 \pm \sqrt{2^2 + 4\cdot 1\cdot 3}}{2\cdot 1} = \frac{-2 \pm 4}{2} \\
\mathrm{Answer:} \qquad -3 \mathrm{~and~} 1
$$



### 8) Find roots:

$$
(x + 7)x + 7 = 5(x + 1) \qquad \mathrm{find~roots} \\
x^2 + 2x + 2 = 0 \\
x_0 = \frac{-2 \pm \sqrt{2^2 - 4\cdot 1\cdot 2}}{2\cdot 1} = \frac{-2 \pm 2i}{2} = -1 \pm i \\
\mathrm{Answer:} \qquad -1 + i \mathrm{~and~} -1 - i
$$



### 9) Find roots:

$$
x^3 - 3x^2 + 3x - 1 = 0 \qquad \mathrm{find~roots} \\
(x-1)^3 = 0 \\
\mathrm{Answer:} \qquad 1 (\mathrm{of~order~} 3)
$$



### 10) Find solutions:

$$
\frac{\mathrm{d}y}{\mathrm{d}x} = 3xy \\
\frac{\mathrm{d}y}{y} = 3x\mathrm{d}x \\
\int \frac{\mathrm{d}y}{y} = \int 3x\mathrm{d}x \\
\ln y + C_y = \frac{3}{2} x^2 + C_x \\
\ln y = \frac{3}{2} x^2 + C \\
y = C\mathrm{e}^{(x^2 \cdot 3/2)}
$$



### 11) Find solutions:

$$
\frac{\mathrm{d}y}{\mathrm{d}x} = x - y \\
\frac{\mathrm{d}y}{\mathrm{d}x} + y = x \\
\frac{\mathrm{d}y}{\mathrm{d}x} + Py = Q, \qquad P = 1, \quad Q = x \\
yt = \int Qt \mathrm{d}x, \qquad t = \mathrm{e}^{\int P \mathrm{d}x} \\
t = \mathrm{e}^{\int 1\mathrm{d}x} = \mathrm{e}^x \\
y\mathrm{e}^x = \int \underbrace{x}_{u} \underbrace{\mathrm{e}^x \mathrm{d}x}_{\mathrm{d}v} = uv - \int v \mathrm{d}u = x \mathrm{e}^x - \int \mathrm{e}^x \mathrm{d}x = x \mathrm{e}^x - \mathrm{e}^x + C \\
\mathrm{e}^x (x-y) = \mathrm{e}^x - C \\
x - y = 1 - \frac{C}{\mathrm{e}^x} \\
y = x - 1 + \frac{C}{\mathrm{e}^x}
$$



## Part B

A husband went out for bread. After 40 minutes he called and said that he is 1.5 km far from home and still heading to the shop. Considering that he moves with constant speed and there’s no people in the shop (no queuing). 

What distance the husband will go after an hour from his first call?

Will the husband arrive in time for dinner?

- gone at 11:00;
- the dinner is at 15:00;
- the shop is 5 km far from home.

If he is to arrive late, how should he increase his speed to be in time? (Assuming he measures and increases his speed right after leaving the house.)

**Solution:** After $40$ minutes ($2/3$ hours) he was $1.5$ km away from home. His constant speed is equal to $\frac{3/2}{2/3} = \frac{9}{4} = 2.25$ km/h.

**Answer 1:** In 1 hour he will cover the distance of $\frac{9}{4} \cdot 1 = \frac{9}{4} = 2.25$ km. Thus, in total, the distance covered will be $1.5 + 2.25 = 3.75$ km.

**Answer 2:** The time required to go to the shop and return is equal to $\frac{(5+5)}{9/4} = \frac{40}{9} \approx 4.44$ hours. From 11:00 to 15:00 he has only $4$ hours. He will be $\sim 0.44$ hours or $\sim 27$ minutes late.

**Answer 3:** The task is ambiguous:

- either we have to compute the new constant speed required for achieving home at 15:00 (in which case: $\frac{10 \mathrm{km}}{4 \mathrm{h}} = 2.5$ km/h),
- or the speed is no longer constant (assuming that the initial speed is $\frac{9}{4} = 2.25$ km/h and the acceleration is constant). In that case:

$$
\ddot{x} = a \\
\dot{x} = v(t) = at + v_0 \\
x = \frac{at^2}{2} + v_0t + x_0
$$

The total distance is $10$ km, we assume that acceleration is constant by absolute value (the husband will accelerate both on the ways to the shop and back). Let's compute $a$ for $x=10$, $x_0 = 0$, $t = 4$ and $v_0 = 2.25$:
$$
10 = \frac{16a}{2} + 2.25 \cdot 4 + 0 \\
10 = 8a + 9 \\
a = \frac{1}{8} = 0.125
$$
**Final answer 3:** depending on the task:

- either constant speed of $2.5 \;km/h$,
- or constant acceleration of $0.125 \;km/h^2$.