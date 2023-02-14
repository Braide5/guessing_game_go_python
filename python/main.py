import random

print("Welcome to the amazing game!")
random_number = random.randint(0, 100)
score = 0

while True:
    print(f"You have {score} points")
    guess = int(input("Enter a number between 0 and 100: "))

    if guess < random_number:
        print("You guessed too low")

    elif guess > random_number:
        print("You guessed too high")

    else:
        print("You guessed correctly")
        break
