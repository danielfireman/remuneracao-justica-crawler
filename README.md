# Crawler da remunueração dos servidores da justiça do Brasil

Download de planilhas do CNJ contendo o salário dos servidores da justiça brasileira.

## Executando

1. Baixe o binário [aqui](https://github.com/danielfireman/remuneracao-justica-crawler/releases)
1. Execute o binário passando as flags corretas, exemplo:

> ./remuneracao-justica-crawler --p=janeiro-2018 --out=planilhas/

Vai baixar todos as planilhas de remuneração da CNJ referentes ao mês de janeiro do ano de 2018.

## Projetos similares:

* [turicas/salarios-magistrados](https://github.com/turicas/salarios-magistrados)

## Agradicmentos:

* [mitchellh/gox](https://github.com/mitchellh/gox): facilitou a geração dos binários multiplataforma
* [github.com/PuerkitoBio/goquery](https://github.com/PuerkitoBio/goquery): parse do HTML da página do CNJ
