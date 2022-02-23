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
<hr>
<br>
<h3>Argumentos de linha de comando:</h3>

<p style="font-size: 1.03em; color: #77f;">O arquivo cmd/print-indented-json.go aceita como primeiro e segundo argumentos 2 strings:</p>

* A primeira string é usada o comprimento como modelo para comprimento da indentação

* A segunda strings fornece apenas seu primeiro caracter como o caracter da intenção

* Ambos os argumentos são opcionais

<br>
<h3>Exemplos:</h3>
<code>go run cmd/print-indented-json.go aaaa</code><br>

* Indentação de 4 espaços com caracter padrão " " (empty space)

<code>go run cmd/print-indented-json.go length .my_indentation</code><br>

* Indentação de 6 espaços com o primeiro caracter de ".my_indentation" (".")