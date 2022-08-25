# Exercícios

## Capítulo 1

1.1 
> log2(128) = 7 etapas

1.2 
> log2(256) = 8 etapas

1.3
> O(log(n)), visto que é possível ser feito utilizando a busca binária.

1.4
> Como é preciso percorrer a agenda inteira, ou seja, todo os elementos, logo sabemos que é O(n).

1.5
> Mesma coisa do caso anterior; precisamos percorrer todos os elementos. O(n).

1.6
> O(n), porque a anotação Big O desconsidera constantes. 

## Capítulo 2

2.1
> Lista ordernada, porque ela leva tempo constante O(1) para inserções, e como visto, o aplicativo terá muito mais inserções do que leituras.

2.2
> Usaria uma lista encadeada, porque tanto para adicionar no começo como para remover do final o tempo de execução seria O(1), basta manter a referência na lista encadeada para o último elemento. E também pelo fato de múltiplas inserções estarem acontecendo. 

2.3
> Como é um acesso aleatório, logicamente usaremos um array. Com a lista encadeada, ele teria que percorrer elemento por elemento até o meio da lista. 

2.4
> O problema do Array com relação as inserções são dois: 
> - Os Arrays tem um tamanho fixo. Caso esse tamanho for ultrapassado, todos os elementos devem ser movidos de um lugar para outro. 
> - Caso essa inserção ocorra no começo do Array, todos os outros elementos devem ser deslocados para a direita. 
> Outro problema é o fato de que para utilizar a lista ordenada os Arrays precisarão estar ordenados. Então a cada nova inserção será necessário ordenar todos os Arrays. 


## Capítulo 3

3.1 
> Existe uma função principal chamada "Sauda" que, por sua vez, irá chamar uma função chamada "Sauda2". Provavelmente a função "Sauda" passa o valor da variável "nome" para a função "Sauda2". 
  
3.2 
> A pilha irá crescer infinitamente até estourar a memória. 

## Capítulo 4

4.1 
> Solução em: [Link](https://github.com/iagosaito/tech-book-exercises/blob/main/entendendo-algoritmos-um-guia-para-programadores-e-outros-curiosos/4.1_recursive_sum_of_elements_in_list.go)

4.2
> Solução em: [Link](https://github.com/iagosaito/tech-book-exercises/blob/main/entendendo-algoritmos-um-guia-para-programadores-e-outros-curiosos/4.2_recursive_count_number_of_elements_in_list.go)

4.3
> Solução em: [Link](https://github.com/iagosaito/tech-book-exercises/blob/main/entendendo-algoritmos-um-guia-para-programadores-e-outros-curiosos/4.3_recursive_highest_number_in_list.go)

4.4 
> O caso base é quando o tamanho da lista é 1. O caso recursivo é dividir a lista em dois e pegar o elemento do meio e aplicar a busca binária novamente. 

4.5
> O(n)

4.6
> O(n)

4.7
> O(1)

4.8
> O(n*n)

Extra
> Criei um algoritmo recurso do quicksort: [Link](https://github.com/iagosaito/tech-book-exercises/blob/main/entendendo-algoritmos-um-guia-para-programadores-e-outros-curiosos/4.5_quicksort.go)

## Capítulo 5

5.1 
> Consistente, porque sempre vai retornar o mesmo valor.  

5.2
> Não é consistente, porque retorna valores diferentes para a mesma entrada. 

5.3
> Não é consistente, porque o índice vazio muda, mesmo que a entrada seja a mesma. 

5.4 
> É consistente, para a mesma entrada, ela devolve o mesmo valor. 

5.5
> a. Pior alternativa. Retornará uma uma tabela hash com colisão em todos os elementos, visto que a função sempre retorna o mesmo resultado. 
> b. Alternativa ruim. Ainda assim haverá muitas colisões, visto que a lista possui 3 pessoas com nome do mesmo tamanho. 
> c. Alternativa média, ele irá garantir um bom espalhamento, mas mesmo assim haverá uma colisão. 
> d. A melhor alternativa, visto que garante o espalhamento da tabela hash.

5.6 
> a. Pior alternativa, haverá colisão em todos os elementos. 
> b. Excelente alternativa, dado que os tamanhos são diferentes, cada função hash hash irá mapear para um espaço diferente. 
> c. Empatado como pior alternativa, porque todos os tamanhos começam com a mesma letra 'A'. 
> d. Péssima alternativa, visto que o cálculo da função irá retornar 0 em todos os resultados. 

