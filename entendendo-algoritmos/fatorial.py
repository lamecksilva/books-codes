def fat(x):
    if x == 1:
        return 1
    else:
        return x * fat(x-1)

numero = int(input("Digite o número para calcular o fatorial recursivamente: "))
print(fat(numero))
