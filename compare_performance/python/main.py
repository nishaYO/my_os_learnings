import time

def perform_blocking_task():
    print("Blocking task started")
    time.sleep(2)
    print("Blocking task completed")

def perform_simple_task():
    print("Simple task started")
    print("Simple task completed")

perform_blocking_task()
perform_simple_task()

# python is blocking and single threaded language.
# this is what makes python really slow.