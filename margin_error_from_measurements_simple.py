import numpy as np

K = 3
ROUND_FOR_S = 2
ROUND_FINAL = 1
DATA = [
    51.2, 
    51.6,
    51.8,
    50.8,
    51.6,
    51.8,
    52.3,
    51.7,
    51.7,
    51.4
]
N = len(DATA)

# mean of all measurements
d_mean = np.round(np.mean(DATA), ROUND_FOR_S)

# standard deviation of one measurement
s = 0
for d in DATA: s += np.power(d - d_mean, 2)
s = np.round(np.sqrt(np.divide(s, N-1)), ROUND_FOR_S)

# if we were doing this 4real, we should now go and check wether
# all data points are deviating from d_mean no more than k*s remove those that deviate more, 
# recalculate (n, d_mean, s) and check again until there are no deviations

# standard deviation of the mean
s_d = np.round(np.divide(s, np.sqrt(N)), ROUND_FINAL)

margin_error = np.round(K * s_d, ROUND_FINAL)
d = np.round(d_mean, ROUND_FINAL)

print(f"Measurement Result: d = ({d} +- {margin_error})")

rel_limit_err = np.round(np.divide(margin_error, d_mean) * 100, ROUND_FINAL)
print(f"Relative limit error: {rel_limit_err} %")