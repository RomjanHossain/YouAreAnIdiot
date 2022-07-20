from tkinter import *
import tkinter as tk
import playsound
from PIL import ImageTk, Image, ImageSequence
import threading
import time
import random
import sys



def disable_event():
    pass

def YouAreAnIdiot():
    global img
    
    root= Tk()
    root.title("Idiot")
    # root.attributes('-toolwindow', True) # bad attribute "-toolwindow": must be -alpha, -topmost, -zoomed, -fullscreen, or -type
    
    # play a sound file 
    playsound.playsound('idiot.mp3', True)


    # root attributes: -alpha, -topmost, -zoomed, -fullscreen, -type
    if sys.platform.startswith('win32') or sys.platform.startswith('cygwin'):
        root.attributes('-toolwindow', True)
    else:
        root.attributes('-alpha', 0.5)

    x= _getRandomX()
    y= _getRandomY()
    # print('this is x: ' + str(x))
    # print('this is y: ' + str(y))
    root.resizable(0,0)
    root.geometry(f'234x171+{x}+{y}')
    
    # change background color in every new window
    
    color = '#%06x' % random.randint(0, 0xFFFFFF)
    root.configure(background=color)



    # Label(root, text="Idiot", font=("Helvetica", 20),background=color ).pack()
    # Label(root, text="You are an idiot", font=("Helvetica", 20),  ).pack()
    # no background color font
    
    Label(root, text="You are an idiot", font=("Franklin Gothic Medium Cond", 23),background=color, anchor='center' ).pack(
        fill=BOTH, expand=1)
    
    


    root.protocol("WM_DELETE_WINDOW", disable_event)
    
    root.mainloop()




values_of_x = []
values_of_y = []
def _getRandomX():
    global values_of_x
    # print(values_of_x)
    x = random.randint(0,999)
    if x not in values_of_x:
        values_of_x.append(x)
        return x
    else:
        return _getRandomX()
        
def _getRandomY():
    global values_of_y
    # print(values_of_y)
    y = random.randint(0,999)
    if y not in values_of_y:
        values_of_y.append(y)
        return y
    else:
        return _getRandomY()




if __name__=='__main__':
    while True:
        thread = threading.Thread(target=YouAreAnIdiot)
        thread.start()
        time.sleep(0.6)
