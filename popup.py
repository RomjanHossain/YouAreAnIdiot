from tkinter import *
import tkinter as tk
import playsound

import threading
import time
import random
import sys

def disable_event():
    pass

def move_window():
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


    x= random.randint(0,999)
    y= random.randint(0,999)
    root.resizable(0,0)
    root.geometry(f'235x200+{x}+{y}')

    # change background color in every second
    
    color = '#%06x' % random.randint(0, 0xFFFFFF)
    root.configure(background=color)
    # root.configure(background='black')



    # Label(root, text="Idiot", font=("Helvetica", 20),background=color ).pack()
    # Label(root, text="You are an idiot", font=("Helvetica", 20),  ).pack()
    # no background color font
    Label(root, text="You are an idiot", font=("Franklin Gothic Medium Cond", 23),background=color, anchor='center' ).pack(
        fill=BOTH, expand=1)

    

    root.protocol("WM_DELETE_WINDOW", disable_event)
    root.mainloop()


if __name__=='__main__':
    while True:
        thread = threading.Thread(target=move_window)
        thread.start()
        time.sleep(2)
