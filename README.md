## Calculo do IRT
### Instrução para executar o programa

Este projecto tem o proposito de facilitar as empresas em calcular o IRT de cada funcionário. Para construir a aplicação a base fundamental foi o conceito de Testing Driven Development, que é abordagem que consistem em criar a aplicação apartir dos testes, boas práticas de pprogramação como o KISS, DRY e SRP.

Para que a aplicação seja executada, é necessário gerar o executável no seu terminal de preferência. Por exemplo: 

Para todas as distro Linux:

```
  go build -o cirt 
```

 Para sistema windows:
```
  go build -o cirt.exe 
```

para executar a aplicação deve executar o seguinte comando:

```
 ./cirt -d [dias úteis de trabalho] -f [número de faltas] -a [subsídio de alimentção] -t [subsídio de transporte] -p [premeo] -s [salário base]
```

  Ex.: 
  
  ```
  ./cirt -d 25 -f 1 -a 1000 -t 1000 -p 2000 -s 800000 
```