import string

print(f'{string.whitespace}')
print(string.punctuation)


tasks = 'get gloves,get mask,give cat vitamins,call ambulance'
print(tasks.split(','))
print(tasks.split())

crypto_list = ['Yeti', 'Bigfoot', 'Loch Ness Monster']
crypto_string = ', '.join(crypto_list)
print('Found and signing book deals:', crypto_string)

setup = "a duck goes into a bar..."
print(setup.replace('duck', 'marmoset'))
print(setup)

print(setup.replace('a ', 'a famous ', 100))

world = " earth "
print(world.strip())
