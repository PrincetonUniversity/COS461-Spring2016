"""
Creates a simple Gantt chart
Adapted from https://bitbucket.org/DBrent/phd/src/1d1c5444d2ba2ee3918e0dfd5e886eaeeee49eec/visualisation/plot_gantt.py
BHC 2014
"""
 
import matplotlib.pyplot as plt
import matplotlib.font_manager as font_manager
from matplotlib import colors as plot_colors 
from pylab import *

import sys

COLORS = plot_colors.cnames.keys() 

def get_color(ind):
    global COLORS
    return COLORS[ind % (len(COLORS))] 

def waterfall(data, out):

    data_items = data.items()
    data_items = sorted(data_items, key = lambda x : x[1][0])
    labels = [x[0] for x in data_items]

    mymax = 0.5*len(labels)+0.5
    pos = arange(0.5,mymax,0.5)
 
    # Initialise plot
 
    fig = plt.figure()
    ax = fig.add_subplot(111)
    # Plot the data
 
    start, end = data[labels[0]]
    ax.barh(0.55, (end - start), left=start, height=0.4, align='center', color = get_color(0))
    for i in range(0,len(labels)-1):
        start, end = data[labels[i+1]]
        ax.barh((i*0.5)+1.0+0.05, (end - start), left=start, height=0.4, align='center', color = get_color(i + 1))
 
    # Format the y-axis
 
    locsy, labelsy = yticks(pos,labels)
    plt.setp(labelsy, fontsize = 10)
 
    # Format the x-axis
 
    ax.axis('tight')
    ax.set_ylim(ymin = -0.1, ymax = mymax)
    ax.grid(color = 'g', linestyle = ':')
 
    labelsx = ax.get_xticklabels()
    plt.setp(labelsx, rotation=30, fontsize=12)
 
    # Format the legend
 
    font = font_manager.FontProperties(size='small')
    ax.legend(loc=1,prop=font)
 
    # Finish up
    ax.invert_yaxis()
    fig.autofmt_xdate()
    plt.savefig(out)
    #plt.show()

def get_data(fname):
    f = open(fname, 'r')
    data = {}
    for line in f.readlines():
        parts = [x.strip() for x in line.split()]
        if len(parts) > 2:
            data[parts[0]] = [float(parts[1]), float(parts[2])]

    f.close()
    return data

if __name__ == "__main__":
    if len(sys.argv) < 3:
        print "usage: python plot_waterfall.py /path/to/data_file plot_output"

    fname = sys.argv[1]
    out = sys.argv[2]
    data = get_data(fname)
    waterfall(data, out)
