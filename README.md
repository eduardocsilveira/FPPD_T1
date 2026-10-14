# FPPD_T1

Trabalho 1 de FPPD (2026/2): problema dos filósofos jantando em Go.

Grupo: Henrique Horch, Felipe Araujo e Eduardo Costa.

## Versões

- `v0` - sequencial
- `v1` - um mutex por garfo, todos pegam o esquerdo e depois o direito (dá deadlock)
- `v2` - trava a mesa inteira para comer (sem deadlock, um come por vez)
- `v3` - pega sempre o garfo de menor número primeiro (sem deadlock, até dois comem juntos)

## Como rodar

Precisa do Go instalado. Dentro da pasta do repositório:

```
go run ./v0
go run ./v1
go run ./v2
go run ./v3
```

Para repetir o experimento (5 execuções de cada versão, saídas em `resultados/saidas/`):

```
scripts/rodar.sh
```

O relatório está em `relatorio/Relatorio-Trabalho1-Filosofos.pdf`.
