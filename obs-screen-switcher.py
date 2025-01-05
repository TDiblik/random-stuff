import time
import pyautogui

SLEEP_TIME = 0.1
CODING_SCREEN = ("CODING_SCREEN", ["f1"])
LEFT_SCREEN = ("LEFT_SCREEN", ["f8"])
RIGH_SCREEN = ("RIGH_SCREEN", ["f10"])


def send_command(command):
    # print("send_command: ", command)
    for i in range(0, 5):  # or 10
        pyautogui.hotkey(command, interval=0.001)
    # print("done")


active_screen = CODING_SCREEN
send_command(CODING_SCREEN[0])
print("obs screen switcher started...")
print("use CTRL+C to kill")
while True:
    pos = pyautogui.position()
    if pos.y > 0 and active_screen != CODING_SCREEN[0]:
        active_screen = CODING_SCREEN[0]
        send_command(CODING_SCREEN[1])

    if pos.y < -3:
        if pos.x <= 462 and active_screen != LEFT_SCREEN[0]:
            active_screen = LEFT_SCREEN[0]
            send_command(LEFT_SCREEN[1])
        if pos.x >= 469 and active_screen != RIGH_SCREEN[0]:
            active_screen = RIGH_SCREEN[0]
            send_command(RIGH_SCREEN[1])

    # print(pyautogui.position())
    time.sleep(SLEEP_TIME)
