def quicksort(array):
    if len(array) < 2:
        return array
    else:
        pivo = array[0] # Caso recursivo
        print(f"Pivo {pivo}")

        menores = [i for i in array[1:] if i <= pivo] # Array com elementos menores que o pivô
        maiores = [i for i in array[1:] if i > pivo] # Array com elementos maiores que o pivô
        return quicksort(menores) + [pivo] + quicksort(maiores)

print("[19, 5, 2, 4, 5, 10, 44]")
print(quicksort([19, 5, 2, 4, 5, 10, 44]))
print("[5,3,6,1,2]")
print(quicksort([5,3,6,1,2]))
