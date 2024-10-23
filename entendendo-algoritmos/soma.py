# Exercício 4.2
def soma(arr):
    tamanho = len(arr)

    if tamanho == 0:
        return None
    elif tamanho == 1:
        return arr[0]
    else:
        return arr[0] + soma(arr[1:])

print(soma([]))
print(soma([5]))
print(soma([1,2]))
print(soma([1,2,3]))

# Exercício 4.3
def mais_alto(arr):
    tamanho = len(arr)

    if tamanho == 0:
        return None
    elif tamanho == 1:
        return arr[0]
    else:
        alto = mais_alto(arr[1:])
        return arr[0] if arr[0] > alto else alto

print(mais_alto([]))
print(mais_alto([3]))
print(mais_alto([5,9]))
print(mais_alto([1,24,7,15]))
