

# t-test formula

The one-sample t-test formula is as follows:
$$
t = {{\bar x - \mu_0} \over {s \over \sqrt{n}}} - (formula 1)\\

\bar x\ is\ the\ mean\ of\ the\ sample. -(formula 2)\\

\mu_0\ is\ the\ expected\ value\ to\ be\ compared. -(formula 3)\\

s\ is\ the\ standard\ deviation\ of\ the\ sample. -(formula 4)\\

n represents the sample size -(formula 5)\\

\sqrt\ n\ is\ the\ correction\ factor. -(formula 6)
$$
In formula 6:

sqrt(n) is the correction factor to adjust the sample standard deviation and to approximate the standard deviation of the reference group for comparison.



(As the sample size n increases, the calculated standard deviation becomes smaller.)

## n and s inversely proportional

Why does the standard deviation become smaller as the sample size n increases ?

(为何样本数增加，标准差会变小)

The reasons are as follows:

When the sample size is smaller, the calculated standard deviation is less stable and more affected by extreme values or deviations.. (样本数愈少就愈不稳定，愈容器受到极端值影响)

It is more likely to be amplified, resulting in a larger value.  (因为不稳定，数量少的样本标准差就同意被放大)

To account for this, a correction is applied by dividing the standard deviation by the square root of n (sqrt(n)). (所以这就是要除上 sqrt(n) 的原因)

This correction helps stabilize the standard deviation and accounts for the effect of sample size on its estimation. (这是要使用 sqrt(n) 去进行修正的原因)

## Standard deviation

Standard deviation also has units.



Certainly! Here's the formula for standard deviation in markdown math format:
$$
The\ standard\ deviation\ formula: \\

s = \sqrt{\frac{{\sum_{i=1}^{n} (x_i - \bar{x})^2}}{n-1}}
$$
In the formulas, \(x_i\) represents individual data points, \(\mu\) represents the population mean, \(\bar{x}\) represents the sample mean, \(N\) represents the population size, and \(n\) represents the sample size. The standard deviation (\(\sigma\) or \(s\)) is calculated by taking the square root of the sum of squared differences between each data point and the mean, divided by the appropriate denominator.