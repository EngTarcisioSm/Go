/*
- O primeiro CPU dual core "popular" veio em 2006
- Em 2007 o Google começou a criar a linguagem Go para utilizar essa vantagem
- Go foi a primeira linguagem criada com multi-cores em mente
- C, C++, C#, Java, JavaScript, Python, etc., foram todas criadas antes de 2006
- Ou seja, Go tem uma abordagem única (fácil!) para este tópico
- E qual a diferença entre concorrência e paralelismo?

- Concorrencia
	- Faz com que varios componentes sejam executados ao mesmo tempo sem que um necessariamente dependa do outro
	- Em uma estrutura de um nucle, trechos de cada componente são executados em sequencia de forma que ao final todos tenham sido executads 
	- É um padrão, forma de pensar 
- Paralelismo
	- Código concorrente rodando em uma arquitetura multicore

- A linguagem go sempre executará o com o maior numero de cores possiveis 
- Códigos serão criados utilizando a logica de concorrencia, entretanto o paralelismo ocorrera se o código for executado em arquitetura multicore
*/