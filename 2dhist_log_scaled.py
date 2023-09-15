import matplotlib.pyplot as plt
import numpy as np
import pandas as pd

df = pd.read_csv('./timeMapAdded.csv')
X = df["diffNext"].values
Y = df["diffPrev"].values

x_space = np.logspace(np.log10(1.0), np.log10(1000000.0), 50)
print('x_space:', x_space)

y_space = np.logspace(np.log10(1.0), np.log10(1000000.0), 50)
print('y_space:', y_space)

plt.hist2d(X, Y, bins=(x_space, y_space), cmap=plt.cm.Blues)
plt.yscale('log')
plt.xscale('log')
plt.show()