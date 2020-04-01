import pandas as pd
import matplotlib.pyplot as plt

def generate_plot(data, file_name):
    data['n_wave/n'] = data['n_wave'] / data['n']
    data.plot(kind='scatter', x='n', y='n_wave/n', color='blue', s=1)
    plt.savefig(file_name)

ks = [2, 3, 10, 100, 400]

for k in ks:
    data = pd.read_csv(f'k_{k}.csv')
    generate_plot(data, f'plot_{k}.png')

for k in ks:
    data = pd.read_csv(f'k_{k}_rep.csv')
    generate_plot(data, f'plot_{k}_rep.png')