# Exercícios

## Capítulo 1

1.1 - Suponha que você tenha uma lista com 128 nomes e esteja fazendo uma pesquina binária. Qual seria o número máximo de etapas que você levaria para encontrar o nome desejado? 

    R: log2(128) = 7 etapas

1.2 - Suponha que você duplique o tamanho da lista. qual seria o número máximo de etapas agora? 

    R: log2(256) = 8 etapas

Forneça o tempo de execução para cada um dos casos a seguir em termos da notação Big O.

1.3 - Você tem um nome e deseja encontrar o número de telefone para esse nome em uma agenda telefônica

    R: O(log(n)), visto que é possível ser feito utilizando a busca binária.

1.4 - Você tem um número de telefone e deseja encontrar o dono dele em uma agenda telefônica. (Dica: Deve procurar pela agenda inteira!!)

    R: Como é preciso percorrer a agenda inteira, ou seja, todo os elementos, logo sabemos que é O(n).

1.5 - Você quer ler o número de cada pessoa na agenda eletrônica.

    R: Mesma coisa do caso anterior; precisamos percorrer todos os elementos. O(n).

1.6 - Você quer ler os números apenas dos nomes que começam com A. (Isso é complicado! Esse algoritmo envolve conceitos que são abordados mais profundamente no capítulo 4. Leia a resposta - você ficara surpreso!)

    R: O(n), porque a anotação Big O desconsidera constantes. 

