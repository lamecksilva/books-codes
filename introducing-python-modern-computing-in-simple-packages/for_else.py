cheeses = ["brie", "gjetost", "harvarti"]

for cheese in cheeses:
    if cheese.startsWith("x"):
        print("I won't eat anything that starts with x")
        break
    else:
        print(cheese)
else:
    print("Didn't find anything that started with x")
