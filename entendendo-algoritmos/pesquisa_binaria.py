def pesquisa_binaria(lista, item):
    baixo = int(0) # Index 0 da lista
    alto = len(lista) - 1 # Index maximo da lista

    while baixo <= alto: # Enquanto os valores estão na condição verdadeira
        meio = int((baixo + alto) / 2) # Meio da lista
        chute = lista[meio]
 
        if chute == item: # Acha o item
            return meio
        if chute > item: # O chute foi alto#
            alto = meio - 1
        else: # O chute foi baixo
            baixo = meio + 1
    return None # Não achou

minha_lista = [1, 3, 5, 7, 9]

print(pesquisa_binaria(minha_lista, 3))
print(pesquisa_binaria(minha_lista, -1))
