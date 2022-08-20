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

> Solução em [Link](entendendo-algoritmos-um-guia-para-programadores-e-outros-curiosos\4.1_recursive_sum_of_elements_in_list.go)











