# quake-log
<br>
<h2><i>Teste Estágio TI - PRAGMA</i><h2>
<hr>
<br>
<h3>Caso não possua o <a href="https://git-scm.com/">Git</a> instalado:</h3>
Em sistemas Debian (Ubuntu, Mint):<br>
<code>sudo apt update && sudo apt upgrade -y && sudo apt install git -y</code>
<hr>
<br>
<h3>Caso não possua o <a href="https://go.dev/">Go</a> instalado:</h3>
Em sistemas Debian (Ubuntu, Mint):<br>
<code>sudo apt update && sudo apt upgrade -y && sudo apt install golang -y</code>
<hr>
<br>
<h3>Para fazer o download, use o comando:</h3>
<code>git clone https://github.com/tomasfn87/quake-log</code>
<hr>
<br>
<h3>Para rodar um dos scripts:</h3>
<code>go run cmd/print-go-structs.go</code><br>
<code>go run cmd/print-json.go</code><br>
<code>go run cmd/print-indented-json.go</code><br>
<code>go run cmd/print-single-indented-json-game-log.go</code><br>
<hr>
<br>
<h3>Argumentos de linha de comando #1:</h3>

<p style="font-size: 0.97em; color: #77f;"><i>O arquivo</i> <b style="font-size: 1.05em; color: #fa3;">cmd/print-indented-json.go</b> <i>aceita como primeiro e segundo argumentos duas strings</i>:</p>

* O comprimento da primeira string será o comprimento da indentação

* O primeiro caracter da segunda string será o caracter da intenção

* Ambos os argumentos são opcionais

<br>
<h3>Exemplos:</h3>
<code>go run cmd/print-indented-json.go aaaa</code><br>

* Indentação de 4 espaços com caracter padrão " " (empty space)

<code>go run cmd/print-indented-json.go length .my_indentation</code><br>

* Indentação de 6 espaços com o primeiro caracter de ".my_indentation" (".")

<br>
<hr>
<h3>Argumentos de linha de comando #2:</h3>
<p style="font-size: 0.97em; color: #77f;"><i>O arquivo</i> <b style="font-size: 1.05em; color: #fa3;">cmd/print-single-indented-json-game-log.go</b> <i>aceita como único argumento um inteiro:</i></p>

* O inteiro é usado para definir qual jogo deverá ser exibido

* Caso o inteiro informado seja menor que 1 ou maior que a quantidade de jogos, o número seja automaticamente ajustado para 1 ou para o último jogo, respectivamente

* Os argumento é opcional: caso não seja informado, será exibido o jogo número 1



<br>
<h3>Exemplos:</h3>
<code>go run cmd/print-single-indented-json-game-log.go 2</code><br>

* Exibe o JSON indentado do segundo jogo

<code>go run cmd/print-single-indented-json-game-log.go 7</code><br>

* Exibe o JSON indentado do sétimo jogo
