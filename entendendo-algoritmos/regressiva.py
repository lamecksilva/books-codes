def regressiva(i):
    print(i)

    if i <= 1:
        return
    else:
        regressiva(i-1)

numero = int(input("Digite um número para recursão: "))
regressiva(numero)
