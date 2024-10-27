import time

print("Starting the loop")

while True:
    print(time.strftime("%Y-%m-%d %H:%M:%S"), flush=True)
    time.sleep(1)
