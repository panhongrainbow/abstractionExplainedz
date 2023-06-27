# T-Test formula

The one-sample t-test formula is as follows:
$$
t = {{\bar x - \mu_0} \over {s \over \sqrt{n}}} - (formula 1)\\

\bar x\ is\ the\ mean\ of\ the\ sample. -(formula 2)\\

\mu_0\ is\ the\ expected\ value\ to\ be\ compared. -(formula 3)\\

s\ is\ the\ standard\ deviation\ of\ the\ sample. s = \sqrt{\frac{{\sum_{i=1}^{n} (x_i - \bar{x})^2}}{n-1}}-(formula 4)\\

n\ represents\ the\ sample\ size -(formula 5)\\

\sqrt\ n\ is\ the\ correction\ factor. -(formula 6)
$$
In `formula 6`:

`sqrt(n)` is `the correction factor`( 縮小) to adjust the sample standard deviation and to `approximate` the standard deviation of `the reference group for comparison`.



(As `the sample size n increases`📈, the calculated `standard deviation becomes smaller`📉.)

因为样本数愈大，数据会开始稳定`

## n and s inversely proportional

Why does the standard deviation become smaller as the sample size n increases ?

(为何样本数增加，标准差会变小？)

The reasons are as follows:

When the sample size is smaller, the calculated standard deviation is less stable and `more affected by extreme values or deviations`. (样本数愈少就愈不稳定，愈容器受到极端值影响)

It is more likely to be amplified, resulting in a larger value.  (因为不稳定，数量少的样本标准差就容意被放大)

To account for this, a correction is applied by dividing the standard deviation by `the square root of n (sqrt(n))`. (所以这就是要除上 sqrt(n) 的原因)

This correction helps stabilize the standard deviation and accounts for the effect of sample size on its estimation. (这是要使用 sqrt(n) 去进行修正的原因，去除 样本数量 对 估计母体标准差时 所造成的误差)

## Standard deviation

Standard deviation also has units. (其实标准差也是有单位的)



The following formula can be observed.

$$
The\ standard\ deviation\ formula\ of\ the\ reference\ group: \\

s = \sqrt{\frac{{\sum_{i=1}^{n} (x_i - {\mu})^2}}{N}} - (formula 7)\\
$$
In the `formulas 7`,

`x_i` represents `individual data points`, `mu` represents `the population mean`.

and `N` represents `the population size`.



Taking scores as an example.

The unit of `(x_i - mu)^2` here is `the square of scores`.

`N` is the correction factor, which has `no units`.

The unit of ((x_i - mu)^2) \over N here is the square of scores.

Therefore, the unit of `sqrt(((x_i - mu)^2) \over N)` here is `the scores`.

(标准差的单位为 分数)


$$
The\ standard\ deviation\ formula\ of\ the\ sample: \\

s = \sqrt{\frac{{\sum_{i=1}^{n} (x_i - \bar{x})^2}}{n-1}} - (formula 8)\\
$$
In the `formulas 8`,

`x_i` represents `individual data points`, `\bar x` represents `the sample mean`.

and `n` represents `the sample size`.



Taking scores as an example.

The unit of `(x_i - \bar x)^2` here is `the square of scores`.

`(n-1)` is the correction factor, which has `no units`.

The unit of `((x_i - \bar x)^2) \over (n-1)` here is `the square of scores`.

Therefore, the unit of `sqrt(((x_i - \bar x)^2) \over (n-1))` here is `the scores`.

(标准差的单位为 分数)

 

Because standard deviation has units, it is possible to normalize it by `removing the units`.

(以下进行去单位化)



In the formulas 1,
$$
t = {{\bar x - \mu_0} \over {s \over \sqrt{n}}} - (formula 1)
$$
Taking scores as an example.

The unit of `(\bar - mu_0)` here is `the scores`.

The unit of `the standard deviation s` here is also `the scores`.

`sqrt(n)` is the correction factor, which has `no units`.

Therefore, the unit of `(\bar - mu_0)` \over s here has `no units`.



The purpose of using standard deviation is to normalize the t-test and `remove the units`.

(所以，使用标准差的目的去把 t 检验进行去单位化)

## Variance

Why do we still need `variance (变异数、方差)` when we already have standard deviation?



To represent `the dispersion(离散度) of a series of numbers`, it is not necessary to rely solely on `standard deviation`; `variance` can also be used. (变异数 和 标准差 都可以表示离散度)



Variance and standard deviation represents different purposes.

When `combining standard deviations`, it is necessary to `square` each individual standard deviation, `sum` them up, and then take `the square root`.

`This approach is more reasonable`. (较合理，用变异数进行合拼比较合理)


$$
S_g = \sqrt{S_1^2+S_2^2+S_3^2} \\
S_g\ is\ the\ combined\ standard\ deviation \\
S_1\, S_2\ and\ S_3\ are\ an\ individual\ and\ independent\ standard\ deviation \\
$$



Therefore, although both variance and standard deviation can be used to represent the dispersion of an array.

`When combining them, variance should be used for the combination`.

(所以虽然 变异数 和 标准差  都可以用来表示数组的离散程度，但进行合拼时，就要用变异数进行合拼)

## The correction factor n-1

### Discuss the reference group

In `formulas 7 and 8`, it can be observed that `one formula` involves dividing by `n−1`, while `the other formula` involves dividing by `n`.



Using an example:

The reference group: `1, 2, 3, 4, 5, 6, 7, 8, 9, 10`

Mean: `5.5`

Using  formula for `the reference group`, `formula 7`:
$$
s = \sqrt{\frac{{\sum_{i=1}^{n} (x_i - {\mu})^2}}{N}} - (formula 7) \\
 = {{(1 - 5.5)^2 + (2 - 5.5)^2 + (3 - 5.5)^2 + (4 - 5.5)^2 + (5 - 5.5)^2 + (6 - 5.5)^2 + (7 - 5.5)^2 + (8 - 5.5)^2 + (9 - 5.5)^2 + (10 - 5.5)^2} \over 10} \\
 = 8.25
$$
`The standard deviation for the reference group` is `8.25`.

### Discuss the Sample

Additionally, there is a sample:

Sample: `1, 4, 7`

Mean: `4`



There are `two methods for calculating` the standard deviation of the sample.

One involves dividing by `n`, while the other involves dividing by `n−1`.



Using the sample formula, `formula 8`, which involves dividing by `n−1`.
$$
s = \sqrt{\frac{{\sum_{i=1}^{n} (x_i - \bar{x})^2}}{n-1}} - (formula 8)\\
= {{(1 - 4)^2 + (4 - 4)^2 + (7 - 4)^2} \over (3-1)} = 9
$$
The other involves dividing by `n`, resulting in the following outcome.
$$
s = {{(1 - 4)^2 + (4 - 4)^2 + (7 - 4)^2} \over 3} = 6
$$


Clearly, `formula 8 is closer to the reference group's standard deviation`, which is why we divide by n−1.

(这也是说，为什么样本的标准差要除上 n - 1 的原因，就是为了要逼近母体的标准差)
